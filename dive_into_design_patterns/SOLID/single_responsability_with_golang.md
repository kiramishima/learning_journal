# Implementing Single Responsability with Golang

```go
// Original without Single Resp applied

type Employee struct {
	name string
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) printTimeSheetReport() {
	/// todo
}
```

We separate the responsability in 2 structs: Employee and TimesheetReport

```golang

type Employee struct {
	name string
}

func (e *Employee) getName() string {
	return e.name
}

// With Single Responsability applied
type TimeSheetReport struct {
	// anything
}

func (sheet *TimeSheetReport) print(employee Employee) {
	// now, we drop printTimeSheetReport method on Employee struct
}

```