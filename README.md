# Ask Me Anything (AMA) Application

![License](https://img.shields.io/github/license/igorjpimenta/AskMeAnything)
![Contributors](https://img.shields.io/github/contributors/igorjpimenta/AskMeAnything)
![Issues](https://img.shields.io/github/issues/igorjpimenta/AskMeAnything)
![Forks](https://img.shields.io/github/forks/igorjpimenta/AskMeAnything)
![Stars](https://img.shields.io/github/stars/igorjpimenta/AskMeAnything)

## Overview

Welcome to the Ask Me Anything (AMA) application! This project is a real-time platform where users can ask questions and receive answers in a live, interactive environment. The backend is built with Go, leveraging WebSocket for real-time communication.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [Acknowledgments](#acknowledgments)
- [Contact](#contact)

## Features

- **Real-time Q&A**: Users can ask and answer questions in real-time.
- **Room-Based Discussions:** Create rooms for different topics or sessions.
- **WebSocket Integration**: Low-latency communication using WebSocket.

## Tech Stack

- **Backend:** Go (Golang)
- **Frontend:** React.js, Tailwind and Typescript
- **WebSockets:** For real-time communication
- **Database:** PostgreSQL (running in Docker)
- **Containerization:** Docker

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Node.js](https://nodejs.org/) and [npm](https://www.npmjs.com/)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/igorjpimenta/AskMeAnything.git
   cd AskMeAnything/backend

2. **Install dependencies**:
    ```bash
    go mod download

3. **Run PostgreSQL in Docker**:
    ```bash
    docker-compose up -d

4. **Run database migrations:**:
    ```bash
    go generate ./...

### Running the Application

1. **Run the backend server**:
    ```bash
    go run ./cmd/main.go

2. The backend server will start on `http://localhost:8080`

### License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- Inspired by an event from Rocketseat, the main backend repo is available [here](https://github.com/rocketseat-education/semana-tech-01-go-react-server/), and frontend [here](https://github.com/rocketseat-education/semana-tech-01-go-react-web).

## Contact
If you have any questions or suggestions, feel free to open an issue or contact the project maintainers.