package add

type Person struct {
	Name string
	Age  int
}

func Add(a *Person) int {
	return a.Age
}
