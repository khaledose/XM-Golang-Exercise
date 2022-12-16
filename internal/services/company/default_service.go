package company

import (
	"xm/companies/internal/adapters/kafka"
	"xm/companies/internal/models"
	"xm/companies/internal/repositories/company"

	"github.com/google/uuid"
)

type DefaultSerice struct {
	CompanyRepository company.IRepository
	KafkaBroker       *kafka.KafkaBroker
}

func NewCompanyService(repo company.IRepository, broker *kafka.KafkaBroker) *DefaultSerice {
	return &DefaultSerice{
		CompanyRepository: repo,
		KafkaBroker:       broker,
	}
}

func (s *DefaultSerice) AddCompany(request models.CompanyRequest) error {
	company := models.Company{
		Id:                uuid.New(),
		Name:              request.Name,
		Description:       request.Description,
		NumberOfEmployees: request.NumberOfEmployees,
		IsRegistered:      request.IsRegistered,
		Type:              request.Type,
	}

	err := s.CompanyRepository.AddCompany(company)
	if err != nil {
		return err
	}

	err = s.KafkaBroker.Produce(models.CompanyMutationTopic, company, models.CompanyNewStatus)
	if err != nil {

		return err
	}

	return nil
}

func (s *DefaultSerice) GetCompany(id uuid.UUID) (*models.Company, error) {
	company, err := s.CompanyRepository.GetCompanyById(id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *DefaultSerice) UpdateCompany(id uuid.UUID, request models.CompanyRequest) error {
	company := &models.Company{
		Name:              request.Name,
		Description:       request.Description,
		Type:              request.Type,
		NumberOfEmployees: request.NumberOfEmployees,
	}

	err := s.CompanyRepository.UpdateCompany(id, company)
	if err != nil {
		return err
	}

	company, err = s.GetCompany(id)
	if err != nil {
		return err
	}

	err = s.KafkaBroker.Produce(models.CompanyMutationTopic, company, models.CompanyUpdateStatus)
	if err != nil {
		return err
	}

	return nil
}

func (s *DefaultSerice) UpdateCompanyStatus(id uuid.UUID, request models.CompanyStatusRequest) error {
	err := s.CompanyRepository.UpdateCompanyStatus(id, request)
	if err != nil {
		return err
	}

	company, err := s.GetCompany(id)
	if err != nil {
		return err
	}

	err = s.KafkaBroker.Produce(models.CompanyMutationTopic, company, models.CompanyUpdateStatus)
	if err != nil {

		return err
	}

	return nil
}

func (s *DefaultSerice) DeleteCompany(id uuid.UUID) error {
	company, err := s.GetCompany(id)
	if err != nil {
		return err
	}

	err = s.CompanyRepository.DeleteCompany(id)
	if err != nil {
		return err
	}

	err = s.KafkaBroker.Produce(models.CompanyMutationTopic, company, models.CompanyDeleteStatus)
	if err != nil {

		return err
	}

	return nil
}
