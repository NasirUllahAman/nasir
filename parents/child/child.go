package child

import "fmt"

func init() {
	fmt.Println("son package is intialized")

}

type Son struct {
	Name string
}

func (s Son) Data(name string) string {
	s.Name = "son: " + name
	return s.Name

}
