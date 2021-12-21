package country

import "grpc-demo/location-service/internal/domain/country"

type CountryServiceIFace interface {
	GetAll() (entiity []country.Country, err error)
	GetById(id int64) (entity country.Country, err error)
}
