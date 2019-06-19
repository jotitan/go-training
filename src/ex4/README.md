## Quatrième exercice, serveur et appels Rest

###Serveur web simple

* Ecrire un serveur web avec la méthode HelloWorld. Mapper la méthode shutdown sur _/stop_. [Lien utile](https://golang.org/pkg/net/http/)
* Ecrire la méthode add5toValue qui prend un paramètre _value_ (qui est un nombre), additionne 5 et le renvoie. La mapper sur _/add5_. [Lien utile](https://golang.org/pkg/strconv/)
* Implémenter la méthode CallURL qui interroge une url et renvoie le résultat en string. [Lien utile](https://golang.org/pkg/io/ioutil/)

###Gestion des timeouts

* Implémenter la méthode wait1Second qui attend une second avant de renvoyer l'heure
* Créer une méthode mappé sur _/getTime_ qui utilise la méthode wait1Second
* Faire en sorte que cette méthode ne traite pas plus de 3 messages en parallèle
* Créer une autre méthode mappée sur _/getTimeWithLimit_ qui reprend le même mécanisme mais attend au maximum 1.5s pour obtenir une réponse.
Passé ce délai, une 500 doit être renvoyée.

### Calcul de panier

Le calcul de panier utilise une méthode longue (imaginons des appels externes) simulé par un sleep (méthode computePrice).

La méthode computeBasket mappée sur "/basket" calcul le cout total d'un panier et renvoie le resultats.

Faites en sorte de calculer le cout le plus rapidement possible a partir de la liste des produits "products".

Info => la variable productsBase est une map avec le prix de chacun des produits

/ ! \ Vous devez absolument utiliser la méthode computePrice qui simule la latence avec le time.Sleep au début.

####Bilan d'utilisation des channels  
* Limiter le débit de traitement
* Gérer des timeout
* Paralléliser des traitements couteux

