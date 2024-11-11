// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "wasi:blobstore/types@0.2.0-draft".
//
// Types used by blobstore
package types

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/protochron/http-password-checker-go/gen/wasi/io/v0.2.1/streams"
)

// InputStream represents the imported type alias "wasi:blobstore/types@0.2.0-draft#input-stream".
//
// See [streams.InputStream] for more information.
type InputStream = streams.InputStream

// OutputStream represents the imported type alias "wasi:blobstore/types@0.2.0-draft#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// ContainerName represents the string "wasi:blobstore/types@0.2.0-draft#container-name".
//
// name of a container, a collection of objects.
// The container name may be any valid UTF-8 string.
//
//	type container-name = string
type ContainerName string

// ObjectName represents the string "wasi:blobstore/types@0.2.0-draft#object-name".
//
// name of an object within a container
// The object name may be any valid UTF-8 string.
//
//	type object-name = string
type ObjectName string

// Timestamp represents the u64 "wasi:blobstore/types@0.2.0-draft#timestamp".
//
// TODO: define timestamp to include seconds since
// Unix epoch and nanoseconds
// https://github.com/WebAssembly/wasi-blob-store/issues/7
//
//	type timestamp = u64
type Timestamp uint64

// ObjectSize represents the u64 "wasi:blobstore/types@0.2.0-draft#object-size".
//
// size of an object, in bytes
//
//	type object-size = u64
type ObjectSize uint64

// Error represents the string "wasi:blobstore/types@0.2.0-draft#error".
//
//	type error = string
type Error string

// ContainerMetadata represents the record "wasi:blobstore/types@0.2.0-draft#container-metadata".
//
// information about a container
//
//	record container-metadata {
//		name: container-name,
//		created-at: timestamp,
//	}
type ContainerMetadata struct {
	_ cm.HostLayout
	// the container's name
	Name ContainerName

	// date and time container was created
	CreatedAt Timestamp
}

// ObjectMetadata represents the record "wasi:blobstore/types@0.2.0-draft#object-metadata".
//
// information about an object
//
//	record object-metadata {
//		name: object-name,
//		container: container-name,
//		created-at: timestamp,
//		size: object-size,
//	}
type ObjectMetadata struct {
	_ cm.HostLayout
	// the object's name
	Name ObjectName

	// the object's parent container
	Container ContainerName

	// date and time the object was created
	CreatedAt Timestamp

	// size of the object, in bytes
	Size ObjectSize
}

// ObjectID represents the record "wasi:blobstore/types@0.2.0-draft#object-id".
//
// identifier for an object that includes its container name
//
//	record object-id {
//		container: container-name,
//		object: object-name,
//	}
type ObjectID struct {
	_         cm.HostLayout
	Container ContainerName
	Object    ObjectName
}

// OutgoingValue represents the imported resource "wasi:blobstore/types@0.2.0-draft#outgoing-value".
//
// A data is the data stored in a data blob. The value can be of any type
// that can be represented in a byte array. It provides a way to write the value
// to the output-stream defined in the `wasi-io` interface.
// Soon: switch to `resource value { ... }`
//
//	resource outgoing-value
type OutgoingValue cm.Resource

// ResourceDrop represents the imported resource-drop for resource "outgoing-value".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingValue) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutgoingValueResourceDrop((uint32)(self0))
	return
}

// OutgoingValueFinish represents the imported static function "finish".
//
// Finalize an outgoing value. This must be
// called to signal that the outgoing value is complete. If the `outgoing-value`
// is dropped without calling `outgoing-value.finalize`, the implementation
// should treat the value as corrupted.
//
//	finish: static func(this: outgoing-value) -> result<_, error>
//
//go:nosplit
func OutgoingValueFinish(this OutgoingValue) (result cm.Result[Error, struct{}, Error]) {
	this0 := cm.Reinterpret[uint32](this)
	wasmimport_OutgoingValueFinish((uint32)(this0), &result)
	return
}

// OutgoingValueNewOutgoingValue represents the imported static function "new-outgoing-value".
//
//	new-outgoing-value: static func() -> outgoing-value
//
//go:nosplit
func OutgoingValueNewOutgoingValue() (result OutgoingValue) {
	result0 := wasmimport_OutgoingValueNewOutgoingValue()
	result = cm.Reinterpret[OutgoingValue]((uint32)(result0))
	return
}

// OutgoingValueWriteBody represents the imported method "outgoing-value-write-body".
//
// Returns a stream for writing the value contents.
//
// The returned `output-stream` is a child resource: it must be dropped
// before the parent `outgoing-value` resource is dropped (or finished),
// otherwise the `outgoing-value` drop or `finish` will trap.
//
// Returns success on the first call: the `output-stream` resource for
// this `outgoing-value` may be retrieved at most once. Subsequent calls
// will return error.
//
//	outgoing-value-write-body: func() -> result<output-stream>
//
//go:nosplit
func (self OutgoingValue) OutgoingValueWriteBody() (result cm.Result[OutputStream, OutputStream, struct{}]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutgoingValueOutgoingValueWriteBody((uint32)(self0), &result)
	return
}

// IncomingValue represents the imported resource "wasi:blobstore/types@0.2.0-draft#incoming-value".
//
// A incoming-value is a wrapper around a value. It provides a way to read the value
// from the input-stream defined in the `wasi-io` interface.
//
// The incoming-value provides two ways to consume the value:
// 1. `incoming-value-consume-sync` consumes the value synchronously and returns the
// value as a list of bytes.
// 2. `incoming-value-consume-async` consumes the value asynchronously and returns
// the
// value as an input-stream.
// Soon: switch to `resource incoming-value { ... }`
//
//	resource incoming-value
type IncomingValue cm.Resource

// ResourceDrop represents the imported resource-drop for resource "incoming-value".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingValue) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_IncomingValueResourceDrop((uint32)(self0))
	return
}

// IncomingValueIncomingValueConsumeAsync represents the imported static function "incoming-value-consume-async".
//
//	incoming-value-consume-async: static func(this: incoming-value) -> result<incoming-value-async-body,
//	error>
//
//go:nosplit
func IncomingValueIncomingValueConsumeAsync(this IncomingValue) (result cm.Result[string, IncomingValueAsyncBody, Error]) {
	this0 := cm.Reinterpret[uint32](this)
	wasmimport_IncomingValueIncomingValueConsumeAsync((uint32)(this0), &result)
	return
}

// IncomingValueIncomingValueConsumeSync represents the imported static function "incoming-value-consume-sync".
//
//	incoming-value-consume-sync: static func(this: incoming-value) -> result<incoming-value-sync-body,
//	error>
//
//go:nosplit
func IncomingValueIncomingValueConsumeSync(this IncomingValue) (result cm.Result[IncomingValueSyncBody, IncomingValueSyncBody, Error]) {
	this0 := cm.Reinterpret[uint32](this)
	wasmimport_IncomingValueIncomingValueConsumeSync((uint32)(this0), &result)
	return
}

// Size represents the imported method "size".
//
//	size: func() -> u64
//
//go:nosplit
func (self IncomingValue) Size() (result uint64) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_IncomingValueSize((uint32)(self0))
	result = (uint64)((uint64)(result0))
	return
}

// IncomingValueAsyncBody represents the imported type alias "wasi:blobstore/types@0.2.0-draft#incoming-value-async-body".
//
// See [InputStream] for more information.
type IncomingValueAsyncBody = InputStream

// IncomingValueSyncBody represents the list "wasi:blobstore/types@0.2.0-draft#incoming-value-sync-body".
//
//	type incoming-value-sync-body = list<u8>
type IncomingValueSyncBody cm.List[uint8]