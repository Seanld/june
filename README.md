# June

A static file server for the [Gemini](https://gemini.circumlunar.space) protocol.

June uses the [Houston](https://git.sr.ht/~seanld/houston) framework under the hood.

## Minimal

- Literally just serves static files from a tree-like directory.
- **Configurable rate-limiting** via token-bucket algorithm.
- Written in Go, which means a lightweight binary, with wide system compatibility.
- By leveraging [Houston](https://git.sr.ht/~seanld/houston) to do the heavy-lifting, there really isn't much code in June.

## Usage

You should have a certificate file and a key file (CA-signed, or self-signed) ready to go, and a directory that contains content you'd like to serve over Gemini.

Assuming your directory structure looked like this:

```
|-> june
|-> cert/
|---> main.crt
|---> main.key
|-> static/
|---> index.gmi
```

You would run:

```
./june -crt='cert/main.crt' -key='cert/main.key'
```

And that will start up June on `0.0.0.0:1965`.

## On Rate-limiting

As previously stated, Houston (and by extension June) use a token-bucket algorithm to manage rate-limiting. It essentially functions by creating a bucket for each new client address that makes a request. For every single request made, a token is taken from the bucket. The default settings are set up to where no more than two requests can be made in a single second. And every second, 2 tokens are added back into the bucket. These two parameters can be adjusted by changing the `-bs` and `-trickle` flags.

In simpler terms, the default parameters mean that only two requests can be made per client per second.

The token-bucket algorithm itself helps with general spam and DoS attacks, but it does not help with DDoS attacks.
