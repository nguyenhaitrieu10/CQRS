package db

import (
	"awesomeProject/schema"
	"context"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take int64) ([]schema.Meow, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertMeow(ctx context.Context, meow schema.Meow) error{
	return impl.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, skip uint64, take int64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}

