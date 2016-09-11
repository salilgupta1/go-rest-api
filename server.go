package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/salilgupta1/go-rest-api/apis"
    "gopkg.in/pg.v4"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)

    options := pg.Options{
        User: "salilgupta",
        Database: "go_api_development",
        Password: "",
        Addr:"localhost:5432",
    }

    // Note: db_connection is a pointer
    db_connection := pg.Connect(&options)
    defer db_connection.Close()

    teams_api := apis.TeamsApi{db_connection}
    players_api := apis.PlayersApi{db_connection}

    router.HandleFunc("/teams", teams_api.GetTeams).Methods("GET")
    router.HandleFunc("/team/{id}", teams_api.GetTeam).Methods("GET")
    router.HandleFunc("/team", teams_api.CreateTeam).Methods("POST")

    router.HandleFunc("/players", players_api.GetPlayers).Methods("GET")
    router.HandleFunc("/player/{id}", players_api.GetPlayer).Methods("GET")
    router.HandleFunc("/player", players_api.CreatePlayer).Methods("POST")

    http.ListenAndServe(":5000", router)
}


// Require/optional arguments
// Handling Missed