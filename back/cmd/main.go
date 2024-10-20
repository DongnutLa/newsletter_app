package main

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/core/services"
	"github.com/DongnutLa/newsletter_app/internal/handlers"
	"github.com/DongnutLa/newsletter_app/internal/middlewares"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/DongnutLa/newsletter_app/internal/server"
	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
	gomail "gopkg.in/mail.v2"
)

func main() {
	utils.LoadConfig()

	conn := repositories.NewMongoDB(context.TODO())
	logger := zerolog.New(os.Stderr)

	host := utils.GetConfig("smtp_host")
	port := utils.GetConfig("smtp_port")
	user := utils.GetConfig("smtp_user")
	pwd := utils.GetConfig("smtp_pwd")
	jwtKey := utils.GetConfig("jwt_key")

	intPort, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("Invalid smtp port %s", err))
	}

	dialer := gomail.NewDialer(host, intPort, user, pwd)
	engine := html.New("./internal/template", ".html")
	// engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	//repositories
	userRepository := repositories.NewUserRepository(context.TODO(), "users", conn.Database, &logger)
	adminRepository := repositories.NewAdminRepository(context.TODO(), "admins", conn.Database, &logger)
	newsletterRepository := repositories.NewNewsletterRepository(context.TODO(), "newsletters", conn.Database, &logger)
	topicRepository := repositories.NewTopicRepository(context.TODO(), "topics", conn.Database, &logger)

	services.MessagingInit()
	eventType := ports.UseBUS
	messaging := services.NewEventMessaging(&logger, eventType)

	//services
	jwtService := services.NewJwtService([]byte(jwtKey), &logger)
	userService := services.NewUserService(context.TODO(), &logger, userRepository, messaging)
	adminService := services.NewAdminService(context.TODO(), &logger, adminRepository, jwtService)
	newsletterService := services.NewNewsletterService(context.TODO(), &logger, newsletterRepository, messaging)
	fileService := services.NewFilesService(context.TODO(), &logger)
	mailService := services.NewMailService(context.TODO(), &logger, dialer, engine)
	topicService := services.NewTopicService(context.TODO(), &logger, topicRepository)

	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	adminHandlers := handlers.NewAdminHandlers(adminService)
	newsletterHandlers := handlers.NewNewsletterHandlers(newsletterService)
	fileHandlers := handlers.NewFileHandlers(fileService)
	topicHandlers := handlers.NewTopicHandlers(topicService)

	//Middlewares
	authMiddleware := middlewares.NewAuthMiddleware(&logger, jwtService)
	fileMiddleware := middlewares.NewFileMiddleware(&logger)

	// ========= Messaging =========
	msgNow := time.Now()
	var wg sync.WaitGroup
	defer wg.Wait()

	handler := handlers.NewEventsHandler(context.TODO(), &logger, mailService, newsletterService)
	handler.Start(&wg)
	defer handler.Stop()
	logger.Info().Msgf("Messaging init time: %dms", time.Since(msgNow).Milliseconds())
	// ========= Messaging =========

	//server
	httpServer := server.NewServer(
		userHandlers,
		adminHandlers,
		newsletterHandlers,
		fileHandlers,
		topicHandlers,
		fileMiddleware,
		authMiddleware,
		engine,
	)
	httpServer.Initialize()
}
