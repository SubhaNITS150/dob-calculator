package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const baseURL = "http://localhost:8090"

// Same DTO as handler
type CreateUserDTO struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========= USER MENU =========")
		fmt.Println("1. List all users")
		fmt.Println("2. Create user")
		fmt.Println("3. Update user")
		fmt.Println("4. Delete user")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			listUsers()

		case "2":
			createUser(reader)

		case "3":
			updateUser(reader)

		case "4":
			deleteUser(reader)

		case "5":
			fmt.Println("Exiting program....")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// ================= LIST USERS =================
func listUsers() {
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var result any
	json.NewDecoder(resp.Body).Decode(&result)

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}

// ================= CREATE USER =================
func createUser(reader *bufio.Reader) {
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter DOB (YYYY-MM-DD): ")
	dob, _ := reader.ReadString('\n')

	req := CreateUserDTO{
		Name: strings.TrimSpace(name),
		Dob:  strings.TrimSpace(dob),
	}

	body, _ := json.Marshal(req)

	resp, err := http.Post(
		baseURL+"/users",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read API response
	var user struct {
		ID  string `json:"id"`
		Name string `json:"name"`
		Dob  string `json:"dob"`
		Age  int    `json:"age"`
	}

	json.NewDecoder(resp.Body).Decode(&user)

	fmt.Println("\nUser created successfully")
	fmt.Println("Name :", user.Name)
	fmt.Println("DOB  :", user.Dob)
	fmt.Println("Age  :", user.Age, "years")
}

// ================= UPDATE USER =================
func updateUser(reader *bufio.Reader) {
	fmt.Print("Enter user ID: ")
	id, _ := reader.ReadString('\n')

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter new DOB (YYYY-MM-DD): ")
	dob, _ := reader.ReadString('\n')

	req := CreateUserDTO{
		Name: strings.TrimSpace(name),
		Dob:  strings.TrimSpace(dob),
	}

	body, _ := json.Marshal(req)

	request, _ := http.NewRequest(
		http.MethodPut,
		baseURL+"/users/"+strings.TrimSpace(id),
		bytes.NewBuffer(body),
	)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err := client.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User updated successfully")
}

// ================= DELETE USER =================
func deleteUser(reader *bufio.Reader) {
	fmt.Print("Enter user ID: ")
	id, _ := reader.ReadString('\n')

	request, _ := http.NewRequest(
		http.MethodDelete,
		baseURL+"/users/"+strings.TrimSpace(id),
		nil,
	)

	client := &http.Client{}
	_, err := client.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User deleted successfully")
}
