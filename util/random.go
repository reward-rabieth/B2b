package util

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func GenerateUUID() string {
	return uuid.New().String()
}
