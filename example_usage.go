package main

import (
	"fmt"
	"log"
)

func main() {
	// Create a new client instance
	client := NewUserAPIClient("http://localhost:8080")

	// Example 1: Add a new user
	fmt.Println("=== Adding a new user ===")
	addReq := AddUserRequest{
		Authorization: "Bearer your-token-here",
		Name:          "john_doe",
		Age:           "25",
	}

	addResp, err := client.AddUser(addReq)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
	} else {
		fmt.Printf("Add user response: %+v\n", addResp)
	}

	// Example 2: Get user information
	fmt.Println("\n=== Getting user information ===")
	getReq := GetUserRequest{
		Authorization: "Bearer your-token-here",
		Name:          "john_doe",
		Delete:        false, // optional parameter
	}

	getResp, err := client.GetUser(getReq)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
	} else {
		fmt.Printf("Get user response: %+v\n", getResp)
	}

	// Example 3: Update user information
	fmt.Println("\n=== Updating user information ===")
	updateReq := UpdateUserRequest{
		Authorization: "Bearer your-token-here",
		Name:          "john_doe",
		Age:           "26",
	}

	updateResp, err := client.UpdateUser(updateReq)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
	} else {
		fmt.Printf("Update user response: %+v\n", updateResp)
	}

	// Example 4: Delete user
	fmt.Println("\n=== Deleting user ===")
	deleteReq := DeleteUserRequest{
		Authorization: "Bearer your-token-here",
		Name:          "john_doe",
	}

	deleteResp, err := client.DeleteUser(deleteReq)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
	} else {
		fmt.Printf("Delete user response: %+v\n", deleteResp)
	}
}