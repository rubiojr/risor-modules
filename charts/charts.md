import { Callout } from 'nextra/components';

# Charts

<Callout type="info" emoji="ℹ️">
This module is not included with Risor.
</Callout>

The `charts` module exposes a simple interface to create charts, powered by the great [go-echarts](https://github.com/go-echarts/go-echarts) library.

## Functions

### bar

```go filename="Function signature"
bar(file string, data map, opts map)
```

Creates a new bar chart.

```go copy filename="Example"
data := {
  "serie A": ["1", "2", "3"],
  "serie B": ["3", "4", "5"],
}

charts.bar(
	"bar.html",
	data,
)
```

The `opts` argument may be a map containing any of the following keys:

| Name   | Type                          | Description                              |
| ------ | ----------------------------- | ---------------------------------------- |
| title  | string                        | The title of the chart                   |
| subtitle | string                      | The subtitle of the chart                |
| xlabels | []string                      | The labels for the x-axis                |

```go copy filename="Example"
data := {
  "serie A": ["1", "2", "3"],
  "serie B": ["3", "4", "5"],
}

charts.bar(
	"bar.html",
	data,
	{
		title: "My awesome bar chart",
		subtitle: "this is a subtitle",
		xlabels: ["one", "two", "three"]
	},
)
```

### line

```go filename="Function signature"
line(file string, data map, opts map)
```

Creates a new line chart.

```go copy filename="Example"
data := {
  "serie A": ["1", "2", "3"],
  "serie B": ["3", "4", "5"],
}

charts.line(
	"line.html",
	data,
)
```

The `opts` argument may be a map containing any of the following keys:

| Name   | Type                          | Description                              |
| ------ | ----------------------------- | ---------------------------------------- |
| title  | string                        | The title of the chart                   |
| subtitle | string                      | The subtitle of the chart                |
| xlabels | []string                      | The labels for the x-axis                |


```go copy filename="Example"
data := {
  "serie A": ["1", "2", "3"],
  "serie B": ["3", "4", "5"],
}

charts.line(
	"line.html",
	data,
	{
		title: "My awesome line chart",
		subtitle: "this is a subtitle",
		xlabels: ["one", "two", "three"]
	},
)
```
