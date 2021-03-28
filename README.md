# Consultar todo de la tabla addresses

```go
storage.NewPostgresDB()
storageAddress := storage.NewPsqlAddress(storage.Pool())
serviceAddress := address.NewService(storageAddress)
ms, err := serviceAddress.GetAll()

if err != nil {
    log.Fatalf("Fatal en storage.GetAll(): %v", err)
}

log.Printf("%s", ms)
```
