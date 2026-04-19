package infrastructure

import (
	"context"
	"errors"

	"github.com/4strodev/4stroblog/site/features/site_meta/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"gorm.io/gorm"
)

type GormSiteMetaRepository struct {
	DB *gorm.DB
}

// Get implements [domain.SiteMetaRepository].
func (r GormSiteMetaRepository) Get(ctx context.Context) (siteMetaModel domain.SiteMeta, err error) {
	err = r.DB.WithContext(ctx).
		First(&siteMetaModel).Error

	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return siteMetaModel, domainerrors.WrapError(errorCode, err)
	}

	siteMeta := domain.SiteMeta{
		FirstStart: siteMetaModel.FirstStart,
	}

	return siteMeta, nil
}

// Save implements [domain.SiteMetaRepository].
func (r GormSiteMetaRepository) Save(ctx context.Context, siteMeta domain.SiteMeta) error {
	siteMetaModel := models.SiteMeta{
		FirstStart: siteMeta.FirstStart,
	}

	err := r.DB.Create(&siteMetaModel).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}

	return nil
}
