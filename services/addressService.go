package services

import (
    "backendtest/repositories"
    "backendtest/models"
)

func GetAddressInfo(addressID string) (*models.Address, error) {
    return repositories.GetAddressByID(addressID)
}
