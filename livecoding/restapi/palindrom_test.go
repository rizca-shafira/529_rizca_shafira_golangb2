package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrom(t *testing.T) {
	result := Palindrom("kasur ini rusak")

	assert.Equal(t, "kasur ini rusak", result, "Not A Palindrom")

	fmt.Println("TestPalindrom Eksekusi terhenti")
}
