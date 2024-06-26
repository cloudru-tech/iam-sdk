// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iam/service_accounts/v1/service_account.proto

package serviceaccountsv1

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

// Validate checks the field values on ServiceAccounts with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ServiceAccounts) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceAccounts with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceAccountsMultiError, or nil if none found.
func (m *ServiceAccounts) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceAccounts) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetServiceAccounts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServiceAccountsValidationError{
						field:  fmt.Sprintf("ServiceAccounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServiceAccountsValidationError{
						field:  fmt.Sprintf("ServiceAccounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceAccountsValidationError{
					field:  fmt.Sprintf("ServiceAccounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ServiceAccountsMultiError(errors)
	}

	return nil
}

// ServiceAccountsMultiError is an error wrapping multiple validation errors
// returned by ServiceAccounts.ValidateAll() if the designated constraints
// aren't met.
type ServiceAccountsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceAccountsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceAccountsMultiError) AllErrors() []error { return m }

// ServiceAccountsValidationError is the validation error returned by
// ServiceAccounts.Validate if the designated constraints aren't met.
type ServiceAccountsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceAccountsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceAccountsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceAccountsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceAccountsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceAccountsValidationError) ErrorName() string { return "ServiceAccountsValidationError" }

// Error satisfies the builtin error interface
func (e ServiceAccountsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceAccounts.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceAccountsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceAccountsValidationError{}

// Validate checks the field values on ServiceAccountID with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ServiceAccountID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceAccountID with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceAccountIDMultiError, or nil if none found.
func (m *ServiceAccountID) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceAccountID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ServiceAccountIDMultiError(errors)
	}

	return nil
}

// ServiceAccountIDMultiError is an error wrapping multiple validation errors
// returned by ServiceAccountID.ValidateAll() if the designated constraints
// aren't met.
type ServiceAccountIDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceAccountIDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceAccountIDMultiError) AllErrors() []error { return m }

// ServiceAccountIDValidationError is the validation error returned by
// ServiceAccountID.Validate if the designated constraints aren't met.
type ServiceAccountIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceAccountIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceAccountIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceAccountIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceAccountIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceAccountIDValidationError) ErrorName() string { return "ServiceAccountIDValidationError" }

// Error satisfies the builtin error interface
func (e ServiceAccountIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceAccountID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceAccountIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceAccountIDValidationError{}

// Validate checks the field values on ServiceAccount with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServiceAccount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceAccount with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServiceAccountMultiError,
// or nil if none found.
func (m *ServiceAccount) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceAccount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NamespaceId

	// no validation rules for Name

	// no validation rules for Email

	// no validation rules for Description

	// no validation rules for UseRefreshTokens

	// no validation rules for Enabled

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceAccountValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceAccountValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceAccountValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceAccountValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceAccountValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceAccountValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceAccountMultiError(errors)
	}

	return nil
}

// ServiceAccountMultiError is an error wrapping multiple validation errors
// returned by ServiceAccount.ValidateAll() if the designated constraints
// aren't met.
type ServiceAccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceAccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceAccountMultiError) AllErrors() []error { return m }

// ServiceAccountValidationError is the validation error returned by
// ServiceAccount.Validate if the designated constraints aren't met.
type ServiceAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceAccountValidationError) ErrorName() string { return "ServiceAccountValidationError" }

// Error satisfies the builtin error interface
func (e ServiceAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceAccountValidationError{}
