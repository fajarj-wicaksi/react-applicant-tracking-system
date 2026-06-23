package route

import (
	"github.com/gin-gonic/gin"
	"ats-backend/internal/interfaces/api/handler"
)

// RegisterTenantRoutes registers the routes for the tenant module
func RegisterTenantRoutes(r *gin.RouterGroup, tenantHandler *handler.TenantHandler) {
	tenants := r.Group("/tenants")
	{
		tenants.POST("", tenantHandler.CreateTenant)
		tenants.GET("", tenantHandler.ListTenants)
		tenants.GET("/:id", tenantHandler.GetTenant)
		tenants.PUT("/:id", tenantHandler.UpdateTenant)
		tenants.DELETE("/:id", tenantHandler.DeleteTenant)
	}
}
