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


## Authentification et API G Suite
* Utilise l'API : https://developers.google.com/admin-sdk/directory/reference/rest/v1/groups/list?apix_params=%7B%22domain%22%3A%22garageisep.com%22%2C%22userKey%22%3A%22theodore.prevot%40garageisep.com%22%7D avec un compte de service.
### Etapes de configuration (Doc Google correspondante : https://support.google.com/a/answer/162106)
* Créer un compte de service sur https://console.cloud.google.com/
* Afficher les détails du compte de service puis afficher les paramètres avancés pour activer le `Client OAuth Google Workspace Marketplace` puis cliquer sur `Afficher la console d'administration google workspace pour autoriser le compte service`
* Aller dans `Sécurité > Contrôle des accès et des données > Commandes des API` puis cliquer sur `Gérer la délégation au niveau du domaine` et ajouter le client API en mettant l'ID client du compte de service créé ainsi que les scopes nécessaires :
	* .../auth/admin.directory.group.readonly
	* .../auth/admin.directory.group
* Attendre une heure minimum
* Générer une clé pour le compte de service via l'onglet `CLES`