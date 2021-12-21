package country

import (
	"gorm.io/gorm"
	"grpc-demo/location-service/internal/domain/country"
)

type CountryRepository struct {
	db *gorm.DB
}

func (r *CountryRepository) FindAll() (entity []country.Country, err error) {

	err = r.db.Find(&entity).Error
	return

}

func (r *CountryRepository) FindById(id int64) (entity country.Country, err error) {
	err = r.db.Find(&entity, id).Error
	return
}

func NewCountryRepository(db *gorm.DB) country.CountryRepositoryIFace {
	return &CountryRepository{
		db: db,
	}
}
