package qif

import (
	"github.com/pkg/errors"
)

// Account contains the fields used for the account metadata block.
type Account interface {
	Name() string
	AccountType() string
	Description() string
}

func (t *account) Name() string {
	return t.name
}

func (t *account) AccountType() string {
	return t.accountType
}

func (t *account) Description() string {
	return t.description
}

type account struct {
	name        string
	accountType string
	description string
}

func (t *account) parseAccountField(line string, config Config) error {
	if line == "" {
		return errors.New("line is empty")
	}

	if line == accountHeader {
		return nil
	}

	switch line[0] {
	case 'N':
		t.name = line[1:]
		return nil

	case 'T':
		t.accountType = line[1:]
		return nil

	case 'D':
		t.description = line[1:]
		return nil

	default:
		return UnsupportedFieldError(errors.Errorf("cannot process line '%s'", line))
	}
}
