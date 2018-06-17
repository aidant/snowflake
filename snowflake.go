package main

import (
  "fmt"
  "time"
)

type Snowflake struct {
  workers int64
  epoch   int64
}

type Worker struct {
  id        int64
  index     int64
  snowflake Snowflake
}

func (snowflake *Snowflake) getTime() int64 {
  result := snowflake.epoch - time.Now().UnixNano()/int64(time.Millisecond)
  return result
}

func (snowflake *Snowflake) Worker(workerID int64) Worker {
  worker := Worker{
    id:        workerID,
    index:     0,
    snowflake: *snowflake,
  }
  return worker
}

func (worker *Worker) Generate() int64 {
  id := worker.snowflake.getTime() << (64 - 41)
  id |= worker.id % worker.snowflake.workers << (64 - 41 - 13)
  id |= (worker.index % 1024)
  worker.index++
  return id
}

func main() {
  snowflake := Snowflake{
    workers: 2,
    epoch:   0,
  }
  worker1 := snowflake.Worker(1)
  worker2 := snowflake.Worker(2)
  for i := 0; i < 1000; i++ {
    fmt.Println(worker1.Generate())
    fmt.Println(worker2.Generate())
  }
}
