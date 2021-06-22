// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-tenant-api/ocp-tenant-api.proto

package ocp_tenant_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Tenant) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	return nil
}

// TenantValidationError is the validation error returned by Tenant.Validate if
// the designated constraints aren't met.
type TenantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantValidationError) ErrorName() string { return "TenantValidationError" }

// Error satisfies the builtin error interface
func (e TenantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantValidationError{}

// Validate checks the field values on CreateTenantV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTenantV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	return nil
}

// CreateTenantV1RequestValidationError is the validation error returned by
// CreateTenantV1Request.Validate if the designated constraints aren't met.
type CreateTenantV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTenantV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTenantV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTenantV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTenantV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTenantV1RequestValidationError) ErrorName() string {
	return "CreateTenantV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTenantV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTenantV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTenantV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTenantV1RequestValidationError{}

// Validate checks the field values on CreateTenantV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTenantV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TenantId

	return nil
}

// CreateTenantV1ResponseValidationError is the validation error returned by
// CreateTenantV1Response.Validate if the designated constraints aren't met.
type CreateTenantV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTenantV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTenantV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTenantV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTenantV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTenantV1ResponseValidationError) ErrorName() string {
	return "CreateTenantV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTenantV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTenantV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTenantV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTenantV1ResponseValidationError{}

// Validate checks the field values on MultiCreateTenantV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateTenantV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCreate() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateTenantV1RequestValidationError{
					field:  fmt.Sprintf("Create[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateTenantV1RequestValidationError is the validation error returned
// by MultiCreateTenantV1Request.Validate if the designated constraints aren't met.
type MultiCreateTenantV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateTenantV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateTenantV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateTenantV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateTenantV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateTenantV1RequestValidationError) ErrorName() string {
	return "MultiCreateTenantV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateTenantV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateTenantV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateTenantV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateTenantV1RequestValidationError{}

// Validate checks the field values on MultiCreateTenantV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateTenantV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SuccessCount

	return nil
}

// MultiCreateTenantV1ResponseValidationError is the validation error returned
// by MultiCreateTenantV1Response.Validate if the designated constraints
// aren't met.
type MultiCreateTenantV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateTenantV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateTenantV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateTenantV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateTenantV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateTenantV1ResponseValidationError) ErrorName() string {
	return "MultiCreateTenantV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateTenantV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateTenantV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateTenantV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateTenantV1ResponseValidationError{}

// Validate checks the field values on DescribeTenantV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeTenantV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TenantId

	return nil
}

// DescribeTenantV1RequestValidationError is the validation error returned by
// DescribeTenantV1Request.Validate if the designated constraints aren't met.
type DescribeTenantV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeTenantV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeTenantV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeTenantV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeTenantV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeTenantV1RequestValidationError) ErrorName() string {
	return "DescribeTenantV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeTenantV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeTenantV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeTenantV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeTenantV1RequestValidationError{}

// Validate checks the field values on DescribeTenantV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeTenantV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTenant()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeTenantV1ResponseValidationError{
				field:  "Tenant",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeTenantV1ResponseValidationError is the validation error returned by
// DescribeTenantV1Response.Validate if the designated constraints aren't met.
type DescribeTenantV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeTenantV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeTenantV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeTenantV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeTenantV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeTenantV1ResponseValidationError) ErrorName() string {
	return "DescribeTenantV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeTenantV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeTenantV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeTenantV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeTenantV1ResponseValidationError{}

// Validate checks the field values on UpdateTenantV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateTenantV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTenant()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateTenantV1RequestValidationError{
				field:  "Tenant",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateTenantV1RequestValidationError is the validation error returned by
// UpdateTenantV1Request.Validate if the designated constraints aren't met.
type UpdateTenantV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTenantV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTenantV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTenantV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTenantV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTenantV1RequestValidationError) ErrorName() string {
	return "UpdateTenantV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTenantV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTenantV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTenantV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTenantV1RequestValidationError{}

// Validate checks the field values on UpdateTenantV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateTenantV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateTenantV1ResponseValidationError is the validation error returned by
// UpdateTenantV1Response.Validate if the designated constraints aren't met.
type UpdateTenantV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTenantV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTenantV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTenantV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTenantV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTenantV1ResponseValidationError) ErrorName() string {
	return "UpdateTenantV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTenantV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTenantV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTenantV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTenantV1ResponseValidationError{}

// Validate checks the field values on ListTenantsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListTenantsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListTenantsV1RequestValidationError is the validation error returned by
// ListTenantsV1Request.Validate if the designated constraints aren't met.
type ListTenantsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTenantsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTenantsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTenantsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTenantsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTenantsV1RequestValidationError) ErrorName() string {
	return "ListTenantsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListTenantsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTenantsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTenantsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTenantsV1RequestValidationError{}

// Validate checks the field values on ListTenantsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListTenantsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetTenants() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTenantsV1ResponseValidationError{
					field:  fmt.Sprintf("Tenants[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListTenantsV1ResponseValidationError is the validation error returned by
// ListTenantsV1Response.Validate if the designated constraints aren't met.
type ListTenantsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTenantsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTenantsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTenantsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTenantsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTenantsV1ResponseValidationError) ErrorName() string {
	return "ListTenantsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListTenantsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTenantsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTenantsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTenantsV1ResponseValidationError{}

// Validate checks the field values on RemoveTenantV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveTenantV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TenantId

	return nil
}

// RemoveTenantV1RequestValidationError is the validation error returned by
// RemoveTenantV1Request.Validate if the designated constraints aren't met.
type RemoveTenantV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTenantV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTenantV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTenantV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTenantV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTenantV1RequestValidationError) ErrorName() string {
	return "RemoveTenantV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveTenantV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTenantV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTenantV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTenantV1RequestValidationError{}

// Validate checks the field values on RemoveTenantV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveTenantV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Removed

	return nil
}

// RemoveTenantV1ResponseValidationError is the validation error returned by
// RemoveTenantV1Response.Validate if the designated constraints aren't met.
type RemoveTenantV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTenantV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTenantV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTenantV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTenantV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTenantV1ResponseValidationError) ErrorName() string {
	return "RemoveTenantV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveTenantV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTenantV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTenantV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTenantV1ResponseValidationError{}
