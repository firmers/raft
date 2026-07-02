/*
 * Copyright (c) 2019, 2026, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package errors

import (
	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/errbase"
	"github.com/cockroachdb/errors/errutil"
	"github.com/cockroachdb/errors/oserror"
	"github.com/cockroachdb/errors/safedetails"
	"github.com/cockroachdb/errors/withstack"
	"github.com/cockroachdb/redact"
)

type StackTrace = errbase.StackTrace

func New(msg string) error {
	return errors.New(msg)
}

func Is(err error, reference error) bool {
	return errors.Is(err, reference)
}

func Errorf(format string, args ...any) error {
	return errutil.NewWithDepthf(1, format, args...)
}

func Safe(v any) redact.SafeValue {
	return safedetails.Safe(v)
}

func WithStack(err error) error {
	return withstack.WithStackDepth(err, 1)
}

func Wrap(err error, msg string) error {
	return errutil.WrapWithDepth(1, err, msg)
}

func IsExist(err error) bool {
	return oserror.IsExist(err)
}

func IsNotExist(err error) bool {
	return oserror.IsNotExist(err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errutil.WrapWithDepthf(1, err, format, args...)
}

func UnwrapOnce(err error) error {
	return errbase.UnwrapOnce(err)
}
