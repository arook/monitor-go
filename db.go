package monitor

import (
  "os"
  "fmt"
  "gosqlite.googlecode.com/hg/sqlite"
)

type Asin struct {
  id int
  asin string
}

type Fetching struct {
  asin *Asin
  dt string
}

type Result struct {
  fetching *Fetching
  seller string
  price float64
  buybox bool
}

type MonitorStore struct {
  conn *sqlite.Conn
}

func NewMonitorStore(filename string) *MonitorStore {
  store := &MonitorStore{}
  conn, err := sqlite.Open(filename)
  if err != nil {
    fmt.Println("Unable to open database: %s", err)
    os.Exit(1);
  }

  //create table
  conn.Exec("CREATE TABLE IF NOT EXISTS asin(id INTEGER PRIMARY KEY AUTOINCREMENT, asin VARCHAR(200));")

  store.conn = conn
  return store
}

func (s *MonitorStore) AddAsin(asin string) {
  insertSql := `insert into asin (asin) values ("` + asin + `");`
  s.conn.Exec(insertSql)
}

func (s *MonitorStore) GetAsinList() []string {
  list := make([]string, 20)
  stmt, err := s.conn.Prepare("select * from Asin limit 10;")
  err = stmt.Exec()
  if err != nil {
    fmt.Println("Error While Selecting: %s", err)
  }

  for stmt.Next() {
    var asin Asin

    err = stmt.Scan(&asin.id, &asin.asin)
    if err != nil {
      fmt.Printf("Error while getting row data:%s\n", err)
      os.Exit(1)
    }
fmt.Println(asin.id, asin.asin)
    list[asin.id] = asin.asin
  }
  return list
}
