package sched

import (
	"context"

	"github.com/reugn/go-quartz/quartz"
	"github.com/risor-io/risor/errz"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type schedObj struct {
	sched quartz.Scheduler
}

func (b *schedObj) SetAttr(name string, value object.Object) error {
	return errz.TypeErrorf("type error: object has no attribute %q", name)
}

func (b *schedObj) IsTruthy() bool {
	return true
}

func (b *schedObj) Cost() int {
	return 0
}

func (b *schedObj) Equals(other object.Object) object.Object {
	so, ok := other.(*schedObj)
	if !ok {
		return object.False
	}
	ok = (*b == *so)
	return object.NewBool(ok)
}

func (b *schedObj) Inspect() string {
	return "sched"
}

func (b *schedObj) Type() object.Type {
	return object.Type("sched")
}

func (b *schedObj) Interface() any {
	return b.sched
}

func (r *schedObj) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.TypeErrorf("type error: unsupported operation for sched: %v", opType)
}

func (s *schedObj) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "stop":
		return object.NewBuiltin("sched.stop", func(ctx context.Context, args ...object.Object) object.Object {
			s.sched.Stop()
			s.sched.Wait(ctx)
			return nil
		}), true
	}

	return nil, false
}
