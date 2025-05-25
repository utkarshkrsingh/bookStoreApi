package initializers

import "github.com/utkarshkrsingh/bookStoreApi/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
