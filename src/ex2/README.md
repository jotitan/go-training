## Second exercice, structures et films

Pour faire passer le test ex_test.go, il faut modifier le fichier structure_cases.go : 
* Modifier la structure Movie et corriger la méthode NewMovie
* Ajouter les fonctions GetTitle, ToJSON, AddActor et GetNbActors a Movie
* Ajouter les fonctions AddMovie, SortByYear et SortByName à Movies [Lien utile](https://golang.org/pkg/sort/)
* Créer une interface JSONable avec une methode ToJSON (renvoie un string) que doit implémenter movies
* Decommenter le test L66 pour tester la précédente fonction