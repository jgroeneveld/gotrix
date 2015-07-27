package migrations

import (
	"sort"
	"strconv"

	"github.com/jgroeneveld/gotrix/lib/db"
	"github.com/jgroeneveld/gotrix/lib/errors"
	"github.com/jgroeneveld/gotrix/lib/logger"
)

func Exec(l logger.Logger, con db.Con, migrations Migrations) error {
	err := setupMigrationsTable(con)
	if err != nil {
		return err
	}

	executed, err := getExecutedMigrations(con)
	if err != nil {
		return err
	}

	cnt := 0
	sort.Sort(migrations)
	for _, migration := range migrations {
		if executed.WasExecuted(migration.ID, migration.Name) {
			l.Printf("Skipping %q", migration)
			continue
		}
		l.Printf("Executing %q", migration)
		err := migration.Fn(con)
		if err != nil {
			return err
		}
		err = markMigrationExecuted(con, migration.ID, migration.Name)
		if err != nil {
			return err
		}
		cnt++
	}

	l.Printf("Executed %d migrations", cnt)
	return nil
}

func setupMigrationsTable(con db.Con) error {
	_, err := con.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER NOT NULL,
			name VARCHAR NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			PRIMARY KEY (id, name)
		)

	`)
	return errors.Wrap(err)
}

func markMigrationExecuted(con db.Con, id int, name string) error {
	_, err := con.Exec("INSERT INTO migrations (id, name) VALUES ($1, $2)", id, name)
	return errors.Wrap(err)
}

func getExecutedMigrations(con db.Con) (executedMigrations, error) {
	list := make(executedMigrations)

	rows, err := con.Query("SELECT id, name FROM migrations")
	if err != nil {
		return nil, errors.Wrap(err)
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		list.Add(id, name)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err)
	}

	return list, nil
}

type executedMigrations map[string]bool

func (m executedMigrations) Add(id int, name string) {
	m[strconv.Itoa(id)+name] = true
}
func (m executedMigrations) WasExecuted(id int, name string) bool {
	return m[strconv.Itoa(id)+name]
}
