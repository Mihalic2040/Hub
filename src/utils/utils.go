package utils

import (
	"bytes"
	"crypto/sha256"
	"reflect"
	"runtime"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
)

// GetFunctionName returns the name of a function
func GetFunctionName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	dotIndex := len(fullName) - 1
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			dotIndex = i
			break
		}
	}
	return fullName[dotIndex+1:]
}

func SplitKey(key string) (string, string, string) {
	if len(key) == 0 || key[0] != '/' {
		return "", "", "Shit"
	}

	key = key[1:]

	i := strings.IndexByte(key, '/')
	if i <= 0 {
		return "", "", "Shit"
	}

	return key[:i], key[i+1:], ""
}

func GeneratePrivateKeyFromString(input string) (crypto.PrivKey, error) {
	// Calculate the SHA256 hash of the input string
	hash := sha256.Sum256([]byte(input))

	// Create a libp2p private key from the hash
	privKey, _, err := crypto.GenerateEd25519Key(bytes.NewReader(hash[:]))
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
