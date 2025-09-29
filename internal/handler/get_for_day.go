package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetForDay(c *gin.Context) {
	userIDStr := c.Query("user_id")
	dateStr := strings.TrimSpace(c.Query("date"))

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events, err := h.srv.GetForDay(userID, dateStr)
	if err != nil {
		logrus.WithError(err).Error("[GetForDay] cannot get events")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)

}
