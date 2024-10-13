package sched

import (
	"context"
	"fmt"
	"time"

	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/logger"
	"github.com/reugn/go-quartz/quartz"
	"github.com/risor-io/risor/object"
)

func Module() *object.Module {
	return object.NewBuiltinsModule(
		"sched", map[string]object.Object{
			"cron":  object.NewBuiltin("sched", Cron),
			"every": object.NewBuiltin("line", Every),
		},
	)
}

// Cron schedules a function to run at a specific time using a cron like expression.
//
// The first argument is the cron expression and the second argument is the function to be scheduled.
func Cron(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.Errorf("missing arguments, 2 required")
	}

	cronLine, err := object.AsString(args[0])
	if err != nil {
		return err
	}

	cfunc, ok := args[1].(*object.Function)
	if !ok {
		return object.Errorf("expected function, got %s", args[1].Type())
	}

	sched := quartz.NewStdScheduler()
	sched.Start(ctx)

	cronTrigger, cerr := quartz.NewCronTrigger(cronLine)
	if cerr != nil {
		return object.Errorf("failed to create cron trigger: %v", cerr)
	}

	//functionJob := job.NewFunctionJob(func(ctx context.Context) error { return cfunc.Call(ctx) })
	functionJob := job.NewFunctionJob(func(_ context.Context) (any, error) {
		res := cfunc.Call(ctx)
		if res.Type() == object.ERROR {
			return nil, res.(*object.Error)
		}

		return res, nil
	})

	// register jobs to scheduler
	sched.ScheduleJob(
		quartz.NewJobDetail(functionJob, quartz.NewJobKey(cronLine)),
		cronTrigger,
	)

	return &schedObj{sched: sched}
}

// Every schedules a function to run every n seconds.
//
// The first argument is the interval in seconds (float)
// The second argument is the function to be scheduled
func Every(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.Errorf("missing arguments, 2 required")
	}

	secs, err := object.AsFloat(args[0])
	if err != nil {
		return err
	}

	cfunc, ok := args[1].(*object.Function)
	if !ok {
		return object.Errorf("expected function, got %s", args[1].Type())
	}

	sched := quartz.NewStdScheduler()
	sched.Start(ctx)

	functionJob := job.NewFunctionJob(func(_ context.Context) (any, error) {
		res := cfunc.Call(ctx)
		if res.Type() == object.ERROR {
			return nil, res.(*object.Error)
		}

		return res, nil
	})

	sched.ScheduleJob(
		quartz.NewJobDetail(functionJob, quartz.NewJobKey(fmt.Sprintf("every-%f", secs))),
		quartz.NewSimpleTrigger(time.Duration(secs)*time.Second),
	)

	return &schedObj{sched: sched}
}

func init() {
	// quartz lib logs by default unfortunately, disable it here
	logger.SetDefault(logger.NewSimpleLogger(nil, logger.LevelOff))
}
