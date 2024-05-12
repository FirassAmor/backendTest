package repositories

import (
    "database/sql"
    "backendtest/models"
	"backendtest/database"
)

func GetAddressByID(addressID string) (*models.Address, error) {
    var address models.Address

    query := "SELECT id, city, state, street1, street2, zip_code FROM address WHERE id = ?"
    err := database.DB.QueryRow(query, addressID).Scan(&address.ID, &address.City, &address.State, &address.Street1, &address.Street2, &address.ZipCode)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil 
        }
        return nil, err
    }

    return &address, nil
}
