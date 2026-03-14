# Gestion de Projet - Projet RED : Legends of Ynov

Ce document détaille l'organisation, la planification et les choix techniques adoptés pour le développement du projet "Legends of Ynov".

---

## 1. Répartition des Tâches

Le projet a été divisé en deux pôles principaux pour optimiser le développement :

#### Backend / Logique du Jeu
-   **Responsable** : [Prénom/Nom de la camarade]
-   **Missions** :
    -   Définition des structures de données (`Character`, `Monster`, `Equipment`).
    -   Implémentation des mécaniques de jeu : combat au tour par tour, gain d'expérience, montée de niveau.
    -   Développement des systèmes annexes : inventaire, marchand, forge.
    -   Gestion de la logique de persistance des données du joueur (PV, or, objets).

#### Frontend / Interface Utilisateur (CLI)
-   **Responsable** : [Ton Prénom/Nom]
-   **Missions** :
    -   Conception de l'affichage des menus principaux et des sous-menus.
    -   Gestion des entrées et des choix de l'utilisateur.
    -   Mise en forme de la sortie terminal avec des couleurs ANSI pour une meilleure lisibilité et immersion.
    -   Garantir une expérience utilisateur fluide et intuitive en ligne de commande.

---

## 2. Calendrier de Réalisation

Le projet a été découpé en sprints hebdomadaires :

-   **Semaine 1 : Initialisation et Structures de Base**
    -   Création du dépôt Git et de l'architecture (`src`, `docs`).
    -   Définition des `structs` pour le personnage et l'équipement (`character.go`).
    -   Implémentation de la création de personnage.

-   **Semaine 2 : Mécaniques de Combat**
    -   Développement de la boucle de combat au tour par tour (`combat.go`).
    -   Intégration des attaques, des monstres et de la condition de victoire/défaite.

-   **Semaine 3 : Systèmes de Jeu et Progression**
    -   Développement des systèmes d'inventaire, de marchand et de forge (`systemes.go`).
    -   Implémentation du système d'XP et de montée de niveau.

-   **Semaine 4 : Finalisation et Polissage**
    -   Amélioration de l'interface (nettoyage d'écran, couleurs).
    -   Phase de tests, débogage et équilibrage.
    -   Rédaction de la documentation (`README.md`, `gestion_projet.md`).

---

## 3. Choix Techniques

#### Langage de Programmation : Go
**Justification** : Le langage Go a été retenu (et imposé par le module) pour sa simplicité, ses performances et son typage statique qui sécurise le code. Sa bibliothèque standard est parfaitement adaptée à la création d'applications en ligne de commande (gestion des entrées/sorties, formatage de chaînes).

#### Structures de Données
-   **Structs** : Le cœur de notre modèle. `Character`, `Monster`, et `Equipment` permettent de regrouper de manière cohérente les attributs de chaque entité du jeu.
-   **Slices** : Idéales pour gérer des collections de taille dynamique comme l'inventaire (`[]string`) et la liste des sorts appris, offrant flexibilité et performance.
-   **Maps** : Utilisées de manière ciblée, notamment dans la fonction de forge (`forger`) pour valider la présence des matériaux requis, ce qui permet une vérification rapide et efficace.