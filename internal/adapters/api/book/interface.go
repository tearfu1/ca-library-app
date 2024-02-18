package book

import (
	"ca-library-app/internal/domain/book"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *book.Author
	GetAll(ctx context.Context, limit, offset int) []*book.Author
	Create(ctx context.Context, dto *book.CreateBookDTO) *book.Author
}
