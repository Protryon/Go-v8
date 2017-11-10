package v8

/*
#include "v8_wrap.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "runtime"

// A compiled JavaScript script.
//
type Script struct {
	self unsafe.Pointer
}

// Compiles the specified script (context-independent).
// 'data' is the Pre-parsing data, as obtained by PreCompile()
// using pre_data speeds compilation if it's done multiple times.
//
func (e *Engine) Compile(code []byte, origin *ScriptOrigin) *Script {
	var originPtr unsafe.Pointer

	if origin != nil {
		originPtr = origin.self
	}

	codePtr := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&code)).Data)
	self := C.V8_Compile(e.self, (*C.char)(codePtr), C.int(len(code)), originPtr)

	if self == nil {
		return nil
	}

	result := &Script{
		self: self,
	}

	runtime.SetFinalizer(result, func(s *Script) {
		if traceDispose {
			println("v8.Script.Dispose()", s.self)
		}
		C.V8_DisposeScript(s.self)
	})

	return result
}

// Runs the script returning the resulting value.
//
func (s *Script) Run() *Value {
	return newValue(C.V8_Script_Run(s.self))
}

// The origin, within a file, of a script.
//
type ScriptOrigin struct {
	self         unsafe.Pointer
	Name         string
	LineOffset   int
	ColumnOffset int
}

func (e *Engine) NewScriptOrigin(name string, lineOffset, columnOffset int) *ScriptOrigin {
	namePtr := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&name)).Data)
	self := C.V8_NewScriptOrigin(e.self, (*C.char)(namePtr), C.int(len(name)), C.int(lineOffset), C.int(columnOffset))

	if self == nil {
		return nil
	}

	result := &ScriptOrigin{
		self:         self,
		Name:         name,
		LineOffset:   lineOffset,
		ColumnOffset: columnOffset,
	}

	runtime.SetFinalizer(result, func(so *ScriptOrigin) {
		if traceDispose {
			println("v8.ScriptOrigin.Dispose()")
		}
		C.V8_DisposeScriptOrigin(so.self)
	})

	return result
}
