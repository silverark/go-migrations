package main

import (
	"flag"
	gm "go-migrations"
	gmStore "go-migrations/store"
	"go-migrations/template/migrations/list"
)

var isRollback *bool

func init() {
	isRollback = flag.Bool("rollback", false, "")
	flag.Parse()
}

func main() {
	if *isRollback {
		gm.Rollback(getMigrationsList())
		return
	}

	gm.Migrate(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateExampleTable{},
		// Register you migrations here
	}
}
