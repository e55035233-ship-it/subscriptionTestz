package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"zadaie/internal/model"
	"zadaie/internal/service"
)

type SubscriptionHandler struct {
	Service *service.SubscriptionService
}

// CreateSubscription godoc
// @Summary Create subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body model.Subscription true "Subscription"
// @Success 201 {object} model.SubscriptionResponse
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub model.Subscription

	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Create(&sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.SubscriptionResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	})
}

// GetAllSubscriptions godoc
// @Summary Get all subscriptions
// @Tags subscriptions
// @Produce json
// @Success 200 {array} model.SubscriptionResponse
// @Router /subscriptions [get]
func (h *SubscriptionHandler) GetAll(c *gin.Context) {
	subs, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]model.SubscriptionResponse, 0, len(subs))

	for _, s := range subs {
		resp = append(resp, model.SubscriptionResponse{
			ID:          s.ID,
			ServiceName: s.ServiceName,
			Price:       s.Price,
			UserID:      s.UserID,
			StartDate:   s.StartDate,
			EndDate:     s.EndDate,
		})
	}

	c.JSON(http.StatusOK, resp)
}

// GetSubscriptionByID godoc
// @Summary Get subscription by id
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} model.SubscriptionResponse
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	sub, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, model.SubscriptionResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	})
}

// UpdateSubscription godoc
// @Summary Update subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param subscription body model.Subscription true "Subscription"
// @Success 200 {object} map[string]string
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var sub model.Subscription

	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Update(uint(id), sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteSubscription godoc
// @Summary Delete subscription
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GetTotalCost godoc
// @Summary Get total cost of subscriptions
// @Tags subscriptions
// @Produce json
// @Param user_id query string true "User ID"
// @Param service_name query string false "Service name"
// @Param start query string false "Start date (YYYY-MM)"
// @Param end query string false "End date (YYYY-MM)"
// @Success 200 {object} map[string]int
// @Router /subscriptions/total [get]
func (h *SubscriptionHandler) GetTotalCost(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	start := c.Query("start")
	end := c.Query("end")

	total, err := h.Service.GetTotalCost(userID, serviceName, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}
