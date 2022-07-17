![license](https://img.shields.io/github/license/SKKUGoon/senateStock)

# Senate Stock Watcher

<p>
Toy project for redis cloud database usage.
</p>

```bash
$ go mod init senateStock
$ go mod download

or 

$ go mod tidy
```

```bash
$ docker build -t skkugoon/senatestock:0.0.0 .

or 

$ docker build --platform amd64 -t skkugoon/senatestock:0.0.0 .
```

## The use of redis
<p>
Directory dataObject stores getting JSON data with the use of go-redis/redis package
</p>