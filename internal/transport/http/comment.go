package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/javing77/go-rest-postgress/internal/comment"
)

type Response struct {
	Message string
}

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type PostCommentRequest struct {
	Slug   string `json:"slog" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertPostCommentRequestToComment(cmt PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   cmt.Slug,
		Author: cmt.Author,
		Body:   cmt.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convertdComment := convertPostCommentRequestToComment(cmt)
	postedComment, err := h.Service.PostComment(r.Context(), convertdComment)
	if err != nil {
		log.Println(err)
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Panicln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Succesfully deleted"}); err != nil {
		panic(err)
	}

}
