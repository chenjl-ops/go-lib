package nanoid

import (
	"crypto/rand"
	"errors"
	"math"
)

var defaultAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	defaultSize = 21
)

func getMask(alphabetSize int) int {
	for i := 1; i <= 8; i++ {
		mask := (2 << uint(i)) - 1
		if mask > alphabetSize-1 {
			return mask
		}
	}
	return 0
}

func Generate(alphabet string, size int) (string, error) {
	chars := []rune(alphabet)

	if len(alphabet) == 0 || len(alphabet) > 255 {
		return "", errors.New("invalid alphabet, must not be empty and no more than 255 characters")
	}
	if size <= 0 {
		return "", errors.New("size must be greater than zero")
	}

	mask := getMask(len(chars))

	ceiArg := 1.6 * float64(mask*size) / float64(len(alphabet))
	step := int(math.Ceil(ceiArg))

	id := make([]rune, size)
	bytes := make([]byte, step)
	for x := 0; ; {
		_, err := rand.Read(bytes)
		if err != nil {
			return "", err
		}
		for y := 0; y < step; y++ {
			currByte := bytes[y] & byte(mask)
			if currByte < byte(len(chars)) {
				id[x] = chars[currByte]
				x++
				if x == size {
					return string(id[:size]), nil
				}
			}
		}
	}
}

func MustGenerate(alphabet string, size int) string {
	id, err := Generate(alphabet, size)
	if err != nil {
		panic(err)
	}
	return id
}

func New(l ...int) (string, error) {
	var size int
	switch {
	case len(l) == 0:
		size = defaultSize
	case len(l) == 1:
		size = l[0]
		if size < 0 {
			return "", errors.New("negative id length")
		}
	default:
		return "", errors.New("unexpected parameter")
	}
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = defaultAlphabet[bytes[i]&63]
	}
	return string(id[:size]), nil
}

func MustNew(l ...int) string {
	id, err := New(l...)
	if err != nil {
		panic(err)
	}
	return id
}
