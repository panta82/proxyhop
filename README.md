# ProxyHop

Super simple command line proxy, suitable for CORS busting

### Usage:

```
proxyhop -p 12345 https://api.someapi.com/v1
```

Then make requests against http://localhost:12345. They will be passed through to https://api.someapi.com/v1. Path and query string will be preserved (eg. http://localhost:12345/a/b/c will translate to https://api.someapi.com/v1/a/b/c)

The proxy will automatically handle your CORS headers, so you can call it from browser just fine.

That's all there is to it.

### Installation

Builds are available for Linux (amd64, x86) and macOS (amd64).

For an easy automated install, paste this into your terminal.

```
curl -sf -L https://raw.githubusercontent.com/panta82/proxyhop/master/misc/install.sh | sudo bash
```

If you don't trust the script, you can manually download one of the [releases](releases), make it executable and copy it somewhere in your `PATH`.

### License

MIT