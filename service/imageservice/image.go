package imageservice

type ImageRepository interface {

}

type Storage interface {

}

type Service struct {
	imageRepo ImageRepository
	storage Storage
}

// func New(imageRepo ImageRepository, storage Storage) Service {
// 	return Service{
// 		imageRepo: imageRepo,
// 		storage: storage,
// 	}
// }

func New() Service {
	return Service{}
}

