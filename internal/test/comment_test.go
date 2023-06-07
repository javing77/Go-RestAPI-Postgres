//go:build integration
// +build integration

package test

import (
	"context"
	"testing"

	"github.com/javing77/go-rest-postgress/internal/comment"
	"github.com/javing77/go-rest-postgress/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("PostComment", func(t *testing.T) {
		db, err := db.NewDataBase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "test",
			Author: "test",
			Body:   "test",
		})

		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "test", newCmt.Slug)

	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := db.NewDataBase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})

}
