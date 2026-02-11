# RWMutex

Read Write mutex is a raeder/writer mutual exclusion locl that allows an arbitrary number of readers to hold the lock concurrently or a single writer to hold the lock exclusively. It allows concurrent reads and exclusive writes.This enables high concurrency for read-heavy workloads while still ensuring data consistency for writes.