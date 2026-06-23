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
	if err := db.AutoMigrate(&domain.Tenant{}, &domain.Role{}, &domain.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 3. Initialize Repositories
	tenantRepo := infraPostgres.NewTenantRepository(db)
	userRepo := infraPostgres.NewUserRepository(db)
	// roleRepo := infraPostgres.NewRoleRepository(db) // Will be used later for seeding/admin

	// 4. Initialize Services
	tenantSvc := service.NewTenantService(tenantRepo)
	authSvc := service.NewAuthService(userRepo)

	// 5. Initialize Handlers
	tenantHandler := handler.NewTenantHandler(tenantSvc)
	authHandler := handler.NewAuthHandler(authSvc)

	// 6. Set up Gin Router
	router := gin.Default()

	// 7. Register Routes
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		tenants := v1.Group("/tenants")
		{
			tenants.POST("", tenantHandler.CreateTenant)
			tenants.GET("", tenantHandler.ListTenants)
			tenants.GET("/:id", tenantHandler.GetTenant)
			tenants.PUT("/:id", tenantHandler.UpdateTenant)
			tenants.DELETE("/:id", tenantHandler.DeleteTenant)
		}
	}

	// 8. Start Server
	log.Println("Server starting on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
