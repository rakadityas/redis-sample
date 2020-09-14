package main

type Prizes struct {
	Prize []Prize `json:"prizes"`
}

type Prize struct {
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}
