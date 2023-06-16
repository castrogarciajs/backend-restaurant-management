package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       *string            `json:"invoive_id"`
	Order_id         *float64           `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"payment_status" validate:"required, eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"payment_due_data"`
	Created_at       time.Time          `json:"create_at"`
	Update_at        time.Time          `json:"updated_at"`
}
