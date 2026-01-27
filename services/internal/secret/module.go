package secret

import "go.uber.org/fx"

var Module = fx.Module("secret",
	fx.Provide(
		fx.Annotate(newSecretServiceInternal,
			fx.As(new(SecretService)),
		),
	),
)
