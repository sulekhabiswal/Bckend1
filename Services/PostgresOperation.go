package Services

import (
	"CMS_PUBSUB_INTEGRATION/IntermediateServices"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostgresOperation(c *fiber.Ctx) error {

	fmt.Println("PostgresOperation started in :", time.Now())

	decodedData, messageId, err := IntermediateServices.PubsubReqValidation(c)
	if err != nil {
		fmt.Println("ERROR:", messageId, "Pub/Sub validation failed:", err)
		return IntermediateServices.SendResponse(c, false, "Pub/Sub validation failed.")
	}

	bookMark, err := IntermediateServices.InsertIntoPostgre(c, decodedData, messageId)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "PostgreSQL insertion failed:", err)
		return IntermediateServices.SendResponse(c, false, "Conditional Validation failed.")
	}

	err = IntermediateServices.InsertService(c, decodedData, bookMark)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "InsertService failed:", err)
		return IntermediateServices.SendResponse(c, false, err.Error())
	}

	return IntermediateServices.SendResponse(c, true, "Data successfully processed.")
}
