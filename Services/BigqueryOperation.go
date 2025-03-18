package Services

import (
	"CMS_PUBSUB_INTEGRATION/IntermediateServices"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func BigquerOperation(c *fiber.Ctx) error {

	fmt.Println("BigqueryOperation started in :", time.Now())

	decodedData, messageId, err := IntermediateServices.PubsubReqValidation(c)
	if err != nil {
		fmt.Println("ERROR:", messageId, "Pub/Sub validation failed:", err)
		return IntermediateServices.SendResponse(c, false, err.Error())
	}

	bookMark, err := IntermediateServices.InsertIntoBigquery(c, decodedData, messageId)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "Column id does not exist", err)
		return IntermediateServices.SendResponse(c, false, err.Error())
	}

	err = IntermediateServices.InsertServiceBigquery(c, decodedData, bookMark)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "InsertService failed:", err)
		return IntermediateServices.SendResponse(c, false, err.Error())
	}

	return IntermediateServices.SendResponse(c, true, "Data successfully processed.")
}
