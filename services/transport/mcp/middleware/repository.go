package middleware

type MiddlewareRepository struct {
	AuthorizationMiddleware *AuthorizationMiddleware
	LoggingMiddleware       *LoggingMiddleware
}

func NewMiddlewareRepository(authorizationMiddleware *AuthorizationMiddleware, loggingMiddleware *LoggingMiddleware) *MiddlewareRepository {
	return &MiddlewareRepository{
		AuthorizationMiddleware: authorizationMiddleware,
		LoggingMiddleware:       loggingMiddleware,
	}
}
