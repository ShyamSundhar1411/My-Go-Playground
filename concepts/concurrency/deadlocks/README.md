# Deadlocks

Deadlock is a condition in an operating system where two or more processes become permanently blocked because each process is waiting for a resource held by another. A deadlock can occur only when four necessary conditions are simultaneously satisfied: **mutual exclusion**, **hold and wait**, **no preemption**, and **circular wait**. Deadlocks can be managed using approaches such as **prevention**, **detection**, and **recovery** algorithms.

## Conditions

- **Mutual Exclusion:** A resource can be used by only one process at a time and cannot be shared simultaneously.
- **Hold and Wait:** A process holds at least one resource while waiting to acquire additional resources held by other processes.
- **No Preemption:** A resource cannot be forcibly taken from a process while it is in use.
- **Circular Wait:** A set of processes exists where each process is waiting for a resource held by the next process in a circular chain.
