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

type UpdateOpts struct {
	Filter  map[string]interface{}
	Payload *map[string]interface{}
	Upsert  bool
}

//go:generate mockery --name=Repository --inpackage=true
type Repository[T any, Q any] interface {
	FindOne(ctx context.Context, opts FindOneOpts, result *T) error
	UpdateOne(ctx context.Context, opts UpdateOpts) (*T, error)
	FindMany(ctx context.Context, opts FindManyOpts, result *[]T, returnCount bool) (*int64, error)
	InsertOne(ctx context.Context, entity T) error
	DeleteOne(ctx context.Context, opts DeleteOpts) (bool, error)
}
