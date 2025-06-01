package consumer

import (
	"context"
	"encoding/json"
	"produkfc/cmd/product/service"
	"produkfc/infrastructure/logger"
	"produkfc/models"

	"github.com/segmentio/kafka-go"
)

type ProductUpdateStockConsumer struct {
	Reader         *kafka.Reader
	productService service.ProductService
}

func NewProductUpdateStockConsumer(brokers []string, topic string, productService service.ProductService) *ProductUpdateStockConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "productfc",
	})

	return &ProductUpdateStockConsumer{
		productService: productService,
		Reader:         reader,
	}

}

func (c *ProductUpdateStockConsumer) Start(ctx context.Context) {
	logger.Logger.Println("[KAFKA] Listening to topic stock.update")

	for {
		message, err := c.Reader.ReadMessage(ctx)
		if err != nil {
			logger.Logger.Println("[KAFKA] Error message", err)
			continue
		}

		// unmarsahl event tp product update stock struct
		var event models.ProductStockUpdateEvent
		err = json.Unmarshal(message.Value, &event)
		if err != nil {
			logger.Logger.Println("[KAFKA] error unmarshal kafka event message", err)
			continue
		}

		// update stock product
		for _, product := range event.Products {
			err = c.productService.DeductProductStockByProductID(ctx, product.ProductID, product.Qty)
			if err != nil {
				logger.Logger.Printf("[KAFKA] error update stock product id #%d", product.ProductID)
				continue
			}

		}
	}
}
