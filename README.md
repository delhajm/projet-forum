[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/J7fVND2K)

# Groupie-tracker

Ce projet vise à offrir un espace où les utilisateurs
peuvent s'inscrire, participer à des discussions, et gérer leur profil personnel.

## Index

- [Index](#index)
- [Project Overview](#project-overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Games Description](#games-description)
- [Author](#author)

## Project Overview

This platform integrates three games: "Guess the Song", "Blind Test", and "Petit Bac", each offering a unique way to test your musical knowledge and reflexes. Players can create accounts, join games, and compete against friends while climbing the scoreboard to music glory.

## Prerequisites

To get started with this project, you'll need:

- [Go](https://go.dev/doc/install) installed on your local machine for backend development.
- [Git](https://git-scm.com/downloads) for version control and collaboration.

## Installation

1. Request access to be added as a collaborator.

2. Clone the repository:
   ```bash
   git clone https://github.com/Decorentin/groupie-tracker
   cd groupie-tracker/

3. Install dependencies and prepare your environment:
    ```bash
    go mod tidy

4. If you have packages problems, you can install them manually
    ```bash
    go get github.com/gosimple/slug
    go get github.com/mattn/go-sqlite3
    go get github.com/gosimple/unidecode

## Getting Started

After installation, follow these steps to run the platform:

1. Start the server:
    ```bash
    go run main/main.go

2. Navigate to http://localhost:8080 on your web browser to access the platform.

3. Sign up or log in to start playing the games.

## Games Description

Guess the Song: Players guess the title of a song from a snippet of its lyrics.
Blind Test: Identify songs as quickly as possible from audio clips.
Petit Bac: A music-themed version of the classic game, challenging players to list music-related terms starting with a given letter.
Each game is designed to test different aspects of your musical knowledge and reaction speed, offering a diverse and engaging experience.

## Author

- [@decorentin] (https://github.com/Decorentin)
- [@Tokennn] (https://github.com/Tokennn)
- [@M4xxes] (https://github.com/M4xxes)
- [@TerminaTorr45] (https://github.com/TerminaTorr45)

