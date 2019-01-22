package ex2_results

import "encoding/json"
import "sort"

// Jsonable is an interface with one method, tojson
type JSONable interface {
	ToJSON() string
}

// Movie represents a movie
type Movie struct {
	Title  string
	Year   int
	Actors []string
}

func (ms Movies) Len() int           { return len(ms) }
func (ms Movies) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }
func (ms Movies) Less(i, j int) bool { return ms[i].Title < ms[j].Title }

func (m MoviesByYear) Len() int           { return len(m) }
func (m MoviesByYear) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MoviesByYear) Less(i, j int) bool { return m[i].Year < m[j].Year }

//MoviesByYear is a type used only for sorting
type MoviesByYear Movies

// Movies is a list of movie
type Movies []Movie

//AddMovie add a new movie in list
func (ms *Movies) AddMovie(movie Movie) {
	*ms = append(*ms, movie)
}

//SortByName sort movies by name
func (ms Movies) SortByName() {
	sort.Sort(ms)
}

//SortByYear sort movies by year
func (ms Movies) SortByYear() {
	sort.Sort(MoviesByYear(ms))
}

//ToJSON return a json representation
func (ms Movies) ToJSON() string {
	if data, err := json.Marshal(ms); err == nil {
		return string(data)
	}
	return "{}"
}

//NewMovie return a new movie
func NewMovie(title string, year int) Movie {
	return Movie{Title: title, Year: year, Actors: make([]string, 0)}
}

// GetTitle return the title of movie
func (movie Movie) GetTitle() string {
	return movie.Title
}

//ToJSON convert movie to json
func (movie Movie) ToJSON() string {
	if data, err := json.Marshal(movie); err == nil {
		return string(data)
	}
	return ""
}

// AddActor add a new actor to movie
func (movie *Movie) AddActor(actor string) {
	(*movie).Actors = append(movie.Actors, actor)
}

// GetNbActors return the number of actors of a movie
func (movie Movie) GetNbActors() int {
	return len(movie.Actors)
}
