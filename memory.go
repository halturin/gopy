// Copyright 2011 Julian Phillips.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

// #include "utils.h"
import "C"

// This file is about as unsafe as you can get ... we are playing tricks on the
// Go runtime in here. :(
import "unsafe"

import "sync"

// Some sizes that we need for various calculations
const (
	upSize   = unsafe.Sizeof(unsafe.Pointer(nil))
	headSize = C.sizeof_PyGC_Head
	baseSize = C.sizeof_PyObject
)

// We need to keep track of things that we have allocated, and the proxies that
// we have created, memLock must be locked whilst using these variables.
var (
	memLock sync.Mutex
	// allocated = make(map[uintptr][]unsafe.Pointer)
	gcProxies = make(map[uintptr]*C.PyObject)
)

func fromGc(g *C.PyGC_Head) *C.PyObject {
	p := unsafe.Pointer(g)
	o := unsafe.Pointer(uintptr(p) + headSize)
	return (*C.PyObject)(o)
}

func goGcMalloc(size uintptr) *C.PyObject {
	// first, lock memLock, and arrange for it to be unlocked on return
	memLock.Lock()
	defer memLock.Unlock()

	g := (*C.PyGC_Head)(unsafe.Pointer(_goMalloc(size + headSize)))
	C.setGcRefs(g, C._PyGC_REFS_UNTRACKED)
	p := fromGc(g)

	// We need to move the original tracked entry to be indexed by the offset
	// address from fromGc
	px := uintptr(unsafe.Pointer(p))
	// gx := uintptr(unsafe.Pointer(g))
	// TODO!
	// allocated[px] = allocated[gx]
	// delete(allocated, gx)

	// We can't access the internals of the GC Module to manipulate the
	// generation counts, so we have to use a proxy object instead.  We just
	// create a bare minimum object, initialise it, and then store it away to be
	// cleaned up later.
	proxy := C._PyObject_GC_Malloc(C.size_t(baseSize))
	if proxy == nil {
		// TODO!
		// delete(allocated, px)
		return nil
	}
	C.GoPyObject_INIT(proxy, c(BaseType))
	gcProxies[px] = proxy

	return p
}

// _goMalloc must be called with memLock already locked.
func _goMalloc(size uintptr) *C.PyObject {
	// We have to use []unsafe.Pointer instead of []byte, otherwise the Go
	// runtime will mark the memory as not containing pointers, and won't use it
	// to pin other Go allocations in memory - in which case we might as well
	// just stick with the Python allocator.
	// n := (size + upSize - 1) / upSize
	// s := make([]unsafe.Pointer, n)
	// p := unsafe.Pointer(&s[0])

	// Go 1.6: Cannot return pointers to Go memory anymore!
	// Instead we just keep a key into the [something] map and use that
	// to get Go memory
	p := C.malloc(C.size_t(size))

	// We need to keep a reference to the allocation, so that the Go runtime
	// doesn't free the memory before we are finished with it.
	// allocated[uintptr(p)] = s

	return (*C.PyObject)(p)
}

func goMalloc(size uintptr) *C.PyObject {
	// first, lock memLock, and arrange for it to be unlocked on return
	memLock.Lock()
	defer memLock.Unlock()

	// now call _goMalloc to do the actual work
	return _goMalloc(size)
}

//export goGenericAlloc
func goGenericAlloc(t unsafe.Pointer, n C.Py_ssize_t) unsafe.Pointer {
	var obj *C.PyObject

	typ := newType((*C.PyObject)(t))
	size := uintptr(C.var_size(c(typ), n))

	if typ.IsGc() {
		obj = goGcMalloc(size)
	} else {
		obj = goMalloc(size)
	}

	if obj == nil {
		raise(MemoryError.ErrV(None))
		return nil
	}

	if typ.o.tp_flags&C.Py_TPFLAGS_HEAPTYPE != 0 {
		typ.Incref()
	}

	if typ.o.tp_itemsize == 0 {
		C.GoPyObject_INIT(obj, c(typ))
	} else {
		C.GoPyObject_INIT_VAR(obj, c(typ), n)
	}

	if typ.IsGc() {
		C.GoPyObject_GC_Track(obj)
	}

	return unsafe.Pointer(obj)
}

//export goGenericFree
func goGenericFree(o unsafe.Pointer) {
	// first, lock memLock, and arrange for it to be unlocked on return
	memLock.Lock()
	defer memLock.Unlock()

	// Remove the entry from allocated, to let the Go runtime reclaim the memory
	// in the next GC run.
	// TODO!
	// delete(allocated, uintptr(o))

	// We need to also delete the proxy, so that the Python GC counts get
	// updated as appropriate.
	proxy, ok := gcProxies[uintptr(o)]
	if ok {
		C.PyObject_GC_Del(unsafe.Pointer(proxy))
		delete(gcProxies, uintptr(o))
	}
}
