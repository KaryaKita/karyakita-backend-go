package main

import (
	"github.com/mygetzu/karyakita-backend-go/handler"
	"github.com/mygetzu/karyakita-backend-go/config"
	"github.com/mygetzu/karyakita-backend-go/controllers"
)

func main() {
	dbsettings, err := config.LoadDBSettings()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db, err := gorm.Open("mysql",
		dbsettings.User+":"+
			dbsettings.Password+"@tcp("+
			dbsettings.Host+":"+
			dbsettings.Port+
			")/"+dbsettings.DBName+
			"?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	router := handler.run(db)
}
