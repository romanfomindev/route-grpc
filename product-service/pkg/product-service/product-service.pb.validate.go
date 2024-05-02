// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/product-service/product-service.proto

package product_service

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
)

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Product) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for CategoryId

	return nil
}

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if m.GetCategoryId() <= 0 {
		return CreateProductRequestValidationError{
			field:  "CategoryId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProductResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProductResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateProductResponseValidationError is the validation error returned by
// CreateProductResponse.Validate if the designated constraints aren't met.
type CreateProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductResponseValidationError) ErrorName() string {
	return "CreateProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductResponseValidationError{}