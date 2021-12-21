package app

import (
	"io"
	"net/http"
	"path"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// этот обработчик принимает только запросы, отправленные методом GETif r.Method != http.MethodGet {
	http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
}

func URLHandler(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodPost:
		GetShortUrl(writer, request)
	case http.MethodGet:
		getFullURL(writer, request)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetShortUrl(writer http.ResponseWriter, request *http.Request) {

	b, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	link := string(b)
	shortLink := GetUrlShort(link)

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(shortLink))
}

func getFullURL(writer http.ResponseWriter, request *http.Request) {

	id := path.Base(request.URL.Path)
	URL, err := GetOriginalUrl(id)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	writer.Header().Set("Location", URL)
	writer.WriteHeader(http.StatusMovedPermanently)

}
