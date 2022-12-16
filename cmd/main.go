package main

import (
	"xm/companies/config"
	"xm/companies/internal/adapters/api"
	"xm/companies/internal/adapters/db"
	"xm/companies/internal/adapters/kafka"
	repo "xm/companies/internal/repositories/company"
	service "xm/companies/internal/services/company"

	"github.com/go-chi/jwtauth"
)

func main() {
	logger := config.NewLogger()
	defer config.CloseLogger(logger)

	conf := config.GetConfig("dev", logger)

	httpServer := api.NewHTTPServer(logger, conf.Server)

	tokenAuth := jwtauth.New("HS256", []byte(conf.Server.Secret), nil)

	kafkaBroker := kafka.NewKafkaBroker(logger, conf.Kafka)

	db := db.NewDatabaseConnection(logger, conf.Database)

	companyRepo := repo.NewCompanyRepository(logger, db)

	companyService := service.NewCompanyService(companyRepo, kafkaBroker)

	api.NewCompanyController(logger, httpServer, companyService, tokenAuth)

	httpServer.Start()
}
