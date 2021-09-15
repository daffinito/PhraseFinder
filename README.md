## Three word phrase finding graphQL API

Returns an array of the 100 most common three word phrases, ordered from most common to least common phrase

Created with https://github.com/99designs/gqlgen

To start the server:

```
go run server.go
```

Then browse to http://localhost:8080 for the GraphQL Playground.

The playground doesn't support file uploads, so currently to test parsing phrases from a file, use curl:

```
curl localhost:8080/query \
  -F operations='{ "query": "mutation ($file: Upload!) { findPhrasesFromFile(file: $file) { text, count } }", "variables": { "file": null } }' \
  -F map='{ "0": ["variables.file"] }' \
  -F 0=@/path/to/file.txt
```