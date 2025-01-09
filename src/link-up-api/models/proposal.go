package models

type Proposal struct {
    CustomerName string  `json:"customer_name" bson:"customer_name"`
    Address      Address `json:"address" bson:"address"`
    Plan         string  `json:"plan" bson:"plan"`
    Email        string  `json:"email" bson:"email"`
    Status       string  `json:"status" bson:"status"`
}
