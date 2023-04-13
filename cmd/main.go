package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/otaviotech/quode/configs"
	"github.com/otaviotech/quode/internal/infra/web/webserver"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	noError(err)

	conn, err := getDbConnection(cfg)
	noError(err)

	bootstrapWebServer(conn, cfg.WebServerPort)
}

func bootstrapWebServer(db *sql.DB, port string) {
	ws := webserver.NewWebserver(port)

	ws.AddHandler(NewCreateAuthorHandler(db))
	ws.AddHandler(NewCreateSubjectHandler(db))

	ws.Start()
}

func getDbConnection(cfg *configs.Conf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s", cfg.DBDriver, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBArgs)
	return sql.Open(cfg.DBDriver, dsn)
}

func noError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
