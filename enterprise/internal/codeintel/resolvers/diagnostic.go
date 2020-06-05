package resolvers

import (
	"context"
	"fmt"

	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/types"
	"github.com/sourcegraph/sourcegraph/internal/api"
	bundles "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
)

type diagnosticResolver struct {
	repo               *types.Repo
	commit             api.CommitID
	diagnostic         bundles.Diagnostic
	collectionResolver *repositoryCollectionResolver
}

var _ graphqlbackend.DiagnosticResolver = &diagnosticResolver{}

func (r *diagnosticResolver) Location(ctx context.Context) (graphqlbackend.LocationResolver, error) {
	treeResolver, err := r.collectionResolver.resolve(ctx, api.RepoID(r.repo.ID), string(r.commit), r.diagnostic.Path)
	if err != nil {
		return nil, err
	}

	if treeResolver == nil {
		return nil, nil
	}

	// TODO(efritz) - translate range to correct commit
	return graphqlbackend.NewLocationResolver(treeResolver, &lsp.Range{
		Start: lsp.Position{Line: r.diagnostic.StartLine, Character: r.diagnostic.EndLine},
		End:   lsp.Position{Line: r.diagnostic.StartCharacter, Character: r.diagnostic.EndCharacter},
	}), nil
}

var severities = map[int]string{
	1: "ERROR",
	2: "WARNING",
	3: "INFORMATION",
	4: "HINT",
}

func (r *diagnosticResolver) Severity(ctx context.Context) (*string, error) {
	if severity, ok := severities[r.diagnostic.Severity]; ok {
		return &severity, nil
	}

	return nil, fmt.Errorf("unknown diagnostic severity %d", r.diagnostic.Severity)
}

func (r *diagnosticResolver) Code(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return &r.diagnostic.Code, nil
}

func (r *diagnosticResolver) Source(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return &r.diagnostic.Source, nil
}

func (r *diagnosticResolver) Message(ctx context.Context) (*string, error) {
	// TODO(efritz) - implement
	return &r.diagnostic.Message, nil
}
