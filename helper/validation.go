package helper

import (
	"errors"
	"final-project/entity"
)

func CheckEmpty(input ...interface{}) error {
	for _, value := range input {
		switch value {
		case "":
			return errors.New("make sure your input not empties")
		case 0:
			return errors.New("make sure your input not zeros")
		case nil:
			return errors.New("make sure your input not nils")
		}
	}

	return nil
}

func CheckCommentInputEmpty(comment entity.CommentInput) error {
	if comment.Message == "" {
		return errors.New("message field must be filled")
	}
	if comment.PhotoID == 0 {
		return errors.New("photo id field must be filled or not zero")
	}

	return nil
}

func CheckSmInputEmpty(sm entity.SocialMediaInput) error {
	if sm.Name == "" {
		return errors.New("name field must be filled")
	}

	if sm.SocialMediaUrl == "" {
		return errors.New("social media url must be filled")
	}

	return nil
}

func CheckPhotoInputEmpty(photo entity.Photo) error {
	if photo.Title == "" {
		return errors.New("title must be filled")
	}
	if photo.Caption == "" {
		return errors.New("caption must be filled")
	}
	if photo.PhotoUrl == "" {
		return errors.New("photo url must be filled")
	}

	return nil

}
