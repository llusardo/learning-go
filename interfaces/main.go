package main

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

// Engineer Type
type Engineer struct {
	Name string
}

// Language implementation
func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

// Age implementation
func (e Engineer) Age() int {
	return 25
}

// Random implementation
func (e Engineer) Random() (string, error) {
	return "random", nil
}

func main() {
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	programmers = append(programmers, elliot)
}
