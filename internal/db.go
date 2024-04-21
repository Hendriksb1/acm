package internal

import (
	"acm/api/pb"
	"errors"
)

type Database interface {
	AddUser(*User) error
	GetUserByChipCardId(id string) (*User, error)
}

// TODO set up postgress
type PostgresDB struct {}

// TODO set up postgress
func (db *PostgresDB) AddUser(*User) error {
	return nil
}

// TODO set up postgress
func (db *PostgresDB) GetUserByChipCardId(id string) (*User, error) {
	return &User{}, nil
}

type MockDB struct {}
const mockChipcardId = "123"

func (m *MockDB) AddUser(u *User) error {
	
	if u.ChipCardID == mockChipcardId {
		return errors.New("user allready exists")
	}
	
	return nil
}

// Implement the Database interface for the mock
func (m *MockDB) GetUserByChipCardId(id string) (*User, error) {

	if id != mockChipcardId {
		return nil, errors.New("user not found")
	}

	// Simulate behavior of fetching user without hitting the database
	return &User{
		Name:         "mockUser",
		ChipCardID:   id,
		AccessRights: int32(pb.AccessLevel_LEVEL_1),
	}, nil
}
