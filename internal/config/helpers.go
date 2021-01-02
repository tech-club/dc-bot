package config

import (
	"fmt"
	"github.com/tech-club/dc-bot/pkg/db"
)

func (dc *dbConfig) GetDSN() string {
	switch dc.Driver {
	case db.DriverMySQL:
		return fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?multiStatements=true&parseTime=true&loc=UTC&collation=utf8mb4_general_ci",
			dc.Username,
			dc.Password,
			dc.Host,
			dc.Port,
			dc.Name,
		)
	}
	return ""
}
