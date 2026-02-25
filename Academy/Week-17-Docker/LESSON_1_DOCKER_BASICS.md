# Week 17: Docker Basics & Multi-Stage Builds

## What is Docker?
It's a way to package your app and all its dependencies (OS, libraries, configs) into a single "Image".
- **Works on my machine:** Works on your machine.
- **Works in Cloud:** Works in Cloud.

## The Dockerfile Recipe
1.  **FROM:** Pick a base OS (e.g., `golang:alpine`).
2.  **COPY:** Move files from your laptop to the container.
3.  **RUN:** Execute commands (e.g., `go build`).
4.  **CMD:** The command to run when the container starts.

## Multi-Stage Builds (The Pro Move)
**Problem:** A Go compiler image is HUGE (800MB). Your binary is TINY (10MB).
**Solution:**
- **Stage 1 (Builder):** Compile the code.
- **Stage 2 (Runner):** Start with a tiny `alpine` image (5MB). Copy **only the binary** from Stage 1.
- **Result:** A 15MB image instead of 800MB.

## Docker Compose
A tool to run multiple containers (App + DB + Redis) together using a YAML file.
