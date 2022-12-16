package company

import (
	"xm/companies/internal/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DefaultRepository implementation
type DefaultRepository struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

// NewCompanyRepository creates a new instance
func NewCompanyRepository(logger *zap.SugaredLogger, conn *gorm.DB) *DefaultRepository {
	return &DefaultRepository{
		db:  conn,
		log: logger,
	}
}

func (r *DefaultRepository) GetCompanyById(id uuid.UUID) (*models.Company, error) {
	r.log.Infof("Retrieving company with id: %s", id.String())
	var company models.Company

	result := r.db.First(&company, id)
	if result.Error != nil {
		r.log.Errorf("Error while trying to get company by id: %s, with error: %s", id.String(), result.Error)
		return nil, result.Error
	}
	r.log.Info("Company found")
	return &company, nil
}

func (r *DefaultRepository) AddCompany(company models.Company) error {
	r.log.Infof("Adding new company: %s", company)

	result := r.db.Create(&company)
	if result.Error != nil {
		r.log.Errorf("Error while trying to add new company, with error: %s", result.Error)
		return result.Error
	}

	r.log.Info("Company added")
	return nil
}

func (r *DefaultRepository) UpdateCompany(id uuid.UUID, company *models.Company) error {
	r.log.Infof("Updating company by id: %s, with details: %s", id.String(), company)

	result := r.db.Model(&models.Company{}).Where("id = ?", id).Updates(company)
	if result.Error != nil {
		r.log.Errorf("Error while trying to update company by id: %s, with error: %s", id.String(), result.Error)
		return result.Error
	}

	r.log.Info("Company updated")
	return nil
}

func (r *DefaultRepository) UpdateCompanyStatus(id uuid.UUID, status models.CompanyStatusRequest) error {
	r.log.Infof("Updating company by id: %s, with details: %s", id.String(), status)

	result := r.db.Model(&models.Company{}).Where("id = ?", id).Update("is_registered", status.IsRegistered)
	if result.Error != nil {
		r.log.Errorf("Error while trying to update company by id: %s, with error: %s", id.String(), result.Error)
		return result.Error
	}

	r.log.Info("Company updated")
	return nil
}

func (r *DefaultRepository) DeleteCompany(id uuid.UUID) error {
	r.log.Infof("Deleting company by id: %s", id.String())

	result := r.db.Delete(&models.Company{}, id)
	if result.Error != nil {
		r.log.Errorf("Error while trying to delete company by id: %s, with error: %s", id.String(), result.Error)
		return result.Error
	}

	r.log.Info("Company deleted")
	return nil
}
