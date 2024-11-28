package hotp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"github.com/chenjl-ops/go-lib/otp"
	"github.com/chenjl-ops/go-lib/otp/internal"
	"math"
	"net/url"
	"strings"
)

const debug = false

var b32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)

func Validate(passcode string, counter uint64, secret string) bool {
	rv, _ := ValidateCustom(passcode, counter, secret, ValidateOpts{
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	return rv
}

func ValidateCustom(passcode string, counter uint64, secret string, opts ValidateOpts) (bool, error) {
	passcode = strings.TrimSpace(passcode)

	if len(passcode) != opts.Digits.Length() {
		return false, otp.ErrValidateInputInvalidLength
	}

	otpstr, err := GenerateCodeCustom(secret, counter, opts)
	if err != nil {
		return false, err
	}

	if subtle.ConstantTimeCompare([]byte(otpstr), []byte(passcode)) == 1 {
		return true, nil
	}
	return false, nil
}

func GenerateCode(secret string, counter uint64) (string, error) {
	return GenerateCodeCustom(secret, counter, ValidateOpts{
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}

func GenerateCodeCustom(secret string, counter uint64, opts ValidateOpts) (passcode string, err error) {
	if opts.Digits == 0 {
		opts.Digits = otp.DigitsSix
	}

	secret = strings.TrimSpace(secret)
	if n := len(secret) % 8; n != 0 {
		secret = secret + strings.Repeat("=", 8-n)
	}

	secret = strings.ToUpper(secret)

	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", otp.ErrValidateSecretInvalidBase32
	}

	buf := make([]byte, 8)
	mac := hmac.New(opts.Algorithm.Hash, secretBytes)
	binary.BigEndian.PutUint64(buf, counter)
	if debug {
		fmt.Printf("counter=%v\n", counter)
		fmt.Printf("buf=%v\n", buf)
	}

	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	l := opts.Digits.Length()
	switch opts.Encoder {
	case otp.EncoderDefault:
		mod := int32(value % int64(math.Pow10(l)))

		if debug {
			fmt.Printf("offset=%v\n", offset)
			fmt.Printf("mod=%v\n", value)
			fmt.Printf("mod=%v\n", mod)
		}
		passcode = opts.Digits.Format(mod)
	case otp.EncoderSteam:
		alphabet := []byte{
			'2', '3', '4', '5', '6', '7', '8', '9',
			'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K',
			'M', 'N', 'P', 'Q', 'R', 'T', 'V', 'W',
			'X', 'Y',
		}
		radix := int64(len(alphabet))

		for i := 0; i < l; i++ {
			digit := value % radix
			value /= radix
			c := alphabet[digit]
			passcode += string(c)
		}
	}
	return
}

func Generate(opts GenerateOpts) (*otp.Key, error) {
	if opts.Issuer == "" {
		return nil, otp.ErrGenerateMissingIssuer
	}

	if opts.AccountName == "" {
		return nil, otp.ErrGenerateMissingAccountName
	}

	if opts.SecretSize == 0 {
		opts.SecretSize = 10
	}

	if opts.Digits == 0 {
		opts.Digits = otp.DigitsSix
	}

	if opts.Rand == nil {
		opts.Rand = rand.Reader
	}

	v := url.Values{}
	if len(opts.Secret) != 0 {
		v.Set("secret", b32NoPadding.EncodeToString(opts.Secret))
	} else {
		secret := make([]byte, opts.SecretSize)
		_, err := opts.Rand.Read(secret)
		if err != nil {
			return nil, err
		}
		v.Set("secret", b32NoPadding.EncodeToString(secret))
	}

	v.Set("issuer", opts.Issuer)
	v.Set("algorithm", opts.Algorithm.String())
	v.Set("digits", opts.Digits.String())

	u := url.URL{
		Scheme:   "otpauth",
		Host:     "hotp",
		Path:     "/" + opts.Issuer + ":" + opts.AccountName,
		RawQuery: internal.EncodeQuery(v),
	}

	return otp.NewKeyFromURL(u.String())
}
