package database

import "url-shortner/database/entity"

func Migration() error {
	database := Database{}
	db, err := database.Connect()
	if err != nil {
		return err
	}
	url := entity.Url{}
	db.AutoMigrate(&url)
	return nil
}
