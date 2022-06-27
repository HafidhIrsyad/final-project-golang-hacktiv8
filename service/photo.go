package service

import (
	"context"
	"final-project/entity"
	"final-project/helper"
	"final-project/repository"
	"fmt"
	"log"
)

type PhotoServiceInterface interface {
	CreatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error)
	GetPhotos(ctx context.Context) ([]*entity.Photo, error)
	GetPhotoById(ctx context.Context, id string) (*entity.Photo, error)
	UpdatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error)
	DeletePhoto(ctx context.Context, id string) error
}

type PhotoService struct {
	photoRepo   repository.PhotoRepoInterface
	commentRepo repository.CommentRepositoryInterface
}

func NewPhotoService(photoRepo repository.PhotoRepoInterface, commentRepo repository.CommentRepositoryInterface) PhotoServiceInterface {
	return &PhotoService{
		photoRepo:   photoRepo,
		commentRepo: commentRepo,
	}
}

func (p PhotoService) CreatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error) {
	//TODO implement me
	if err := helper.CheckPhotoInputEmpty(photo); err != nil {
		return nil, err
	}

	newPhoto, err := p.photoRepo.CreatePhoto(ctx, id, photo)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return newPhoto, nil
}

func (p PhotoService) GetPhotos(ctx context.Context) ([]*entity.Photo, error) {
	//TODO implement me
	photos, err := p.photoRepo.GetPhotos(ctx)

	if err != nil {
		fmt.Println("errGetAllPhotoSer", err.Error())
		return nil, err
	}

	return photos, nil
}

func (p PhotoService) GetPhotoById(ctx context.Context, id string) (*entity.Photo, error) {
	//TODO implement me
	photoId, err := p.photoRepo.GetPhotoById(ctx, id)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return photoId, nil
}

func (p PhotoService) UpdatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error) {
	//TODO implement me
	if err := helper.CheckPhotoInputEmpty(photo); err != nil {
		return nil, err
	}

	updatePhoto, err := p.photoRepo.UpdatePhoto(ctx, id, photo)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return updatePhoto, nil
}

func (p PhotoService) DeletePhoto(ctx context.Context, id string) error {
	err := p.photoRepo.DeletePhoto(ctx, id)

	if err != nil {
		fmt.Println("err-del-ser", err.Error())
		return err
	}

	return nil
}
