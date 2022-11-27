package token

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	tkn, _ := Sign(fmt.Sprintf("%d%d", 1, 123))
	fmt.Println("[token]:", tkn)
	//tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJYLVVzZXItSWQiOiI4MjQ2NDA2MDIyMjQxMjMiLCJleHAiOjIwMjkwMzk5NTcsImlhdCI6MTY2OTAzOTk1NywibmJmIjoxNjY5MDM5OTU3fQ.OPOSmQn3bFm6vuUoblzUHS2dO9zxEOFFALps-lRvySc"
	identityKey, _ := Parse(tkn, config.key)

	fmt.Println("[Key]:", identityKey)

}
