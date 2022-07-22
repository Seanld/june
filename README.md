# June

A static file server for the [Gemini](https://gemini.circumlunar.space) protocol.

June uses the [Houston](https://git.sr.ht/~seanld/houston) framework under the hood.

## Minimal

- Literally just serves static files from a tree-like directory.
- **Configurable rate-limiting** via token-bucket algorithm.
- Written in Go, which means a lightweight binary, with wide system compatibility.
- By leveraging [Houston](https://git.sr.ht/~seanld/houston) to do the heavy-lifting, there really isn't much code in June.
