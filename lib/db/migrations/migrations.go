package migrations

import (
	"fmt"

	"gotrix/lib/db"
)

type Migration struct {
	ID   int
	Name string
	Fn   MigrateFunc
}

type MigrateFunc func(con db.Con) error

func (m *Migration) String() string {
	return fmt.Sprintf("%d %s", m.ID, m.Name)
}

type Migrations []*Migration

func (m *Migrations) AddFunc(id int, name string, migration MigrateFunc) {
	*m = append(*m, &Migration{
		ID:   id,
		Name: name,
		Fn:   migration,
	})
}

func (m *Migrations) Add(id int, name string, migration string) {
	m.AddFunc(id, name, func(con db.Con) error {
		_, err := con.Exec(migration)
		return err
	})
}

func (m Migrations) Len() int {
	return len(m)
}

func (m Migrations) Less(i, j int) bool {
	if m[i].ID == m[j].ID {
		return m[i].Name < m[j].Name
	}

	return m[i].ID < m[j].ID
}

func (m Migrations) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
