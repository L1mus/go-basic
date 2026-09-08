package gobasic

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTagNEmbeddingStruct(t *testing.T) {
	tagNEmbeddingStruct()
}

type Audit struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsActive bool   `json:"is_active"`
	Audit
}

func tagNEmbeddingStruct() {
	user1 := User{
		ID:        1,
		Name:      "Limus",
		Email:     "lim@mail.com",
		Password:  "Password123",
		IsActive:  true,
		CreatedAt: "2026:09:07",
		UpdatedAt: "2026:09:08",
	}
	user2 := User{
		ID:        2,
		Name:      "John",
		Email:     "john@mail.com",
		Password:  "Password123",
		IsActive:  true,
		CreatedAt: "2026:09:06",
		UpdatedAt: "2026:09:09",
	}

	fmt.Printf(`
	Name : %s
	Akun dibuat : %s
	`, user1.Name, user1.CreatedAt)

	v, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("failed to encoding to JSON")
	}

	v1, err := json.Marshal(user2)
	if err != nil {
		fmt.Println("failed to encoding to JSON")
	}

	fmt.Println(string(v))
	fmt.Println(string(v1))
}
