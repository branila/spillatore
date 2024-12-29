package types

import "github.com/branila/spillatore/utils"

type UserId string

type UserStats map[UserId]int

type Database struct {
	Counter int         `json:"counter"`
	Stats   []UserStats `json:"stats"`
}

func (d Database) String() string {
	return utils.PrettifyObject(d)
}
