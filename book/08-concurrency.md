# 8. Goroutines, channels, and context

Starting a goroutine is easy enough to make you suspicious. Put `go` in front of a call and the function runs concurrently. The runtime schedules goroutines onto operating-system threads. Now comes the work. Who waits for the result? Who stops the goroutine? What happens if one worker finishes while another is blocked trying to send?

A goroutine shares memory with the rest of the process. If one goroutine writes a variable while another accesses it, you need synchronization. A mutex is usually a good fit for shared state such as our catalog. Use a channel to send work or deliver results, especially when sending transfers ownership. Sending a pointer through a channel doesn't make simultaneous mutations through that pointer safe.

## A batch with a clear end

In `examples/cmd/parallel`, two workers read title indexes from a jobs channel. Each owns a different result slot for each job and only reads the input slice. The caller closes the jobs channel after sending the last index and waits for the workers before reading the result. This preserves the input order even when a later title finishes first.

```go
for range workers {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := range jobs {
            result[i] = strings.ToUpper(titles[i])
        }
    }()
}
for i := range titles {
    jobs <- i
}
close(jobs)
wg.Wait()
```

A send on an unbuffered channel waits for a receiver. A buffered channel can accept sends until it fills. A buffer changes when the sender waits. You still need an answer for what happens after the buffer fills. The sender that knows there are no more results closes the channel. A send to a closed channel panics. A receive can use `v, ok := <-ch` to detect that the channel is closed and drained. `range` over a channel stops at that point. A nil channel blocks forever on send or receive. In a `select` it disables its case.

## A caller gets to leave

Pass `context.Context` as the first parameter to an operation that can wait. HTTP handlers can use `r.Context()`, which the server cancels when the request's lifetime ends. Derive deadlines with `context.WithTimeout` and defer the cancel function to release resources promptly. Cancellation is a signal, so every worker has to observe it. A library that blocks forever without reading the context doesn't become cancellable because we named a variable `ctx`.

```go
select {
case jobs <- job:
    // The worker can now receive job.
case <-ctx.Done():
    return ctx.Err()
}
```

If multiple cases are ready, a select may choose any of them. For a worker pool, pair cancellation-aware sends with cancellation-aware work and a wait for the started goroutines. Close an output channel only after all its senders are done. Give external I/O its own deadline and use context-aware APIs. Avoid putting optional configuration or database handles in context values. A request ID is a reasonable value because it belongs to the request's lifetime.

In Go modules using language version 1.22 or later, loop variables declared by the loop have per-iteration scope. An existing outer variable assigned by the loop is still shared. Check what the closure actually references, especially when maintaining older modules.

**Try it.** Add a context to the parallel example. Cancel it while jobs are outstanding. Stop sending, close jobs, and wait before returning. Decide whether to return partial results and test that contract with the race detector.
