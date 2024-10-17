package ports

import "context"

type FindOneOpts struct {
	Filter map[string]interface{}
}
type FindManyOpts struct {
	Filter map[string]interface{}
	Skip   int64
	Take   int64
}
type DeleteOpts struct {
	Filter map[string]interface{}
}

type Repository[T any, Q any] interface {
	FindOne(ctx context.Context, opts FindOneOpts, result *T) error
	FindMany(ctx context.Context, opts FindManyOpts, result *[]T, returnCount bool) (*int64, error)
	InsertOne(ctx context.Context, entity T) error
	DeleteOne(ctx context.Context, opts DeleteOpts) (bool, error)
}
