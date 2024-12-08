// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *CreateTelegramAccountOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.AccountID.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "account_id",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *CreateTelegramAccountReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.PhoneNumber.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "phone_number",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *Error) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.TraceID.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "trace_id",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.SpanID.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "span_id",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ErrorStatusCode) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Response.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "Response",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *SetTelegramAccountCodeOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.AccountID.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "account_id",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *SetTelegramAccountCodeReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[0-9]{3,6}$"],
		}).Validate(string(s.Code)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s SpanID) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        regexMap["[[:xdigit:]]{16}"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s TelegramAccountID) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        regexMap["^[0-9]{7,15}$"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s TraceID) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        regexMap["[[:xdigit:]]{32}"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}
