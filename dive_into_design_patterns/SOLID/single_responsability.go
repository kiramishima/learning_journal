package solid

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

// With Single Responsability applied
type TimeSheetReport struct {
	// anything
}

func (sheet *TimeSheetReport) print(employee Employee) {
	// now, we can remove printTimeSheetReport method in Employee struct
}
