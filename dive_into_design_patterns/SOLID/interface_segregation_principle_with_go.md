Before

```go
type CloudProvider interface {
    storeFile(name string)
    getFile(name string)
    createServer(region string)
    listServer(region string)
    getCDNAddress()
}

type Amazon struct {
    CloudProvider
}

type Dropbox struct {
    CloudProvider
}
```

After

```go
type CloudHostingProvider interface {
    createServer(region string)
    listServer(region string)
}

type CDNProvider interface {
    getCDNAddress()
}

type CloudStorageProvider interface {
    storeFile(name string)
    getFile(name string)
}

type Amazon struct {
    CloudHostingProvider
    CDNProvider
    CloudStorageProvider
}

type Dropbox struct {
    CloudStorageProvider
}
```