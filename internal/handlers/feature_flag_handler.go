package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kashyapkrlucky/ff-go-src/internal/messaging"
	"github.com/kashyapkrlucky/ff-go-src/internal/models"
	"github.com/kashyapkrlucky/ff-go-src/internal/repositories"
)

// FeatureFlagHandler handles API requests
type FeatureFlagHandler struct {
	Repo      *repositories.FeatureFlagRepo
	Publisher *messaging.Publisher
}

// NewFeatureFlagHandler creates a new handler
func NewFeatureFlagHandler(repo *repositories.FeatureFlagRepo) *FeatureFlagHandler {
	return &FeatureFlagHandler{Repo: repo}
}

// GetAllFlags handles GET /api/flags
func (h *FeatureFlagHandler) GetAllFlags(c *gin.Context) {
	flags, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flags"})
		return
	}
	c.JSON(http.StatusOK, flags)
}

// CreateFlag handles POST /api/flags
func (h *FeatureFlagHandler) CreateFlag(c *gin.Context) {
	var flag models.FeatureFlag
	if err := c.ShouldBindJSON(&flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.Repo.Create(flag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flag"})
		return
	}
	// Publish the feature flag creation message to RabbitMQ
	if err := h.Publisher.PublishFlagChange(flag.ID, "created"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to propagate flag change"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Feature flag created"})
}

// UpdateFlag updates a feature flag
func (h *FeatureFlagHandler) UpdateFlag(c *gin.Context) {
	var flag models.FeatureFlag
	if err := c.ShouldBindJSON(&flag); err != nil {
		log.Println("Invalid request payload:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.Repo.Update(flag); err != nil {
		log.Println("Failed to update feature flag:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feature flag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feature flag updated successfully"})
}

// DeleteFlag removes a feature flag
func (h *FeatureFlagHandler) DeleteFlag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feature flag ID"})
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		log.Println("Failed to delete feature flag:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feature flag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feature flag deleted successfully"})
}
