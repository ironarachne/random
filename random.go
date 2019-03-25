package random

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
)

// Item returns a random string from an array of strings
func Item(items []string) string {
	return items[rand.Intn(len(items))]
}

// ItemFromThresholdMap returns a random weighted string
func ItemFromThresholdMap(items map[string]int) string {
	result := ""
	ceiling := 0
	start := 0
	var thresholds = make(map[string]int)

	for item, weight := range items {
		ceiling += weight
		thresholds[item] = start
		start += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, threshold := range thresholds {
		if threshold <= randomValue {
			result = item
		}
	}

	return result
}

// ItemInCollection checks to see if a string is in an array of strings
func ItemInCollection(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// SeedFromString uses a string to seed the random number generator
func SeedFromString(source string) {
	h := md5.New()
	io.WriteString(h, source)
	seed := binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))
}
