package service

import (
	"context"
	"final-project/entity"
	"final-project/helper"
	"final-project/repository"
)

type SocialMediaServiceInterface interface {
	CreateSocialMedia(ctx context.Context, id string, sm entity.SocialMediaInput) (*entity.SocialMedia, error)
	GetAllSocialMedia(ctx context.Context) ([]*entity.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id string) (*entity.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, id string, sm entity.SocialMediaInput) (*entity.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id string) error
}

type SocialMediaService struct {
	socialMediaService repository.SocialMediaRepoInterface
}

func NewSocialMediaService(socialMediaService repository.SocialMediaRepoInterface) SocialMediaServiceInterface {
	return &SocialMediaService{socialMediaService: socialMediaService}
}

func (sv SocialMediaService) CreateSocialMedia(ctx context.Context, id string, sm entity.SocialMediaInput) (*entity.SocialMedia, error) {
	//TODO implement me
	var socialMedia entity.SocialMedia

	if err := helper.CheckSmInputEmpty(sm); err != nil {
		return nil, err
	}

	socialMedia.Name = sm.Name
	socialMedia.SocialMediaUrl = sm.SocialMediaUrl

	newSm, err := sv.socialMediaService.CreateSocialMedia(ctx, id, socialMedia)

	if err != nil {
		return nil, err
	}

	return newSm, nil
}

func (sv SocialMediaService) GetAllSocialMedia(ctx context.Context) ([]*entity.SocialMedia, error) {
	//TODO implement me
	getSmAll, err := sv.socialMediaService.GetAllSocialMedia(ctx)

	if err != nil {
		return nil, err
	}

	return getSmAll, nil
}

func (sv SocialMediaService) GetSocialMediaById(ctx context.Context, id string) (*entity.SocialMedia, error) {
	//TODO implement me
	getSmById, err := sv.socialMediaService.GetSocialMediaById(ctx, id)

	if err != nil {
		return nil, err
	}

	return getSmById, nil
}

func (sv SocialMediaService) UpdateSocialMedia(ctx context.Context, id string, sm entity.SocialMediaInput) (*entity.SocialMedia, error) {
	//TODO implement me
	var socialMedia entity.SocialMedia

	if err := helper.CheckSmInputEmpty(sm); err != nil {
		return nil, err
	}

	socialMedia.Name = sm.Name
	socialMedia.SocialMediaUrl = sm.SocialMediaUrl

	updateSm, err := sv.socialMediaService.UpdateSocialMedia(ctx, id, socialMedia)

	if err != nil {
		return nil, err
	}

	return updateSm, nil
}

func (sv SocialMediaService) DeleteSocialMedia(ctx context.Context, id string) error {
	//TODO implement me
	err := sv.socialMediaService.DeleteSocialMedia(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
