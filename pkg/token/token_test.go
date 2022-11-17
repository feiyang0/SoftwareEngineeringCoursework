package token

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	tkn, _ := Sign(fmt.Sprintf("%d%d", 1, 123))
	fmt.Println("[token]:", tkn)
	identityKey, _ := Parse(tkn, config.key)

	fmt.Println("[Key]:", identityKey)

	role := "2"
	if identityKey[0:1] < role {
		fmt.Println("[err]: 没有权限")
	}
}
