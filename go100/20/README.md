#Fonctions variadiques

Une fonction variadique accepte un nombre variable de paramètres.
Le type de l’argument dont le nombre varie est précédé par …
Cela ne peut concerner que le dernier argument de la fonction, qui sera un slice à l'intérieur de la fonction.

    func fonction(b ...string) {
        fmt.Printf("%v, type is %T", b, b)
    }

    fonction("a", "b", "c", "hello") // [a b c hello], type is []string
