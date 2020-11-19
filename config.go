package orm

type DatabaseConfig struct {
	Host           string
	User           string
	Password       string
	Database       string
	Port           string
	MaxConnections int
}
