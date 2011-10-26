package main

import (
  //"gob"
  "os"
  "sync"
  "log"
)

const (
  saveTimeout = 10e9
  saveQueueLength = 1000
)

type Store interface {
  Put(url, key *string) os.Error
  Get(key, url *string) os.Error
}

type MonitorStore struct {
  mu sync.RWMutex
  count int
  save chan record
}

type record struct {
  Key, URL string
}

func NewMonitorStore(filename string) *MonitorStore {
  s := &MonitorStore{}
  if filename != "" {
    s.save = make(chan record, saveQueueLength)
    if err := s.load(filename); err != nil {
      log.Println("URLStore:", err)
    }
    //go s.saveLoop(filename)
  }
  return s
}

func (s *MonitorStore) load(filename string) os.Error {
  return nil
}
