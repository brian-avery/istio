package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingAuthorizationSchema(t *testing.T) {
	_, err := getBasicToken("foo")
	if err == nil {
		t.Fatalf("Expected auth token to indicate missing authorization schema but it did not.")
	}
}

func TestTokenNotBasic(t *testing.T) {
	if _, err := getBasicToken("bar foo"); err == nil {
		t.Fatalf("Expected auth token to indicate that it was of an invalid type but it did not.")
	}
}

func TestValidTokenSchema(t *testing.T) {
	if _, err := getBasicToken("Basic foo"); err != nil {
		t.Fatalf("Expected auth token to be marked valid but it was not.")
	}
}

func TestInvalidBase64Encoding(t *testing.T) {
	if _, _, err := getTokenSegments("foo"); err == nil {
		t.Fatalf("Expected error concerning invalid base64 encoding but it was absent.")
	}
}

func TestBase64TooManySegments(t *testing.T) {
	if _, _, err := getTokenSegments("dXNlcjpwYXNzd29yZDpmb29iYXI="); err == nil {
		t.Fatalf("Expected error concerning too many segments but it was missing.")
	}
}

func TestValidBase64Encoding(t *testing.T) {
	user, password, err := getTokenSegments("dXNlcjpwYXNzd29yZA==")
	if err != nil {
		t.Fatalf("Expected token to be acceptable but got an error.")
	}

	assert.Equal(t, user, "user")
	assert.Equal(t, password, "password")
}
