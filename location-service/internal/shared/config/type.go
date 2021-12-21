package config

import "grpc-demo/location-service/internal/shared/database"

type ApplicationConfig struct {
	Database database.DatabaseConfig `mapstructure:"database"`
}
