# Week 1: The Chaos of Concurrency

## The "Baby" Explanation: The Kitchen

Imagine you are cooking in a kitchen.

**1. Serial Execution (The Normal Way)**
There is **one chef** (The CPU). He chops onions, *then* boils water, *then* fries the egg. He does one thing at a time. It is slow, but safe. Nothing burns.

**2. Concurrency (The Manager)**
The chef is smart. He puts the water to boil. While waiting for it to boil, he chops onions. He is switching tasks. He is still just **one person**, but he is managing multiple tasks at once.
*In Go, this is like having one CPU core running multiple Goroutines.*

**3. Parallelism (The Team)**
Now we hire a **second chef**. One chef chops onions. The other chef fries eggs. They are working at the exact same time.
*In Go, this is multiple CPU cores running multiple Goroutines.*

---

## The Problem: The Shared Knife (Race Condition)

Imagine both chefs need the **one huge knife** (Shared Memory/Variable).

1.  Chef A reaches for the knife.
2.  Chef B reaches for the knife at the exact same time.
3.  They bump heads. The knife falls. A foot is lost.

In code, this is a **Race Condition**.
Two goroutines try to change the same variable at the same time. The computer gets confused and writes the wrong data.

---

## Assignment 1: Break It

We are not going to write "correct" code today. We are going to write **broken** code. You cannot fix a bug you do not understand.

**Task:**
1.  Create a file `main.go` inside this folder.
2.  Create a variable `counter = 0`.
3.  Write a function that adds `1` to `counter` 1000 times using a loop.
4.  Launch that function **two times** using `go routine` (the `go` keyword).
5.  Wait for them to finish (use a `time.Sleep` for now, we will learn proper waiting later).
6.  Print the `counter`.

**Expected Result:**
If math works, 1000 + 1000 = 2000.
But if you do this right... it will **NOT** be 2000.

**Go write the code. Run it multiple times. Tell me what happens.**
