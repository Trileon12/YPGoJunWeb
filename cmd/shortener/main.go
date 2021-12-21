package main

//TODO: Ментор: почему такое странное подключение internal пакета? почему не работает подлючение черех github
import (
	"../../internal/app"
	"net/http"
)

func main() {

	app.InitializeUrlRepository()

	http.HandleFunc("/", app.URLHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println("Fatal error ", err)
	}

}
