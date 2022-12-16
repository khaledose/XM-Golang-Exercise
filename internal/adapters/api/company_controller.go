package api

import (
	"encoding/json"
	"net/http"
	"xm/companies/internal/models"
	"xm/companies/internal/services/company"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CompanyController struct {
	CompanyService company.IService
	TokenAuth      *jwtauth.JWTAuth
	log            *zap.SugaredLogger
}

func NewCompanyController(logger *zap.SugaredLogger, server *HTTPServer, cs company.IService, tokenAuth *jwtauth.JWTAuth) *CompanyController {
	cc := &CompanyController{
		CompanyService: cs,
		TokenAuth:      tokenAuth,
		log:            logger,
	}

	server.Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/company/{id}", cc.handleGetCompany)
		r.Post("/company", cc.handleAddCompany)
		r.Patch("/company/{id}/status", cc.handleUpdateCompanyStatus)
		r.Patch("/company/{id}", cc.handleUpdateCompany)
		r.Delete("/company/{id}", cc.handleDeleteCompany)
	})

	server.Router.Group(func(r chi.Router) {
		r.Post("/token", cc.handleGetToken)
	})

	return cc
}

func (cc *CompanyController) handleGetCompany(w http.ResponseWriter, r *http.Request) {
	companyId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		cc.log.Errorf("Error while trying to parse id from url, with error: %s", err)
		panic(err)
	}

	company, err := cc.CompanyService.GetCompany(companyId)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(company)
	if err != nil {
		cc.log.Errorf("Error while trying to encode: %s, with error: %s", company, err)
		panic(err)
	}
}

func (cc *CompanyController) handleAddCompany(w http.ResponseWriter, r *http.Request) {
	var companyRequest models.CompanyRequest

	err := json.NewDecoder(r.Body).Decode(&companyRequest)
	if err != nil {
		cc.log.Errorf("Error while trying to decode: %s, with error: %s", companyRequest, err)
		panic(err)
	}

	err = cc.CompanyService.AddCompany(companyRequest)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (cc *CompanyController) handleUpdateCompanyStatus(w http.ResponseWriter, r *http.Request) {
	companyId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		cc.log.Errorf("Error while trying to parse id from url, with error: %s", err)
		panic(err)
	}

	var companyStatus models.CompanyStatusRequest

	err = json.NewDecoder(r.Body).Decode(&companyStatus)
	if err != nil {
		cc.log.Errorf("Error while trying to decode: %s, with error: %s", companyStatus, err)
		panic(err)
	}

	cc.log.Infof("Parsed incoming request to: %s", companyStatus)

	err = cc.CompanyService.UpdateCompanyStatus(companyId, companyStatus)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (cc *CompanyController) handleUpdateCompany(w http.ResponseWriter, r *http.Request) {
	companyId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		cc.log.Errorf("Error while trying to parse id from url, with error: %s", err)
		panic(err)
	}

	var companyRequest models.CompanyRequest

	err = json.NewDecoder(r.Body).Decode(&companyRequest)
	if err != nil {
		cc.log.Errorf("Error while trying to decode: %s, with error: %s", companyRequest, err)
		panic(err)
	}

	cc.log.Infof("Parsed incoming request to: %s", companyRequest)

	err = cc.CompanyService.UpdateCompany(companyId, companyRequest)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (cc *CompanyController) handleDeleteCompany(w http.ResponseWriter, r *http.Request) {
	companyId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		cc.log.Errorf("Error while trying to parse id from url, with error: %s", err)
		panic(err)
	}

	err = cc.CompanyService.DeleteCompany(companyId)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (cc *CompanyController) handleGetToken(w http.ResponseWriter, r *http.Request) {
	var authRequest models.AuthRequest

	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		panic(err)
	}
	_, tokenString, _ := cc.TokenAuth.Encode(map[string]interface{}{"email": authRequest.Email, "passowrd": authRequest.Password})

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tokenString)
	if err != nil {
		panic(err)
	}
}
