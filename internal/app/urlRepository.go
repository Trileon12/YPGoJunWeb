package app

import (
	"errors"
)

var URLs map[string]string
var URLsRevers map[string]string

func InitializeUrlRepository() {
	URLs = map[string]string{}
}

const shortUrlLen = 6

func GetUrlShort(originalURL string) string {

	if shortUrl, isExists := URLsRevers[originalURL]; isExists {
		return shortUrl
	}

	found := false
	shortUrl := RandString(shortUrlLen)

	for _, found = URLs[shortUrl]; !found; {
		shortUrl = RandString(shortUrlLen)
	}

	URLs[shortUrl] = originalURL

	return shortUrl
}

func GetOriginalUrl(shortURL string) (string, error) {

	if originalUrl, isExists := URLs[shortURL]; isExists {
		return originalUrl, nil
	} else {
		return "", errors.New("URL не найден")
	}
}
