package ex2

// TODO : create structure Movie with public attributes title, year and list of actors (string list)
type Movie struct{}

type Movies []Movie

func (ms *Movies) AddMovie(movie Movie) {}

// Sort movies by year
func (ms Movies) SortByYear() {}

// Sort movies by name
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
func (movie Movie) ToJson() string {
	return ""
}

// AddActor add a new actor to movie
func (movie *Movie) AddActor(actor string) {}

func (movie Movie) GetNbActors() int {
	return 0
}
