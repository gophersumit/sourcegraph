package resolvers

import (
	"context"
	"fmt"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
)

type diagnosticResolver struct {
	// TODO
}

var _ graphqlbackend.DiagnosticResolver = &diagnosticResolver{}

func (r *diagnosticResolver) Location(ctx context.Context) (graphqlbackend.LocationResolver, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("D unimplemented")
}

func (r *diagnosticResolver) Severity(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("E unimplemented")
}

func (r *diagnosticResolver) Code(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("F unimplemented")
}

func (r *diagnosticResolver) Source(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("G unimplemented")
}

func (r *diagnosticResolver) Message(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("H unimplemented")
}
