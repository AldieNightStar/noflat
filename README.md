# No flat - library
* Uses _flat_ data such as `get` `set` by keys to put collections and groups of data

# Init
```go
nf := noflat.Init(DataMap)
```

# Group
```go
G := nf.Group("name")
G.Get("name")
G.Set("name", "Hello!")
```

# SubGroups
```go

```