package totp

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/chenjl-ops/go-lib/otp"
	"github.com/chenjl-ops/go-lib/otp/hotp"
	"github.com/chenjl-ops/go-lib/otp/internal"
	"math"
	"net/url"
	"strconv"
	"time"
)

var b32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)

// Validate ...
func Validate(passcode string, secret string) bool {
	rv, _ := ValidateCustom(passcode, secret, time.Now().UTC(), ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	return rv
}

func ValidateCustom(passcode string, secret string, t time.Time, opts ValidateOpts) (bool, error) {
	if opts.Period == 0 {
		opts.Period = 30
	}

	counters := make([]uint64, 0)
	counter := int64(math.Floor(float64(t.Unix()) / float64(opts.Period)))

	counters = append(counters, uint64(counter))
	for i := 1; i <= int(opts.Skew); i++ {
		counters = append(counters, uint64(counter+int64(i)))
		counters = append(counters, uint64(counter-int64(i)))
	}

	for _, counter := range counters {
		rv, err := hotp.ValidateCustom(passcode, counter, secret, hotp.ValidateOpts{
			Digits:    opts.Digits,
			Algorithm: opts.Algorithm,
			Encoder:   opts.Encoder,
		})
		if err != nil {
			return false, err
		}

		if true == rv {
			return true, nil
		}
	}
	return false, nil
}

// GenerateCode ...
func GenerateCode(secret string, t time.Time) (string, error) {
	return GenerateCodeCustom(secret, t, ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}

// GenerateCodeCustom ...
func GenerateCodeCustom(secret string, t time.Time, opts ValidateOpts) (passcode string, err error) {
	if opts.Period == 0 {
		opts.Period = 30
	}
	counter := uint64(math.Floor(float64(t.Unix()) / float64(opts.Period)))
	passcode, err = hotp.GenerateCodeCustom(secret, counter, hotp.ValidateOpts{
		Digits:    opts.Digits,
		Algorithm: opts.Algorithm,
		Encoder:   opts.Encoder,
	})

	if err != nil {
		return "", err
	}
	return passcode, nil
}

func Generate(opts GenerateOpts) (*otp.Key, error) {
	if opts.Issuer == "" {
		return nil, otp.ErrGenerateMissingIssuer
	}

	if opts.AccountName == "" {
		return nil, otp.ErrGenerateMissingAccountName
	}

	if opts.Period == 0 {
		opts.Period = 30
	}

	if opts.SecretSize == 0 {
		opts.SecretSize = 20
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
	v.Set("period", strconv.FormatUint(uint64(opts.Period), 10))
	v.Set("algorithm", opts.Algorithm.String())
	v.Set("digits", opts.Digits.String())

	u := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + opts.Issuer + ":" + opts.AccountName,
		RawQuery: internal.EncodeQuery(v),
	}

	return otp.NewKeyFromURL(u.String())
}
