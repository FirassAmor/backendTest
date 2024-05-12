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


// GetAllPersons retrieves a list of persons from the database with pagination support
func GetAllPersons(offset, limit int) ([]*models.Person, int, error) {
    var persons []*models.Person

    // Query to get the total count of persons
    var totalCount int
    err := database.DB.QueryRow("SELECT COUNT(*) FROM person").Scan(&totalCount)
    if err != nil && err != sql.ErrNoRows {
        return nil, 0, err
    }

    // Check if the offset exceeds the total count
    if offset >= totalCount {
        return persons, totalCount, nil
    }

    // If the limit exceeds the remaining count, adjust the limit
    if limit > totalCount-offset {
        limit = totalCount - offset
    }

    query := "SELECT id, name, age FROM person LIMIT ? OFFSET ?"
    rows, err := database.DB.Query(query, limit, offset)
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()

    for rows.Next() {
        var person models.Person
        if err := rows.Scan(&person.ID, &person.Name, &person.Age); err != nil {
            return nil, 0, err
        }
        persons = append(persons, &person)
    }

    return persons, totalCount, nil
}
