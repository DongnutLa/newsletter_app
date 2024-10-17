package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/DongnutLa/newsletter_app/internal/core/services"
	"github.com/DongnutLa/newsletter_app/internal/handlers"
	"github.com/DongnutLa/newsletter_app/internal/middlewares"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/DongnutLa/newsletter_app/internal/server"
	"github.com/DongnutLa/newsletter_app/internal/utils"
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

	//repositories
	userRepository := repositories.NewUserRepository(context.TODO(), "users", conn.Database, &logger)
	adminRepository := repositories.NewAdminRepository(context.TODO(), "admins", conn.Database, &logger)
	newsletterRepository := repositories.NewNewsletterRepository(context.TODO(), "newsletters", conn.Database, &logger)

	//services
	jwtService := services.NewJwtService([]byte(jwtKey), &logger)
	userService := services.NewUserService(context.TODO(), &logger, userRepository)
	adminService := services.NewAdminService(context.TODO(), &logger, adminRepository, jwtService)
	newsletterService := services.NewNewsletterService(context.TODO(), &logger, newsletterRepository, dialer)
	fileService := services.NewFilesService(context.TODO(), &logger)

	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	adminHandlers := handlers.NewAdminHandlers(adminService)
	newsletterHandlers := handlers.NewNewsletterHandlers(newsletterService)
	fileHandlers := handlers.NewFileHandlers(fileService)

	//Middlewares
	authMiddleware := middlewares.NewAuthMiddleware(&logger, jwtService)
	fileMiddleware := middlewares.NewFileMiddleware(&logger)

	//server
	httpServer := server.NewServer(
		userHandlers,
		adminHandlers,
		newsletterHandlers,
		fileHandlers,
		fileMiddleware,
		authMiddleware,
	)
	httpServer.Initialize()
}
