package main

import (
	"log"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = uint64(len(charset))

// url_id is snowflake id which is generate when user is saving in the database
func convertURL(URL_ID uint64) string {
	if URL_ID == 0 {
		return string(charset[0])
	}
	var chars []byte
	for URL_ID > 0 {
		rem := URL_ID % base
		chars = append(chars, charset[rem])
		URL_ID /= 10
	}

	for low, high := 0, len(chars)-1; low < high; low, high = low+1, high-1 {
		chars[low], chars[high] = chars[high], chars[low]
	}

	return string(chars)
}

func main() {
	var URL uint64 = 29843724962314
	shortURL := convertURL(URL)
	log.Printf("base 62 short url are %s", shortURL)
}
