package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ats-backend/internal/application/port"
)

type AdminHandler struct {
	adminService port.AdminService
}

func NewAdminHandler(adminService port.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// GET /api/v1/admin/tenants
func (h *AdminHandler) ListTenants(c *gin.Context) {
	tenants, err := h.adminService.ListTenants(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list tenants"})
		return
	}
	c.JSON(http.StatusOK, tenants)
}

// GET /api/v1/admin/tenants/:id/stats
func (h *AdminHandler) GetTenantStats(c *gin.Context) {
	tenantID := c.Param("id")
	stats, err := h.adminService.GetTenantStats(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// GET /api/v1/admin/monitoring
func (h *AdminHandler) GetSystemMonitoring(c *gin.Context) {
	monitoring, err := h.adminService.GetSystemMonitoring(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get monitoring data"})
		return
	}
	c.JSON(http.StatusOK, monitoring)
}
