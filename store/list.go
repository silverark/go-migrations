package store

import "github.com/jmoiron/sqlx"

type Migratable interface {
	GetName() string
	Up(*sqlx.DB)
	Down(*sqlx.DB)
}

var list []Migratable

func RegisterMigrations(migs []Migratable) {
	list = migs
}

func GetMigrationsList() []Migratable {
	return list
}

func FindMigration(name string) Migratable {
	for _, m := range list {
		if m.GetName() == name {
			return m
		}
	}
	return nil
}