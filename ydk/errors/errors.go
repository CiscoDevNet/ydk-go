// Package errors provides built-in errors specified in
// YANG RFC 6020 and YDK Go APIs.
//
// YANG Development Kit Copyright 2017 Cisco Systems. All rights reserved.
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package errors

// Represents YDK Go error types
type Y_ERROR_TYPE int

const (
	Y_ERROR_TYPE_NONE Y_ERROR_TYPE = iota
	Y_ERROR_TYPE_ERROR
	Y_ERROR_TYPE_CLIENT_ERROR
	Y_ERROR_TYPE_SERVICE_PROVIDER_ERROR
	Y_ERROR_TYPE_SERVICE_ERROR
	Y_ERROR_TYPE_ILLEGAL_STATE_ERROR
	Y_ERROR_TYPE_INVALID_ARGUMENT_ERROR
	Y_ERROR_TYPE_OPERATION_NOTSUPPORTED_ERROR
	Y_ERROR_TYPE_MODEL_ERROR
)

// State represents the error state
type State struct {
	Private interface{}
}

// CState represents the error state
type CState struct {
	Private interface{}
}

type CError interface {
	Error() string
}

// YError is the basic error type in Go
type YError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YError) Error() string {
	return "YError:" + e.Msg
}

// YClientError is the error for client.
type YClientError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YClientError) Error() string {
	return "YClientError:" + e.Msg
}

// YServiceProviderError is the error for service provider.
type YServiceProviderError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YServiceProviderError) Error() string {
	return "YServiceProviderError:" + e.Msg
}

// YServiceError is the error for service.
type YServiceError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YServiceError) Error() string {
	return "YServiceError:" + e.Msg
}

// YIllegalStateError is raised when an operation/service is invoked on an object that is not in the right state.
type YIllegalStateError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YIllegalStateError) Error() string {
	return "YIllegalStateError:" + e.Msg
}

// YInvalidArgumentError is raised when there is an invalid argument.
type YInvalidArgumentError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YInvalidArgumentError) Error() string {
	return "YInvalidArgumentError:" + e.Msg
}

// YOperationNotSupportedError is raised for an unsupported operation.
type YOperationNotSupportedError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YOperationNotSupportedError) Error() string {
	return "YOperationNotSupportedError:" + e.Msg
}

// YModelError is raised when a model constraint is violated.
type YModelError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YModelError) Error() string {
	return "YModelError:" + e.Msg
}

// YModelError is the error for core.
type YCoreError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YCoreError) Error() string {
	return "YCoreError:" + e.Msg
}

// YCodecError encapsualtes the validation errors for codec service.
type YCodecError struct {
	Msg string
}

// Error satisfies the error interface
// Returns the error message (string)
func (e *YCodecError) Error() string {
	return "YCodecError" + e.Msg
}
