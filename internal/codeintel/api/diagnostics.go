package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
	bundles "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
)

// TODO(efritz) - document
func (api *codeIntelAPI) Diagnostics(ctx context.Context, file string, uploadID int) ([]bundles.Diagnostic, error) {
	fmt.Printf("IN API\n")

	dump, exists, err := api.db.GetDumpByID(ctx, uploadID)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetDumpByID")
	}
	if !exists {
		return nil, ErrMissingDump
	}

	pathInBundle := strings.TrimPrefix(file, dump.Root)
	bundleClient := api.bundleManagerClient.BundleClient(dump.ID)

	fmt.Printf("PATH IN BUNDLE: %s", pathInBundle)

	diagnostics, err := bundleClient.Diagnostics(ctx, pathInBundle)
	log15.Warn(fmt.Sprintf("QQQ: %v - %v\n", diagnostics, err))
	if err != nil {
		if err == client.ErrNotFound {
			log15.Warn("Bundle does not exist")
			return nil, nil
		}
		return nil, errors.Wrap(err, "bundleClient.Diagnostics")
	}

	return diagnostics, nil
}
