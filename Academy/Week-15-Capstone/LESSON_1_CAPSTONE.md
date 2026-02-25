# Week 15: The Final Capstone (Load Balancer Stress Test)

## The Objective
Use `db-stress` (Phase 4 Tool) to attack `load-balancer` (Phase 2 Tool).

## The Setup
1.  **Target:** A Load Balancer with 3 backends.
2.  **Weapon:** `db-stress` running 50 concurrent workers.
3.  **Chaos:** Killing a backend node mid-test.

## The Result
- **Resilience:** The LB's Health Check detected the dead node.
- **Routing:** Traffic automatically shifted to the remaining 2 nodes.
- **Performance:** Sustained ~1,500 RPS.

## The Lesson
A system is only "Production Ready" if it can survive a component failure without crashing the whole system. This is **Fault Tolerance**.
