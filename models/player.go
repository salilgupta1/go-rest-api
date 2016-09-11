package models

type Player struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    Position    string      `json:"position"`
    Nationality string      `json:"nationality"`
    Age         int         `json:"age"`
    TeamID      int         `json:"team_id"`
    Team        *Team       `json:"team"`
}