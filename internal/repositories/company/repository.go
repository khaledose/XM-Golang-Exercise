package company

import (
	"xm/companies/internal/models"

	"github.com/google/uuid"
)

type IRepository interface {
	GetCompanyById(id uuid.UUID) (*models.Company, error)
	AddCompany(company models.Company) error
	UpdateCompany(id uuid.UUID, company *models.Company) error
	UpdateCompanyStatus(id uuid.UUID, status models.CompanyStatusRequest) error
	DeleteCompany(id uuid.UUID) error
}
