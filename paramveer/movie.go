package main

type Movie struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Director  string  `json:"director"`
	IMDb      float64 `json:"imdb"`
	Hollywood string  `json:"hollywood"`
	Bollywood string  `json:"bollywood"`
}
