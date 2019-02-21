// Code generated by webidl-bind. DO NOT EDIT.

package htmlevent

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/patch"
)

// using following types:
// domcore.Event
// javascript.Promise
// patch.FormData

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

// dictionary: TrackEventInit
type TrackEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Track      *Union
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *TrackEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Track.JSValue()
	out.Set("track", value3)
	return out
}

// TrackEventInitFromJS is allocating a new
// TrackEventInit object and copy all values from
// input javascript object
func TrackEventInitFromJS(value js.Wrapper) *TrackEventInit {
	input := value.JSValue()
	var out TrackEventInit
	var (
		value0 bool   // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool   // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool   // javascript: boolean {composed Composed composed}
		value3 *Union // javascript: Union {track Track track}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	if input.Get("track").Type() != js.TypeNull {
		value3 = UnionFromJS(input.Get("track"))
	}
	out.Track = value3
	return &out
}

// dictionary: FormDataEventInit
type FormDataEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	FormData   *patch.FormData
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FormDataEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.FormData.JSValue()
	out.Set("formData", value3)
	return out
}

// FormDataEventInitFromJS is allocating a new
// FormDataEventInit object and copy all values from
// input javascript object
func FormDataEventInitFromJS(value js.Wrapper) *FormDataEventInit {
	input := value.JSValue()
	var out FormDataEventInit
	var (
		value0 bool            // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool            // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool            // javascript: boolean {composed Composed composed}
		value3 *patch.FormData // javascript: FormData {formData FormData formData}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = patch.FormDataFromJS(input.Get("formData"))
	out.FormData = value3
	return &out
}

// dictionary: PopStateEventInit
type PopStateEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	State      js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PopStateEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.State
	out.Set("state", value3)
	return out
}

// PopStateEventInitFromJS is allocating a new
// PopStateEventInit object and copy all values from
// input javascript object
func PopStateEventInitFromJS(value js.Wrapper) *PopStateEventInit {
	input := value.JSValue()
	var out PopStateEventInit
	var (
		value0 bool     // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool     // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool     // javascript: boolean {composed Composed composed}
		value3 js.Value // javascript: any {state State state}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = input.Get("state")
	out.State = value3
	return &out
}

// dictionary: HashChangeEventInit
type HashChangeEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	OldURL     string
	NewURL     string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *HashChangeEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.OldURL
	out.Set("oldURL", value3)
	value4 := _this.NewURL
	out.Set("newURL", value4)
	return out
}

// HashChangeEventInitFromJS is allocating a new
// HashChangeEventInit object and copy all values from
// input javascript object
func HashChangeEventInitFromJS(value js.Wrapper) *HashChangeEventInit {
	input := value.JSValue()
	var out HashChangeEventInit
	var (
		value0 bool   // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool   // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool   // javascript: boolean {composed Composed composed}
		value3 string // javascript: USVString {oldURL OldURL oldURL}
		value4 string // javascript: USVString {newURL NewURL newURL}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("oldURL")).String()
	out.OldURL = value3
	value4 = (input.Get("newURL")).String()
	out.NewURL = value4
	return &out
}

// dictionary: PageTransitionEventInit
type PageTransitionEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Persisted  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PageTransitionEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Persisted
	out.Set("persisted", value3)
	return out
}

// PageTransitionEventInitFromJS is allocating a new
// PageTransitionEventInit object and copy all values from
// input javascript object
func PageTransitionEventInitFromJS(value js.Wrapper) *PageTransitionEventInit {
	input := value.JSValue()
	var out PageTransitionEventInit
	var (
		value0 bool // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool // javascript: boolean {composed Composed composed}
		value3 bool // javascript: boolean {persisted Persisted persisted}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("persisted")).Bool()
	out.Persisted = value3
	return &out
}

// dictionary: ErrorEventInit
type ErrorEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Message    string
	Filename   string
	Lineno     uint
	Colno      uint
	Error      js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ErrorEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Message
	out.Set("message", value3)
	value4 := _this.Filename
	out.Set("filename", value4)
	value5 := _this.Lineno
	out.Set("lineno", value5)
	value6 := _this.Colno
	out.Set("colno", value6)
	value7 := _this.Error
	out.Set("error", value7)
	return out
}

// ErrorEventInitFromJS is allocating a new
// ErrorEventInit object and copy all values from
// input javascript object
func ErrorEventInitFromJS(value js.Wrapper) *ErrorEventInit {
	input := value.JSValue()
	var out ErrorEventInit
	var (
		value0 bool     // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool     // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool     // javascript: boolean {composed Composed composed}
		value3 string   // javascript: DOMString {message Message message}
		value4 string   // javascript: USVString {filename Filename filename}
		value5 uint     // javascript: unsigned long {lineno Lineno lineno}
		value6 uint     // javascript: unsigned long {colno Colno colno}
		value7 js.Value // javascript: any {error Error _error}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("message")).String()
	out.Message = value3
	value4 = (input.Get("filename")).String()
	out.Filename = value4
	value5 = (uint)((input.Get("lineno")).Int())
	out.Lineno = value5
	value6 = (uint)((input.Get("colno")).Int())
	out.Colno = value6
	value7 = input.Get("error")
	out.Error = value7
	return &out
}

// dictionary: PromiseRejectionEventInit
type PromiseRejectionEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Promise    *javascript.Promise
	Reason     js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PromiseRejectionEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Promise.JSValue()
	out.Set("promise", value3)
	value4 := _this.Reason
	out.Set("reason", value4)
	return out
}

// PromiseRejectionEventInitFromJS is allocating a new
// PromiseRejectionEventInit object and copy all values from
// input javascript object
func PromiseRejectionEventInitFromJS(value js.Wrapper) *PromiseRejectionEventInit {
	input := value.JSValue()
	var out PromiseRejectionEventInit
	var (
		value0 bool                // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                // javascript: boolean {composed Composed composed}
		value3 *javascript.Promise // javascript: Promise {promise Promise promise}
		value4 js.Value            // javascript: any {reason Reason reason}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = javascript.PromiseFromJS(input.Get("promise"))
	out.Promise = value3
	value4 = input.Get("reason")
	out.Reason = value4
	return &out
}

// interface: TrackEvent
type TrackEvent struct {
	domcore.Event
}

// TrackEventFromJS is casting a js.Wrapper into TrackEvent.
func TrackEventFromJS(value js.Wrapper) *TrackEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &TrackEvent{}
	ret.Value_JS = input
	return ret
}

func NewTrackEvent(_type string, eventInitDict *TrackEventInit) (_result *TrackEvent) {
	_klass := js.Global().Get("TrackEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *TrackEvent // javascript: TrackEvent _what_return_name
	)
	_converted = TrackEventFromJS(_returned)
	_result = _converted
	return
}

// Track returning attribute 'track' with
// type Union (idl: Union).
func (_this *TrackEvent) Track() *Union {
	var ret *Union
	value := _this.Value_JS.Get("track")
	if value.Type() != js.TypeNull {
		ret = UnionFromJS(value)
	}
	return ret
}

// interface: FormDataEvent
type FormDataEvent struct {
	domcore.Event
}

// FormDataEventFromJS is casting a js.Wrapper into FormDataEvent.
func FormDataEventFromJS(value js.Wrapper) *FormDataEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FormDataEvent{}
	ret.Value_JS = input
	return ret
}

func NewFormDataEvent(_type string, eventInitDict *FormDataEventInit) (_result *FormDataEvent) {
	_klass := js.Global().Get("FormDataEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FormDataEvent // javascript: FormDataEvent _what_return_name
	)
	_converted = FormDataEventFromJS(_returned)
	_result = _converted
	return
}

// FormData returning attribute 'formData' with
// type patch.FormData (idl: FormData).
func (_this *FormDataEvent) FormData() *patch.FormData {
	var ret *patch.FormData
	value := _this.Value_JS.Get("formData")
	ret = patch.FormDataFromJS(value)
	return ret
}

// interface: PopStateEvent
type PopStateEvent struct {
	domcore.Event
}

// PopStateEventFromJS is casting a js.Wrapper into PopStateEvent.
func PopStateEventFromJS(value js.Wrapper) *PopStateEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PopStateEvent{}
	ret.Value_JS = input
	return ret
}

func NewPopStateEvent(_type string, eventInitDict *PopStateEventInit) (_result *PopStateEvent) {
	_klass := js.Global().Get("PopStateEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *PopStateEvent // javascript: PopStateEvent _what_return_name
	)
	_converted = PopStateEventFromJS(_returned)
	_result = _converted
	return
}

// State returning attribute 'state' with
// type Any (idl: any).
func (_this *PopStateEvent) State() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("state")
	ret = value
	return ret
}

// interface: HashChangeEvent
type HashChangeEvent struct {
	domcore.Event
}

// HashChangeEventFromJS is casting a js.Wrapper into HashChangeEvent.
func HashChangeEventFromJS(value js.Wrapper) *HashChangeEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &HashChangeEvent{}
	ret.Value_JS = input
	return ret
}

func NewHashChangeEvent(_type string, eventInitDict *HashChangeEventInit) (_result *HashChangeEvent) {
	_klass := js.Global().Get("HashChangeEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *HashChangeEvent // javascript: HashChangeEvent _what_return_name
	)
	_converted = HashChangeEventFromJS(_returned)
	_result = _converted
	return
}

// OldURL returning attribute 'oldURL' with
// type string (idl: USVString).
func (_this *HashChangeEvent) OldURL() string {
	var ret string
	value := _this.Value_JS.Get("oldURL")
	ret = (value).String()
	return ret
}

// NewURL returning attribute 'newURL' with
// type string (idl: USVString).
func (_this *HashChangeEvent) NewURL() string {
	var ret string
	value := _this.Value_JS.Get("newURL")
	ret = (value).String()
	return ret
}

// interface: PageTransitionEvent
type PageTransitionEvent struct {
	domcore.Event
}

// PageTransitionEventFromJS is casting a js.Wrapper into PageTransitionEvent.
func PageTransitionEventFromJS(value js.Wrapper) *PageTransitionEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PageTransitionEvent{}
	ret.Value_JS = input
	return ret
}

func NewPageTransitionEvent(_type string, eventInitDict *PageTransitionEventInit) (_result *PageTransitionEvent) {
	_klass := js.Global().Get("PageTransitionEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *PageTransitionEvent // javascript: PageTransitionEvent _what_return_name
	)
	_converted = PageTransitionEventFromJS(_returned)
	_result = _converted
	return
}

// Persisted returning attribute 'persisted' with
// type bool (idl: boolean).
func (_this *PageTransitionEvent) Persisted() bool {
	var ret bool
	value := _this.Value_JS.Get("persisted")
	ret = (value).Bool()
	return ret
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

// interface: ErrorEvent
type ErrorEvent struct {
	domcore.Event
}

// ErrorEventFromJS is casting a js.Wrapper into ErrorEvent.
func ErrorEventFromJS(value js.Wrapper) *ErrorEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ErrorEvent{}
	ret.Value_JS = input
	return ret
}

func NewErrorEvent(_type string, eventInitDict *ErrorEventInit) (_result *ErrorEvent) {
	_klass := js.Global().Get("ErrorEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ErrorEvent // javascript: ErrorEvent _what_return_name
	)
	_converted = ErrorEventFromJS(_returned)
	_result = _converted
	return
}

// Message returning attribute 'message' with
// type string (idl: DOMString).
func (_this *ErrorEvent) Message() string {
	var ret string
	value := _this.Value_JS.Get("message")
	ret = (value).String()
	return ret
}

// Filename returning attribute 'filename' with
// type string (idl: USVString).
func (_this *ErrorEvent) Filename() string {
	var ret string
	value := _this.Value_JS.Get("filename")
	ret = (value).String()
	return ret
}

// Lineno returning attribute 'lineno' with
// type uint (idl: unsigned long).
func (_this *ErrorEvent) Lineno() uint {
	var ret uint
	value := _this.Value_JS.Get("lineno")
	ret = (uint)((value).Int())
	return ret
}

// Colno returning attribute 'colno' with
// type uint (idl: unsigned long).
func (_this *ErrorEvent) Colno() uint {
	var ret uint
	value := _this.Value_JS.Get("colno")
	ret = (uint)((value).Int())
	return ret
}

// Error returning attribute 'error' with
// type Any (idl: any).
func (_this *ErrorEvent) Error() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("error")
	ret = value
	return ret
}

// interface: PromiseRejectionEvent
type PromiseRejectionEvent struct {
	domcore.Event
}

// PromiseRejectionEventFromJS is casting a js.Wrapper into PromiseRejectionEvent.
func PromiseRejectionEventFromJS(value js.Wrapper) *PromiseRejectionEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseRejectionEvent{}
	ret.Value_JS = input
	return ret
}

func NewPromiseRejectionEvent(_type string, eventInitDict *PromiseRejectionEventInit) (_result *PromiseRejectionEvent) {
	_klass := js.Global().Get("PromiseRejectionEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := eventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *PromiseRejectionEvent // javascript: PromiseRejectionEvent _what_return_name
	)
	_converted = PromiseRejectionEventFromJS(_returned)
	_result = _converted
	return
}

// Promise returning attribute 'promise' with
// type javascript.Promise (idl: Promise).
func (_this *PromiseRejectionEvent) Promise() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("promise")
	ret = javascript.PromiseFromJS(value)
	return ret
}

// Reason returning attribute 'reason' with
// type Any (idl: any).
func (_this *PromiseRejectionEvent) Reason() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("reason")
	ret = value
	return ret
}
