package todo_test

import (
	"os"
	"testing"

	"github.com/Fallenstedt/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, go %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
  l := todo.List{}

  taskName := "New test"
  l.Add(taskName)

  l.Complete(1)

  if !l[0].Done {
    t.Errorf("New task should be completed")
  }

}

func TestDelete (t *testing.T) {
  l := todo.List{}

  tasks := []string{"a", "b", "c"}
  for _, v := range tasks {
    l.Add(v)
  }

  l.Delete(2)

  if len(l) != 2 {
    t.Errorf("Expected list length %d, got %d instead", 2, len(l))
  }

  if l[1].Task != tasks[2] {
    t.Errorf("Expected %q got %q insead", tasks[2], l[1].Task)
  }
}

func TestUpdate (t *testing.T) {
  l := todo.List{}

  l.Add("Task name")

  l.Update(1, "Updated name")

  if l[0].Task != "Updated name" {
    t.Errorf("Expected %q got %q instead", "Updated name", l[0].Task)
  }
}

func TestSaveGet(t *testing.T) {
  l1 := todo.List{}
  l2 := todo.List{}
  taskName := "New Task"
  l1.Add(taskName)

  if l1[0].Task != taskName {
    t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
  }

  tf, err := os.CreateTemp("", "")
  if err != nil {
    t.Fatalf("Error creating temp file: %s", err)
  }

  defer os.Remove(tf.Name())
  if err := l1.Save(tf.Name()); err != nil {
    t.Fatalf("Error saving lisdescribedt to file: %s", err)
  }

  if err := l2.Get(tf.Name()); err != nil {
    t.Fatalf("Error getting list from file: %s", err)
  }

  if l1[0].Task != l2[0].Task {
    t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
  }

}