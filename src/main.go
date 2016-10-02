package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	//	"html/template"
	//	"log"
	"math/rand"
	//	"net/http"
	//	"strings"
)

var Decks map[int][]card

var NoRole card
var Werewolf card
var Mystic card
var Robber card
var Tanner card
var Drunk card
var DG card
var Insomniac card
var Seer card
var APSeer card
var ParaInv card
var Witch card
var Troublemaker card
var Minion card
var Alpha card
var Mason card
var Villager card

// Different Decks depending on player count
func start() {
	NoRole = card{"NOROLE", "xxx"}
	Werewolf = card{"Werewolf", "ww"}
	Mystic = card{"Mystic Wolf", "ww"}
	Robber = card{"Robber", "v"}
	Tanner = card{"Tanner", "t"}
	Drunk = card{"Drunk", "v"}
	DG = card{"Doppleganger", "v"}
	Insomniac = card{"Insomniac", "v"}
	Seer = card{"Seer", "v"}
	APSeer = card{"Apprentice Seer", "v"}
	ParaInv = card{"Paranormal Investigator", "v"}
	Witch = card{"Witch", "v"}
	Troublemaker = card{"Troublemaker", "v"}
	Minion = card{"Minion", "v"}
	Alpha = card{"Alpha Wolf", "ww"}
	Mason = card{"Mason", "v"}
	Villager = card{"Villager", "v"}
	Decks = make(map[int][]card)
	Decks[3] = []card{Werewolf, Werewolf, Robber, Tanner, Drunk, DG}
	Decks[4] = []card{Werewolf, Werewolf, Robber, Tanner, Drunk, DG, Troublemaker}
	Decks[5] = []card{Werewolf, Seer, Insomniac, Mystic, APSeer, ParaInv, Witch, Troublemaker}
	Decks[6] = []card{Werewolf, Werewolf, Villager, Seer, Robber, Troublemaker, Tanner, Drunk, Minion}
	Decks[7] = []card{Werewolf, Werewolf, Villager, Seer, Robber, Troublemaker, Tanner, Minion, Mason, Mason}
	Decks[8] = []card{Werewolf, Werewolf, Villager, Villager, Seer, Robber, Troublemaker, Tanner, Minion, Mason, Mason}
	Decks[9] = []card{Mystic, Werewolf, Robber, Troublemaker, Tanner, Minion, DG, Mason, Mason, ParaInv, Seer, Drunk}
	Decks[10] = []card{Robber, Troublemaker, Tanner, Minion, DG, Alpha, Mystic, APSeer, ParaInv, Mason, Mason, Seer, Drunk}

}

/*func login(w http.ResponseWriter, r *http.Request) {
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
}*/
func main() {
	/*http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
	//user:pass@tcp/hello
	db, err := sql.open("mysql", "")

	start()
	players := []player{player{NoRole, "John"}, player{NoRole, "James"}, player{NoRole, "Jimmy"}}
	for j := 0; j < 10; j++ {
		middle := AssignRoles(players)
		for i := 0; i < len(players); i++ {
			fmt.Printf("%s has %s card\n", players[i].name, players[i].role)
		}
		fmt.Printf("%v is the deck\n", middle)
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

func (p *player) CheckResults() string {
	//killed player was a werewolf, village victory
	if p.role.team == "ww" {
		return "v"
		//killed player was a tanner, tanner victory
	} else if p.role.team == "t" {
		return "t"
		//killed player was something else, werewolf victory
	} else {
		return "w"
	}

}

func SwapCards(x, y card) (card, card) {
	return y, x
}

//Called at the start of the round to give players their cards TODO
func AssignRoles(players []player) []card {
	var deck []card
	var middle []card
	if len(players) < 3 {
		fmt.Printf("OI you  don't have enough players! %d", len(players))
		return middle
	}
	deck = Decks[len(players)]
	CardOrder := rand.Perm(len(deck))
	//Assigns each player a random card #THIS SHOULD MAKE NEW CARDS INSTEAD
	//OF SCRAMBLING THE PRESETS
	for i := 0; i < len(players); i++ {
		players[i].role = card{deck[CardOrder[i]].name, deck[CardOrder[i]].team}
	}
	//Puts the rest of the cards in the middle
	//TODO -- Fix this shit
	middle = deck[len(players):]
	return middle
}
