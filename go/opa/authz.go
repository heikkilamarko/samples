package main

import (
	"context"
	_ "embed"

	"github.com/open-policy-agent/opa/rego"
)

//go:embed authz.rego
var authzRego string

type AuthZ struct {
	query rego.PreparedEvalQuery
}

type AuthZInput struct {
	Action string `json:"action"`
	Object string `json:"object"`
	Token  any    `json:"token"`
}

func NewAuthZ(ctx context.Context) (*AuthZ, error) {
	q, err := rego.New(
		rego.Query("data.authz.allow"),
		rego.Module("authz.rego", authzRego),
	).PrepareForEval(ctx)

	if err != nil {
		return nil, err
	}

	return &AuthZ{q}, nil
}

func (az *AuthZ) Authorize(ctx context.Context, input any) (bool, error) {
	r, err := az.query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return false, err
	}
	return r.Allowed(), nil
}
