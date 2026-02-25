# Capstone: The Layer 7 Load Balancer

**Goal:** Build a simulation of a Load Balancer (like Nginx) that routes traffic based on Consistent Hashing.

## Architecture
1.  **The LB (Load Balancer):** Entry point.
    - Holds the `ConsistentHash` Ring.
    - Receives HTTP Requests.
    - Routes them to the correct "Backend".
2.  **The Backends:** 3 separate HTTP Servers running on different ports (8081, 8082, 8083).
3.  **The Client:** Sends requests to the LB.

## Components to Build

### 1. `backend/main.go`
- A simple HTTP server.
- Accepts a command line flag `-port`.
- Endpoint `/`: Returns "Hello from Server [Port]".

### 2. `lb/main.go`
- A Proxy Server running on port `8000`.
- **Startup:**
    - Initialize Consistent Hash Ring.
    - Add "http://localhost:8081", "http://localhost:8082", "http://localhost:8083".
- **Handle Request:**
    - Extract IP or User-ID from request.
    - Call `ring.Get(userID)`.
    - **Reverse Proxy:** Forward the request to that backend URL.
    - Write response back to client.

## Steps
1.  Create `backend/main.go`.
2.  Create `lb/main.go` (Copy `ring.go` logic here or import it).
3.  Run 3 backends in 3 terminals.
4.  Run LB in 1 terminal.
5.  Curl the LB and watch the traffic distribute.
