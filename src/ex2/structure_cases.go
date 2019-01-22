package ex2

//Movie TODO : create structure Movie with public attributes title, year and list of actors (string list)
type Movie struct{}

//Movies is a list of movies
type Movies []Movie

//AddMovie add a new movie in list
func (ms *Movies) AddMovie(movie Movie) {}

//SortByYear sort movies by year
func (ms Movies) SortByYear() {}

//SortByName sort movies by name
func (ms Movies) SortByName() {}

//NewMovie create a new Movie
func NewMovie(title string, year int) Movie {
	return Movie{}
}

// GetTitle return the title of movie
func (movie Movie) GetTitle() string {
	return ""
}

//ToJson convert movie to json
func (movie Movie) ToJSON() string {
	return ""
}

// AddActor add a new actor to movie
func (movie *Movie) AddActor(actor string) {}

// GetNbActors return the number of actors of a movie
func (movie Movie) GetNbActors() int {
	return 0
}
