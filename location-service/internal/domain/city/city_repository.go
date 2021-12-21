package city

type CityRepositoryIFace interface {
	FindAll() (entity []City, err error)
	FindById(id int64) (entity City, err error)
}
