package user_service_test

import (
	"testing"
	"time"

	m "github.com/VanessaPellegrini/CMS_headless/app/models"
	user_service "github.com/VanessaPellegrini/CMS_headless/app/services/user.service"
)

func TestCreate(t *testing.T) {

	user := m.User{
		FirstName: "Carlos",
		Email:     "Carlos@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user_service.Create(user)

	if err != nil {
		t.Error("Error creating user")
		t.Fail()
	} else {
		t.Log("User created")
	}

}

func TestRead(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
