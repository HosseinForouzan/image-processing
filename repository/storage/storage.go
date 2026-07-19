package storage

type LocalStorage struct {
	root string
}

func New(root string) LocalStorage {
	return LocalStorage{root: root}
}