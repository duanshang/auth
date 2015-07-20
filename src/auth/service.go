package main

import (
	"errors"
	"golang.org/x/net/context"
	"regexp"
	"strings"
)

import (
	. "proto"
)

const (
	SERVICE = "[AUTH]"
)
const (
	_port = ":50006"
)

var (
	ERROR_METHOD_NOT_SUPPORTED = errors.New("method not supoorted")
)

var (
	uuid_regexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

type server struct {
}

func (s *server) init() {
}

func (s *server) Auth(ctx context.Context, cert *Auth_Certificate) (*Auth_Result, error) {
	switch cert.Type {
	case Auth_UUID:
		if uuid_regexp.MatchString(strings.ToLower(string(cert.Proof))) {
			return &Auth_Result{true, nil}, nil
		} else {
			return &Auth_Result{false, nil}, nil
		}
	case Auth_PLAIN:
	case Auth_TOKEN:
	case Auth_FACEBOOK:
	default:
	}
	return nil, nil
}
