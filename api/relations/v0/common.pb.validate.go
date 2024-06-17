// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: relations/v0/common.proto

package v0

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

// Validate checks the field values on Relationship with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Relationship) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Relationship with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelationshipMultiError, or
// nil if none found.
func (m *Relationship) ValidateAll() error {
	return m.validate(true)
}

func (m *Relationship) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationshipValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationshipValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RelationshipMultiError(errors)
	}

	return nil
}

// RelationshipMultiError is an error wrapping multiple validation errors
// returned by Relationship.ValidateAll() if the designated constraints aren't met.
type RelationshipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationshipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationshipMultiError) AllErrors() []error { return m }

// RelationshipValidationError is the validation error returned by
// Relationship.Validate if the designated constraints aren't met.
type RelationshipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationshipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationshipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationshipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationshipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationshipValidationError) ErrorName() string { return "RelationshipValidationError" }

// Error satisfies the builtin error interface
func (e RelationshipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationship.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationshipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationshipValidationError{}

// Validate checks the field values on SubjectReference with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SubjectReference) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectReference with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubjectReferenceMultiError, or nil if none found.
func (m *SubjectReference) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectReference) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubjectReferenceValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubjectReferenceValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubjectReferenceValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Relation != nil {
		// no validation rules for Relation
	}

	if len(errors) > 0 {
		return SubjectReferenceMultiError(errors)
	}

	return nil
}

// SubjectReferenceMultiError is an error wrapping multiple validation errors
// returned by SubjectReference.ValidateAll() if the designated constraints
// aren't met.
type SubjectReferenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectReferenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectReferenceMultiError) AllErrors() []error { return m }

// SubjectReferenceValidationError is the validation error returned by
// SubjectReference.Validate if the designated constraints aren't met.
type SubjectReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectReferenceValidationError) ErrorName() string { return "SubjectReferenceValidationError" }

// Error satisfies the builtin error interface
func (e SubjectReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectReferenceValidationError{}

// Validate checks the field values on RequestPagination with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RequestPagination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestPagination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RequestPaginationMultiError, or nil if none found.
func (m *RequestPagination) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestPagination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Limit

	if m.ContinuationToken != nil {
		// no validation rules for ContinuationToken
	}

	if len(errors) > 0 {
		return RequestPaginationMultiError(errors)
	}

	return nil
}

// RequestPaginationMultiError is an error wrapping multiple validation errors
// returned by RequestPagination.ValidateAll() if the designated constraints
// aren't met.
type RequestPaginationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestPaginationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestPaginationMultiError) AllErrors() []error { return m }

// RequestPaginationValidationError is the validation error returned by
// RequestPagination.Validate if the designated constraints aren't met.
type RequestPaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestPaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestPaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestPaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestPaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestPaginationValidationError) ErrorName() string {
	return "RequestPaginationValidationError"
}

// Error satisfies the builtin error interface
func (e RequestPaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestPaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestPaginationValidationError{}

// Validate checks the field values on ResponsePagination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResponsePagination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResponsePagination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResponsePaginationMultiError, or nil if none found.
func (m *ResponsePagination) ValidateAll() error {
	return m.validate(true)
}

func (m *ResponsePagination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ContinuationToken

	if len(errors) > 0 {
		return ResponsePaginationMultiError(errors)
	}

	return nil
}

// ResponsePaginationMultiError is an error wrapping multiple validation errors
// returned by ResponsePagination.ValidateAll() if the designated constraints
// aren't met.
type ResponsePaginationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResponsePaginationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResponsePaginationMultiError) AllErrors() []error { return m }

// ResponsePaginationValidationError is the validation error returned by
// ResponsePagination.Validate if the designated constraints aren't met.
type ResponsePaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponsePaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponsePaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponsePaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponsePaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponsePaginationValidationError) ErrorName() string {
	return "ResponsePaginationValidationError"
}

// Error satisfies the builtin error interface
func (e ResponsePaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponsePagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponsePaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponsePaginationValidationError{}

// Validate checks the field values on ObjectReference with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ObjectReference) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ObjectReference with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ObjectReferenceMultiError, or nil if none found.
func (m *ObjectReference) ValidateAll() error {
	return m.validate(true)
}

func (m *ObjectReference) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetType() == nil {
		err := ObjectReferenceValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetType(); a != nil {

	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := ObjectReferenceValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ObjectReferenceMultiError(errors)
	}

	return nil
}

// ObjectReferenceMultiError is an error wrapping multiple validation errors
// returned by ObjectReference.ValidateAll() if the designated constraints
// aren't met.
type ObjectReferenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ObjectReferenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ObjectReferenceMultiError) AllErrors() []error { return m }

// ObjectReferenceValidationError is the validation error returned by
// ObjectReference.Validate if the designated constraints aren't met.
type ObjectReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectReferenceValidationError) ErrorName() string { return "ObjectReferenceValidationError" }

// Error satisfies the builtin error interface
func (e ObjectReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObjectReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectReferenceValidationError{}

// Validate checks the field values on ObjectType with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ObjectType) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ObjectType with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ObjectTypeMultiError, or
// nil if none found.
func (m *ObjectType) ValidateAll() error {
	return m.validate(true)
}

func (m *ObjectType) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Name

	if len(errors) > 0 {
		return ObjectTypeMultiError(errors)
	}

	return nil
}

// ObjectTypeMultiError is an error wrapping multiple validation errors
// returned by ObjectType.ValidateAll() if the designated constraints aren't met.
type ObjectTypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ObjectTypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ObjectTypeMultiError) AllErrors() []error { return m }

// ObjectTypeValidationError is the validation error returned by
// ObjectType.Validate if the designated constraints aren't met.
type ObjectTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectTypeValidationError) ErrorName() string { return "ObjectTypeValidationError" }

// Error satisfies the builtin error interface
func (e ObjectTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObjectType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectTypeValidationError{}
