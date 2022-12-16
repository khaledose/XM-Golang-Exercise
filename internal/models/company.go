package models

import "github.com/google/uuid"

type Company struct {
	Id                uuid.UUID `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name" gorm:"unique;not null"`
	Description       string    `json:"description"`
	NumberOfEmployees uint      `json:"numberOfEmployees" gorm:"not null"`
	IsRegistered      bool      `json:"isRegistered" gorm:"not null"`
	Type              string    `json:"type" gorm:"not null"`
}

type CompanyRequest struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	NumberOfEmployees uint   `json:"numberOfEmployees,omitempty"`
	IsRegistered      bool   `json:"isRegistered,omitempty"`
	Type              string `json:"type,omitempty"`
}

type CompanyStatusRequest struct {
	IsRegistered bool `json:"isRegistered"`
}

const (
	Corporations       = "CORPORATION"
	NonProfit          = "NON_PROFIT"
	Cooperative        = "COOPERATIVE"
	SoleProprietorship = "SOLE_PROPRIETORSHIP"
)

const AppName = "XM_GOLANG_EXCERCISE"
