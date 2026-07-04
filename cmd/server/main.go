package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
)

type UserDb struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var users = []UserDb{}
var user = []User{
	{
		Id:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "123-456-7890",
	},
	{
		Id:    3,
		Name:  "Jane3 Smith",
		Email: "jane.smith@example.com",
		Phone: "098-765-4321",
	},
	{
		Id:    4,
		Name:  "Jane4 Smith",
		Email: "jane.smith@example.com",
		Phone: "098-765-4321",
	},
	{
		Id:    4,
		Name:  "Jane4 Smith",
		Email: "jane4.smith@example.com",
		Phone: "098-765-4321",
	},
}

func main() {
	fmt.Println(users)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("POST /users", createUsers)
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("GET /users/{id}", getUserById)
	mux.HandleFunc("PATCH /users/{id}", updateUserById)
	mux.HandleFunc("DELETE /users/{id}", userDeleted)

	fmt.Println("Server is running on port:", 5000)
	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome my server")
}
func createUsers(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	page := r.URL.Query().Get("page")
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user UserDb
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	users = append(users, user)

	fmt.Println(users)
	fmt.Fprintln(w, "User created successfully", page, user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user, _ := json.Marshal(user)
	w.Write(user)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, u := range user {
		if u.Id == id {
			userJson, _ := json.Marshal(u)
			w.Header().Set("Content-Type", "application/json")
			w.Write(userJson)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var updatedUser User
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	for i, u := range user {
		if u.Id == id {
			user[i].Name = updatedUser.Name
			user[i].Email = updatedUser.Email
			user[i].Phone = updatedUser.Phone

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, "User updated successfully")
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)

}

func userDeleted(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for i, u := range user {
		if u.Id == id {
			user = slices.Delete(user, i, i+1)
			fmt.Fprintln(w, "User deleted successfully")
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
