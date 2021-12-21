package country

type CountryRepositoryIFace interface {
	FindAll() (entity []Country, err error)
	FindById(id int64) (entity Country, err error)
}
