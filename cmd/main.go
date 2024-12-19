package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("GET /", templ.Handler(home()))

	assets := http.FileServer(http.Dir("assets/"))
	http.Handle("GET /assets/", http.StripPrefix("/assets/", assets))

	hand := new_hand()
	shoe := new_shoe(6)
	http.HandleFunc("GET /basic_play", func(w http.ResponseWriter, r *http.Request) {
		hand.deal(&shoe)
		basic_play(hand, shoe).Render(r.Context(), w)
	})
	http.HandleFunc("POST /basic_play/hit", func(w http.ResponseWriter, r *http.Request) {
		hand.hit(&shoe)
		show_hand(hand).Render(r.Context(), w)
	})

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
