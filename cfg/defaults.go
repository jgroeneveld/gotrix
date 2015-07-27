package cfg

var configDefaults = map[string]config{
	"production": config{
		ApplicationName: "gotrix",
		DatabaseURL:     "-",
		Port:            "",
	},
	"development": config{
		ApplicationName: "gotrix_development",
		DatabaseURL:     "postgres://127.0.0.1:5432/gotrix_development?sslmode=disable",
		Port:            "3000",
	},
	"test": config{
		ApplicationName: "gotrix_test",
		DatabaseURL:     "postgres://127.0.0.1:5432/gotrix_test?sslmode=disable",
		Port:            "",
	},
}
