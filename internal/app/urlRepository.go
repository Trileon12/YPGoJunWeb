package app

import (
	"errors"
)

var URLs map[string]string
var URLsRevers map[string]string
var HostShortURLs string

func InitializeUrlRepository() {
	URLs = map[string]string{}
	HostShortURLs = "http://localhost:8080/"
}

const shortUrlLen = 6

func GetUrlShort(originalURL string) string {

	if shortUrl, isExists := URLsRevers[originalURL]; isExists {
		return getHostShortUrl(shortUrl)
	}

	shortUrl := getUnicUrl()
	URLs[shortUrl] = originalURL

	return getHostShortUrl(shortUrl)
}

func GetOriginalUrl(shortURL string) (string, error) {

	if originalUrl, isExists := URLs[shortURL]; isExists {
		return originalUrl, nil
	} else {
		return "", errors.New("URL не найден")
	}
}

func getUnicUrl() string {

	found := false
	shortUrl := RandString(shortUrlLen)

	for _, found = URLs[shortUrl]; found; {
		shortUrl = RandString(shortUrlLen)
	}
	return shortUrl
}

func getHostShortUrl(shortUrl string) string {
	return HostShortURLs + shortUrl
}
