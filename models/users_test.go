package models

import (
	"fmt"
	"testing"
	"time"
)

func testingUserService() (*UserService, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "test_site"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}

	us.db.LogMode(false)

	// Clear the users table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "test",
		Email: "ddd@dsdsd.ru",
	}

	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected ID > 0. Got %d", user.ID)
	}

	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Got %s", user.CreatedAt)
	}

	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected UpdatedAt to be recent. Got %s", user.UpdatedAt)
	}
}
