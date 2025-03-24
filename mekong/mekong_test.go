package mekong

import (
	"fmt"
	"testing"
)

func TestMekong(t *testing.T) {
	me := Mekong{
		UserName: "xxxxx",
		Password: "xxxxx",
	}
	fmt.Println(me)

	data, err := me.SendSMS("xxxx", "this is test from go", "000000000", 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
