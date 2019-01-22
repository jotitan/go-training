package ex2_test

import (
	. "ex2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ex2", func() {
	Describe("Manage structure with movies use cases", func() {
		Context("I create a movie", func() {
			movie := NewMovie("Title 1", 2010)
			It("Should exist", func() {
				Expect(movie.GetTitle()).To(Equal("Title 1"))
			})
			It("Should be jsonable", func() {
				Expect(movie.ToJSON()).To(Equal("{\"Title\":\"Title 1\",\"Year\":2010,\"Actors\":[]}"))
			})
		})
		Context("I add actors", func() {
			movie := NewMovie("Title 1", 2010)
			movie.AddActor("Actor 1")
			movie.AddActor("Actor 2")
			It("Have now 2 actors inside", func() {
				Expect(movie.GetNbActors()).To(Equal(2))
			})
		})
	})
	Describe("Manage list of movies", func() {
		movies := createMoviesList()
		Context("When I count movies", func() {
			It("Must found 8 movies in list", func() {
				Expect(len(movies)).To(Equal(8))
			})
		})
		Context("When I sort by name", func() {
			movies.SortByName()
			It("Casino Royale is the first movie", func() {
				Expect("Casino Royale").To(Equal(movies[0].GetTitle()))
			})
			It("Skyfall is the last movie", func() {
				Expect("Skyfall").To(Equal(movies[7].GetTitle()))
			})
			It("Goldeneye is the fourth movie", func() {
				Expect("Goldeneye").To(Equal(movies[3].GetTitle()))
			})
		})

		Context("When I sort by year", func() {
			movies2 := createMoviesList()
			movies2.SortByYear()
			It("From Russia with love is the first movie", func() {
				Expect("From Russia with Love").To(Equal(movies2[0].GetTitle()))
			})
			It("Skyfall is the last movie", func() {
				Expect("Skyfall").To(Equal(movies2[7].GetTitle()))
			})
			It("Casino Royale is the seventh movie", func() {
				Expect("Casino Royale").To(Equal(movies2[6].GetTitle()))
			})
			It("Live or let Die is the third movie", func() {
				Expect("Live or let Die").To(Equal(movies2[2].GetTitle()))
			})
		})
		// TODO : uncomment when creating JSONable interface
		/*Context("I used interface to generate a nice json of movies", func() {
			movies := createMoviesList()
			It("Get a json represntation of movies", func() {
				Ω(len(getJSON(movies))).Should(BeNumerically(">", 100))
			})
		})*/
	})
})

func createMoviesList() Movies {
	movies := Movies(make([]Movie, 0))
	movies.AddMovie(NewMovie("Goldeneye", 1995))
	movies.AddMovie(NewMovie("From Russia with Love", 1963))
	movies.AddMovie(NewMovie("Skyfall", 2012))
	movies.AddMovie(NewMovie("Licence to Kill", 1989))
	movies.AddMovie(NewMovie("Diamonds are forever", 1971))
	movies.AddMovie(NewMovie("Live or let Die", 1973))
	movies.AddMovie(NewMovie("Octopussy", 1983))
	movies.AddMovie(NewMovie("Casino Royale", 2006))
	return movies
}

func getJSON(obj JSONable) string {
	return obj.ToJSON()
}
