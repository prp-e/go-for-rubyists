# 12. Shipping and maintaining a Go service

The first time I wrote Hello World in Go, the interesting bit was that I could run the executable after the compiler had gone home. Now the executable has become a server. Before handing it to anyone, we have a few questions to answer about what happens when it is slow, loses a dependency, or receives a signal while serving a request.

## Build a release you can identify

```sh
cd examples
go test ./...
go vet ./...
go test -race ./...
go build -trimpath -o readinglist ./cmd/readinglist
go version -m ./readinglist
```

Build for the platform that will run the binary. For this standard-library example you can set `GOOS=linux GOARCH=amd64 CGO_ENABLED=0` before `go build` from a Unix shell. Native dependencies can change the build requirements. Run tests on the target platform or a suitable CI runner. Building a binary for another operating system doesn't execute it. Record the source commit, toolchain version, and build commands for each release.

Set timeouts around network edges. Our HTTP server sets read header, read, write, and idle timeouts, and the handler caps request bodies. Choose values for a real workload by considering slow clients and long legitimate requests. Streaming endpoints often need a different timeout policy. Start a shutdown budget on SIGTERM, and close external resources after active handlers finish. If the budget expires, force close and report the error.

## Make the next incident easier to understand

Use structured logs with `log/slog`. Record method, route, status, duration, and a request ID where applicable. Avoid logging bodies, credentials, and authorization headers. Keep liveness simple so an overloaded database doesn't make a process supervisor restart every instance. Readiness asks whether an instance should receive traffic and can include short dependency checks. Track request counts, latency, and errors as separate observations. A single average can look fine while a small group waits ten seconds.

Use a benchmark to compare operations after you have a question. Run `go test -bench . -benchmem` in the relevant package with real benchmark functions. Use CPU and memory profiles when a measurable bottleneck needs a location. Keep profiling interfaces private. For parsing boundaries, seed a fuzz test with empty, malformed, and ordinary input. Try `go test -fuzz FuzzName -fuzztime 10s` in the package that has that target. Keep minimized failures as regression cases. Benchmark results depend on hardware and workload, which makes claims such as Go being a fixed number of times faster than Ruby about as useful as comparing cars without saying where you're driving.

## Where to go next

The capstone is a reasonable place to try an SQL-backed store, migrations, authentication, or a small web UI. Add one change at a time and preserve the request behavior you already tested. Read the [Go specification](https://go.dev/ref/spec) when the exact language rule matters, the [package documentation](https://pkg.go.dev/std) when you need an API contract, and the [Go security guidance](https://go.dev/security/) when planning updates and dependency checks.

**Try it.** Build a binary and run it outside the source directory. Start a request, send a termination signal, and observe shutdown. Then explain to another Rubyist why there are two different contexts involved.
