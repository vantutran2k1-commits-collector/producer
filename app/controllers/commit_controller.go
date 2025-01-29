package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vantutran2k1-commits-collector/producer/app/constants"
	"github.com/vantutran2k1-commits-collector/producer/app/services"
	"net/http"
	"time"
)

type CommitController struct {
	CommitService services.CommitService
}

func NewCommitController(commitService services.CommitService) CommitController {
	return CommitController{CommitService: commitService}
}

func (c *CommitController) Collect(ctx *gin.Context) {
	var fromTime *time.Time

	fromTimeStr := ctx.Query("from_time")
	if fromTimeStr != "" {
		parsedTime, err := time.Parse(constants.DateTimeFormat, fromTimeStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fromTime = &parsedTime
	}
	commits, err := c.CommitService.Collect(fromTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": commits})
}
