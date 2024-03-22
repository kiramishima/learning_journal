Before

```go
type BudgetReport struct {
    database MySQLDatabase
}

type MySQLDatabase struct {

}
```

After

```go
type Database interface {
    insert()
    update()
    delete()
}

type BudgetReport struct {
    database Database
}

type MySQL struct {}
var _ Database = (*MySQL)(nil) // Debemos implementar sus metodos, solo como demostracion no lo haremos :P

type MongoDB struct {}
var _ Database = (*MongoDB)(nil) // Debemos implementar sus metodos
```