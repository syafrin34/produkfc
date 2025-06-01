package consumer

import (
	"context"
	"encoding/json"
	"log"
	"produkfc/cmd/product/service"
	"produkfc/models"

	"github.com/segmentio/kafka-go"
)

type ProductRollbackStockConsumer struct {
	Reader         *kafka.Reader
	ProductService service.ProductService
}

func NewProductRollbackStockConsumer(brokers []string, topic string, productService service.ProductService) *ProductRollbackStockConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "productfc",
	})

	return &ProductRollbackStockConsumer{
		Reader:         reader,
		ProductService: productService,
	}
}

func (c *ProductRollbackStockConsumer) Start(ctx context.Context) {
	log.Println("[KAFKA] Listening to topic 'stock.rollback' ")
	for {
		message, err := c.Reader.ReadMessage(ctx)
		if err != nil {
			continue
		}

		var event models.ProductStockUpdateEvent
		err = json.Unmarshal(message.Value, &event)
		if err != nil {
			continue
		}
		// looping based on product stock update event
		for _, product := range event.Products {
			//rollback stock
			err = c.ProductService.AddProductStockByProductID(ctx, product.ProductID, product.Qty)
			if err != nil {
				continue
			}
		}

	}
}
