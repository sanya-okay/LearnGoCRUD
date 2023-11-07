package database

import (
	"LearnGoCRUD/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PersonsDB struct {
	db *sqlx.DB
}

func (d *PersonsDB) CreatePerson(person models.Person) error {
	_, err := d.db.Exec("INSERT INTO person (name, age, license) VALUES ($1, $2, $3)", person.Name, person.Age, person.License)
	if err != nil {
		return err
	}
	return nil
}

func (d *PersonsDB) GetAllPersons() ([]models.Person, error) {
	rows, err := d.db.Query("SELECT id, name, age, license FROM person") //что лежит в rows?
	if err != nil {
		return nil, err
	}
	persons := make([]models.Person, 0)
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.Id, &person.Name, &person.Age, &person.License); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, rows.Err()
}

func (d *PersonsDB) DeletePerson(name string) error {
	_, err := d.db.Exec("DELETE FROM person WHERE name=$1", name)
	if err != nil {
		return err
	}
	return fmt.Errorf("person %s was successfully deleted or it was not found in database", name)
}

func (d *PersonsDB) UpdatePerson(person models.Person) error {
	_, err := d.db.Exec("UPDATE person SET name=$1, age=$2, license=$3 WHERE id=$4", person.Name, person.Age, person.License, person.Id)
	if err != nil {
		return err
	}
	return fmt.Errorf("person name = %s was successfully update or it was not found in database", person.Name)
}

func NewDB(db *sqlx.DB) *PersonsDB {
	return &PersonsDB{db}
}
