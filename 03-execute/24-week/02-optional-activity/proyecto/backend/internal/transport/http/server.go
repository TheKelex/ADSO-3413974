package http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"

	"workshop/internal/config"
	"workshop/internal/repository"
	"workshop/internal/usecase"
)

// Dependency is everything the API needs from the outside world.
type Dependency struct {
	Database *sql.DB
	Config   config.Config
	Now      func() time.Time
	NewID    func() string
}

// NewServer builds the HTTP server: it wires the adapters into the use cases,
// registers the routes and applies the middleware. It performs no query of its
// own and holds no business rule.
func NewServer(dependency Dependency) *http.Server {
	now := dependency.Now
	if now == nil {
		now = time.Now
	}
	newID := dependency.NewID
	if newID == nil {
		newID = func() string { return uuid.NewString() }
	}
	settings := dependency.Config
	timeout := settings.DatabaseTimeout

	userRepository := repository.NewUserRepository(dependency.Database, timeout)
	customerRepository := repository.NewCustomerRepository(dependency.Database, timeout)
	vehicleRepository := repository.NewVehicleRepository(dependency.Database, timeout)
	technicianRepository := repository.NewTechnicianRepository(dependency.Database, timeout)
	orderRepository := repository.NewServiceOrderRepository(dependency.Database, timeout)
	assignmentRepository := repository.NewAssignmentRepository(dependency.Database, timeout)
	diagnosticRepository := repository.NewDiagnosticRepository(dependency.Database, timeout)
	interventionRepository := repository.NewInterventionRepository(dependency.Database, timeout)
	warrantyRepository := repository.NewWarrantyRepository(dependency.Database, timeout)

	issuer := NewTokenIssuer(settings.TokenSecret, settings.TokenTTL)
	loginLimiter := NewLoginRateLimiter(settings.LoginMaxAttempt, settings.LoginWindow, now)

	authHandler := NewAuthHandler(usecase.NewAuthenticateUser(userRepository, issuer, now))
	customerHandler := NewCustomerHandler(usecase.NewCustomerUseCase(customerRepository, newID, now))
	vehicleHandler := NewVehicleHandler(usecase.NewVehicleUseCase(vehicleRepository, customerRepository, newID, now))
	technicianHandler := NewTechnicianHandler(usecase.NewTechnicianUseCase(technicianRepository))
	orderUseCase := usecase.NewServiceOrderUseCase(orderRepository, vehicleRepository, assignmentRepository, technicianRepository, newID, now)
	orderHandler := NewServiceOrderHandler(orderUseCase)
	assignmentHandler := NewAssignmentHandler(
		usecase.NewAssignmentUseCase(assignmentRepository, orderRepository, technicianRepository, newID, now),
	)
	diagnosticHandler := NewDiagnosticHandler(usecase.NewDiagnosticUseCase(
		diagnosticRepository, orderRepository, assignmentRepository, technicianRepository, newID, now,
	))
	interventionHandler := NewInterventionHandler(usecase.NewInterventionUseCase(
		interventionRepository, orderRepository, assignmentRepository, technicianRepository, newID, now,
	))
	warrantyHandler := NewWarrantyHandler(usecase.NewWarrantyUseCase(warrantyRepository, interventionRepository, newID, now))
	timelineHandler := NewTimelineHandler(usecase.NewTimelineUseCase(
		vehicleRepository, orderRepository, diagnosticRepository, interventionRepository, warrantyRepository,
	))
	dashboardHandler := NewDashboardHandler(usecase.NewDashboardUseCase(orderRepository, technicianRepository))

	protected := http.NewServeMux()
	protected.HandleFunc("GET /api/customer", customerHandler.List)
	protected.HandleFunc("POST /api/customer", customerHandler.Create)
	protected.HandleFunc("GET /api/vehicle", vehicleHandler.List)
	protected.HandleFunc("POST /api/vehicle", vehicleHandler.Create)
	protected.HandleFunc("GET /api/vehicle/{vehicleId}/timeline", timelineHandler.Build)
	protected.HandleFunc("GET /api/technician", technicianHandler.List)
	protected.HandleFunc("GET /api/service-order", orderHandler.List)
	protected.HandleFunc("POST /api/service-order", orderHandler.Create)
	protected.HandleFunc("GET /api/service-order/{serviceOrderId}", orderHandler.Find)
	protected.HandleFunc("GET /api/service-order/{serviceOrderId}/transition", orderHandler.ListTransition)
	protected.HandleFunc("POST /api/service-order/{serviceOrderId}/status", orderHandler.Advance)
	protected.HandleFunc("GET /api/service-order/{serviceOrderId}/assignment", assignmentHandler.Find)
	protected.HandleFunc("POST /api/service-order/{serviceOrderId}/assignment", assignmentHandler.Assign)
	protected.HandleFunc("GET /api/service-order/{serviceOrderId}/diagnostic", diagnosticHandler.Find)
	protected.HandleFunc("POST /api/service-order/{serviceOrderId}/diagnostic", diagnosticHandler.Record)
	protected.HandleFunc("GET /api/service-order/{serviceOrderId}/intervention", interventionHandler.List)
	protected.HandleFunc("POST /api/service-order/{serviceOrderId}/intervention", interventionHandler.Register)
	protected.HandleFunc("GET /api/warranty", warrantyHandler.List)
	protected.HandleFunc("POST /api/warranty", warrantyHandler.Issue)
	protected.HandleFunc("GET /api/dashboard", dashboardHandler.Build)

	root := http.NewServeMux()
	root.HandleFunc("GET /api/health", health)
	root.Handle("POST /api/session", loginLimiter.Wrap(http.HandlerFunc(authHandler.SignIn)))
	root.Handle("/api/", authMiddleware(issuer, now)(protected))

	return &http.Server{
		Addr:              ":" + settings.HTTPPort,
		Handler:           chain(root, corsMiddleware(settings.AllowedOrigin)),
		ReadTimeout:       settings.RequestTimeout,
		ReadHeaderTimeout: settings.RequestTimeout,
		WriteTimeout:      settings.RequestTimeout,
		IdleTimeout:       2 * settings.RequestTimeout,
	}
}

func health(writer http.ResponseWriter, _ *http.Request) {
	respond(writer, http.StatusOK, map[string]string{"status": "ok"})
}
