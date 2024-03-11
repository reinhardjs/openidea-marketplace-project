package hashing

import (
	"bytes"
	"errors"

	"golang.org/x/crypto/argon2"
)

type Hashing interface {
	GenerateHash(password []byte) (*hashSalt, error)
	Compare(hash, salt, password []byte) error
}

// HashSalt struct used to store
// generated hash and salt used to
// generate the hash.
type hashSalt struct {
	Hash, Salt []byte
}

type argon2idHash struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	salt   []byte
}

// NewArgon2idHash constructor function for
// Argon2idHash.
func NewArgon2idHash(time uint32, salt []byte, memory uint32, threads uint8, keyLen uint32) *argon2idHash {
	return &argon2idHash{
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
		salt:    salt,
	}
}

// GenerateHash using the password and provided salt.
// If not salt value provided fallback to random value
// generated of a given length.
func (a *argon2idHash) GenerateHash(password []byte) (*hashSalt, error) {
	// Generate hash
	hash := argon2.IDKey(password, a.salt, a.time, a.memory, a.threads, a.keyLen)
	// Return the generated hash and salt used for storage.
	return &hashSalt{Hash: hash, Salt: a.salt}, nil
}

// Compare generated hash with store hash.
func (a *argon2idHash) Compare(hash, salt, password []byte) error {
	// Generate hash for comparison.
	hashSalt, err := a.GenerateHash(password)
	if err != nil {
		return err
	}
	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if !bytes.Equal(hash, hashSalt.Hash) {
		return errors.New("hash doesn't match")
	}
	return nil
}
