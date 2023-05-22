// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
)

func TestSpanLinkSlice(t *testing.T) {
	es := NewSpanLinkSlice()
	assert.Equal(t, 0, es.Len())
	es = newSpanLinkSlice(&[]*otlptrace.Span_Link{})
	assert.Equal(t, 0, es.Len())

	emptyVal := NewSpanLink()
	testVal := generateTestSpanLink()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestSpanLink(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestSpanLinkSlice_CopyTo(t *testing.T) {
	dest := NewSpanLinkSlice()
	// Test CopyTo to empty
	NewSpanLinkSlice().CopyTo(dest)
	assert.Equal(t, NewSpanLinkSlice(), dest)

	// Test CopyTo larger slice
	generateTestSpanLinkSlice().CopyTo(dest)
	assert.Equal(t, generateTestSpanLinkSlice(), dest)

	// Test CopyTo same size slice
	generateTestSpanLinkSlice().CopyTo(dest)
	assert.Equal(t, generateTestSpanLinkSlice(), dest)
}

func TestSpanLinkSlice_EnsureCapacity(t *testing.T) {
	es := generateTestSpanLinkSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestSpanLinkSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestSpanLinkSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestSpanLinkSlice(), es)
}

func TestSpanLinkSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestSpanLinkSlice()
	dest := NewSpanLinkSlice()
	src := generateTestSpanLinkSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestSpanLinkSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestSpanLinkSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestSpanLinkSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanLinkSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanLinkSlice()
	emptySlice.RemoveIf(func(el SpanLink) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestSpanLinkSlice()
	pos := 0
	filtered.RemoveIf(func(el SpanLink) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpanLinkSlice_Sort(t *testing.T) {
	es := generateTestSpanLinkSlice()
	es.Sort(func(a, b SpanLink) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) < uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b SpanLink) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) > uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestSpanLinkSlice() SpanLinkSlice {
	es := NewSpanLinkSlice()
	fillTestSpanLinkSlice(es)
	return es
}

func fillTestSpanLinkSlice(es SpanLinkSlice) {
	*es.orig = make([]*otlptrace.Span_Link, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlptrace.Span_Link{}
		fillTestSpanLink(newSpanLink((*es.orig)[i]))
	}
}
