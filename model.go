package main

import (
	"github.com/Tusharpaul231/RSS-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ApiKey    string `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID: 	  user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		ApiKey:   user.ApiKey,
	}
}

type Feed struct {
	ID       uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Url      string `json:"url"`
	UserID   uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed{
	return Feed{
		ID: 	  dbFeed.ID,
		Username: dbFeed.Username,
		Email:    dbFeed.Email,
		Url:       dbFeed.Url,
		UserID:   dbFeed.UserID,
	}
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = databaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}