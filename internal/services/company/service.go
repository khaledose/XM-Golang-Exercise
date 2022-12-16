package company

import (
	"xm/companies/internal/models"

	"github.com/google/uuid"
)

type IService interface {
	AddCompany(request models.CompanyRequest) error
	GetCompany(id uuid.UUID) (*models.Company, error)
	UpdateCompany(id uuid.UUID, request models.CompanyRequest) error
	UpdateCompanyStatus(id uuid.UUID, request models.CompanyStatusRequest) error
	DeleteCompany(id uuid.UUID) error
}
