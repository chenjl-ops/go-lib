package rsa

import (
	"fmt"
	"testing"
)

func TestGenerateRSAKey(t *testing.T) {
	_, _, err := NewGenerateRSAKey()
	if err != nil {
		t.Error(err)
	}
}

func TestLoadPublicKeyFromString(t *testing.T) {

}

func TestLoadPrivateKeyFromString(t *testing.T) {

}

func TestSignRSA(t *testing.T) {
	privateKey, _, err := NewGenerateRSAKey()
	if err != nil {
		t.Error(err)
	}
	privateKeys, err := LoadPrivateKeyFromString(privateKey)
	if err != nil {
		t.Error(err)
	}
	sign, err := SignRSA("hello", privateKeys)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("sign: ", sign)

}

func TestVerifyRSAFromString(t *testing.T) {
	privateKey, publicKey, err := NewGenerateRSAKey()
	if err != nil {
		t.Error(err)
	}
	privateKeys, err := LoadPrivateKeyFromString(privateKey)
	if err != nil {
		t.Error(err)
	}
	publicKeys, err := LoadPublicKeyFromString(publicKey)
	if err != nil {
		t.Error(err)
	}
	sign, err := SignRSA("hello", privateKeys)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("sign: ", sign)
	isOk := VerifyRSA(publicKeys, "hello", sign)
	if isOk == true {
		t.Log("verify success")
	} else {
		t.Error("verify fail")
	}

}

func TestVerifyRSA(t *testing.T) {
	_, _, err := NewGenerateRSAKey()
	if err != nil {
		t.Error(err)
	}
	privateKeys, err := LoadPrivateKey("/tmp/private.pem")
	if err != nil {
		t.Error(err)
	}
	publicKeys, err := LoadPublicKey("/tmp/public.pem")
	if err != nil {
		t.Error(err)
	}
	sign, err := SignRSA("hello", privateKeys)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("sign: ", sign)
	isOk := VerifyRSA(publicKeys, "hello", sign)
	if isOk == true {
		t.Log("verify success")
	} else {
		t.Error("verify fail")
	}

}
