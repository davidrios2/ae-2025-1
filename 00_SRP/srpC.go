package _0_SRP

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Name  string
	Email string
}

// UserManager handles user data management.
type UserManager struct {
	users []User
}

// NewUserManager creates a new UserManager.
func NewUserManager(users []User) *UserManager {
	return &UserManager{users: users}
}

// AddUser adds a new user to the list.
func (um *UserManager) AddUser(name, email string) {
	user := User{Name: name, Email: email}
	um.users = append(um.users, user)
}

func (um *UserManager) GetUsers() []User {
	return um.users
}

type UserListView struct {
	output io.Writer
}

// NewUserListView creates a new UserListView.
func NewUserListView(output io.Writer) *UserListView {
	return &UserListView{output: output}
}

// Display displays the user list.
func (ulv *UserListView) Display(users []User) error {
	_, err := ulv.output.Write([]byte("User List:\n")) // Use the io.Writer
	if err != nil {
		return fmt.Errorf("error writing to output: %w", err)
	}
	for _, user := range users {
		_, err := ulv.output.Write([]byte(fmt.Sprintf("- %s: %s\n", user.Name, user.Email)))
		if err != nil {
			return fmt.Errorf("error writing to output: %w", err)
		}
	}
	return nil
}

// UserStorage handles saving and loading user data.
type UserStorage struct {
	filename string
}

// NewUserStorage creates a new UserStorage.
func NewUserStorage(filename string) *UserStorage {
	return &UserStorage{filename: filename}
}

func (us *UserStorage) SaveUsers(users []User) error {
	file, err := os.Create(us.filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err) // Wrap for more context
	}
	defer file.Close()

	for _, user := range users {
		_, err := file.WriteString(fmt.Sprintf("%s,%s\n", user.Name, user.Email))
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	return nil
}

// LoadUsers loads user data from a file.
func (us *UserStorage) LoadUsers() ([]User, error) {
	file, err := os.Open(us.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil // Return empty slice, not an error, if file doesn't exist
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var users []User
	var content []byte
	content, err = io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Println("Skipping invalid line:", line) // Handle invalid lines
			continue
		}
		users = append(users, User{Name: parts[0], Email: parts[1]})
	}
	return users, nil
}

func SrpC() {
	storage := NewUserStorage("users.txt")
	initialUsers, err := storage.LoadUsers()
	if err != nil {
		fmt.Println("Error loading users:", err) // important: handle error
		initialUsers = []User{}                  // Initialize to empty if loading fails
	}

	manager := NewUserManager(initialUsers)
	view := NewUserListView(os.Stdout) // Use os.Stdout

	manager.AddUser("Alice", "alice@example.com")
	manager.AddUser("Bob", "bob@example.com")

	usersToDisplay := manager.GetUsers()
	err = view.Display(usersToDisplay)
	if err != nil {
		fmt.Println("Error displaying users:", err)
	}

	err = storage.SaveUsers(usersToDisplay)
	if err != nil {
		fmt.Println("Error saving users:", err) // important: handle error
	}
	fmt.Println("Final user list:")
	for _, user := range manager.GetUsers() {
		fmt.Printf("- %s: %s\n", user.Name, user.Email)
	}
}
