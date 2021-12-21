package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	countryInfra "grpc-demo/location-service/internal/infrastucture/postgresql/country"
	countryInterface "grpc-demo/location-service/internal/interfaces"
	config "grpc-demo/location-service/internal/shared/config"
	countryUseCase "grpc-demo/location-service/internal/usecase/country"
	"os"
	"os/signal"
	"time"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetConfigFile("location-service/resources/config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := &config.ApplicationConfig{}
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	spew.Dump(conf)

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		conf.Database.Host,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	countryRepo := countryInfra.NewCountryRepository(db)
	countryService := countryUseCase.NewCountryService(countryRepo)
	countryHandler := countryInterface.NewCountryHandler(countryService)

	server := echo.New()

	group := server.Group("/country")
	{
		group.GET("", countryHandler.GetAll)
		group.GET("/:id", countryHandler.GetById)
	}

	go func() {
		if err := server.Start(":9090"); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}

}
