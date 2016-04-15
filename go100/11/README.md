#Defer

Une déclaration defer reporte l'exécution d'une fonction jusqu'à ce que la fonction environnante retourne.

Les arguments de l'appel différé sont évalués immédiatement, mais l'appel de fonction n'est pas exécuté jusqu'à ce que la fonction environnante retourne.

Les appels de fonction différés sont poussés sur une pile. Quand une fonction retourne, ses appels différés sont exécutées dans l'ordre dernier entré, premier sorti.