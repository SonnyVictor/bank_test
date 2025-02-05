package util

import (
	"fmt"
	"strings"
	"testing"
	// "github.com/sonnyvictok/bank_test/util"
)

func TestPassword(t *testing.T) {
	fields := strings.Fields("Vinh dep trai")

	fmt.Println(fields[0])
	fmt.Println(strings.ToLower(fields[0]))

	// password := "vinhdeptraiquadi"
	// input, _ := HashPassword("vinhdeptraiquadi")
	// fmt.Println("input", input)

	// test := CheckPassword(password, input)

	// fmt.Println("test", test)
	// require.Equal(t, password, input)

	// password := RandomString(6)

	// hashedPassword1, err := HashPassword(password)
	// require.NoError(t, err)
	// require.NotEmpty(t, hashedPassword1)

	// err = CheckPassword(password, hashedPassword1)
	// require.NoError(t, err)

	// wrongPassword := RandomString(6)
	// err = CheckPassword(wrongPassword, hashedPassword1)
	// require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// hashedPassword2, err := HashPassword(password)
	// require.NoError(t, err)
	// require.NotEmpty(t, hashedPassword2)
	// require.NotEqual(t, hashedPassword1, hashedPassword2)
}
