package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("a@gmail.com", "hunter2001")
	assert.Nil(t, err)
	assert.NotNil(t, user.EncryptedPassword)
}

func TestUserPassword(t *testing.T) {
	pw := "hunter2001"
	user, err := NewUser("a@gmail.com", pw)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(pw))
	assert.False(t, user.ValidatePassword("hunter2005"))
}
