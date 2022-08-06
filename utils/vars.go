package utils

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
