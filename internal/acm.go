package internal

// User represents a user with a chip card
type User struct {
    Name         string
    ChipCardID   string
    AccessRights int32
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
