package engineer

import "fmt"

type Engineer struct {
	Name string
}

func (m *Engineer) Do(work string) {
	fmt.Printf("Engineer %s %s\n", m.Name, work)
}
