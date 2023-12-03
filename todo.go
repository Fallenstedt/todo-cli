package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)



type (
  item struct {
    Task string
    Done bool
    CreatedAt time.Time
    CompletedAt time.Time
  }

  List []item

  IList interface {
    Add()
    Complete(i int) error
    Delete(i int) error
    Update(i int, task string) error
    Save(filename string) error
    Get(filename string) error
  }
)

func (l *List) Add(task string) {
  t := item {
    Task: task,
    Done: false,
    CreatedAt: time.Now(),
    CompletedAt: time.Time{},
  }
  *l = append(*l, t)
}

func (l *List) Complete(i int) error {
  ls := *l

  if err := l.indexCheck(i); err != nil {
    return err
  }

  ls[i-1].Done = true
  ls[i-1].CompletedAt = time.Now()

  return nil
}

func (l *List) Delete(i int) error {
  ls := *l

  if err := l.indexCheck(i); err != nil {
    return err
  }

  *l = append(ls[:i-1], ls[i:]...)
  return nil
}

func (l *List) Update(i int, task string) error {
  ls := *l

  if err := l.indexCheck(i); err != nil {
    return err
  }
  ls[i-1].Task = task

  *l = ls

  return nil
}

func (l* List) Save(filename string) error {
  js, err := json.Marshal(l)
  if err != nil {
    return err
  }
  return os.WriteFile(filename, js, 0644)
}


func (l *List) Get(filename string) error {
  file, err := os.ReadFile(filename)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      return nil
    }
  }
  if len(file) == 0 {
    return nil
  }
  return json.Unmarshal(file, l)
}

func (l *List) String() string {
  formatted := ""

  for k, t := range *l {
    prefix := " "
    if t.Done {
      prefix = "X "
    }

    formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
  }

  return formatted
}

func (l *List) indexCheck(i int) error {
  ls := *l
  if i<= 0 || i > len(ls) {
    return fmt.Errorf("item %d does not exist", i)
  }
  return nil
}