package postgres

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/jpynicolas/hello-world/pkg/config"
	_ "github.com/lib/pq"
)

const connString = "postgres://{{.User}}:{{.Pass}}@{{Host}}:{{Port}}/{{DbName}}@sslmode=disable"

func buildConnectionStringOrPanic(cnf config.Postgres) string {
	sb := strings.Builder{}
	tmp := template.Must(template.New("ConnString").Parse(connString))
	if err := tmp.Execute(&sb, cnf); err != nil {
		panic(err)
	}
	return sb.String()
}
func NewPostgres(cnf config.Postgres) (*sql.DB, error) {
	conn := buildConnectionStringOrPanic(cnf)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Failed to ping the database")
		return nil, fmt.Errorf("can not ping the database: %w", err)
	}
	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)
	return db, nil
}
