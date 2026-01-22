package offer

import "go.uber.org/fx"

var Module = fx.Module("offer",
	fx.Provide(NewOfferRepository),
	fx.Provide(NewOfferTool),
)
