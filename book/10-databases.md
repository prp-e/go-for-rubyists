# 10. Databases and transactions

Rails can get you from a model class to a working database before you've thought about a connection pool. In Go the standard library gives you `database/sql`, an API for working with pools, queries, and transactions. Add a driver for your database. Driver registration and connection strings depend on the driver. SQL syntax and placeholder forms depend on the database. The example service in this book uses memory, so you can read this chapter and run the rest of the book without installing a database.

## Open the pool once

A `*sql.DB` represents a pool suitable for concurrent use. Create it at startup, configure its limits for the deployment, and close it after the server stops. `sql.Open` may create the handle without connecting. Use `PingContext` with a startup deadline to check that the database is reachable. Watch pool wait counts and durations with `db.Stats()` before guessing why a query appears slow.

Here's a read using PostgreSQL placeholder syntax. The query text keeps the input separate from the SQL.

```go
func findBook(ctx context.Context, db *sql.DB, id int64) (Book, error) {
    var b Book
    err := db.QueryRowContext(ctx,
        "SELECT id, title, read FROM books WHERE id = $1", id,
    ).Scan(&b.ID, &b.Title, &b.Read)
    if err != nil {
        return Book{}, fmt.Errorf("find book %d: %w", id, err)
    }
    return b, nil
}
```

Check for a missing row with `errors.Is(err, sql.ErrNoRows)`. For many rows, use `QueryContext`, defer `rows.Close()`, scan inside `for rows.Next()`, and check `rows.Err()` after iteration. SQL NULL needs an explicit representation, such as `sql.NullString` at the database boundary. Don't let storage nullability leak into unrelated parts of the app without deciding what null means to them.

## Keep an operation on the same transaction

If two writes must succeed together, call `BeginTx`, defer rollback, issue both calls through the returned `*sql.Tx`, and check the result of `Commit`. A rollback after a successful commit has no effect. Mixing a pool call into the middle of a transaction can leave that call outside the transaction and even deadlock a small pool. For example, marking a book as read and inserting a corresponding activity entry should either both commit or both roll back.

Use constraints to enforce durable invariants such as uniqueness across simultaneous requests. Choose transaction isolation for the invariant you're protecting, and handle any required retry at the whole-transaction boundary. A read followed by a write does not automatically prevent another transaction acting on the same old value. Test real database behavior against the same engine you deploy, especially constraints and rollbacks. Fake repositories are good for checking HTTP decisions and cannot prove SQL semantics.

**Try it.** Sketch a query that lists books by increasing ID and a transaction that marks a book read and records when it happened. Identify what should happen if the second statement fails.
