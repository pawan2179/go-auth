package db

type Storage struct {
	UserRespository UserRespository
}

func NewStorage() *Storage {
	return &Storage{
		UserRespository: &UserRespositoryImpl{},
	}
}
