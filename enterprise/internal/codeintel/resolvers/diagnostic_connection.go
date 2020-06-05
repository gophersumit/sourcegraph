package resolvers

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/types"
	"github.com/sourcegraph/sourcegraph/internal/api"
	bundles "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
)

type diagnosticConnectionResolver struct {
	repo        *types.Repo
	commit      api.CommitID
	diagnostics []bundles.Diagnostic
}

var _ graphqlbackend.DiagnosticConnectionResolver = &diagnosticConnectionResolver{}

func (r *diagnosticConnectionResolver) Nodes(ctx context.Context) ([]graphqlbackend.DiagnosticResolver, error) {
	collectionResolver := &repositoryCollectionResolver{
		commitCollectionResolvers: map[api.RepoID]*commitCollectionResolver{},
	}

	var resolvers []graphqlbackend.DiagnosticResolver
	for _, diagnostic := range r.diagnostics {
		resolvers = append(resolvers, &diagnosticResolver{
			repo:               r.repo,
			commit:             r.commit,
			diagnostic:         diagnostic,
			collectionResolver: collectionResolver,
		})
	}

	return resolvers, nil
}

func (r *diagnosticConnectionResolver) TotalCount(ctx context.Context) (int32, error) {
	return int32(len(r.diagnostics)), nil
}

func (r *diagnosticConnectionResolver) PageInfo(ctx context.Context) (*graphqlutil.PageInfo, error) {
	return graphqlutil.HasNextPage(false), nil
}
