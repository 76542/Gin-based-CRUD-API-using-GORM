// implements logic and performs database operations
package services

import (
	"errors"
	"example/Crud/config"
	"example/Crud/logger"
	"example/Crud/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ITEM interface {
	CreateItem(item *models.Item) (models.Item, error)
	FindItem(id int) (models.Item, error)
	FindAllItems() ([]models.Item, error)
	FindItemByCategory(category string) ([]models.Item, error)
}

type item struct {
	db *gorm.DB
}

func NewItem(db *gorm.DB) *item {
	return &item{
		db: config.DB,
	}
}

var Item = NewItem(config.DB)

func (i *item) CreateItem(item *models.Item) (*models.Item, error) {
	logger.CustomLogger.Debug("Creating Item : ", item)
	if err := i.db.Create(item).Error; err != nil {
		logger.CustomLogger.Error("Error creating item : ", err)
		return nil, err
	}
	logger.CustomLogger.Info("Iten created successfully: ", item)
	return item, nil

}

func GetItems() (*models.Item, error) {
	//variable to hold the item
	var item models.Item
	logger.CustomLogger.Debug("Fetching all items")
	if err := config.DB.Find(&item).Error; err != nil {
		logger.CustomLogger.Error("Failed to fetch items", err)
		return nil, err
	}
	logger.CustomLogger.Info("Items fetched successfully")
	return &item, nil
}

func GetItemByID(id uint) (*models.Item, error) {
	//variable to hold the item
	var item models.Item
	logger.CustomLogger.Debug("Fetching items by ID: ", id)
	if err := config.DB.First(&item, id).Error; err != nil {
		//checks if the error occurs bcoz the item is not present in the list
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.CustomLogger.Error("Item not found with id : ", id)
			return nil, err
		}
		//for all other errors
		logger.CustomLogger.Error("Item not found with id : ", err)
		return nil, err
	}
	logger.CustomLogger.Info("Item fetched successfully")
	return &item, nil
}

func UpdateItem(id uint, updatedItem *models.Item) (*models.Item, error) {
	//variable to hold the existing item
	var item models.Item
	logger.CustomLogger.Debug("Updating item by ID: ", id)
	if err := config.DB.First(&item, id).Error; err != nil {
		//checks if the error occurs bcoz the item is not present in the list
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.CustomLogger.Error("Item not found with id : ", id)
			return nil, errors.New("item not found")
		}
		logger.CustomLogger.Error("Failed to find item to update : ", err)
		return nil, err
	}
	item.Name = updatedItem.Name
	item.Price = updatedItem.Price

	//saves the updated name in the database
	if err := config.DB.Save(&item).Error; err != nil {
		logger.CustomLogger.Error("Failed to update item : ", err)
		return nil, err
	}

	logger.CustomLogger.Info("Item updated successfully: ", item)
	return &item, nil
}

func DeleteItem(id uint) error {
	logger.CustomLogger.Debug("Deleting item by ID: ", id)
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.CustomLogger.Error("Item not found with id : ", id)
			return errors.New("item not found")
		}
		logger.CustomLogger.Error("Failed to delete item: ", err)
		return err
	}
	logger.CustomLogger.Info("Successfully deleted item with id: ", id)
	return nil
}

type mockItem struct {
	mock.Mock
}

func (m *mockItem) CreateItem(item *models.Item) (*models.Item, error) {
	mockItem := models.Item{
		Name: item.Name,
	}
	return &mockItem, nil

}
