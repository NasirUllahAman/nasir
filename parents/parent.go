package parents

import "fmt"

func intit() {
	fmt.Println("father package is intlize")

}

type Father struct {
	Name string
}

func (f Father) Data(name string) string {
	f.Name = "father: " + name

	return f.Name

}
