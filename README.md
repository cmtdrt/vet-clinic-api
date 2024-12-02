<p align="center">
API en langage GO pour une clinique vétérinaire qui s'occupe principalement des chats. 
Cette API permettra de gérer les informations des patients (chats), les consultations vétérinaires, et les traitements administrés.
</p>
<p align="center">
Toutes les requêtes POSTMAN sont disponibles dans le projet et sont prêtes à être importées:
</p>
<p align="center">
  <img src="https://github.com/user-attachments/assets/040469f5-3f89-4448-a5fa-5f6eeae1bc8d" alt="image" height="350px">
</p>


**Objectifs:**

- Structurer l’API avec des routes organisées par fonctionnalité.
- Utiliser GORM pour les opérations **CRUD** (Create, Read, Update, Delete) sur les patients, les consultations et les traitements.
- Implémenter une architecture modulaire avec des repositories et une configuration centralisée.
_________________________________________________________________________
**Exigences Fonctionnelles:**

- Gestion des Chats (Patients)
  - Création d’un chat : permet de créer un nouvel enregistrement de chat avec des informations comme le nom, l’âge, la race et le poids.
  - Récupération des chats : permet de récupérer tous les chats ou un chat spécifique par ID.
  - Mise à jour d’un chat : permet de mettre à jour les informations d’un chat existant.
  - Suppression d’un chat : permet de supprimer un chat de la base de données.
- Gestion des Consultations
  - Création d’une consultation : enregistre une nouvelle consultation pour un chat donné, avec des informations comme la date, le motif et le vétérinaire en charge.
  - Historique des consultations : permet de récupérer l'historique complet des consultations pour un chat spécifique.
- Gestion des Traitements
  - Création d’un traitement : ajoute un traitement prescrit (ex : un médicament) pour une consultation donnée.
  - Récupération des traitements : permet de consulter tous les traitements prescrits pour une consultation.
_________________________________________________________________________
**Challenge Avancé:**

- Historique des soins pour un chat :
  - Ajoutez une route /cats/{id}/history qui récupère l’historique complet des soins pour un chat : consultations et traitements associés.
- Filtrage des consultations :
  - Permettez un filtrage des consultations par date, vétérinaire, ou motif.
- Gestion des relations avancées avec GORM :
  - Pour chaque consultation, chargez également les traitements associés.
