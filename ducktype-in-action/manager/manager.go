package manager

import "fmt"

type Manager struct {
	Name string
}

func (m *Manager) Do(work string) {
	fmt.Printf("Manager %s %s\n", m.Name, work)
}
