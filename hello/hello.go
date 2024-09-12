package hello

import (
	"context"
	"fmt"

	"github.com/risor-io/risor/object"
)

func Require(funcName string, count int, args []object.Object) *object.Error {
	nArgs := len(args)
	if nArgs != count {
		if count == 1 {
			return object.Errorf(
				fmt.Sprintf("type error: %s() takes exactly 1 argument (%d given)",
					funcName, nArgs))
		}
		return object.Errorf(
			fmt.Sprintf("type error: %s() takes exactly %d arguments (%d given)",
				funcName, count, nArgs))
	}
	return nil
}

func World(ctx context.Context, args ...object.Object) object.Object {
	if err := Require("hello.world", 0, args); err != nil {
		return err
	}

	fmt.Println("Hello world!")
	return nil
}

func Module() *object.Module {
	return object.NewBuiltinsModule(
		"hello", map[string]object.Object{
			"world": object.NewBuiltin("world", World),
		},
	)
}
