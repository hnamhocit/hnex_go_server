package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	defaultMemory      uint32 = 64 * 1024 // 64 MB
	defaultIterations  uint32 = 3         // 3 iterations
	defaultParallelism uint8  = 2         // 2 threads
	defaultSaltLength  uint32 = 16        // 16-byte salt
	defaultKeyLength   uint32 = 32        // 32-byte hash output
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, defaultSaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate random salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultIterations,
		defaultMemory,
		defaultParallelism,
		defaultKeyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultMemory,
		defaultIterations,
		defaultParallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("invalid argon2id hash format: expected 6 parts, got %d", len(parts))
	}

	if parts[1] != "argon2id" {
		return false, fmt.Errorf("unsupported argon2 variant: %s", parts[1])
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, fmt.Errorf("invalid version format in hash: %w", err)
	}
	if version != argon2.Version {
		log.Printf("Warning: Argon2 version mismatch. Hash version: %d, Current library version: %d", version, argon2.Version)
	}

	var memory, iterations uint32
	var parallelism uint8
	paramParts := strings.Split(parts[3], ",")
	if len(paramParts) != 3 {
		return false, fmt.Errorf("invalid parameter format in hash: expected 3 parts, got %d", len(paramParts))
	}
	if _, err := fmt.Sscanf(paramParts[0], "m=%d", &memory); err != nil {
		return false, fmt.Errorf("invalid memory parameter in hash: %w", err)
	}
	if _, err := fmt.Sscanf(paramParts[1], "t=%d", &iterations); err != nil {
		return false, fmt.Errorf("invalid iterations parameter in hash: %w", err)
	}
	if _, err := fmt.Sscanf(paramParts[2], "p=%d", &parallelism); err != nil {
		return false, fmt.Errorf("invalid parallelism parameter in hash: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt from hash: %w", err)
	}
	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash from hash string: %w", err)
	}

	computedHash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(storedHash)),
	)

	if subtle.ConstantTimeCompare(storedHash, computedHash) == 1 {
		return true, nil
	}
	return false, nil
}
