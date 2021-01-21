package Singleton

/*
	Singleton is a pattern that ensure only one instance of type exists in a program
and provide global access to it
	Example: Database connection
*/

import (
	"gorm.io/gorm"
)
var (
	SingletonDB *gorm.DB
)
