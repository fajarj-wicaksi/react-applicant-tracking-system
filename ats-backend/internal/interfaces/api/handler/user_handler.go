package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"ats-backend/internal/application/dto"
	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type UserHandler struct {
	userRepo port.UserRepository
}

func NewUserHandler(userRepo port.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// POST /api/v1/users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(c.GetString("tenantID"))
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	// Check if email already exists
	existingUser, _ := h.userRepo.GetByEmail(c.Request.Context(), req.Email)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := &domain.User{
		TenantID:     tenantID,
		RoleID:       roleID,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Password:     string(hashedPassword),
	}

	if err := h.userRepo.Create(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Re-fetch to get role name populated or just return basic info
	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
}

// GET /api/v1/users
func (h *UserHandler) ListUsers(c *gin.Context) {
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	users, err := h.userRepo.ListByTenant(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	var responses []dto.UserDetailsResponse
	for _, u := range users {
		roleName := ""
		if u.Role != nil {
			roleName = u.Role.Name
		}
		responses = append(responses, dto.UserDetailsResponse{
			ID:        u.ID.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			RoleName:  roleName,
			CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	c.JSON(http.StatusOK, responses)
}

// GET /api/v1/users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	user, err := h.userRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if user.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	roleName := ""
	if user.Role != nil {
		roleName = user.Role.Name
	}

	c.JSON(http.StatusOK, dto.UserDetailsResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RoleName:  roleName,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

// PUT /api/v1/users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if user.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Email != "" && req.Email != user.Email {
		// Check uniqueness
		existingUser, _ := h.userRepo.GetByEmail(c.Request.Context(), req.Email)
		if existingUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
			return
		}
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}
	if req.RoleID != "" {
		roleID, err := uuid.Parse(req.RoleID)
		if err == nil {
			user.RoleID = roleID
		}
	}

	if err := h.userRepo.Update(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DELETE /api/v1/users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	user, err := h.userRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if user.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := h.userRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}
