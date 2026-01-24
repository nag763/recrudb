package offer

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("offer",
	fx.Provide(NewOfferRepository),
	injection.AsMcpToolProvider(NewStoreOfferTool),
)
