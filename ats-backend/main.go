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
	dsn := "host=localhost user=postgres password=postgres dbname=talentflow port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Auto-Migrate Schemas
	log.Println("Migrating database schemas...")
	if err := db.AutoMigrate(&domain.Tenant{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 3. Initialize Repositories
	tenantRepo := infraPostgres.NewTenantRepository(db)

	// 4. Initialize Services
	tenantSvc := service.NewTenantService(tenantRepo)

	// 5. Initialize Handlers
	tenantHandler := handler.NewTenantHandler(tenantSvc)

	// 6. Set up Gin Router
	router := gin.Default()

	// 7. Register Routes
	v1 := router.Group("/api/v1")
	{
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
