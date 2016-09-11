package apis

import (
    "net/http"
    "encoding/json"
    "fmt"
    "strconv"

    "github.com/salilgupta1/go-rest-api/models"
    "github.com/gorilla/mux"
    "gopkg.in/pg.v4"
)

type PlayersApi struct{
    DB *pg.DB
}

func(api *PlayersApi) GetPlayers(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")

    var players []models.Player
    db_err := api.DB.Model(&players).Select()

    if db_err != nil {
        http.Error(res, db_err.Error(), http.StatusInternalServerError)
    }

    json, json_err := json.Marshal(players)

    if json_err != nil {
        http.Error(res, json_err.Error(), http.StatusInternalServerError)
    }

    fmt.Fprint(res, string(json))
}

func(api *PlayersApi) GetPlayer(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")

    player_id := mux.Vars(req)["id"]

    var player models.Player

    db_err := api.DB.Model(&player).
                Column("player.*", "Team").
                Where("id = ?", player_id).
                Select()

    if db_err != nil {
        res.WriteHeader(http.StatusNotFound)
        fmt.Fprint(res, string("Player not found"))
    }

    json, json_err := json.Marshal(player)

    if json_err != nil {
        http.Error(res, json_err.Error(), http.StatusInternalServerError)
    }

    fmt.Fprint(res, string(json))
}


func(api *PlayersApi) CreatePlayer(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")

    name := req.FormValue("name")
    position := req.FormValue("position")
    nationality := req.FormValue("nationality")
    age, _ := strconv.Atoi(req.FormValue("age"))
    team_id, _ := strconv.Atoi(req.FormValue("teamID"))

    player := models.Player{
        Name: name,
        Position: position,
        Nationality: nationality,
        Age: age,
        TeamID: team_id,
    }

    db_err := api.DB.Create(&player)

    if db_err != nil {
        http.Error(res, db_err.Error(), http.StatusInternalServerError)
    } else {
        res.WriteHeader(http.StatusOK)
    }
}
