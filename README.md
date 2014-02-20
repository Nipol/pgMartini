
pgMartini
=========

Martini middleware for lib/pg (https://github.com/lib/pq)

## Usage
pgMartini uses [lib/pg](http://godoc.org/github.com/lib/pq)
package to [database/sql](http://golang.org/pkg/database/sql/)

~~~ go
// main.go
package main

import (
  "github.com/codegangsta/martini"
  "github.com/Nipol/pgMartini"
  "database/sql"
)

func main() {
  m := martini.Classic()

  m.Use(pgmartini.PG("user_id", "password", "address", "port", "dbname", "sslmode"))

  m.Get("/", func(db *sql.DB) string {
    var plus string
    db.QueryRow("SELECT 1+1 as plus").Scan(&plus)
    return plus
  })
}
~~~