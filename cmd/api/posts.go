package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/lib/pq"
	"github.com/sanjbh/social/internal/store"
)

type CreatePostsPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostsPayload

	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		fmt.Println(string(debug.Stack()))
		return
	}

	post := store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		// Tags:    payload.Tags,
		Tags: pq.Array(payload.Tags),
	}

	if err := app.store.Posts.Create(r.Context(), &post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		fmt.Println(string(debug.Stack()))
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		fmt.Println(string(debug.Stack()))
		return
	}
}
