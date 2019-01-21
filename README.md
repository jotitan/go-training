# Projet de training fournissant des exemples Go à coder

## Pré-requis

Pas grand chose, mais il le faut quand même : 
* Installer [Go](https://golang.org/dl/)
* Avoir un IDE : 
    * [Visual studio code](https://code.visualstudio.com/Download) avec le [plugin associé](https://code.visualstudio.com/docs/languages/go)
    * Intellij ultimate avec le [plugin associé](https://plugins.jetbrains.com/plugin/9568-go)
    * [JetBrains Goland](https://www.jetbrains.com/go/) pour les chanceux
* Configurer son environnement : 
    * GOPATH pointe sur les répertoires où se trouvent les sources : 
        * Les sources chargées sur internet (comme un repo maven)
        * Le répertoire ou se trouve le projet
        * Les répertoires sont séparés par : sous linux et par ; sous windows
    * GOPATH dans son IDE : le même, pour qu'il puisse aider à l'autocompletion

## Organisation

Pour chaque exercice, deux packages package : 
* Un package principal (ex*) avec : 
    * Un fichier de structure avec les fonctions à implementer (_cases.go)
    * Un fichier de test (ex*_test.go)
    * Un fichier lanceur de test (ginkgo) (ex*_suite_test.go)
* Un package avec une implémentation correcte (ex*_results)

## Lancement des tests

* Pour lancer les tests, dans un terminal, mettez vous dans le package (par exemple ex1) et faites : 
    * go test
* Si vous souhaitez modifier le test pour utiliser la version réponse : 
    * Modifier l'import en ajoutant "_results" au nom du package pour passer sur la solution

## Rappel des commandes importantes

* Construire un executable : 
    * Un fichier avec une méthode main : _go build mon_fichier.go_
    * Un répertoire avec une méthode main : _go build mon_folder_
* Executer un fichier go
    * _go run mon_fichier.go_
* Lancer les tests dans un package
    * go test