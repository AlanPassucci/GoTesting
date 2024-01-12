package server

import (
	"gotesting/04-TestsFuncionales/prey"
	"gotesting/04-TestsFuncionales/shark"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		// Request
		// - Get request body
		var reqBody request
		if err := context.ShouldBindJSON(&reqBody); err != nil {
			context.JSON(http.StatusBadRequest, response{Success: false})
			return
		}

		// Process
		// - Configure position and speed of the shark
		sharkPosition := []float64{reqBody.XPosition, reqBody.YPosition}
		h.shark.Configure([2]float64(sharkPosition), reqBody.Speed)

		// Response
		context.JSON(http.StatusOK, response{Success: true})
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		// Request
		// - Get request body
		var reqBody request
		if err := context.ShouldBindJSON(&reqBody); err != nil {
			context.JSON(http.StatusBadRequest, response{Success: false})
			return
		}

		// Process
		// - Set speed of the prey
		h.prey.SetSpeed(reqBody.Speed)

		// Response
		context.JSON(http.StatusOK, response{Success: true})
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		// Process
		// - Simulate a hunt
		err, timeToCatch := h.shark.Hunt(h.prey)
		if err != nil {
			context.JSON(http.StatusBadRequest, response{Success: false, Message: err.Error()})
			return
		}

		// Response
		context.JSON(http.StatusOK, response{Success: true, Time: timeToCatch})
	}
}
