#Les valeurs d'interface
En coulisse, les valeurs d'interface peuvent être considérés comme un n-uplet d'une valeur et un type concret:

(value, type)
Une valeur d'interface détient une valeur d'un type concret spécifique sous-jacent.

Appeler une méthode sur une valeur d'interface exécute la méthode du même nom sur son type sous-jacent.

##Les valeurs d'interface avec les valeurs nulle sous-jacente
Si la valeur concrète, à l'intérieur même de l'interface est nulle, la méthode sera appelée avec un récepteur nul.

Dans certaines langues, cela déclencherait une exception de pointeur nul, mais en Go il est courant d'écrire des méthodes qui gèrent avec élégance d'être appelé avec un récepteur nul (comme avec la méthode M dans cet exemple.)

Notez que la valeur de l'interface qui contient une valeur concrète nulle est en soi non-nulle.

##Les valeurs nulle d'interface
Une valeur nulle d'interface détient ni valeur, ni type concret.

Appeler une méthode sur une interface nulle est une erreur d'exécution, car il n'y a pas de type à l'intérieur du tuple d'interface pour indiquer quelle méthode concrète appeler.
