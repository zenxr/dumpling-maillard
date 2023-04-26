# Dumpling Maillard

A dumpling becomes unsavory after too much Maillard reaction.

## ToDo

- [ ] containerization
- [ ] postgres integration, via container
- [x] go environment variables
    - postgres connection info
- [x] define schema
    - fields: app name, what environment restored from, tenants, date, and who restored
- [x] simple API interaction
    - hardcoded models at first, and useable for testing
    - todo: single entrypoint for each "router"
- [ ] basic web page showing postgres data, via template
- [ ] basic authentication
- [ ] tests for dumps


## Structure

- [x] Want multiple routers

## Notes

Resources: https://www.youtube.com/watch?v=gJx6gODwOwM

DB interaction:

Use `sql#DB.Exec` with placeholder params to build queries, properly applies escapes.

```go
email, loginTime := "human@example.com", time.Now()
result, err := db.Exec("INSERT INTO UserAccount VALUES ($1, $2)", email, loginTime)
if err != nil {
  panic(err)
}
```

```go
// Generic Example
func remove[T comparable](slice []T, i int) []T {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
```
