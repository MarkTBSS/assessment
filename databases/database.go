package databases

import "os"

func DatabaseURL() string {
	return "host=" + os.Getenv("DATABASE_HOST") +
		" port=" + os.Getenv("DATABASE_PORT") +
		" user=" + os.Getenv("DATABASE_USERNAME") +
		" password=" + os.Getenv("DATABASE_PASSWORD") +
		" dbname=" + os.Getenv("DATABASE_DATABASE") +
		" sslmode=" + os.Getenv("DATABASE_SSL_MODE")
}
