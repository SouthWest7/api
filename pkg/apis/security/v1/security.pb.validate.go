// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/security/v1/security.proto

package v1

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

// Validate checks the field values on CertificateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CertificateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CertificateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CertificateRequestMultiError, or nil if none found.
func (m *CertificateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CertificateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Csr

	// no validation rules for ValidityDuration

	if len(errors) > 0 {
		return CertificateRequestMultiError(errors)
	}

	return nil
}

// CertificateRequestMultiError is an error wrapping multiple validation errors
// returned by CertificateRequest.ValidateAll() if the designated constraints
// aren't met.
type CertificateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CertificateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CertificateRequestMultiError) AllErrors() []error { return m }

// CertificateRequestValidationError is the validation error returned by
// CertificateRequest.Validate if the designated constraints aren't met.
type CertificateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateRequestValidationError) ErrorName() string {
	return "CertificateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateRequestValidationError{}

// Validate checks the field values on CertificateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CertificateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CertificateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CertificateResponseMultiError, or nil if none found.
func (m *CertificateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CertificateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CertificateResponseMultiError(errors)
	}

	return nil
}

// CertificateResponseMultiError is an error wrapping multiple validation
// errors returned by CertificateResponse.ValidateAll() if the designated
// constraints aren't met.
type CertificateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CertificateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CertificateResponseMultiError) AllErrors() []error { return m }

// CertificateResponseValidationError is the validation error returned by
// CertificateResponse.Validate if the designated constraints aren't met.
type CertificateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateResponseValidationError) ErrorName() string {
	return "CertificateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateResponseValidationError{}
