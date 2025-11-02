## Go Healthcheck Service
A simple Go-based healthcheck service, built and deployed in multiple phases to practice containerization, orchestration, monitoring, and CI/CD.

### 📌 Project Goals
This project is designed to:
Learn and apply Go development best practices
Containerize and orchestrate services with Docker & Kubernetes
Implement monitoring and logging
Automate testing and deployment via CI/CD


### 🚀 Progress Tracker  

| Phase | Description | Status |
|-------|-------------|--------|
| 1️⃣ | **Building the app locally** — Clone/init project, install dependencies, run locally, test with curl/Postman | ![Completed](https://img.shields.io/badge/status-completed-brightgreen?style=flat-square) |
| 2️⃣ | **Dockerizing the app** — Create clean Dockerfile, build image, connect to Postgres via Docker Compose | ![In Progress](https://img.shields.io/badge/status-in--progress-yellow?style=flat-square) |
| 3️⃣ | **Container Orchestration** — Deploy app + DB via Docker Compose & Kubernetes | ![In Progress](https://img.shields.io/badge/status-in--progress-yellow?style=flat-square) |
| 4️⃣ | **Monitoring & Logs** — Add Prometheus metrics, structured JSON logs, integrate with Loki/ELK/Grafana | ![Pending](https://img.shields.io/badge/status-pending-lightgrey?style=flat-square) |
| 5️⃣ | **CI/CD** — Use GitHub Actions to run tests, build & push image, deploy to dev env | ![Pending](https://img.shields.io/badge/status-pending-lightgrey?style=flat-square) |


### Tech Stack
### Languages & Frameworks
Go
### Tools & Platforms
Docker
PostgreSQL
Kubernetes
Prometheus
Grafana
GitHub Actions

### 📖 How to Run Locally
#### Clone repository
git clone https://github.com/username/go_healthcheck.git
cd go_healthcheck

### Install dependencies
go mod tidy

### Run the app
go run main.go

### Test endpoint:
curl http://localhost:8080/healthz
