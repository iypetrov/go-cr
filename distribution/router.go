package distribution

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func makeDistributionHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var e Error
			if errors.As(err, &e) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(e.StatusCode)
				json.NewEncoder(w).Encode(
					Errors{
						Errors: []Error{e},
					},
				)
			}
		}
	}
}

func Router(reg *Registry) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", reg.healthcheck)

		// Content Discovery
		r.Get("/{name}/tags/list", makeDistributionHandler(reg.listTags))
		r.Get("/{name}/manifests/{reference}", makeDistributionHandler(reg.getManifest))

		// Content Management
		r.Head("/{name}/blobs/{digest}", makeDistributionHandler(reg.checkBlobExists))
		r.Get("/{name}/blobs/{digest}", makeDistributionHandler(reg.getBlob))
		r.Post("/{name}/blobs/uploads/", makeDistributionHandler(reg.startBlobUpload))
		r.Patch("/{name}/blobs/uploads/{uuid}", makeDistributionHandler(reg.uploadBlobChunk))
		r.Put("/{name}/blobs/uploads/{uuid}", makeDistributionHandler(reg.completeBlobUpload))
		r.Delete("/{name}/blobs/{digest}", makeDistributionHandler(reg.deleteBlob))
		r.Put("/{name}/manifests/{reference}", makeDistributionHandler(reg.putManifest))
		r.Delete("/{name}/manifests/{reference}", makeDistributionHandler(reg.deleteManifest))

		// Registry catalog
		r.Get("/_catalog", makeDistributionHandler(reg.getCatalog))
	}
}
