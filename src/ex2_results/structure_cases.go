package ex2_results

import "encoding/json"
import "sort"

type Movie struct {
	Title  string
	Year   int
	Actors []string
}

func (m Movies) Len() int           { return len(m) }
func (m Movies) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Movies) Less(i, j int) bool { return m[i].Title < m[j].Title }

func (m MoviesByYear) Len() int           { return len(m) }
func (m MoviesByYear) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MoviesByYear) Less(i, j int) bool { return m[i].Year < m[j].Year }

type MoviesByYear Movies
type Movies []Movie

//AddMovie
func (ms *Movies) AddMovie(movie Movie) {
	*ms = append(*ms, movie)
}

func (ms Movies) SortByName() {
	sort.Sort(ms)
}

func (ms Movies) SortByYear() {
	sort.Sort(MoviesByYear(ms))
}

//NewMovie return a new movie
func NewMovie(title string, year int) Movie {
	return Movie{Title: title, Year: year, Actors: make([]string, 0)}
}

// GetTitle return the title of movie
func (movie Movie) GetTitle() string {
	return movie.Title
}

//ToJson convert movie to json
func (movie Movie) ToJson() string {
	if data, err := json.Marshal(movie); err == nil {
		return string(data)
	}
	return ""
}

// AddActor add a new actor to movie
func (movie *Movie) AddActor(actor string) {
	(*movie).Actors = append(movie.Actors, actor)
}

func (movie Movie) GetNbActors() int {
	return len(movie.Actors)
}