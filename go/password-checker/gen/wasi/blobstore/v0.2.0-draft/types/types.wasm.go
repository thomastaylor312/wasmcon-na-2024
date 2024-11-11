// Code generated by wit-bindgen-go. DO NOT EDIT.

package types

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:blobstore@0.2.0-draft".

//go:wasmimport wasi:blobstore/types@0.2.0-draft [resource-drop]outgoing-value
//go:noescape
func wasmimport_OutgoingValueResourceDrop(self0 uint32)

//go:wasmimport wasi:blobstore/types@0.2.0-draft [static]outgoing-value.finish
//go:noescape
func wasmimport_OutgoingValueFinish(this0 uint32, result *cm.Result[Error, struct{}, Error])

//go:wasmimport wasi:blobstore/types@0.2.0-draft [static]outgoing-value.new-outgoing-value
//go:noescape
func wasmimport_OutgoingValueNewOutgoingValue() (result0 uint32)

//go:wasmimport wasi:blobstore/types@0.2.0-draft [method]outgoing-value.outgoing-value-write-body
//go:noescape
func wasmimport_OutgoingValueOutgoingValueWriteBody(self0 uint32, result *cm.Result[OutputStream, OutputStream, struct{}])

//go:wasmimport wasi:blobstore/types@0.2.0-draft [resource-drop]incoming-value
//go:noescape
func wasmimport_IncomingValueResourceDrop(self0 uint32)

//go:wasmimport wasi:blobstore/types@0.2.0-draft [static]incoming-value.incoming-value-consume-async
//go:noescape
func wasmimport_IncomingValueIncomingValueConsumeAsync(this0 uint32, result *cm.Result[string, IncomingValueAsyncBody, Error])

//go:wasmimport wasi:blobstore/types@0.2.0-draft [static]incoming-value.incoming-value-consume-sync
//go:noescape
func wasmimport_IncomingValueIncomingValueConsumeSync(this0 uint32, result *cm.Result[IncomingValueSyncBody, IncomingValueSyncBody, Error])

//go:wasmimport wasi:blobstore/types@0.2.0-draft [method]incoming-value.size
//go:noescape
func wasmimport_IncomingValueSize(self0 uint32) (result0 uint64)