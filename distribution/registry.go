package distribution

import (
	"net/http"

	"github.com/iypetrov/go-cr/logger"
)

// OCI distribution spec can be found here
// https://github.com/opencontainers/distribution-spec/blob/main/spec.md
type Registry struct {
	storage  *Storage
	metadata *Metadata
	log      logger.Logger
}

func NewRegistry(
	storage *Storage,
	metadata *Metadata,
	log logger.Logger,
) *Registry {
	return &Registry{
		storage:  storage,
		metadata: metadata,
		log:      log,
	}
}

func (reg *Registry) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (reg *Registry) listTags(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) getManifest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) checkBlobExists(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) getBlob(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) startBlobUpload(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) uploadBlobChunk(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) completeBlobUpload(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) deleteBlob(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) putManifest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) deleteManifest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (reg *Registry) getCatalog(w http.ResponseWriter, r *http.Request) error {
	return nil
}
