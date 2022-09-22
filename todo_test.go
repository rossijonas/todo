package todo

import (
	"testing"

	"github.com/rossijonas/todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "NewTask"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "NewTask"
	l.add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}
