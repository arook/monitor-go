package main

import (
  //"gob"
  "os"
)

type Store interface {
  Put(url, key *string) os.Error
  Get(key, url *string) os.Error
}
