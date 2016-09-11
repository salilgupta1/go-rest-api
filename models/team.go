package models

type Team struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    League      string      `json:"league"`
    Players     []*Player   `json:"players"`
}
