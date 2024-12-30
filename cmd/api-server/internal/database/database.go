package database

type Database interface {
	Ping() error
	Close() error
}
