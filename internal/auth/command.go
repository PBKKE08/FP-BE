package auth

import "context"

type InginLoginCommand interface {
	Run(ctx context.Context, dto InginLoginDTO) error
}

type InginLogoutCommand interface {
	Run(ctx context.Context, dto InginLogoutDTO) error
}

type InginRegisterCommand interface {
	Run(ctx context.Context, dto InginRegisterDTO) error
}

type InginVerifikasiEmail interface {
	Run(ctx context.Context, dto InginVerifikasiEmailDTO) error
}
