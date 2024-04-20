package main

import (
    "fmt"
)

// Define the access level constants
const (
    NoAccess = iota
    Level1
    Level2
    Admin
)

// User represents a user with a chip card
type User struct {
    Name         string
    ChipCardID   string
    AccessRights int
}

// AccessControlSystem represents the access control system
type AccessControlSystem struct {
    Users map[string]User // Map of chip card ID to User
}

// NewAccessControlSystem creates a new access control system
func NewAccessControlSystem() *AccessControlSystem {
    return &AccessControlSystem{
        Users: make(map[string]User),
    }
}

// AddUser adds a new user to the access control system
func (acs *AccessControlSystem) AddUser(name, chipCardID string, accessRights int) {
    user := User{Name: name, ChipCardID: chipCardID, AccessRights: accessRights}
    acs.Users[chipCardID] = user
}

// CheckAccess checks if a user with the given chip card ID has access rights to a door
func (acs *AccessControlSystem) CheckAccess(chipCardID string, doorLevel int) bool {
    user, ok := acs.Users[chipCardID]
    if !ok {
        fmt.Println("Access denied: Unknown chip card ID")
        return false
    }

    if user.AccessRights >= doorLevel {
        fmt.Printf("Access granted: Welcome, %s!\n", user.Name)
        return true
    }

    fmt.Println("Access denied: Insufficient access rights")
    return false
}

func main() {
    // Create a new access control system
    acs := NewAccessControlSystem()

    // Add some users with chip cards
    acs.AddUser("Alice", "123456789", Level1)
    acs.AddUser("Bob", "987654321", Level2)
    acs.AddUser("Charlie", "543216789", Admin)

    // Simulate accessing a door with chip cards
    doorLevel := Level1 // Change this to test different door access levels
    fmt.Println("Trying to access a door with access level", doorLevel)

    // Alice tries to access the door
    _ = acs.CheckAccess("123456789", doorLevel)

    // Bob tries to access the door
    _ = acs.CheckAccess("987654321", doorLevel)

    // Charlie tries to access the door
    _ = acs.CheckAccess(" ", doorLevel)

    // Unknown chip card tries to access the door
    _ = acs.CheckAccess("000000000", doorLevel)
}