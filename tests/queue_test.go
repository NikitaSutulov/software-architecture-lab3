package tests

import (
	"testing"

	"github.com/NikitaSutulov/software-architecture-lab3/painter"
)

func TestMessageQueueEmpty(t *testing.T) {
	mq := &painter.MessageQueue{}
	if !mq.Empty() {
		t.Errorf("expected empty message queue, got non-empty")
	}
}

func TestMessageQueuePush(t *testing.T) {
	mq := &painter.MessageQueue{}
	op1 := new(MockOperation)
	mq.Push(op1)
	if mq.Empty() {
		t.Errorf("expected non-empty message queue, got empty")
	}
}

func TestMessageQueuePull(t *testing.T) {
	mq := &painter.MessageQueue{}
	op1 := new(MockOperation)
	mq.Push(op1)
	op2 := mq.Pull()
	if op1 != op2 {
		t.Errorf("expected %v, got %v", op1, op2)
	}
	if !mq.Empty() {
		t.Errorf("expected empty message queue after pull, got non-empty")
	}
}
