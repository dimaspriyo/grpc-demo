package country

import "grpc-demo/location-service/internal/domain/country"

func toCountryResponse(entity country.Country) CountryResponse {
	return CountryResponse{
		Id:         entity.CountryId,
		Country:    entity.Country,
		LastUpdate: entity.LastUpdate,
	}
}
