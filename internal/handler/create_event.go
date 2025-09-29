package handler

import (
	"L2.18-calendar/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (h *Handler) Create(c *gin.Context) {
	var dto Response
	err := validator.BindJSON(&dto, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("[Validator] Invalid JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные входные данные"})
		return
	}
	date, err := time.Parse("2006-01-02", dto.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, expected YYYY-MM-DD"})
		return
	}
	e := Event{
		UserID: dto.UserID,
		Date:   date,
		Text:   dto.Text,
	}
	if e.Date.IsZero() {
		logrus.WithError(err).Warn("[Validator] Invalid date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные входные данные"})
		return
	}
	srvReq := e.ToSrv()
	err = h.srv.Create(&srvReq)
	if err != nil {
		logrus.WithError(err).Error("[Create] something went wrong")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, "create success")
}
