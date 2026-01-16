# Concurrency

**Concurrency** is the ability of a system to handle and execute multiple tasks simultaneously. Importantly, these tasks **do not necessarily execute at the exact same time**; they can be interleaved, scheduled, or paused by the system.

---

## Concurrency vs Parallelism

**Parallelism** refers to the **simultaneous execution** of multiple tasks using multiple cores or processors.

| Aspect | Concurrency | Parallelism |
|--------    |------------|-------------|
| Definition | Handling multiple tasks at the same time (interleaving) | Executing multiple tasks at the same exact time |
| Execution | Can occur on a single core | Requires multiple cores/CPUs |
| CPU Behavior | Tasks are scheduled via context switching | Tasks truly run in parallel |
| Relationship | All parallel processes are concurrent, but not all concurrent processes are parallel |

**Key takeaway:** A concurrent system can simulate parallelism on a single-core machine, but true parallelism requires multiple cores.


---

## Go and Concurrency

Go provides **first-class support for concurrency** via:
- **Goroutines:** Lightweight threads managed by the Go runtime  
- **Channels:** Safe communication and synchronization between goroutines  
- **Mutexes / RWMutex / sync.Cond:** Fine-grained control over shared state  

These primitives allow building **safe, high-performance concurrent systems** without relying on low-level OS threads directly.
