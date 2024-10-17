package server

import (
	"log"

	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	//We will add every new Handler here
	userHandlers       ports.UserHandlers
	adminHandlers      ports.AdminHandlers
	newsletterHandlers ports.NewsletterHandlers
	topicHandlers      ports.TopicHandlers
	filesHandlers      ports.FilesHandlers
	fileMiddleware     fiber.Handler
	authMiddleware     fiber.Handler
	//middlewares ports.Middlewares
}

func NewServer(
	uHandlers ports.UserHandlers,
	aHandlers ports.AdminHandlers,
	nHandlers ports.NewsletterHandlers,
	fHandlers ports.FilesHandlers,
	tHandlers ports.TopicHandlers,
	fileMiddleware fiber.Handler,
	authMiddleware fiber.Handler,
) *Server {
	return &Server{
		userHandlers:       uHandlers,
		adminHandlers:      aHandlers,
		newsletterHandlers: nHandlers,
		filesHandlers:      fHandlers,
		topicHandlers:      tHandlers,
		fileMiddleware:     fileMiddleware,
		authMiddleware:     authMiddleware,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	usersRoute := v1.Group("/users")
	usersRoute.Get("", s.authMiddleware, s.userHandlers.ListUsers)
	usersRoute.Get("/register", s.userHandlers.RegisterToNewsletter)
	usersRoute.Get("/unregister", s.userHandlers.UnregisterToNewsletter)

	newslettersRoute := v1.Group("/newsletter")
	newslettersRoute.Get("/", s.authMiddleware, s.newsletterHandlers.ListNewsletters)
	newslettersRoute.Post("/", s.authMiddleware, s.newsletterHandlers.CreateNewsletter)
	newslettersRoute.Post("/send", s.authMiddleware, s.newsletterHandlers.SendNewsletter)
	newslettersRoute.Post("/schedule", s.authMiddleware, s.newsletterHandlers.ScheduleNewsletter)

	adminsRoutes := v1.Group("/admin")
	adminsRoutes.Post("/login", s.adminHandlers.Login)

	filesRoutes := v1.Group("/files")
	filesRoutes.Post("/upload", s.authMiddleware, s.fileMiddleware, s.filesHandlers.SaveFile)

	topicsRoutes := v1.Group("/topics")
	topicsRoutes.Get("/", s.topicHandlers.ListTopics)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
