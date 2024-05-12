package repositories

import (
    "database/sql"
    "backendtest/models"
	"backendtest/database" 

)


func GetPersonByID(personID string) (*models.Person, error) {
    var person models.Person

    query := "SELECT id, name, age FROM person WHERE id = ?"
    err := database.DB.QueryRow(query, personID).Scan(&person.ID, &person.Name, &person.Age)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return &person, nil
}

func CreatePerson(person *models.Person) error {
    _, err := database.DB.Exec("INSERT INTO person (name, age) VALUES (?, ?)", person.Name, person.Age)
    if err != nil {
        return err
    }
    return nil
}