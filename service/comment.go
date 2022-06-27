package service

import (
	"context"
	"final-project/entity"
	"final-project/helper"
	"final-project/repository"
	"fmt"
)

type CommentServiceInterface interface {
	CreateComment(ctx context.Context, id string, input entity.CommentInput) (*entity.Comment, error)
	GetAllComment(ctx context.Context) ([]*entity.Comment, error)
	GetCommentById(ctx context.Context, id string) (*entity.Comment, error)
	UpdateComment(ctx context.Context, id string, input entity.CommentInput) (*entity.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}

type CommentService struct {
	commentRepo repository.CommentRepositoryInterface
}

func NewCommentService(commentRepo repository.CommentRepositoryInterface) CommentServiceInterface {
	return &CommentService{commentRepo: commentRepo}
}

func (c CommentService) CreateComment(ctx context.Context, id string, input entity.CommentInput) (*entity.Comment, error) {
	//TODO implement me
	var comment entity.Comment

	if err := helper.CheckCommentInputEmpty(input); err != nil {
		return nil, err
	}

	comment.Message = input.Message
	comment.PhotoID = input.PhotoID

	newComment, err := c.commentRepo.CreateComment(ctx, id, comment)

	if err != nil {
		fmt.Println("err-service-comm", err.Error())
		return nil, err
	}

	return newComment, nil
}

func (c CommentService) GetAllComment(ctx context.Context) ([]*entity.Comment, error) {
	//TODO implement me
	getAll, err := c.commentRepo.GetAllComment(ctx)

	if err != nil {
		return nil, err
	}

	return getAll, nil
}

func (c CommentService) GetCommentById(ctx context.Context, id string) (*entity.Comment, error) {
	//TODO implement me
	getById, err := c.commentRepo.GetCommentById(ctx, id)

	if err != nil {
		return nil, err
	}

	return getById, nil
}

func (c CommentService) UpdateComment(ctx context.Context, id string, input entity.CommentInput) (*entity.Comment, error) {
	//TODO implement me
	var comment entity.Comment

	if err := helper.CheckCommentInputEmpty(input); err != nil {
		return nil, err
	}

	comment.Message = input.Message

	updateComment, err := c.commentRepo.UpdateComment(ctx, id, comment)

	if err != nil {
		return nil, err
	}

	return updateComment, nil
}

func (c CommentService) DeleteComment(ctx context.Context, id string) error {
	//TODO implement me
	if err := c.commentRepo.DeleteComment(ctx, id); err != nil {
		return err
	}

	return nil
}
