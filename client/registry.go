package client

import (
	"github.com/ory/fosite"
	"github.com/ory/hydra/x"
)

type registry interface {
	x.RegistryWriter
	Registry
}

type Registry interface {
	ClientValidator() *Validator
	ClientManager() Manager
	ClientHasher() fosite.Hasher
}

type Configuration interface {
	DefaultClientScope() []string
	GetSubjectTypesSupported() []string
}
