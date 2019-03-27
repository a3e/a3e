package aci

import (
	"github.com/souz9/errlist"
)

func aggregateErrs(errs []error) error {
	return errlist.Error(errs)
}
