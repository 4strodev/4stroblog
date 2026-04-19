package domain

import "context"

type SiteMetaRepository interface {
	Save(context.Context, SiteMeta) error
	Get(context.Context) (SiteMeta, error)
}
