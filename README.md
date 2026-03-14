# 🛡️ Projet RED : Legends of Ynov

Bienvenue dans **Legends of Ynov**, un RPG textuel (CLI) développé en langage Go.
Ce projet a été réalisé dans le cadre du module de programmation Go à Ynov Campus.

## 📋 Fonctionnalités

Le jeu propose une aventure complète en ligne de commande incluant :

- **Création de Personnage** : Choix du nom et de la classe (Humain, Elfe, Nain) influençant les statistiques (PV).
- **Système de Combat** : Combats au tour par tour contre des monstres avec :
    - Attaques physiques et magiques.
    - Gestion des coups critiques et de l'aléatoire.
    - *Scaling* : Les ennemis deviennent plus forts à mesure que vous gagnez des niveaux.
- **Progression (RPG)** :
    - Système d'expérience (XP) et montée de niveau.
    - Amélioration des PV Max et restauration de la santé.
- **Gestion d'Inventaire** :
    - Stockage d'objets (limité à 10 places).
    - Utilisation de potions.
- **Économie & Crafting** :
    - **Marchand** : Achat de potions, sorts et matériaux.
    - **Forgeron** : Fabrication d'équipements (Tête, Torse, Pieds) à partir de ressources pour booster ses stats.
- **Interface Immersive** : Utilisation de couleurs ANSI et nettoyage d'écran pour une meilleure lisibilité.

## 🚀 Prérequis

- **Go** (version 1.16 ou supérieure recommandée) installé sur votre machine.
- Un terminal supportant les couleurs ANSI (PowerShell, Bash, Zsh, Terminal macOS, etc.).

## 🎮 Comment lancer le jeu

1. Clonez le dépôt ou téléchargez les fichiers sources.
2. Ouvrez un terminal dans le dossier `src` du projet.
3. Exécutez la commande suivante :

```bash
go run .
```

*(Alternativement : `go run main.go character.go systemes.go combat.go`)*

## 🛠️ Structure du Projet

- **main.go** : Point d'entrée, gestion du menu principal et de l'interface utilisateur.
- **character.go** : Définition des structures (Personnage, Équipement) et statistiques.
- **combat.go** : Logique de combat, IA des monstres et tours de jeu.
- **systemes.go** : Gestion de l'inventaire, du marchand et de la forge.

## 👥 Auteurs

- **Frontend / UI** : [Ton Prénom/Nom]
- **Backend / Logique** : [Prénom/Nom de ta camarade]

---
*Projet réalisé pour Ynov Campus.*