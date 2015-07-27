package migrations

func init() {
	Migrations.Add(1, "Create Expenses Table", `
		CREATE TABLE expenses (
			id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL
		)
	`)
}
