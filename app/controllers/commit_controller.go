package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vantutran2k1-commits-collector/producer/app/payloads"
	"github.com/vantutran2k1-commits-collector/producer/app/services"
	"net/http"
)

type CommitController struct {
	CommitService services.CommitService
}

func NewCommitController(commitService services.CommitService) CommitController {
	return CommitController{CommitService: commitService}
}

func (c *CommitController) Collect(ctx *gin.Context) {
	payload := payloads.CommitPayload{
		Sha:    "sha 1",
		NodeId: "node 1",
		Commit: payloads.Commit{
			Author: payloads.User{
				Name:  "name 1",
				Email: "email 1",
				Date:  "2025-01-01T00:00:00",
			},
			Committer: payloads.User{
				Name:  "name 2",
				Email: "email 2",
				Date:  "2025-01-01T00:00:00",
			},
			Message: "message 1",
		},
	}

	err := c.CommitService.Send(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": payload})
}
