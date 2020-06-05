package api

import (
	"context"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
	bundles "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
)

// TODO(efritz) - document
func (api *codeIntelAPI) Diagnostics(ctx context.Context, file string, uploadID int) ([]bundles.Diagnostic, error) {
	dump, exists, err := api.db.GetDumpByID(ctx, uploadID)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetDumpByID")
	}
	if !exists {
		return nil, ErrMissingDump
	}

	pathInBundle := strings.TrimPrefix(file, dump.Root)
	bundleClient := api.bundleManagerClient.BundleClient(dump.ID)

	diagnostics, err := bundleClient.Diagnostics(ctx, pathInBundle)
	if err != nil {
		if err == client.ErrNotFound {
			log15.Warn("Bundle does not exist")
			return nil, nil
		}
		return nil, errors.Wrap(err, "bundleClient.Diagnostics")
	}

	return diagnostics, nil
}
