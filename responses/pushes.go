package responses

import (
	"github.com/mitsuse/pushbullet-go/requests"
)

type Push struct {
	Active                  bool    `json:"active"`
	Body                    string  `json:"body"`
	Created                 float64 `json:"created"`
	Direction               string  `json:"direction"`
	Dismissed               bool    `json:"dismissed"`
	Iden                    string  `json:"iden"`
	Modified                float64 `json:"modified"`
	RecieverEmail           string  `json:"reciever_email"`
	RecieverEmailNormalized string  `json:"reciever_email_normalized"`
	RecieverIden            string  `json:"reciever_iden"`
	SenderName              string  `json:"sender_name"`
	Title                   string  `json:"title"`
	Type                    string  `json:"type"`
}

type Note struct {
	*Push
	*requests.Note
}

type Link struct {
	*Push
	*requests.Link
}

type Address struct {
	*Push
	*requests.Address
}

type Checklist struct {
	*Push
	*requests.Checklist
}

type File struct {
	*Push
	*requests.File
}
