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
	Werewolf     card = card{"Werewolf", "ww"}
	Mystic       card = card{"Mystic Wolf", "ww"}
	Robber       card = card{"Robber", "v"}
	Tanner       card = card{"Tanner", "t"}
	Drunk        card = card{"Drunk", "v"}
	DG           card = card{"Doppleganger", "v"}
	Insomniac    card = card{"Insomniac", "v"}
	Seer         card = card{"Seer", "v"}
	APSeer       card = card{"Apprentice Seer", "v"}
	ParaInv      card = card{"Paranormal Investigator", "v"}
	Witch        card = card{"Witch", "v"}
	Troublemaker card = card{"Troublemaker", "v"}
	Minion       card = card{"Minion", "v"}
	Alpha        card = card{"Alpha Wolf", "ww"}
	Mason        card = card{"Mason", "v"}
)

// Different Decks depending on player count
func init() {
	Decks[3] = []card{Werewolf, Werewolf, Robber, Tanner, Drunk, DG}
	Decks[4] = []card{Werewolf, Werewolf, Robber, Tanner, Drunk, DG, Troublemaker}
	Decks[5] = []card{Werewolf, Seer, Insomniac, Mystic, APSeer, ParaInv, Witch, Troublemaker}
	Decks[6] = []card{Werewolf, Werewolf, Villager, Seer, Robber, Troublemaker, Tanner, Drunk, Minion}
	Decks[7] = []card{Werewolf, Werewolf, Villager, Seer, Robber, Troublemaker, Tanner, Minion, Mason, Mason}
	Decks[8] = []card{Werewolf, Werewolf, Villager, Villager, Seer, Robber, Troublemaker, Tanner, Minion, Mason, Mason}
	Decks[9] = []card{Mystic, Werewolf, Robber, Troublemaker, Tanner, Minion, DG, Mason, Mason, ParaInv, Seer, Drunk}
	Decks[10] = []card{Robber, Troublemaker, Tanner, Minion, DG, Alpha, Mystic, APSeer, ParaInv, Mason, Mason, Seer, Drunk}
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

func SwapCards(x, y card) (card, card) {
	return y, x
}

//Called at the start of the round to give players their cards TODO
func AssignRoles(players []player) {
	var deck []card
	deck = Decks[len(players)]
}
