package book

import (
	"ca-library-app/internal/domain/book"
)

type bookStorage struct {
}

func NewStorage() book.Storage {
	return &bookStorage{}
}

func (bs *bookStorage) GetOne(uuid string) *book.Author {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) []*book.Author {
	return nil
}
func (bs *bookStorage) Create(book *book.Author) *book.Author {
	return nil
}
func (bs *bookStorage) Delete(book *book.Author) error {
	return nil
}
