package types

import "github.com/branila/spillatore/utils"

type Chat struct {
	Id int `json:"id"`
}

func (c Chat) String() string {
	return utils.PrettifyObject(c)
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

func (m Message) String() string {
	return utils.PrettifyObject(m)
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

func (u Update) String() string {
	return utils.PrettifyObject(u)
}
