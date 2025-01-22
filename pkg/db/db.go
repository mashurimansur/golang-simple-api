package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct {
	DB       *sql.DB
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewMySQLConnection(host string, port string, username, password, database string) *MySQLConnection {
	return &MySQLConnection{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func (m *MySQLConnection) Connect() error {
	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)

	// Open database connection
	var err error
	m.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}

	// Configure connection pool settings
	m.DB.SetMaxOpenConns(25)
	m.DB.SetMaxIdleConns(25)
	m.DB.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection with ping
	if err = m.DB.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	return nil
}

func (m *MySQLConnection) Close() error {
	if m.DB != nil {
		return m.DB.Close()
	}
	return nil
}
