package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/javing77/go-rest-postgress/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

// Transform de cmtRow response
// into a comment.Commeent
func converCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Slug.String,
		Body:   c.Body.String,
	}
}

func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {

	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id =$1`,
		uuid,
	)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid: %w", err)
	}

	return converCommentRowToComment(cmtRow), nil
}
