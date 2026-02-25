# Week 19: Advanced Service Discovery

## The Goal
A standard Kubernetes `Service` acts as a Layer 4 Load Balancer. It hides the individual Pod IPs behind a single Virtual IP (VIP).
**The Problem:** If you are building a Layer 7 Load Balancer (like an API Gateway or a Consistent Hashing proxy), you *need* to know the individual Pod IPs to route traffic intelligently (e.g., sticky sessions).

## The Solution: Headless Services
By setting `clusterIP: None` in a K8s Service, you create a "Headless Service."
K8s stops providing a VIP. Instead, when you do a DNS lookup on the service name, K8s DNS returns **all the individual IP addresses** of the pods backing that service.

## The Architecture
1. **Headless Service:** Exposes Pod IPs via DNS.
2. **The Go Control Loop:** 
   - Every 5 seconds, perform `net.LookupHost("backend-headless")`.
   - Compare the newly discovered IPs against the current Consistent Hash Ring.
   - **Remove** dead IPs.
   - **Add** new IPs.

## Why this makes you a Senior Engineer
You didn't just write an app that runs *on* Kubernetes; you wrote an app that interacts *with* Kubernetes architecture dynamically. You built a self-healing, auto-scaling traffic router.
