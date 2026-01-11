package auth

import (
	"server/internal/service"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(&sAuth{})
}
