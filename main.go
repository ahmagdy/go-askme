package main

import (
	"fmt"
	"github.com/bashmohandes/go-askme/answer"
	"github.com/bashmohandes/go-askme/question"
	"github.com/bashmohandes/go-askme/user"
	"go.uber.org/fx"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bashmohandes/go-askme/models"
	"github.com/bashmohandes/go-askme/web/askme"
	"github.com/bashmohandes/go-askme/web/framework"
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fx.New(
		framework.Module,
		user.Module,
		question.Module,
		answer.Module,
		askme.Module,
		fx.Provide(newConfig),
		fx.Provide(newFileProvider),
		fx.Invoke(invoke),
		)

	if err := app.Err(); err != nil{
		log.Fatalln(err)
	}

}

func invoke(config *framework.Config, app *askme.App)error{
	if err:= migrateDB(config); err != nil{
		return err
	}

	return app.Start()
}

func newFileProvider(config *framework.Config) framework.FileProvider {
	return packr.NewBox(config.PublicFolder)
}

func newConfig() *framework.Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Incorrect 'PORT' format: %v\n", err)
	}
	debug, err := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	if err != nil {
		log.Fatalf("Incorrect 'DEBUG_MODE' format: %v\n", err)
	}
	sessionMaxLife, err := time.ParseDuration(os.Getenv("SESSION_MAX_LIFE_TIME"))
	if err != nil {
		log.Fatalf("Incorrect 'SESSION_MAX_LIFE_TIME' format: %v\n", err)
	}

	config := &framework.Config{
		Debug:              debug,
		PublicFolder:       os.Getenv("PUBLIC_FOLDER"),
		Port:               port,
		SessionMaxLifeTime: sessionMaxLife,
		SessionCookie:      os.Getenv("SESSION_COOKIE"),
		PostgresUser:       os.Getenv("POSTGRES_USER"),
		PostgresPassword:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:         os.Getenv("POSTGRES_DB"),
		PostgresHost:       os.Getenv("POSTGRES_HOST"),
		OktaClient:         os.Getenv("OKTA_CLIENT_ID"),
		OktaSecret:         os.Getenv("OKTA_CLIENT_SECRET"),
		OktaIssuer:         os.Getenv("OKTA_ISSUER"),				
	}	

	linkedInIdp := os.Getenv("OKTA_SOCIAL_LINKEDIN_IDP")
	facebookIdp := os.Getenv("OKTA_SOCIAL_FACEBOOK_IDP")

	if linkedInIdp != "" {
		config.OktaSocialIdps = append(config.OktaSocialIdps, framework.OktaSocialIdp{
			ID: linkedInIdp,
			Name: "LINKEDIN",
		})
	}

	if facebookIdp != "" {
		config.OktaSocialIdps = append(config.OktaSocialIdps, framework.OktaSocialIdp{
			ID: facebookIdp,
			Name: "FACEBOOK",
		})
	}

	return config
}

func migrateDB(config *framework.Config) error {
	log.Print("Auto Migration Starting")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresHost, 5432, config.PostgresDB)
	db, err := gorm.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatalf("%v", err)
	}
	db.LogMode(config.Debug)
	db.AutoMigrate(&models.User{}, &models.Question{}, &models.Answer{})
	log.Print("Auto Migration Ended")
	return nil
}
