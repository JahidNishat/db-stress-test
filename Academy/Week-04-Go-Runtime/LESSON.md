# Week 4: The Go Runtime & Scheduler

## The G-M-P Model

Go's secret weapon is its scheduler. It doesn't map 1 Goroutine to 1 OS Thread. Instead, it uses **M:N Scheduling**.

-   **G (Goroutine):** The smallest unit of execution. It has a small stack (starts at 2KB).
-   **M (Machine):** An actual OS Thread managed by the kernel.
-   **P (Processor):** A logical resource. Usually `P == number of CPU cores`. You need a `P` to run a `G` on an `M`.

### Why is this fast?
1.  **Low Memory:** 1 million goroutines = 2GB RAM. 1 million OS threads = ~1,000GB+ RAM.
2.  **Fast Switching:** Switching between goroutines happens in user-space (no expensive syscalls).
3.  **Work Stealing:** If one `P` is idle, it "steals" goroutines from another `P` to keep all CPU cores busy.

---

## Task: Visualization

1.  Run the code:
    ```bash
    cd Academy/Week-04-Go-Runtime
    go run main.go
    ```
2.  View the trace:
    ```bash
    go tool trace trace.out
    ```
    *(Note: This will open a browser window. Click on "View trace".)*

### What to look for:
-   How many "Procs" (P) do you see?
-   Can you see the goroutines being moved between different processors?
-   Look at the "Goroutines" row—how many are "Runnable" vs "Running"?
