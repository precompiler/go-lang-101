package singleton

import (
	"fmt"
)

type databaseHelper struct{}

var instance *databaseHelper

func GetInstance() *databaseHelper {
	// This is not thread safe
	if instance == nil {
		instance = new(databaseHelper)
	}
	return instance
}

func (d *databaseHelper) CreateConnection() {
	fmt.Println("Create a database connection")
}

func (d *databaseHelper) CloseConnection() {
	fmt.Println("Close a database connection")
}
