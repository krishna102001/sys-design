package main

import (
	"log"
	"strings"
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
		URL_ID /= base
	}

	for low, high := 0, len(chars)-1; low < high; low, high = low+1, high-1 {
		chars[low], chars[high] = chars[high], chars[low]
	}

	return string(chars)
}

func getOriginalURL(shortURL string) uint64 {
	var URL_ID uint64
	for i := 0; i < len(shortURL); i++ {
		pos := strings.IndexByte(charset, shortURL[i])
		if pos == -1 {
			log.Printf("invalid character found")
			continue
		}
		URL_ID = URL_ID*base + uint64(pos)
	}
	return URL_ID
}

func main() {
	var URL_ID uint64 = 29843724962314
	shortURL := convertURL(URL_ID)
	log.Printf("base 62 short url are %s", shortURL)
	originalURL_ID := getOriginalURL(shortURL)
	log.Printf("original url %d", originalURL_ID)
}
