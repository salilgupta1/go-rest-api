package apis

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/salilgupta1/go-rest-api/models"
    "github.com/gorilla/mux"
    "gopkg.in/pg.v4"
    "gopkg.in/pg.v4/orm"

)

type TeamsApi struct{
    DB *pg.DB
}

func(api *TeamsApi) GetTeams(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")

    var teams []models.Team
    db_err := api.DB.Model(&teams).Select()

    if db_err != nil {
        http.Error(res, db_err.Error(), http.StatusInternalServerError)
    }

    json, json_err := json.Marshal(teams)

    if json_err != nil {
        http.Error(res, json_err.Error(), http.StatusInternalServerError)
    }

    fmt.Fprint(res, string(json))
}

func(api *TeamsApi) GetTeam(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")

    team_id := mux.Vars(req)["id"]

    var team models.Team

    db_err := api.DB.Model(&team).Column("team.*", "Players").
                    Relation("Players", func(q *orm.Query) *orm.Query {
                        return q.Where("team_id = ?", team_id)
                    }).
                    Where("id = ?", team_id).
                    Select()

    if db_err != nil {
        res.WriteHeader(http.StatusNotFound)
        fmt.Fprint(res, string("Team not found"))
    }

    json, json_err := json.Marshal(team)

    if json_err != nil {
        http.Error(res, json_err.Error(), http.StatusInternalServerError)
    }

    fmt.Fprint(res, string(json))
}


func(api *TeamsApi) CreateTeam(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")

    name := req.FormValue("name")
    league := req.FormValue("league")

    team := models.Team{ Name: name, League: league }

    db_err := api.DB.Create(&team)

    if db_err != nil {
        http.Error(res, db_err.Error(), http.StatusInternalServerError)
    } else {
        res.WriteHeader(http.StatusOK)
    }
}
