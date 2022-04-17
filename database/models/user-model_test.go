package userModel

import (
	"api/constants"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelShouldBeInvalidUsername(t *testing.T) {
	// GIVEN a fake username
	fakeRandomUsername := "randomUserName"

	// WHEN I call the function with the fake username
	got, err := GetByUserName(fakeRandomUsername)

	// THEN I should be able to see an error message
	assert.EqualError(t, err, fmt.Sprintf("%s!", constants.InvalidUserName))
	// AND I should not be able to see any result
	assert.Empty(t, got)
}

func TestUserModelShouldNotThrowError(t *testing.T) {
	// GIVEN a random fake username
	fakeUsername := "istartedtoliveasme"

	// WHEN I call the function with the fake username
	got, err := GetByUserName(fakeUsername)

	// THEN I should not be able to see an error message
	assert.NoError(t, err)
	// AND I should be able to have a result
	assert.NotEmpty(t, got)
}

func TestUserModelShouldSeeResultValue(t *testing.T) {
	// GIVEN a random fake username
	fakeUsername := "istartedtoliveasme"

	// WHEN I call the function with the fake username
	got, _ := GetByUserName(fakeUsername)

	// THEN I should be able to see the email containing the fake username
	assert.Contains(t, got.Email, fakeUsername)
}
