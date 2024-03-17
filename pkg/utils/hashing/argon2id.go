package hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

/**
 * From: https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
**/

type Hashing interface {
	GenerateHashFromPassword(password string) (string, error)
	Compare(hash, password string) (bool, error)
}

type argon2idHash struct {
	// time represents the number of
	// passed over the specified memory (iterations).
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen  uint32
	saltLen uint32
}

// NewArgon2idHash constructor function for
// Argon2idHash.
func NewArgon2idHash(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) *argon2idHash {
	return &argon2idHash{
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
		saltLen: saltLen,
	}
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (a *argon2idHash) GenerateHashFromPassword(password string) (encodedHash string, err error) {
	salt, err := randomSecret(a.saltLen)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, a.time, a.memory, a.threads, a.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.time, a.threads, b64Salt, b64Hash)

	return encodedHash, nil
}

// Compare generated hash with store hash.
func (a *argon2idHash) Compare(encodedHash, password string) (match bool, err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return
	}

	// Generate hash for comparison.
	generatedHash := argon2.IDKey([]byte(password), salt, p.time, p.memory, p.threads, p.keyLen)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, generatedHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *argon2idHash, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("ErrInvalidHash")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("ErrIncompatibleVersion")
	}

	p = &argon2idHash{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.time, &p.threads)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLen = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLen = uint32(len(hash))

	return p, salt, hash, nil
}
