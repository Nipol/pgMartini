package pgmartini

import (
  "log"
  "strings"

  _ "github.com/lib/pq"
  "database/sql"
  

  "github.com/codegangsta/martini"
)

func PG(id, password, address, port, dbname, sslmode string) martini.Handler {
  turl := []string{"postgres://", id, ":", password, "@", address, ":", port, "/", dbname, "?sslmode=", sslmode}
  url  := strings.Join(turl, "")

  return func(c martini.Context) {
    db, err := sql.Open("postgres", url)
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close();
    c.Map(db)
    c.Next()
  }
}