package services

import (
    "backendtest/repositories"
    "backendtest/models"
)

func GetPersonInfo(personID string) (*models.Person, error) {
    return repositories.GetPersonByID(personID)
}

func GetAllPersons(offset, limit int) ([]*models.Person, int, error) {
    return repositories.GetAllPersons(offset, limit)
}

func CreatePerson(person *models.Person) error {
    err := repositories.CreatePerson(person)
    if err != nil {
        return err
    }
    return nil
}