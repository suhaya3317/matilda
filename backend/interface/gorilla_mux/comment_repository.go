package gorilla_mux

import "net/http"

type MuxCommentRepository struct {
	GorillaMuxHandler
}

func (repo *MuxCommentRepository) Find(r *http.Request, key string) string {
	return repo.Get(r, key)
}
