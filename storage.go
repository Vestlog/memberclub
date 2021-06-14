package main

import (
	"errors"
	"fmt"
	"time"
)

type Storage interface {
	SaveMember(string, string) error
	GetAllMembers() []Member
}

type SliceStorage struct {
	Storage []Member
}

func (ss *SliceStorage) NewID() int {
	return len(ss.Storage) + 1
}

func (ss *SliceStorage) SaveMember(name string, email string) error {
	member := Member{
		ID:           ss.NewID(),
		Name:         name,
		Email:        email,
		Registration: Date(),
	}
	for _, e := range ss.Storage {
		if member.Email == e.Email {
			return errors.New("user with this email already exists")
		}
	}
	ss.Storage = append(ss.Storage, member)
	return nil
}

func (ss *SliceStorage) GetAllMembers() []Member {
	return ss.Storage
}

func CreateSliceStorage() *SliceStorage {
	return &SliceStorage{
		Storage: make([]Member, 0, 10),
	}
}

func Date() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%02d.%02d.%04d", d, m, y)
}
