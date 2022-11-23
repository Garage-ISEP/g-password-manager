Démarre l'app
S'authentifie -> Middleware lambda cognito récupère les groupes du user

Le user récupère tous les noms de MDP pour ses groupes -> 

Si il a besoin d'en voir un, Lambda decode (vérifie droit de décodage -> renvoie le secret décodé) 

Pour créer un nv mdp:
Le user spécifie les groupes pour lesquels les MDP est dispo :
 * Nom
 * Secret
 * Email
 * ...
 * Sk = Id
 * Pk = PK_SECRET
 * WritingGroups = ["DSI"]
 * ReadingGroups = ["DSI", "CoderLab"]


