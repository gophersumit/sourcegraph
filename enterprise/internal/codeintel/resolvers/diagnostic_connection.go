package resolvers

import (
	"context"
	"fmt"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
)

type diagnosticConnectionResolver struct {
	// TODO
}

var _ graphqlbackend.DiagnosticConnectionResolver = &diagnosticConnectionResolver{}

func (r *diagnosticConnectionResolver) Nodes(ctx context.Context) ([]graphqlbackend.DiagnosticResolver, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("A unimplemented")
}

func (r *diagnosticConnectionResolver) TotalCount(ctx context.Context) (int32, error) {
	// TODO(efritz) - implement
	return 0, fmt.Errorf("B unimplemented")
}

func (r *diagnosticConnectionResolver) PageInfo(ctx context.Context) (*graphqlutil.PageInfo, error) {
	// TODO(efritz) - implement
	return nil, fmt.Errorf("C unimplemented")
}
