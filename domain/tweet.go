package domain

import (
	"context"

	"github.com/fadedreams/xclone"
	"github.com/fadedreams/xclone/uuid"
)

type TweetService struct {
	TweetRepo xclone.TweetRepo
}

func NewTweetService(tr xclone.TweetRepo) *TweetService {
	return &TweetService{
		TweetRepo: tr,
	}
}

func (ts *TweetService) All(ctx context.Context) ([]xclone.Tweet, error) {
	return ts.TweetRepo.All(ctx)
}

func (ts *TweetService) GetByParentID(ctx context.Context, id string) ([]xclone.Tweet, error) {
	return ts.TweetRepo.GetByParentID(ctx, id)
}

func (ts *TweetService) Create(ctx context.Context, input xclone.CreateTweetInput) (xclone.Tweet, error) {
	currentUserID, err := xclone.GetUserIDFromContext(ctx)
	if err != nil {
		return xclone.Tweet{}, xclone.ErrUnauthenticated
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return xclone.Tweet{}, err
	}

	tweet, err := ts.TweetRepo.Create(ctx, xclone.Tweet{
		Body:   input.Body,
		UserID: currentUserID,
	})
	if err != nil {
		return xclone.Tweet{}, err
	}

	return tweet, nil
}

func (ts *TweetService) GetByID(ctx context.Context, id string) (xclone.Tweet, error) {
	if !uuid.Validate(id) {
		return xclone.Tweet{}, xclone.ErrInvalidUUID
	}

	return ts.TweetRepo.GetByID(ctx, id)
}

func (ts *TweetService) Delete(ctx context.Context, id string) error {
	currentUserID, err := xclone.GetUserIDFromContext(ctx)
	if err != nil {
		return xclone.ErrUnauthenticated
	}

	if !uuid.Validate(id) {
		return xclone.ErrInvalidUUID
	}

	tweet, err := ts.TweetRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if !tweet.CanDelete(xclone.User{ID: currentUserID}) {
		return xclone.ErrForbidden
	}

	return ts.TweetRepo.Delete(ctx, id)
}

func (ts *TweetService) CreateReply(ctx context.Context, parentID string, input xclone.CreateTweetInput) (xclone.Tweet, error) {
	currentUserID, err := xclone.GetUserIDFromContext(ctx)
	if err != nil {
		return xclone.Tweet{}, xclone.ErrUnauthenticated
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return xclone.Tweet{}, err
	}

	if !uuid.Validate(parentID) {
		return xclone.Tweet{}, xclone.ErrInvalidUUID
	}

	if _, err := ts.TweetRepo.GetByID(ctx, parentID); err != nil {
		return xclone.Tweet{}, xclone.ErrNotFound
	}

	tweet, err := ts.TweetRepo.Create(ctx, xclone.Tweet{
		Body:     input.Body,
		UserID:   currentUserID,
		ParentID: &parentID,
	})
	if err != nil {
		return xclone.Tweet{}, err
	}

	return tweet, nil
}
