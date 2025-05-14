package eval

import (
	"fmt"

	"github.com/psidh/Ganges/src/object"
)

var builtins = map[string]*object.Builtin{
	"dairghya": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `dairghya` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"pratham": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be Array. got=%s", args[0].Type())
			}

			arr := args[0].(*object.Array)

			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"antha": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be Array. got=%s", args[0].Type())
			}

			arr := args[0].(*object.Array)

			if len(arr.Elements) > 0 {
				return arr.Elements[len(arr.Elements)-1]
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)

			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"set": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			s := &object.Set{Elements: make(map[object.HashKey]object.Object)}

			for _, arg := range args {
				hashable, ok := arg.(object.Hashable)
				if !ok {
					return newError("unusable as hash key: %s", arg.Type())
				}

				key := hashable.HashKey()
				s.Elements[key] = arg
			}
			return s
		},
	},
	"has": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			setObj, ok := args[0].(*object.Set)
			if !ok {
				return newError("first argument must be SET, got %s", args[0].Type())
			}

			hashable, ok := args[1].(object.Hashable)

			if !ok {
				return newError("unusable as hash key: %s", args[1].Type())
			}

			_, exists := setObj.Elements[hashable.HashKey()]

			if exists {
				return SATYA
			}
			return ASATYA

		},
	},
	"add": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, wanted=2", len(args))
			}

			setObj, ok := args[0].(*object.Set)

			if !ok {
				return newError("first argument must be SET, got %s", args[0].Type())
			}

			hashable, ok := args[1].(object.Hashable)

			if !ok {
				return newError("unusable as hash key: %s", args[1].Type())
			}

			key := hashable.HashKey()

			setObj.Elements[key] = args[1]
			return setObj
		},
	},
	"remove": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			setObj, ok := args[0].(*object.Set)
			if !ok {
				return newError("first argument must be SET, got %s", args[0].Type())
			}
			hashable, ok := args[1].(object.Hashable)
			if !ok {
				return newError("unusable as hash key: %s", args[1].Type())
			}
			key := hashable.HashKey()
			delete(setObj.Elements, key)
			return setObj
		},
	},
}
