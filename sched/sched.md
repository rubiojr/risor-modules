import { Callout } from 'nextra/components';

# Sched

<Callout type="info" emoji="ℹ️">
This module is not included with Risor.
</Callout>

The `sched` module exposes a simple interface to schedule tasks, powered by [go-quartz](https://github.com/reugn/go-quartz) library.

## Functions

### cron

```go filename="Function signature"
cron(cronline string, fn func)
```

Creates a new cron job.

```go copy filename="Example"
// Run every second
s := sched.cron("*/1 * * * * *", func() {
	print("hello world!")
})
time.sleep(10)
```

### every

```go filename="Function signature"
every(cronline string, fn func)
```

Creates a new job that runs every second.

```go copy filename="Example"
// Run every second
s := sched.every(1, func() {
	print("hello world!")
})
time.sleep(10)
```

```go copy filename="Example"
// Run every 500 milliseconds
s := sched.every(0.5, func() {
	print("hello world!")
})
time.sleep(10)
```
