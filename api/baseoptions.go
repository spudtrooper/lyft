// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type BaseOption struct {
	f func(*baseOptionImpl)
	s string
}

func (o BaseOption) String() string { return o.s }

type BaseOptions interface {
	Token() string
	HasToken() bool
}

func BaseToken(token string) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.BaseToken(string %+v)}", token)}
}
func BaseTokenFlag(token *string) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.BaseToken(string %+v)}", token)}
}

type baseOptionImpl struct {
	token     string
	has_token bool
}

func (b *baseOptionImpl) Token() string  { return b.token }
func (b *baseOptionImpl) HasToken() bool { return b.has_token }

func makeBaseOptionImpl(opts ...BaseOption) *baseOptionImpl {
	res := &baseOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeBaseOptions(opts ...BaseOption) BaseOptions {
	return makeBaseOptionImpl(opts...)
}
