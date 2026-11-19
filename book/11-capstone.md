# 11. A reading list that can take a request

We've done enough exercises involving twelve numbers and a rectangle. Let's write a little service you could actually use while deciding which book to read next. It stores a title, assigns an ID, and remembers whether you've finished it. The examples in this chapter run together in `examples/cmd/readinglist` and `examples/internal/readinglist`.

## Decide the contract first

`GET /books` lists books in creation order. The empty response is an empty JSON array. `POST /books` accepts a JSON object with a title, returns 201 with the created book, and sets Location to its resource URL. `GET /books/{id}` retrieves one record. `PUT /books/{id}/read` marks it read and returns the record, even if it was already read. Repeating that PUT leaves the same state. Repeating POST creates a new record.

The storage type keeps a slice of records behind a read/write mutex. Each ID is the record's one-based index because our contract has no delete operation. If you add deletion, give IDs an independent sequence and change the lookup strategy before reusing the old code. The list method returns its own slice so an HTTP encoder can use the snapshot after the lock is released. A book contains strings, a boolean, and an integer, so copying the record is enough here. A new slice field would make that worth revisiting.

The handler parses request data, calls the store, and translates the result into HTTP. It trims titles through the store, rejects blank titles and titles over 200 Unicode code points, and limits JSON bodies to 4096 bytes. The server expects `Content-Type: application/json` for creation. The API returns JSON for application errors without echoing the parser's internals. The service binds to loopback and stores data in memory. Restarting it clears the list. There's no authentication built into this exercise.

## Take it for a walk

```sh
cd examples
go run ./cmd/readinglist
```

In another terminal:

```sh
curl -i http://127.0.0.1:8080/books
curl -i -H 'Content-Type: application/json' \
  -d '{"title":"The Dispossessed"}' \
  http://127.0.0.1:8080/books
curl -i http://127.0.0.1:8080/books/1
curl -i -X PUT http://127.0.0.1:8080/books/1/read
```

Try posting a blank title and fetching ID 999. The first should give 400 and the second should give 404. Create two books and check the list order. Press Ctrl+C in the server terminal. The command stops accepting new connections and waits up to five seconds for active requests. It uses a fresh context for shutdown because the signal context has already been canceled.

## Follow one POST through the code

Start in `internal/readinglist/http.go`. The POST route calls `createBook`. The media type check happens before decoding so a client knows how to correct an unsupported request. The first decode answers whether the body begins with an object we can use. The second decode answers whether that's the whole body. If you remove the second decode and post two objects in a row, the first object can still be accepted. Try it in the test before putting the check back.

Next follow `store.Add`. It trims the title and counts code points before acquiring the lock. This work depends only on the input, so there's no reason to hold the lock through it. Inside the lock, constructing the Book and appending it form one operation. The caller receives a value copy. Read back `store.List` and look for the same ownership choice in its return value.

Finally the handler writes Location before its 201 status, and the encoder writes the book. If a client disconnects after the append, the book still exists even if the client never gets a response. Retrying a POST can therefore create a duplicate. A production API that promises safe retries needs a contract such as a client-supplied idempotency key and storage that protects that promise across concurrent requests and restarts. Notice how the question began in HTTP and ended back at the storage layer.

## What would change with persistence?

Start by writing down the operations the handler really needs before introducing an interface. For this service that's add, list, find, and mark read. An SQL version has additional expected failures, so a signature returning only a record and a boolean would hide information. Give the operation a context and a way to surface a storage error. Use a missing-record error for a 404 and treat a database failure as a server error. Keep internal database details in logs.

Keep your existing HTTP tests and add integration tests for the new implementation. Check creation order, simultaneous inserts, and what happens when an operation fails halfway through. If you change the contract, change its tests deliberately. It is perfectly reasonable to leave the memory version in the book as a little program you can run on a train.

## Give the bad paths tests of their own

The HTTP tests drive the routes without an external listener. They check the creation status and Location, lookup, repeated marking, an empty collection, and invalid requests, including oversized bodies. The store tests call it concurrently to exercise ID assignment and snapshot isolation. A lock around the ID increment alone would still leave the slice exposed. In our store the lock covers the whole append and publishes a complete record.

**Try it.** Add a `GET /books?read=true` filter. Specify the behavior for an invalid boolean and verify that filtering doesn't rearrange the list. Consider what changes if you implement pagination and allow deletion.
