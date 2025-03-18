package IntermediateServices

import (
	"github.com/gofiber/fiber/v2"
)


type Response struct {
	Status            string `json:"status"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
}


func SendResponse(c *fiber.Ctx, success bool, description string) error {
	response := Response{
		Status:            "success",
		StatusCode:        0,
		StatusDescription: description,
	}

	if !success {
		response.Status = "failed"
		response.StatusCode = -1
		c.SendStatus(fiber.StatusBadRequest)
	} else {
		c.SendStatus(fiber.StatusOK)
	}

	return c.JSON(response)
}
