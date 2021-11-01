package persistence

import (
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/models"
	"gorm.io/gorm"
)

// CreateCustomer creates and persists a new customer.
func CreateCustomer(db *gorm.DB, customer models.Customer) (models.Customer, error) {
	if err := db.Model(&models.Customer{}).Create(&customer).Error; err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

// GetCustomerByUsername returns a customer based on the username provided as argument.
func GetCustomerByUsername(db *gorm.DB, application, service, username string) (models.Customer, error) {
	var result models.Customer

	err := db.Model(&models.Customer{}).
		Where("handle = ? AND application = ? AND service = ?", username, application, service).
		First(&result).Error

	if err != nil {
		return models.Customer{}, err
	}
	return result, nil
}

// GetCustomerByCustomerID returns a customer based on the customer id provided as argument.
func GetCustomerByCustomerID(db *gorm.DB, application, service, id string) (models.Customer, error) {
	var result models.Customer
	err := db.Model(&models.Customer{}).
		Where("customer_id = ? AND application = ? AND service = ?", id, application, service).
		First(&result).Error

	if err != nil {
		return models.Customer{}, err
	}
	return result, nil
}
