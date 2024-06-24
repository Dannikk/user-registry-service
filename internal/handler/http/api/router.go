package api

import (
	"net/http"
	"user_registry/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddRoutes(eng *gin.Engine) {
	eng.GET("/ping", h.Ping)
	eng.POST("/sign/hmacsha512", h.SignHMAC)
	eng.POST("/postgres/users", h.CreateUser)
	eng.POST("/redis/incr", h.IncrementRedis)
}

func (h Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h Handler) SignHMAC(c *gin.Context) {
	var ctx = c.Request.Context()
	textKey := &entity.TextKey{}

	if err := c.BindJSON(textKey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.uc.Sign(ctx, textKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"hex_code": res})
	}
}

func (h Handler) CreateUser(c *gin.Context) {
	var ctx = c.Request.Context()
	user := &entity.User{}

	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.uc.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": res})
	}
}

func (h Handler) IncrementRedis(c *gin.Context) {
	var ctx = c.Request.Context()
	keyValue := &entity.KeyValue{}

	if err := c.BindJSON(keyValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.uc.Increment(ctx, keyValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"value": res})
	}
}
