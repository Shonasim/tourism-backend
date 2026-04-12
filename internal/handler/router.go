package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	userHandler        *UserHandler
	tourHandler        *TourHandler
	bookingHandler     *BookingHandler
	paymentHandler     *PaymentHandler
	reviewHandler      *ReviewHandler
	destinationHandler *DestinationHandler
}

func NewRouter(
	userHandler *UserHandler,
	tourHandler *TourHandler,
	bookingHandler *BookingHandler,
	paymentHandler *PaymentHandler,
	reviewHandler *ReviewHandler,
	destinationHandler *DestinationHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // логирование запросов
	r.Use(middleware.Recoverer) // восстановление после panic

	// Users
	r.Post("/auth/register", userHandler.Register)
	r.Get("/users", userHandler.GetAll)
	r.Get("/users/{id}", userHandler.GetByID)

	// Tours
	r.Get("/tours", tourHandler.GetAll)
	r.Get("/tours/{id}", tourHandler.GetByID)
	r.Post("/tours", tourHandler.Create)
	r.Put("/tours/{id}", tourHandler.Update)
	r.Delete("/tours/{id}", tourHandler.Delete)

	// Bookings
	r.Post("/bookings", bookingHandler.Create)
	r.Get("/bookings", bookingHandler.GetAll)
	r.Get("/bookings/user/{id}", bookingHandler.GetByUserID)
	r.Put("/bookings/{id}/status", bookingHandler.UpdateStatus)
	r.Delete("/bookings/{id}", bookingHandler.Delete)

	// Payments
	r.Post("/payments", paymentHandler.Create)
	r.Get("/payments/{id}", paymentHandler.GetByID)
	r.Put("/payments/{id}/status", paymentHandler.UpdateStatus)

	// Reviews
	r.Post("/reviews", reviewHandler.Create)
	r.Get("/reviews/tour/{id}", reviewHandler.GetByTourID)
	r.Delete("/reviews/{id}", reviewHandler.Delete)

	// Destinations
	r.Post("/destinations", destinationHandler.Create)
	r.Get("/destinations", destinationHandler.GetAll)
	r.Get("/destinations/{id}", destinationHandler.GetByID)
	r.Delete("/destinations/{id}", destinationHandler.Delete)

	return r
}
