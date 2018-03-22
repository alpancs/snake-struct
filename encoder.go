package snaky

import (
	"encoding/json"
	"regexp"
	"strings"
)

var patternPreKey = regexp.MustCompile(`"\w*":`)
var patternCamel = regexp.MustCompile(`[^A-Z][A-Z]+`)

func Marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	return []byte(patternPreKey.ReplaceAllStringFunc(string(data), preKeySnakeCaser)), err
}

func preKeySnakeCaser(preKey string) string {
	l, r := 1, len(preKey)-2
	return preKey[:l] + keySnakeCaser(preKey[l:r]) + preKey[r:]
}

func keySnakeCaser(key string) string {
	return patternCamel.ReplaceAllStringFunc(" "+key, snakeCaser)[2:]
}

func snakeCaser(s string) string {
	return s[:1] + "_" + strings.ToLower(s[1:])
}
