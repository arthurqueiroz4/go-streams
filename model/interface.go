package model

type DataMapper[T any] interface {
	Map([]string) (*T, error)
}
