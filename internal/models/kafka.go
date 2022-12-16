package models

import "time"

type KafkaEvent struct {
	Id        string      `json:"id"`
	Status    string      `json:"type"`
	Source    string      `json:"source"`
	TimeStamp time.Time   `json:"timeStamp"`
	Data      interface{} `json:"data"`
}

const (
	CompanyMutationTopic = "company.mutated"
)

const (
	CompanyNewStatus    = "NEW"
	CompanyUpdateStatus = "UPDATE"
	CompanyDeleteStatus = "DELETE"
)
