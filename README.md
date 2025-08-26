# Notely - CI/CD Pipeline Project

This repository contains the "Notely" application, a Go web server that was used as the foundation for the ["Learn CI/CD with GitHub Actions, Docker and Go"](https://www.boot.dev/) course on boot.dev.

The primary goal of this project was not to build the application itself, but to design and implement a complete Continuous Integration (CI) and Continuous Deployment (CD) pipeline for it using modern DevOps tools.

## The CI/CD Pipeline

A full CI/CD pipeline was built using GitHub Actions to automate the testing, building, and deployment of the application.

### Continuous Integration (CI)

The CI pipeline (`ci.yml`) is triggered on every pull request to the `main` branch. It runs a series of checks to ensure code quality and stability before any code is merged.

**CI Jobs:**
*   **Tests:**
    *   Checks out the code.
    *   Runs the full suite of unit tests using `go test`.
    *   Performs a security scan on the codebase using `gosec` to identify potential vulnerabilities.
*   **Style:**
    *   Ensures all Go code is formatted correctly according to `gofmt` standards.
    *   Runs a comprehensive linting check using `staticcheck` to enforce code quality and best practices.

### Continuous Deployment (CD)

The CD pipeline (`cd.yml`) is triggered on every push to the `main` branch, automatically deploying the new version of the application.

**Deployment Steps:**
1.  **Build:** Compiles the Go application into a production-ready binary.
2.  **Authenticate with GCP:** Securely authenticates with Google Cloud Platform using stored secrets.
3.  **Build & Push Docker Image:**
    *   Builds a lightweight, production-ready Docker image using the provided `Dockerfile`.
    *   Pushes the new image to Google Artifact Registry.
4.  **Database Migration:** Runs database migrations using `goose` to ensure the database schema is up-to-date before deploying the new code.
5.  **Deploy to Cloud Run:** Deploys the newly built Docker image to Google Cloud Run, making the updated application live.

## Key DevOps & CI/CD Skills Learned

*   **GitHub Actions:**
    *   Authored complex CI/CD workflows from scratch using YAML.
    *   Configured jobs to run on specific triggers (pull requests, pushes).
    *   Used community actions (`actions/checkout`, `actions/setup-go`, etc.) to build efficient pipelines.
    *   Managed secrets for securely authenticating with cloud providers.
*   **Docker:**
    *   Wrote a `Dockerfile` to containerize a Go application for consistent, portable deployments.
    *   Built and pushed Docker images to a container registry (Google Artifact Registry).
*   **Cloud Deployment (Google Cloud Run):**
    *   Automated the deployment of a containerized application to a serverless platform.
    *   Managed application configuration and scaling settings through deployment commands.
*   **Automated Quality Gates:**
    *   Integrated automated testing, linting, formatting, and security scanning into the development lifecycle to maintain high code quality.
*   **Database Migrations:**
    *   Incorporated automated database schema migrations into the deployment process to ensure consistency between the application and its database.