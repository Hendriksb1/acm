package internal

import (
	"database/sql"
	"errors"

	"golang.org/x/exp/slog"
)

// Database interface defines methods for interacting with the database.
type Database interface {
	AddUser(*User) error
	DeleteUserByChipCardId(id string) error
	GetUserByChipCardId(id string) (*User, error)
}

// Postgres represents a PostgreSQL database.
type Postgres struct {
	db *sql.DB
}

// NewPostgresDB creates a new instance of Postgres and initializes the connection to the database.
func NewPostgresDB(dbConnString string, log *slog.Logger) (*Postgres, error) {
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Ping the database to ensure the connection is valid
	err = db.Ping()
	if err != nil {
		log.Error("failed to ping database")
		return nil, err
	}

	log.Info("connected to the database")
	return &Postgres{db: db}, nil
}

// InnitUserTable initializes the Users table in the database if it doesn't exist.
func (p *Postgres) InnitUserTable(log *slog.Logger) error {
	// Create Users table if it doesn't exist
	_, err := p.db.Exec(`CREATE TABLE IF NOT EXISTS Users (
        ChipCardID TEXT PRIMARY KEY,
        Name TEXT,
        AccessRights INT
    )`)
	if err != nil {
		log.Error("failed to create user table database")
		return err
	}

	log.Info("users table created or already exists")
	return nil
}

// AddUser inserts a new user into the database.
func (p *Postgres) AddUser(u *User) error {
	// Execute the SQL statement to insert a new user into the "Users" table
	_, err := p.db.Exec("INSERT INTO Users (ChipCardID, Name, AccessRights) VALUES ($1, $2, $3)", u.ChipCardID, u.Name, u.AccessRights)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByChipCardId deletes a user from the database based on the chip card ID.
func (p *Postgres) DeleteUserByChipCardId(id string) error {
	// Execute the SQL DELETE statement to remove the user with the specified ChipCardID
	_, err := p.db.Exec("DELETE FROM Users WHERE ChipCardID = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByChipCardId retrieves a user from the database based on the chip card ID.
func (p *Postgres) GetUserByChipCardId(id string) (*User, error) {
	// Execute the SQL SELECT statement to retrieve the user with the specified ChipCardID
	row := p.db.QueryRow("SELECT ChipCardID, Name, AccessRights FROM Users WHERE ChipCardID = $1", id)

	// Create a new User object to store the result
	user := &User{}

	// Scan the row to extract the values into the User object
	err := row.Scan(&user.ChipCardID, &user.Name, &user.AccessRights)
	if err != nil {
		// Check if the user was not found
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		// Return other errors
		return nil, err
	}

	return user, nil
}


