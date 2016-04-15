#Switch

Le switch est très classique :

    switch expression {
    case a:
        ...
	case b:
        ...
	default:
        ...
    }

Un cas est terminé automatiquement, sauf s'il fini par une déclaration fallthrough.

Le switch évalue les cas de haut en bas, en s'arrêtant quand un cas correspond.

(Par exemple,

    switch i {
    case 0:
    case f():
    }
n'appelle pas f si i==0.)