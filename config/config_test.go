package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	InitConfig()
	c := GetConfig()
	fmt.Printf("%+v\n", c)
}
