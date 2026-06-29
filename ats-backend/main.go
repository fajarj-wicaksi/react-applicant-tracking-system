package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ats-backend/internal/application/service"
	"ats-backend/internal/domain"
	infraPostgres "ats-backend/internal/infrastructure/postgres"
	"ats-backend/internal/interfaces/api/handler"
	"ats-backend/internal/interfaces/api/middleware"
)

func main() {
	// 1. Initialize Database (Hardcoded for now, should use env vars)
	dsn := "host=localhost user=postgres password=1234 dbname=talentflow port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Auto-Migrate Schemas
	log.Println("Migrating database schemas...")
	if err := db.AutoMigrate(
		&domain.Tenant{},
		&domain.Role{},
		&domain.User{},
		&domain.Position{},
		&domain.Candidate{},
		&domain.Application{},
		&domain.BillingPlan{},
		&domain.Subscription{},
		&domain.Interview{},
		&domain.Task{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 3. Initialize Repositories
	tenantRepo := infraPostgres.NewTenantRepository(db)
	userRepo := infraPostgres.NewUserRepository(db)
	candidateRepo := infraPostgres.NewCandidateRepository(db)
	applicationRepo := infraPostgres.NewApplicationRepository(db)
	positionRepo := infraPostgres.NewPositionRepository(db)
	interviewRepo := infraPostgres.NewInterviewRepository(db)
	taskRepo := infraPostgres.NewTaskRepository(db)

	// 4. Initialize Services
	tenantSvc := service.NewTenantService(tenantRepo)
	authSvc := service.NewAuthService(userRepo)
	adminSvc := service.NewAdminService(db)

	// 5. Initialize Handlers
	tenantHandler := handler.NewTenantHandler(tenantSvc)
	authHandler := handler.NewAuthHandler(authSvc)
	candidateHandler := handler.NewCandidateHandler(candidateRepo, applicationRepo, positionRepo)
	adminHandler := handler.NewAdminHandler(adminSvc)
	userHandler := handler.NewUserHandler(userRepo)
	interviewHandler := handler.NewInterviewHandler(interviewRepo)
	taskHandler := handler.NewTaskHandler(taskRepo)

	// 6. Set up Gin Router
	router := gin.Default()

	// 7. Register Routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes (require JWT)
		protected := v1.Group("")
		protected.Use(middleware.RequireAuth())
		{
			// Tenants
			tenants := protected.Group("/tenants")
			{
				tenants.POST("", tenantHandler.CreateTenant)
				tenants.GET("", tenantHandler.ListTenants)
				tenants.GET("/:id", tenantHandler.GetTenant)
				tenants.PUT("/:id", tenantHandler.UpdateTenant)
				tenants.DELETE("/:id", tenantHandler.DeleteTenant)
			}

			// Candidates
			candidates := protected.Group("/candidates")
			{
				candidates.POST("", candidateHandler.CreateCandidate)
				candidates.GET("", candidateHandler.ListCandidates)
				candidates.GET("/:id", candidateHandler.GetCandidate)
				candidates.PUT("/:id", candidateHandler.UpdateCandidate)
				candidates.DELETE("/:id", candidateHandler.DeleteCandidate)
			}

			// Positions
			positions := protected.Group("/positions")
			{
				positions.POST("", candidateHandler.CreatePosition)
				positions.GET("", candidateHandler.ListPositions)
				positions.GET("/:id", candidateHandler.GetPosition)
				positions.PUT("/:id", candidateHandler.UpdatePosition)
				positions.DELETE("/:id", candidateHandler.DeletePosition)
			}

			// Applications & Pipeline
			applications := protected.Group("/applications")
			{
				applications.POST("", candidateHandler.CreateApplication)
				applications.PATCH("/:id/stage", candidateHandler.UpdateApplicationStage)
			}
			protected.GET("/pipeline", candidateHandler.GetPipeline)

			// Admin Panel
			admin := protected.Group("/admin")
			{
				admin.GET("/tenants", adminHandler.ListTenants)
				admin.GET("/tenants/:id/stats", adminHandler.GetTenantStats)
				admin.GET("/monitoring", adminHandler.GetSystemMonitoring)
			}

			// Users
			users := protected.Group("/users")
			{
				users.POST("", userHandler.CreateUser)
				users.GET("", userHandler.ListUsers)
				users.GET("/:id", userHandler.GetUser)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}

			// Interviews
			interviews := protected.Group("/interviews")
			{
				interviews.POST("", interviewHandler.CreateInterview)
				interviews.GET("", interviewHandler.ListInterviews)
				interviews.GET("/:id", interviewHandler.GetInterview)
				interviews.PUT("/:id", interviewHandler.UpdateInterview)
				interviews.DELETE("/:id", interviewHandler.DeleteInterview)
			}

			// Tasks
			tasks := protected.Group("/tasks")
			{
				tasks.POST("", taskHandler.CreateTask)
				tasks.GET("", taskHandler.ListTasks)
				tasks.GET("/:id", taskHandler.GetTask)
				tasks.PUT("/:id", taskHandler.UpdateTask)
				tasks.DELETE("/:id", taskHandler.DeleteTask)
			}
		}
	}

	// 8. Start Server
	log.Println("Server starting on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
