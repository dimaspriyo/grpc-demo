package country

import "grpc-demo/location-service/internal/domain/country"

type CountryService struct {
	repo country.CountryRepositoryIFace
}

func (s *CountryService) GetAll() (entity []country.Country, err error) {
	entity, err = s.repo.FindAll()
	return
}

func (s *CountryService) GetById(id int64) (entity country.Country, err error) {
	entity, err = s.repo.FindById(id)
	return
}

func NewCountryService(repo country.CountryRepositoryIFace) CountryServiceIFace {
	return &CountryService{
		repo: repo,
	}
}
