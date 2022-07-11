package model

type Book struct{
	ID       uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Author     string    `json:"author,omitempty"`
	Quantity    uint64    `json:"quantity,omitempty"`	
}