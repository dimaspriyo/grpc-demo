package user

type UserRepositoryIFace interface {
	FindById(id int64) (entity *UserEntity, err error)
	Save(entity *UserEntity) (err error)
	FindAll() entity[]
}
