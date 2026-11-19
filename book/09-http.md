# 9. From Sinatra to net/http

One of the things I liked about Sinatra was how quickly I could get from a thought to a request in a terminal. Let's keep that part. In Sinatra, `get "/books" do ... end` installs a route. In Go we register a function with a multiplexer. The function receives the request and a writer for the response.

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintln(w, id)
})
```

Method and wildcard patterns are available in Go 1.22. A GET pattern also matches HEAD. If another method targets a registered path, the mux can reply with 405 and an Allow header. Convert path values to the type your application needs. A positive decimal ID that doesn't exist and a path segment that can't be parsed are two different questions. Our service answers them with 404 and 400 respectively.

## JSON and the response boundary

An exported field such as `Title string` can be serialized. A tag `json:"title"` picks its JSON name. Give incoming JSON a dedicated struct so clients can only set the fields the endpoint accepts. Someone posting a new book doesn't get to pick the server's next identifier.

Use `http.MaxBytesReader` before decoding a request. Then use `json.Decoder.DisallowUnknownFields` if the endpoint has a closed input schema. Decode a second time and require EOF to reject a second value after the first. Decide what to do about null and missing fields. A string field's zero value alone doesn't distinguish all the cases. The capstone treats missing, null, and blank title as invalid. Standard decoding still accepts duplicate object keys, with later values generally winning for our string field. A protocol requiring duplicate-key rejection needs additional validation.

Set response headers before writing a status or any body bytes. The first write defaults to 200 if you haven't chosen a status. Use 201 for a created resource and send its URL in Location. Return a stable shape for application errors. Routes handled directly by the standard mux can use its plain-text error responses, so clients should handle both when calling our example.

## Test a handler without starting a server

Use `httptest.NewRequest` to set up a request and `httptest.NewRecorder` to capture a response. Call `mux.ServeHTTP(recorder, req)` when route matching or path variables are part of the behavior. Call `recorder.Result()` and check status, headers, and decoded body. `httptest.NewServer` starts a real local listener when you need to test a client. Close it when the test ends.

An HTTP client needs a time budget too. Use `http.NewRequestWithContext` and an `http.Client` with a timeout. Check the HTTP status separately from the transport error. Close a successful response's body after reading what you need. Reuse the client across requests so its transport can reuse connections.

**Try it.** Test the route above for an ID of `42`, a missing path segment, and a POST. Predict which calls reach the handler.
