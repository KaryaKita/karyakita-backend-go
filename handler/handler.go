package handler

import (
	"github.com/mygetzu/karyakita-backend-go/routers"
)

var db *gorm.DB 

func run(_db *gorm.DB) *gin.Engine {
	db = _db
	return createRoutes()
}