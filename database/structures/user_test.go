package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_ComparePasswordShouldMatch(t *testing.T) {
	// GIVEN the random user password
	randomPassword := "myRandomPassword"
	// AND added to the user data
	user := User{
		Password: randomPassword,
	}

	// WHEN I call the function with the identical password
	got := user.ComparePassword(randomPassword)

	// THEN I should be able to see that it matches by having a value of TRUE
	assert.True(t, got)
}

func TestUser_ComparePasswordShouldNotMatch(t *testing.T) {
	// GIVEN the user data
	user := User{
		Password: "randomPassword",
	}

	// WHEN I call the function with the un-identical password
	got := user.ComparePassword("anotherRandomPassword")

	// THEN I should be able to see that it's NOT matching by having a value of FALSE
	assert.False(t, got)
}
