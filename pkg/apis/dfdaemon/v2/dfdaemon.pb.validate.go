// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/dfdaemon/v2/dfdaemon.proto

package dfdaemon

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SyncPiecesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SyncPiecesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncPiecesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SyncPiecesRequestMultiError, or nil if none found.
func (m *SyncPiecesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncPiecesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := SyncPiecesRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetInterestedPieceNumbers()) < 1 {
		err := SyncPiecesRequestValidationError{
			field:  "InterestedPieceNumbers",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SyncPiecesRequestMultiError(errors)
	}

	return nil
}

// SyncPiecesRequestMultiError is an error wrapping multiple validation errors
// returned by SyncPiecesRequest.ValidateAll() if the designated constraints
// aren't met.
type SyncPiecesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncPiecesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncPiecesRequestMultiError) AllErrors() []error { return m }

// SyncPiecesRequestValidationError is the validation error returned by
// SyncPiecesRequest.Validate if the designated constraints aren't met.
type SyncPiecesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncPiecesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncPiecesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncPiecesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncPiecesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncPiecesRequestValidationError) ErrorName() string {
	return "SyncPiecesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SyncPiecesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncPiecesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncPiecesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncPiecesRequestValidationError{}

// Validate checks the field values on SyncPiecesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SyncPiecesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncPiecesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SyncPiecesResponseMultiError, or nil if none found.
func (m *SyncPiecesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncPiecesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PieceNumber

	if len(errors) > 0 {
		return SyncPiecesResponseMultiError(errors)
	}

	return nil
}

// SyncPiecesResponseMultiError is an error wrapping multiple validation errors
// returned by SyncPiecesResponse.ValidateAll() if the designated constraints
// aren't met.
type SyncPiecesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncPiecesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncPiecesResponseMultiError) AllErrors() []error { return m }

// SyncPiecesResponseValidationError is the validation error returned by
// SyncPiecesResponse.Validate if the designated constraints aren't met.
type SyncPiecesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncPiecesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncPiecesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncPiecesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncPiecesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncPiecesResponseValidationError) ErrorName() string {
	return "SyncPiecesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SyncPiecesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncPiecesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncPiecesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncPiecesResponseValidationError{}

// Validate checks the field values on DownloadPieceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadPieceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadPieceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadPieceRequestMultiError, or nil if none found.
func (m *DownloadPieceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadPieceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := DownloadPieceRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PieceNumber

	if len(errors) > 0 {
		return DownloadPieceRequestMultiError(errors)
	}

	return nil
}

// DownloadPieceRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadPieceRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadPieceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadPieceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadPieceRequestMultiError) AllErrors() []error { return m }

// DownloadPieceRequestValidationError is the validation error returned by
// DownloadPieceRequest.Validate if the designated constraints aren't met.
type DownloadPieceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadPieceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadPieceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadPieceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadPieceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadPieceRequestValidationError) ErrorName() string {
	return "DownloadPieceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadPieceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadPieceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadPieceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadPieceRequestValidationError{}

// Validate checks the field values on DownloadPieceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadPieceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadPieceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadPieceResponseMultiError, or nil if none found.
func (m *DownloadPieceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadPieceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPiece() == nil {
		err := DownloadPieceResponseValidationError{
			field:  "Piece",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPiece()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownloadPieceResponseValidationError{
					field:  "Piece",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownloadPieceResponseValidationError{
					field:  "Piece",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPiece()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownloadPieceResponseValidationError{
				field:  "Piece",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DownloadPieceResponseMultiError(errors)
	}

	return nil
}

// DownloadPieceResponseMultiError is an error wrapping multiple validation
// errors returned by DownloadPieceResponse.ValidateAll() if the designated
// constraints aren't met.
type DownloadPieceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadPieceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadPieceResponseMultiError) AllErrors() []error { return m }

// DownloadPieceResponseValidationError is the validation error returned by
// DownloadPieceResponse.Validate if the designated constraints aren't met.
type DownloadPieceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadPieceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadPieceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadPieceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadPieceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadPieceResponseValidationError) ErrorName() string {
	return "DownloadPieceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadPieceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadPieceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadPieceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadPieceResponseValidationError{}

// Validate checks the field values on DownloadTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadTaskRequestMultiError, or nil if none found.
func (m *DownloadTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDownload() == nil {
		err := DownloadTaskRequestValidationError{
			field:  "Download",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDownload()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownloadTaskRequestValidationError{
					field:  "Download",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownloadTaskRequestValidationError{
					field:  "Download",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDownload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownloadTaskRequestValidationError{
				field:  "Download",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DownloadTaskRequestMultiError(errors)
	}

	return nil
}

// DownloadTaskRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadTaskRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadTaskRequestMultiError) AllErrors() []error { return m }

// DownloadTaskRequestValidationError is the validation error returned by
// DownloadTaskRequest.Validate if the designated constraints aren't met.
type DownloadTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadTaskRequestValidationError) ErrorName() string {
	return "DownloadTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadTaskRequestValidationError{}

// Validate checks the field values on DownloadTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadTaskResponseMultiError, or nil if none found.
func (m *DownloadTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ContentLength

	if m.GetPiece() == nil {
		err := DownloadTaskResponseValidationError{
			field:  "Piece",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPiece()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownloadTaskResponseValidationError{
					field:  "Piece",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownloadTaskResponseValidationError{
					field:  "Piece",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPiece()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownloadTaskResponseValidationError{
				field:  "Piece",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DownloadTaskResponseMultiError(errors)
	}

	return nil
}

// DownloadTaskResponseMultiError is an error wrapping multiple validation
// errors returned by DownloadTaskResponse.ValidateAll() if the designated
// constraints aren't met.
type DownloadTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadTaskResponseMultiError) AllErrors() []error { return m }

// DownloadTaskResponseValidationError is the validation error returned by
// DownloadTaskResponse.Validate if the designated constraints aren't met.
type DownloadTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadTaskResponseValidationError) ErrorName() string {
	return "DownloadTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadTaskResponseValidationError{}

// Validate checks the field values on UploadTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UploadTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadTaskRequestMultiError, or nil if none found.
func (m *UploadTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTask() == nil {
		err := UploadTaskRequestValidationError{
			field:  "Task",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UploadTaskRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UploadTaskRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UploadTaskRequestValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UploadTaskRequestMultiError(errors)
	}

	return nil
}

// UploadTaskRequestMultiError is an error wrapping multiple validation errors
// returned by UploadTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadTaskRequestMultiError) AllErrors() []error { return m }

// UploadTaskRequestValidationError is the validation error returned by
// UploadTaskRequest.Validate if the designated constraints aren't met.
type UploadTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadTaskRequestValidationError) ErrorName() string {
	return "UploadTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadTaskRequestValidationError{}

// Validate checks the field values on StatTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StatTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatTaskRequestMultiError, or nil if none found.
func (m *StatTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StatTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := StatTaskRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StatTaskRequestMultiError(errors)
	}

	return nil
}

// StatTaskRequestMultiError is an error wrapping multiple validation errors
// returned by StatTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type StatTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatTaskRequestMultiError) AllErrors() []error { return m }

// StatTaskRequestValidationError is the validation error returned by
// StatTaskRequest.Validate if the designated constraints aren't met.
type StatTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatTaskRequestValidationError) ErrorName() string { return "StatTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e StatTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatTaskRequestValidationError{}

// Validate checks the field values on DeleteTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskRequestMultiError, or nil if none found.
func (m *DeleteTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := DeleteTaskRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTaskRequestMultiError(errors)
	}

	return nil
}

// DeleteTaskRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskRequestMultiError) AllErrors() []error { return m }

// DeleteTaskRequestValidationError is the validation error returned by
// DeleteTaskRequest.Validate if the designated constraints aren't met.
type DeleteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskRequestValidationError) ErrorName() string {
	return "DeleteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskRequestValidationError{}
