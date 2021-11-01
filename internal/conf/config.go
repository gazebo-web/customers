package conf

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

// Database contains the config for initializing an SQL database.
type Database struct {
	// Username is the database username.
	Username string `env:"CUSTOMERS_DATABASE_USERNAME,required"`

	// Password is the database password.
	Password string `env:"CUSTOMERS_DATABASE_PASSWORD,required"`

	// Host is host on which the SQL server instance is running.
	Host string `env:"CUSTOMERS_DATABASE_HOST,required"`

	// Port is the TPC/IP network port on which the target SQL server is listening for connections.
	Port uint `env:"CUSTOMERS_DATABASE_PORT,required"`

	// Name is the name of the database for the connection.
	Name string `env:"CUSTOMERS_DATABASE_NAME,required"`

	// Charset is the name of the set of characters that are legal in a string.
	// Defaults to UTF-8.
	Charset string `env:"CUSTOMERS_DATABASE_CHARSET" envDefault:"utf8"`
}

// ToDSN converts the Database config into a valid MySQL data source name.
func (db Database) ToDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		db.Username, db.Password, db.Host, db.Port, db.Name, db.Charset,
	)
}

// Parse fills Database data from an external source.
func (db *Database) Parse() error {
	return env.Parse(db)
}

// Config contains the needed config to start the Customers HTTP server.
type Config struct {
	// Database contains the configuration needed to open an SQL connection.
	Database Database
}

// Parse fills Config data from an external source.
func (c *Config) Parse() error {
	if err := c.Database.Parse(); err != nil {
		return err
	}
	if err := env.Parse(c); err != nil {
		return err
	}
	return nil
}
