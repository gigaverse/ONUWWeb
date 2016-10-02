package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var Decks map[int][]card

const (
	Werewolf card = card{"Werewolf", "ww"}
)

func init() {
	Decks[3] = []card{Werewolf}
	Decks[4] = []card{Werewolf}

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Objects used during Werewolf Game
type card struct {
	name string
	team string
}

//TODO Add info from server
type player struct {
	role card
	name string
}

func (p *player) CheckRole() string {
	//killed player was a werewolf, village victory
	if p.card.team == "ww" {
		return "w"
		//killed player was a tanner, tanner victory
	} else if p.card.team == "t" {
		return "t"
		//killed player was something else, werewolf victory
	} else {
		return "v"
	}

}

func AssignRoles(players []player) {
	roles := []string{""}
}
