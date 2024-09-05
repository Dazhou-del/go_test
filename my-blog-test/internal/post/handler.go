package post

import (
	"net/http"
)

type PostHandler struct {
	postService IPostService
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	post, err := h.postService.GetPostById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(post))
}

func NewPostHandler(postService IPostService) *PostHandler {
	return &PostHandler{postService: postService}
}
