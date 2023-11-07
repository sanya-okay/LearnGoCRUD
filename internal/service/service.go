package service

import (
	"LearnGoCRUD/internal/models"
	"fmt"
	"math/rand"
)

type PersonDB interface {
	CreatePerson(person models.Person) error
	GetAllPersons() ([]models.Person, error)
	UpdatePerson(person models.Person) error
	DeletePerson(name string) error
}

type Service struct {
	PersonDB PersonDB
}

func (s *Service) CreatePerson(person models.Person) error {
	//тут бизнес логика (checkBoss, ValidationPerson) и вызов бд
	if rand.Intn(2) == 0 {
		return fmt.Errorf("boss is angry")
	}
	if !person.ValidationPerson() {
		return fmt.Errorf("person is not validation")
	}
	err := s.PersonDB.CreatePerson(person)
	if err != nil {
		return fmt.Errorf("service CreatePerson error:%s", err)
	}
	return nil
}

func (s *Service) GetAllPersons() ([]models.Person, error) {
	persons, err := s.PersonDB.GetAllPersons()
	if err != nil {
		return nil, fmt.Errorf("service GetAllPersons error:%s", err)
	}
	return persons, nil

}

func (s *Service) DeletePerson(name string) error {
	err := s.PersonDB.DeletePerson(name)
	if err == nil {
		return fmt.Errorf("service DeletePerson error")
	}
	return err
}

func (s *Service) UpdatePerson(person models.Person) error {
	err := s.PersonDB.UpdatePerson(person)
	if err != nil {
		return err
	}
	return fmt.Errorf("service UpdatePerson error:%s", err)
}

func NewService(service PersonDB) *Service {
	return &Service{
		PersonDB: service,
	}
}
