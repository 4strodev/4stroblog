package infrastructure

import (
	"context"

	"github.com/4strodev/4stroblog/site/features/site_meta/domain"
)

type InMemorySiteMetaRepository struct {
	SiteMeta domain.SiteMeta
}

func (r *InMemorySiteMetaRepository) Save(ctx context.Context, sitemeta domain.SiteMeta) error {
	r.SiteMeta = sitemeta
	return nil
}

func (r *InMemorySiteMetaRepository) Get(ctx context.Context) (domain.SiteMeta, error) {
	return r.SiteMeta, nil
}
