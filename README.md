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
- [License](#license)
- [Acknowledgments](#acknowledgments)
- [Contact](#contact)

## Features

- **Real-time Q&A**: Get real-time questions from the community and answer them live.
- **Room-Based Discussions:** Create rooms for different topics or sessions.
- **Real-Time Updates**: Receive real-time updates on message creation, reactions, and status changes via WebSockets.
- **Room Sharing**: Share the room URL to allow others to join and participate.
- **Ownership & Management**: Each room is managed by an owner who can mark and unmark messages as answered.

## Tech Stack

- **Frontend**: React, TypeScript, Tailwind CSS
- **Backend:** Go (Golang)
- **WebSockets:** For real-time communication
- **Database:** PostgreSQL (running in Docker)
- **Containerization:** Docker

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Node.js and npm](https://nodejs.org/) (for frontend development)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/igorjpimenta/AskMeAnything.git
   cd AskMeAnything
    ```

2. **Set the environment variables**:
    ```bash
    echo 'DATABASE_HOST="localhost"
    DATABASE_PORT=5432
    DATABASE_NAME="your-database-name"
    DATABASE_USER="your-database-user"
    DATABASE_PASSWORD="your-database-password"
    PGADMIN_USER="your-pgadmin-user"
    PGADMIN_PASSWORD="your-pgadmin-password"
    VITE_API_URL="http://localhost:8080"
    VITE_WS_URL="ws://localhost:8080"' > .env

3. **Install backend dependencies**:
    ```bash
    cd backend
    go mod download
    ```

4. **Run PostgreSQL in Docker**:
    ```bash
    docker-compose up -d
    ```

5. **Run database migrations:**:
    ```bash
    go generate ./...
    ```

6. **Install frontend dependencies**:
    ```bash
    cd ../frontend
    npm install -f
    ```

### Running the Application

1. **Run the backend server**:
    ```bash
    cd backend
    go run ./cmd/main.go
    ```
    The backend server will start on `http://localhost:8080` and pgAdmin on `http://localhost:8081`.

2. **Run the frontend development server**:
    ```bash
    cd ../frontend
    npm run dev
    ```
    The frontend application will start on `http://localhost:3000`.

## Contributing
We welcome contributions from the community! Here's how you can get involved:

1. Fork the repository: Click the "Fork" button at the top right of this page.

2. Clone your fork:
    ```bash
    git clone https://github.com/yourusername/AskMeAnything.git
    ```

3. Create a branch:
    ```bash
    git checkout -b feature/your-feature-name
    ```

4. Make your changes: Write clear, concise commit messages.

5. Push to your fork:
    ```bash
    git push origin feature/your-feature-name
    ```

6. Submit a pull request: Describe your changes in the PR and link any relevant issues.

### Contribution Guidelines
- Follow the existing code style.
- Write tests for new features or bug fixes.
- Keep your pull requests small and focused on a single issue or feature.
- Provide clear documentation for new features.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- Inspired by an event from Rocketseat, the main backend repo is available [here](https://github.com/rocketseat-education/semana-tech-01-go-react-server/), and frontend [here](https://github.com/rocketseat-education/semana-tech-01-go-react-web).

## Contact
May you have questions or suggestions, feel free to open an issue or contact the project maintainers.