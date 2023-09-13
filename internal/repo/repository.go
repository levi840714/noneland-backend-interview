package repo

import "github.com/google/wire"

var RepositoryProviderSet = wire.NewSet(
	NewRiskRepo,
)
