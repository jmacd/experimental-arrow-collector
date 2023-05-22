// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package plog

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
)

func TestLogRecordSlice(t *testing.T) {
	es := NewLogRecordSlice()
	assert.Equal(t, 0, es.Len())
	es = newLogRecordSlice(&[]*otlplogs.LogRecord{})
	assert.Equal(t, 0, es.Len())

	emptyVal := NewLogRecord()
	testVal := generateTestLogRecord()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestLogRecord(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestLogRecordSlice_CopyTo(t *testing.T) {
	dest := NewLogRecordSlice()
	// Test CopyTo to empty
	NewLogRecordSlice().CopyTo(dest)
	assert.Equal(t, NewLogRecordSlice(), dest)

	// Test CopyTo larger slice
	generateTestLogRecordSlice().CopyTo(dest)
	assert.Equal(t, generateTestLogRecordSlice(), dest)

	// Test CopyTo same size slice
	generateTestLogRecordSlice().CopyTo(dest)
	assert.Equal(t, generateTestLogRecordSlice(), dest)
}

func TestLogRecordSlice_EnsureCapacity(t *testing.T) {
	es := generateTestLogRecordSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestLogRecordSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestLogRecordSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestLogRecordSlice(), es)
}

func TestLogRecordSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestLogRecordSlice()
	dest := NewLogRecordSlice()
	src := generateTestLogRecordSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestLogRecordSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestLogRecordSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestLogRecordSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestLogRecordSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewLogRecordSlice()
	emptySlice.RemoveIf(func(el LogRecord) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestLogRecordSlice()
	pos := 0
	filtered.RemoveIf(func(el LogRecord) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestLogRecordSlice_Sort(t *testing.T) {
	es := generateTestLogRecordSlice()
	es.Sort(func(a, b LogRecord) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) < uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b LogRecord) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) > uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestLogRecordSlice() LogRecordSlice {
	es := NewLogRecordSlice()
	fillTestLogRecordSlice(es)
	return es
}

func fillTestLogRecordSlice(es LogRecordSlice) {
	*es.orig = make([]*otlplogs.LogRecord, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlplogs.LogRecord{}
		fillTestLogRecord(newLogRecord((*es.orig)[i]))
	}
}
