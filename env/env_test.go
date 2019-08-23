package env

import (
	"fmt"
	"testing"
)

func Test_Env(t *testing.T) {
	dict, err := Load(".env", "cc.env", "conf.env")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	host, _ := dict.GetBool("IS")
	fmt.Println(host)
}
