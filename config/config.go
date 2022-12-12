package config

// DBConnection contains the connection details to the db
// in production ready service, these details would be sent through environment variables
type DBConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func GetDBConnection() DBConnection {
	return DBConnection{
		Host:     "localhost",
		Port:     55000,
		User:     "postgres",
		Password: "Aa123456",
		DBName:   "onboarding_db",
	}
}
