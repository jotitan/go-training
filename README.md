# Projet de training fournissant des exemples Go à coder

## Installation

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
* Installation les deux librairies pour faire les tests (Ginkgo et Gomega) : 
    * go get github.com/onsi/ginkgo/ginkgo
    * go get github.com/onsi/gomega/...

## Atelier

Il y a 4 packages avec des exercices : 
* ex1 : les basiques
* ex2 : les structures
* ex3 : les channels
* ex4 : les serveurs

Pour chaque exercice, vous devez : 
* Lire le README :)
* Lancer les tests dans *ex_test* : 
  * Soit en allant dans ex_test puis go test
  * Soit en lançant go test ./ex_test
* Corriger l'implémentation, quasiment vide par défaut jusqu'à ce que les tests soient verts
* Si vous avez besoin d'aide, il y a un fichier qui fonctionne (le ex_results), mais le mieux est de ne pas regarder

Si vous voulez simplifier la gestion des tests en évitant que tous tournent à chaque fois : 
* Initialiser le mode de test avec la commande suivante (lancer depuis src) : 
    * go mod tidy
* Lorsque vous lancez les tests, il faut utiliser ginkgo avec un petit hook (en étant dans un répertoire, ex1 par exemple) : 
    * ginkgo --afterSuiteHook="go run ../launchers/next_case.go (ginkgo-suite-passed) (ginkgo-suite-name)"
* A tout moment vous pouvez revenir au cas normal en lancant : 
    * go run launchers\next_case.go reset

### Arboresence

Pour chaque exercice, il y a trois packages : 
* Un package de test ex*_test : 
    * Un fichier de test (ex*_test.go)
    * Un fichier lanceur de test (ginkgo) (ex*_suite_test.go)
* Un package avec les fonctions à implémentater à faire (ex*/_cases.go)
* Un package avec une implémentation correcte (ex*_results/_cases.go)

## Lancement des tests

* Pour lancer les tests, dans un terminal, mettez vous dans le package (par exemple ex1) et faites : 
    * _go test_ ou _ginkgo test_
    * _ginkgo watch_ => relance les tests à chaque modification de code
* Si vous souhaitez modifier le test pour utiliser la version réponse : 
    * Modifier l'import du fichier*_test.go en ajoutant "_results" au nom du package pour passer sur la solution

## Rappel des commandes utiles

* Construire un executable : 
    * Un fichier avec une méthode main : _go build mon_fichier.go_
    * Un répertoire avec une méthode main : _go build mon_folder_
* Executer un fichier go
    * _go run mon_fichier.go_
* Lancer les tests dans un package
    * go test
* Lancer un watcher ginkgo pour relancer les tests quand le code change
    * ginkgo watch _folder_