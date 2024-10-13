package sched

import (
	"context"
	"testing"
	"time"

	"github.com/risor-io/risor/object"
	"github.com/stretchr/testify/assert"
)

func TestCron(t *testing.T) {
	ctx := context.Background()

	var executed bool
	callFn := func(ctx context.Context, fn *object.Function, args []object.Object) (object.Object, error) {
		executed = true
		return nil, nil
	}
	ctx = object.WithCallFunc(ctx, callFn)
	var fn *object.Function

	// Schedule the function using a cron expression
	cronExpr := object.NewString("*/1 * * * * *") // Every second
	sched := Cron(ctx, cronExpr, fn)
	assert.False(t, object.IsError(sched))
	assert.NotNil(t, sched)
	assert.Equal(t, "sched", string(sched.Type()))

	// Wait for a few seconds to allow the cron job to execute
	assert.False(t, executed)
	time.Sleep(1 * time.Second)
	assert.True(t, executed)

	// Stop the scheduler
	stopFn, ok := sched.GetAttr("stop")
	assert.True(t, ok)
	stopFn.(*object.Builtin).Call(ctx)
}

func TestEvery(t *testing.T) {
	ctx := context.Background()

	var executed bool
	callFn := func(ctx context.Context, fn *object.Function, args []object.Object) (object.Object, error) {
		executed = true
		return nil, nil
	}
	ctx = object.WithCallFunc(ctx, callFn)
	var fn *object.Function

	// Schedule the function to run every 1 seconds
	interval := object.NewFloat(1)
	sched := Every(ctx, interval, fn)
	assert.False(t, object.IsError(sched))
	assert.Equal(t, "sched", string(sched.Type()))

	// Wait for a few seconds to allow the job to execute
	time.Sleep(100 * time.Millisecond)
	assert.False(t, executed)
	time.Sleep(1 * time.Second)
	assert.True(t, executed)

	stopFn, ok := sched.GetAttr("stop")
	assert.True(t, ok)
	stopFn.(*object.Builtin).Call(ctx)
}

func TestEqual(t *testing.T) {
	callFn := func(ctx context.Context, fn *object.Function, args []object.Object) (object.Object, error) {
		return nil, nil
	}
	ctx := object.WithCallFunc(context.Background(), callFn)
	var fn *object.Function

	cronExpr := object.NewString("* * * * * *")
	sched := Cron(ctx, cronExpr, fn)
	sched2 := Cron(ctx, cronExpr, fn)

	assert.True(t, sched.Equals(sched).(*object.Bool).Value())
	assert.False(t, sched.Equals(sched2).(*object.Bool).Value())
}

func TestInvalidCron(t *testing.T) {
	callFn := func(ctx context.Context, fn *object.Function, args []object.Object) (object.Object, error) {
		return nil, nil
	}
	ctx := object.WithCallFunc(context.Background(), callFn)
	var fn *object.Function

	// Schedule the function using an invalid cron expression
	cronExpr := object.NewString("invalid-cron")
	sched := Cron(ctx, cronExpr, fn)

	assert.True(t, object.IsError(sched))
}
