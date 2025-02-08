package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/sonnyvictok/bank_test/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {

	randomString := util.RandomString(32)

	randomPaseto, err := NewPasetoMaker(randomString)

	require.NoError(t, err)

	fmt.Println(randomPaseto)
	duration := time.Minute
	// issuedAt := time.Now()
	// expiredAt := issuedAt.Add(duration)
	createToken, payload, err := randomPaseto.CreateToken("vinh", duration)
	fmt.Println("createtoken", createToken)
	fmt.Println("payload", payload)

	//check token verify
	check, err := randomPaseto.VerifyToken(createToken)

	fmt.Println("check token", check)

}
