package main

import (
	"context"
	_ "embed"
	"errors"

	"github.com/open-policy-agent/opa/rego"
	"github.com/samber/lo"
)

//go:embed authz.rego
var authzRego string

type AuthZ struct {
	query rego.PreparedEvalQuery
}

type AuthZInput struct {
	Permission string `json:"permission"`
	Token      any    `json:"token"`
}

type AuthZResult struct {
	Allow       bool
	Permissions []string
}

func NewAuthZ(ctx context.Context) (*AuthZ, error) {
	q, err := rego.New(
		rego.Query("allow=data.authz.allow permissions=data.authz.permissions"),
		rego.Module("authz.rego", authzRego),
	).PrepareForEval(ctx)

	if err != nil {
		return nil, err
	}

	return &AuthZ{q}, nil
}

func (az *AuthZ) Authorize(ctx context.Context, input any) (*AuthZResult, error) {
	result, err := az.query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("undefined authz result")
	}

	bindings := result[0].Bindings

	allow, ok := bindings["allow"].(bool)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'allow' must be bool")
	}

	p, ok := bindings["permissions"].([]any)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'permissions' must be []string")
	}

	permissions, ok := lo.FromAnySlice[string](p)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'permissions' must be []string")
	}

	return &AuthZResult{allow, permissions}, nil
}
