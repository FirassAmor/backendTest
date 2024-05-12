package services

import (
    "backendtest/repositories"
    "backendtest/models"
)

func GetPersonInfo(personID string) (*models.Person, error) {
    return repositories.GetPersonByID(personID)
}

func CreatePerson(person *models.Person) error {
    err := repositories.CreatePerson(person)
    if err != nil {
        return err
    }
    return nil
}