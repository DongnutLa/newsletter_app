package server

import (
	"log"

	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	fiber "github.com/gofiber/fiber/v2"
)

type Server struct {
	//We will add every new Handler here
	userHandlers       ports.UserHandlers
	adminHandlers      ports.AdminHandlers
	newsletterHandlers ports.NewsletterHandlers
	filesHandlers      ports.FilesHandlers
	fileMiddleware     fiber.Handler
	//middlewares ports.Middlewares
}

func NewServer(uHandlers ports.UserHandlers, aHandlers ports.AdminHandlers, nHandlers ports.NewsletterHandlers, fHandlers ports.FilesHandlers, fileMiddleware fiber.Handler) *Server {
	return &Server{
		userHandlers:       uHandlers,
		adminHandlers:      aHandlers,
		newsletterHandlers: nHandlers,
		filesHandlers:      fHandlers,
		fileMiddleware:     fileMiddleware,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	usersRoute := v1.Group("/users")
	usersRoute.Get("/register", s.userHandlers.RegisterToNewsletter)
	usersRoute.Get("/unregister", s.userHandlers.UnregisterToNewsletter)

	newslettersRoute := v1.Group("/newsletter")
	newslettersRoute.Post("/send", s.newsletterHandlers.SendNewsletter)
	newslettersRoute.Post("/schedule", s.newsletterHandlers.ScheduleNewsletter)

	adminsRoutes := v1.Group("/admin")
	adminsRoutes.Post("/login", s.adminHandlers.Login)

	filesRoutes := v1.Group("/files")
	filesRoutes.Post("/upload", s.fileMiddleware, s.filesHandlers.SaveFile)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
