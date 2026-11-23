# cashier: a sample project structure

```
cashier/
├── cmd/cashier/main.go      the program you run, it only wires things together
└── internal/
    ├── product/             product catalog (domain)
    └── order/               creates orders, needs the catalog through a small interface
```

- A thin `main` uses the `run() error` pattern.
- `order` defines its own `Catalog` interface, holding only the method it uses.
- `internal/` keeps these packages from being imported outside the `cashier` folder.

```bash
go run ./06-struktur-proyek/cashier/cmd/cashier
go test ./06-struktur-proyek/...
```
