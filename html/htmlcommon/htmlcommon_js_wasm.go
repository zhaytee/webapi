// Code generated by webidl-bind. DO NOT EDIT.

package htmlcommon

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// domcore.Event

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// enum: WorkerType
type WorkerType int

const (
	ClassicWorkerType WorkerType = iota
	ModuleWorkerType
)

var workerTypeToWasmTable = []string{
	"classic", "module",
}

var workerTypeFromWasmTable = map[string]WorkerType{
	"classic": ClassicWorkerType, "module": ModuleWorkerType,
}

// JSValue is converting this enum into a java object
func (this *WorkerType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this WorkerType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(workerTypeToWasmTable) {
		return workerTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// WorkerTypeFromJS is converting a javascript value into
// a WorkerType enum value.
func WorkerTypeFromJS(value js.Value) WorkerType {
	key := value.String()
	conv, ok := workerTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: FrameRequestCallback
type FrameRequestCallbackFunc func(time float64)

// FrameRequestCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FrameRequestCallback js.Func

func FrameRequestCallbackToJS(callback FrameRequestCallbackFunc) *FrameRequestCallback {
	if callback == nil {
		return nil
	}
	ret := FrameRequestCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 float64 // javascript: double time
		)
		_p0 = (args[0]).Float()
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func FrameRequestCallbackFromJS(_value js.Value) FrameRequestCallbackFunc {
	return func(time float64) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := time
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: OnBeforeUnloadEventHandlerNonNull
type OnBeforeUnloadEventHandlerFunc func(event *domcore.Event) *string

// OnBeforeUnloadEventHandler is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type OnBeforeUnloadEventHandler js.Func

func OnBeforeUnloadEventHandlerToJS(callback OnBeforeUnloadEventHandlerFunc) *OnBeforeUnloadEventHandler {
	if callback == nil {
		return nil
	}
	ret := OnBeforeUnloadEventHandler(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *domcore.Event // javascript: Event event
		)
		_p0 = domcore.EventFromJS(args[0])
		_returned := callback(_p0)
		_converted := _returned
		return _converted
	}))
	return &ret
}

func OnBeforeUnloadEventHandlerFromJS(_value js.Value) OnBeforeUnloadEventHandlerFunc {
	return func(event *domcore.Event) (_result *string) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := event.JSValue()
		_args[0] = _p0
		_end++
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *string // javascript: DOMString
		)
		if _returned.Type() != js.TypeNull {
			__tmp := (_returned).String()
			_converted = &__tmp
		}
		_result = _converted
		return
	}
}

// callback: OnErrorEventHandlerNonNull
type OnErrorEventHandlerFunc func(event *Union, source *string, lineno *uint, colno *uint, _error js.Value) interface{}

// OnErrorEventHandler is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type OnErrorEventHandler js.Func

func OnErrorEventHandlerToJS(callback OnErrorEventHandlerFunc) *OnErrorEventHandler {
	if callback == nil {
		return nil
	}
	ret := OnErrorEventHandler(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Union   // javascript: Union event
			_p1 *string  // javascript: DOMString source
			_p2 *uint    // javascript: unsigned long lineno
			_p3 *uint    // javascript: unsigned long colno
			_p4 js.Value // javascript: any _error
		)
		_p0 = UnionFromJS(args[0])
		if len(args) > 1 {
			__tmp := (args[1]).String()
			_p1 = &__tmp
		}
		if len(args) > 2 {
			__tmp := (uint)((args[2]).Int())
			_p2 = &__tmp
		}
		if len(args) > 3 {
			__tmp := (uint)((args[3]).Int())
			_p3 = &__tmp
		}
		if len(args) > 4 {
			_p4 = args[4]
		}
		_returned := callback(_p0, _p1, _p2, _p3, _p4)
		_converted := _returned
		return _converted
	}))
	return &ret
}

func OnErrorEventHandlerFromJS(_value js.Value) OnErrorEventHandlerFunc {
	return func(event *Union, source *string, lineno *uint, colno *uint, _error js.Value) (_result interface{}) {
		var (
			_args [5]interface{}
			_end  int
		)
		_p0 := event.JSValue()
		_args[0] = _p0
		_end++
		if source != nil {
			_p1 := source
			_args[1] = _p1
			_end++
		}
		if lineno != nil {
			_p2 := lineno
			_args[2] = _p2
			_end++
		}
		if colno != nil {
			_p3 := colno
			_args[3] = _p3
			_end++
		}
		if _error.Type() != js.TypeUndefined {
			_p4 := _error
			_args[4] = _p4
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted js.Value // javascript: any
		)
		_converted = _returned
		_result = _converted
		return
	}
}

// interface: BeforeUnloadEvent
type BeforeUnloadEvent struct {
	domcore.Event
}

// BeforeUnloadEventFromJS is casting a js.Wrapper into BeforeUnloadEvent.
func BeforeUnloadEventFromJS(value js.Wrapper) *BeforeUnloadEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &BeforeUnloadEvent{}
	ret.Value_JS = input
	return ret
}

// ReturnValue returning attribute 'returnValue' with
// type string (idl: DOMString).
func (_this *BeforeUnloadEvent) ReturnValue() string {
	var ret string
	value := _this.Value_JS.Get("returnValue")
	ret = (value).String()
	return ret
}

// SetReturnValue setting attribute 'returnValue' with
// type string (idl: DOMString).
func (_this *BeforeUnloadEvent) SetReturnValue(value string) {
	input := value
	_this.Value_JS.Set("returnValue", input)
}