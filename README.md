# ProxyHop

Super simple command line proxy, suitable for CORS busting

Usage:

```
proxyhop -p 12345 https://api.someapi.com/v1
```

Then make requests against `localhost:12345`. Path will be appended (eg. http://localhost:12345/a/b/c will translate to https://api.someapi.com/v1/a/b/c)

Query string and headers should all work.

That's all there is to it.

##### NOTE: Still WIP