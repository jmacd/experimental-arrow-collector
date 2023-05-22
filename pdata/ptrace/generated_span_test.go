// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestSpan_MoveTo(t *testing.T) {
	ms := generateTestSpan()
	dest := NewSpan()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpan(), ms)
	assert.Equal(t, generateTestSpan(), dest)
}

func TestSpan_CopyTo(t *testing.T) {
	ms := NewSpan()
	orig := NewSpan()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSpan()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSpan_TraceID(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.TraceID(data.TraceID([16]byte{})), ms.TraceID())
	testValTraceID := pcommon.TraceID(data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestSpan_SpanID(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.SpanID(data.SpanID([8]byte{})), ms.SpanID())
	testValSpanID := pcommon.SpanID(data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func TestSpan_TraceState(t *testing.T) {
	ms := NewSpan()
	internal.FillTestTraceState(internal.TraceState(ms.TraceState()))
	assert.Equal(t, pcommon.TraceState(internal.GenerateTestTraceState()), ms.TraceState())
}

func TestSpan_ParentSpanID(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.SpanID(data.SpanID([8]byte{})), ms.ParentSpanID())
	testValParentSpanID := pcommon.SpanID(data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetParentSpanID(testValParentSpanID)
	assert.Equal(t, testValParentSpanID, ms.ParentSpanID())
}

func TestSpan_Name(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, "", ms.Name())
	ms.SetName("test_name")
	assert.Equal(t, "test_name", ms.Name())
}

func TestSpan_Kind(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, SpanKind(otlptrace.Span_SpanKind(0)), ms.Kind())
	testValKind := SpanKind(otlptrace.Span_SpanKind(3))
	ms.SetKind(testValKind)
	assert.Equal(t, testValKind, ms.Kind())
}

func TestSpan_StartTimestamp(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.Timestamp(0), ms.StartTimestamp())
	testValStartTimestamp := pcommon.Timestamp(1234567890)
	ms.SetStartTimestamp(testValStartTimestamp)
	assert.Equal(t, testValStartTimestamp, ms.StartTimestamp())
}

func TestSpan_EndTimestamp(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.Timestamp(0), ms.EndTimestamp())
	testValEndTimestamp := pcommon.Timestamp(1234567890)
	ms.SetEndTimestamp(testValEndTimestamp)
	assert.Equal(t, testValEndTimestamp, ms.EndTimestamp())
}

func TestSpan_Attributes(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpan_DroppedAttributesCount(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}

func TestSpan_Events(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, NewSpanEventSlice(), ms.Events())
	fillTestSpanEventSlice(ms.Events())
	assert.Equal(t, generateTestSpanEventSlice(), ms.Events())
}

func TestSpan_DroppedEventsCount(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, uint32(0), ms.DroppedEventsCount())
	ms.SetDroppedEventsCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedEventsCount())
}

func TestSpan_Links(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, NewSpanLinkSlice(), ms.Links())
	fillTestSpanLinkSlice(ms.Links())
	assert.Equal(t, generateTestSpanLinkSlice(), ms.Links())
}

func TestSpan_DroppedLinksCount(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, uint32(0), ms.DroppedLinksCount())
	ms.SetDroppedLinksCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedLinksCount())
}

func TestSpan_Status(t *testing.T) {
	ms := NewSpan()
	fillTestStatus(ms.Status())
	assert.Equal(t, generateTestStatus(), ms.Status())
}

func generateTestSpan() Span {
	tv := NewSpan()
	fillTestSpan(tv)
	return tv
}

func fillTestSpan(tv Span) {
	tv.orig.TraceId = data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.SpanId = data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
	internal.FillTestTraceState(internal.NewTraceState(&tv.orig.TraceState))
	tv.orig.ParentSpanId = data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.Name = "test_name"
	tv.orig.Kind = otlptrace.Span_SpanKind(3)
	tv.orig.StartTimeUnixNano = 1234567890
	tv.orig.EndTimeUnixNano = 1234567890
	internal.FillTestMap(internal.NewMap(&tv.orig.Attributes))
	tv.orig.DroppedAttributesCount = uint32(17)
	fillTestSpanEventSlice(newSpanEventSlice(&tv.orig.Events))
	tv.orig.DroppedEventsCount = uint32(17)
	fillTestSpanLinkSlice(newSpanLinkSlice(&tv.orig.Links))
	tv.orig.DroppedLinksCount = uint32(17)
	fillTestStatus(newStatus(&tv.orig.Status))
}
