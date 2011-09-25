// Copyright 2011 Julian Phillips.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is automatically generated.  To regenerate:
//   ./gen_exc.py > exc.go

package py

// #include "utils.h"
import "C"

type _exc struct {
	BaseException             *BaseException
	Exception                 *BaseException
	StopIteration             *BaseException
	GeneratorExit             *BaseException
	StandardError             *BaseException
	ArithmeticError           *BaseException
	LookupError               *BaseException
	AssertionError            *BaseException
	AttributeError            *BaseException
	EOFError                  *BaseException
	FloatingPointError        *BaseException
	EnvironmentError          *BaseException
	IOError                   *BaseException
	OSError                   *BaseException
	ImportError               *BaseException
	IndexError                *BaseException
	KeyError                  *BaseException
	KeyboardInterrupt         *BaseException
	MemoryError               *BaseException
	NameError                 *BaseException
	OverflowError             *BaseException
	RuntimeError              *BaseException
	NotImplementedError       *BaseException
	SyntaxError               *BaseException
	IndentationError          *BaseException
	TabError                  *BaseException
	ReferenceError            *BaseException
	SystemError               *BaseException
	SystemExit                *BaseException
	TypeError                 *BaseException
	UnboundLocalError         *BaseException
	UnicodeError              *BaseException
	UnicodeEncodeError        *BaseException
	UnicodeDecodeError        *BaseException
	UnicodeTranslateError     *BaseException
	ValueError                *BaseException
	ZeroDivisionError         *BaseException
	BufferError               *BaseException
	MemoryErrorInst           *BaseException
	RecursionErrorInst        *BaseException
	Warning                   *BaseException
	UserWarning               *BaseException
	DeprecationWarning        *BaseException
	PendingDeprecationWarning *BaseException
	SyntaxWarning             *BaseException
	RuntimeWarning            *BaseException
	FutureWarning             *BaseException
	ImportWarning             *BaseException
	UnicodeWarning            *BaseException
	BytesWarning              *BaseException
}

func _get_exceptions() _exc {
	return _exc{
		newException(C.PyExc_BaseException),
		newException(C.PyExc_Exception),
		newException(C.PyExc_StopIteration),
		newException(C.PyExc_GeneratorExit),
		newException(C.PyExc_StandardError),
		newException(C.PyExc_ArithmeticError),
		newException(C.PyExc_LookupError),
		newException(C.PyExc_AssertionError),
		newException(C.PyExc_AttributeError),
		newException(C.PyExc_EOFError),
		newException(C.PyExc_FloatingPointError),
		newException(C.PyExc_EnvironmentError),
		newException(C.PyExc_IOError),
		newException(C.PyExc_OSError),
		newException(C.PyExc_ImportError),
		newException(C.PyExc_IndexError),
		newException(C.PyExc_KeyError),
		newException(C.PyExc_KeyboardInterrupt),
		newException(C.PyExc_MemoryError),
		newException(C.PyExc_NameError),
		newException(C.PyExc_OverflowError),
		newException(C.PyExc_RuntimeError),
		newException(C.PyExc_NotImplementedError),
		newException(C.PyExc_SyntaxError),
		newException(C.PyExc_IndentationError),
		newException(C.PyExc_TabError),
		newException(C.PyExc_ReferenceError),
		newException(C.PyExc_SystemError),
		newException(C.PyExc_SystemExit),
		newException(C.PyExc_TypeError),
		newException(C.PyExc_UnboundLocalError),
		newException(C.PyExc_UnicodeError),
		newException(C.PyExc_UnicodeEncodeError),
		newException(C.PyExc_UnicodeDecodeError),
		newException(C.PyExc_UnicodeTranslateError),
		newException(C.PyExc_ValueError),
		newException(C.PyExc_ZeroDivisionError),
		newException(C.PyExc_BufferError),
		newException(C.PyExc_MemoryErrorInst),
		newException(C.PyExc_RecursionErrorInst),
		newException(C.PyExc_Warning),
		newException(C.PyExc_UserWarning),
		newException(C.PyExc_DeprecationWarning),
		newException(C.PyExc_PendingDeprecationWarning),
		newException(C.PyExc_SyntaxWarning),
		newException(C.PyExc_RuntimeWarning),
		newException(C.PyExc_FutureWarning),
		newException(C.PyExc_ImportWarning),
		newException(C.PyExc_UnicodeWarning),
		newException(C.PyExc_BytesWarning),
	}
}