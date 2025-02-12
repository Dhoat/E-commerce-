package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

// GenerateShortURL creates a unique short URL hash
func GenerateShortURL(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:6] // Take first 6 characters
	return shortURL
}
