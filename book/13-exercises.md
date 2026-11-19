# Appendix A. Exercise notes

## 1–3. Get comfortable with the small stuff

For Hello World, `fmt.Println("go", "ruby")` prints a space between the words and appends a newline. `fmt.Printf("%s/%s\n", "go", "ruby")` prints a slash instead of the space. This is also a good time to see what `go vet` says about a missing format argument.

For FizzBuzz, parse the single argument with `strconv.Atoi`, check its error and the bounds, then call `label` from the loop. Avoid a loop that can overflow by using `for i := range limit` and calling `label(i + 1)`. Fifteen should exercise the combined branch. One should exercise only the default branch.

For word counts, increment `counts[word]++` after allocating the map. Use an ordinary slice of keys and `sort.Strings(keys)`. For the aliasing exercise, start with `source := []int{1, 2, 3}`, slice `left := source[:2]`, then append 99 to `left`. You should see 99 at index 2 of source as well. Set a full slice bound with `source[:2:2]` and repeat. Now append requires new backing storage.

## 4–7. Follow the contracts

For line counting, trim each token with `strings.TrimSpace` before testing it. Count only when it has content. Keep the final scanner error check after the loop. A read failure is a different result from an empty file.

For Rename, validate into a temporary string before assigning `b.Title`. A map element can't be addressed to satisfy the pointer receiver. Use `b := byID[id]`, call `b.Rename`, then write back on success. Check the lookup boolean first.

For interfaces, implement a small type with a `Write` method that returns zero and a sentinel error. Check that `errors.Is` still matches. For generics, declare `type ISBN string` and use a `map[ISBN]string`. An untyped constant can be inferred as an ISBN at the call site. An existing string variable can be converted with `ISBN(raw)`.

For testing FizzBuzz, use cases with distinct names or let the subtest name include the input. A case for 30 helps check the combined branch beyond the first multiple.

## 8–10. Think about boundaries

For concurrency, put job submission in a helper or use a labeled break to stop submitting on cancellation. Either way, close the jobs channel and call `wg.Wait()` before returning. Prefer a controllable blocking worker in the test to guessing how long to sleep. Document whether a canceled call can deliver partial results.

For routing, `GET /books/42` reaches the handler. `GET /books` without a collection route returns 404. `POST /books/42` returns 405. The wildcard extracts text, so a nonnumeric ID will still reach the handler unless you parse it there.

For the database, use `ORDER BY id` for the list. Mark-read and activity insertion must use the same transaction object. If insertion fails, return the error and let rollback discard both changes. A concurrent service also needs a policy for activity on repeat requests.

## 11–12. Exercise the whole service

For filtering, check whether the query parameter is present before parsing it. If you accept only the strings true and false, avoid accidentally accepting more spellings. Filter an existing snapshot in one pass to preserve its order. Include tests for no filter, false, true, and invalid text.

For deployment, the signal context indicates that the process should stop serving. The fresh shutdown context gives in-flight requests a bounded amount of time to finish. Using the canceled signal context for both makes the shutdown budget expire immediately.


## 13–14. Test the web app through the router

For the health route, call `w.WriteHeader(http.StatusNoContent)` without writing a body. Build an `httptest` request for GET, then one for POST. Use the actual Chi router so the test exercises method matching.

For blog search with literal matching, try SQLite's `instr(title, ?) > 0` with the search string as the bound argument. It treats percent and underscore as text. Decide how an empty query should behave and test the empty result, multiple matches, and a title containing a literal percent.
