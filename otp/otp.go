/*
refer to: https://github.com/pquerna/otp

*/

package otp

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"image"
	"net/url"
	"strconv"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Digits int

const (
	DigitsSix   Digits = 6
	DigitsEight Digits = 8
)

type Algorithm int

const (
	AlgorithmSHA1 Algorithm = iota
	AlgorithmSHA256
	AlgorithmSHA512
	AlgorithmMD5
)

type Encoder string

const (
	EncoderDefault Encoder = ""
	EncoderSteam   Encoder = "steam"
)

// ErrValidateSecretInvalidBase32 Error when attempting to convert the secret from base32 to raw bytes.
var ErrValidateSecretInvalidBase32 = errors.New("decoding of secret as base32 failed")

// ErrValidateInputInvalidLength The user provided passcode length was not expected.
var ErrValidateInputInvalidLength = errors.New("input length unexpected")

// ErrGenerateMissingIssuer When generating a Key, the Issuer must be set.
var ErrGenerateMissingIssuer = errors.New("issuer must be set")

// ErrGenerateMissingAccountName When generating a Key, the Account Name must be set.
var ErrGenerateMissingAccountName = errors.New("AccountName must be set")

// NewKeyFromURL creates a new Key from an TOTP or HOTP url.
//
// The URL format is documented here:
//
// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func NewKeyFromURL(orig string) (*Key, error) {
	s := strings.TrimSpace(orig)

	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	return &Key{
		orig: s,
		url:  u,
	}, nil
}

func (k *Key) String() string {
	return k.orig
}

// Type return "hotp" or "totp"
func (k *Key) Type() string {
	return k.url.Host
}

func (k *Key) URL() string {
	return k.url.String()
}

// Image returns an QR-Code image of the specified width and height,
// suitable for use by many clients like Google-Authenricator
// to enroll a user's TOTP/HOTP key.
func (k *Key) Image(width int, height int) (image.Image, error) {
	b, err := qr.Encode(k.orig, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	b, err = barcode.Scale(b, width, height)

	if err != nil {
		return nil, err
	}

	return b, nil

	//var png []byte
	//png, err := qrcode.Encode(k.orig, qrcode.Medium, 256)
	//if err != nil {
	//	return nil, err
	//}

}

// Issuer returns the name of the issuing organization.
func (k *Key) Issuer() string {
	q := k.url.Query()

	issuer := q.Get("issuer")

	if issuer != "" {
		return issuer
	}

	p := strings.TrimPrefix(k.url.Path, "/")
	i := strings.Index(p, ":")

	if i == -1 {
		return ""
	}

	return p[:i]
}

func (k *Key) AccountName() string {
	p := strings.TrimPrefix(k.url.Path, "/")
	i := strings.Index(p, ":")

	if i == -1 {
		return p
	}

	return p[i+1:]
}

func (k *Key) Secret() string {
	q := k.url.Query()
	return q.Get("secret")
}

func (k *Key) Period() uint64 {
	q := k.url.Query()

	if u, err := strconv.ParseUint(q.Get("period"), 10, 64); err == nil {
		return u
	}

	return 30
}

func (k *Key) Digits() Digits {
	q := k.url.Query()

	if u, err := strconv.ParseUint(q.Get("digits"), 10, 64); err == nil {
		return Digits(u)
	}

	return DigitsSix
}

func (k *Key) Algorithm() Algorithm {
	q := k.url.Query()

	a := strings.ToLower(q.Get("algorithm"))
	switch a {
	case "md5":
		return AlgorithmMD5
	case "sha256":
		return AlgorithmSHA256
	case "sha512":
		return AlgorithmSHA512
	default:
		return AlgorithmSHA1
	}
}

func (k *Key) Encoder() Encoder {
	q := k.url.Query()

	a := strings.ToLower(q.Get("encoder"))
	switch a {
	case "steam":
		return EncoderSteam
	default:
		return EncoderDefault
	}
}

func (a Algorithm) String() string {
	switch a {
	case AlgorithmSHA1:
		return "SHA1"
	case AlgorithmSHA256:
		return "SHA256"
	case AlgorithmSHA512:
		return "SHA512"
	case AlgorithmMD5:
		return "MD5"
	}
	panic("unreached")
}

func (a Algorithm) Hash() hash.Hash {
	switch a {
	case AlgorithmSHA1:
		return sha1.New()
	case AlgorithmSHA256:
		return sha256.New()
	case AlgorithmSHA512:
		return sha512.New()
	case AlgorithmMD5:
		return md5.New()
	}
	panic("unreached")
}

func (d Digits) String() string {
	return fmt.Sprintf("%d", d)
}

func (d Digits) Format(in int32) string {
	f := fmt.Sprintf("%%0%dd", d)
	return fmt.Sprintf(f, in)
}

func (d Digits) Length() int {
	return int(d)
}
