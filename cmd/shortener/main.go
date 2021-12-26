package main

//TODO: Ментор: почему такое странное подключение internal пакета? почему не работает подлючение черех github
import (
	"../../internal/app"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	r.Post("/", app.GetShortUrl)
	r.Get("/{ID}", app.GetFullURLByFullUrl)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		println("Fatal error ", err)
	}

}
