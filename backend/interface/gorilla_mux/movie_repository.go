package gorilla_mux

import "net/http"

type MovieMuxRepository struct {
	GorillaMuxHandler
}

func (repo *MovieMuxRepository) Find(r *http.Request, key string) string {
	return repo.Get(r, key)
}
