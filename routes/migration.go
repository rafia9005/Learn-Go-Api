package routes

import "Learn-Go-Api/database"

func RunMigrate(dataModel interface{}) {
	database.DB.AutoMigrate(dataModel)
}
