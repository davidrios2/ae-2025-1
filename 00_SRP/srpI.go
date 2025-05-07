package _0_SRP

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Consumer struct {
	Name  string
	Email string
}
type UserHandler struct {
	users    []Consumer
	filename string
}

func NewUserHandler(filename string) *UserHandler {
	return &UserHandler{
		filename: filename,
		users:    []Consumer{},
	}
}

func (uh *UserHandler) AddUser(name, email string) error {
	user := Consumer{Name: name, Email: email}
	uh.users = append(uh.users, user)
	uh.displayUserList()
	return uh.saveUsersToStorage()
}

func (uh *UserHandler) displayUserList() {
	fmt.Println("User List:")
	for _, user := range uh.users {
		fmt.Printf("- %s: %s\n", user.Name, user.Email)
	}
}

func (uh *UserHandler) saveUsersToStorage() error {
	file, err := os.Create(uh.filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	for _, user := range uh.users {
		_, err := file.WriteString(fmt.Sprintf("%s,%s\n", user.Name, user.Email))
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	fmt.Println("Users saved to users.txt")
	return nil
}

func (uh *UserHandler) GetUsers() []Consumer {
	return uh.users
}

func (uh *UserHandler) LoadUsers() ([]Consumer, error) {
	file, err := os.Open(uh.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Consumer{}, nil // Return empty slice, not an error, if file doesn't exist
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var users []Consumer
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
		users = append(users, Consumer{Name: parts[0], Email: parts[1]})
	}
	return users, nil
}

func SrpI() {
	handler := NewUserHandler("users.txt")
	initialUsers, err := handler.LoadUsers()
	if err != nil {
		fmt.Println("Error loading users:", err)
	}
	handler.users = initialUsers

	err = handler.AddUser("Alice", "alice@example.com")
	if err != nil {
		fmt.Println("Error adding user", err)
	}
	err = handler.AddUser("Bob", "bob@example.com")
	if err != nil {
		fmt.Println("Error adding user", err)
	}

	fmt.Println("Final user list:")
	for _, user := range handler.GetUsers() {
		fmt.Printf("- %s: %s\n", user.Name, user.Email)
	}
}
