package hotp

import (
	"io"

	"github.com/chenjl-ops/go-lib/otp"
)

// ValidateOpts provides options for ValidateCustom().
type ValidateOpts struct {
	// Digits as part of the input. Defaults to 6.
	Digits otp.Digits
	// Algorithm to use for HMAC. Defaults to SHA1.
	Algorithm otp.Algorithm
	// Encoder to use for output code.
	Encoder otp.Encoder
}

// GenerateOpts provides options for .Generate()
type GenerateOpts struct {
	// Name of the issuing Organization/Company.
	Issuer string
	// Name of the User's Account (eg, email address)
	AccountName string
	// Size in size of the generated Secret. Defaults to 10 bytes.
	SecretSize uint
	// Secret to store. Defaults to a randomly generated secret of SecretSize.  You should generally leave this empty.
	Secret []byte
	// Digits to request. Defaults to 6.
	Digits otp.Digits
	// Algorithm to use for HMAC. Defaults to SHA1.
	Algorithm otp.Algorithm
	// Reader to use for generating HOTP Key.
	Rand io.Reader
}
