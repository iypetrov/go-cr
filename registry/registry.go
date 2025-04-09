package registry

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/go-cr/logger"
)

// OCI distribution spec can be found here
// https://github.com/opencontainers/distribution-spec/blob/main/spec.md
type OCIDistribution struct {
	log logger.Logger
}

func New(log logger.Logger) *OCIDistribution {
	return &OCIDistribution{
		log: log,
	}
}

func (ocid *OCIDistribution) Router() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", ocid.healthcheck)

		// Content Discovery
		r.Get("/{name}/tags/list", ocid.listTags)
		r.Get("/{name}/manifests/{reference}", ocid.getManifest)

		// Content Management
		r.Head("/{name}/blobs/{digest}", ocid.checkBlobExists)
		r.Get("/{name}/blobs/{digest}", ocid.getBlob)
		r.Post("/{name}/blobs/uploads/", ocid.startBlobUpload)
		r.Patch("/{name}/blobs/uploads/{uuid}", ocid.uploadBlobChunk)
		r.Put("/{name}/blobs/uploads/{uuid}", ocid.completeBlobUpload)
		r.Delete("/{name}/blobs/{digest}", ocid.deleteBlob)
		r.Put("/{name}/manifests/{reference}", ocid.putManifest)
		r.Delete("/{name}/manifests/{reference}", ocid.deleteManifest)

		// Registry catalog
		r.Get("/_catalog", ocid.getCatalog)
	}
}

func (ocid *OCIDistribution) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (ocid *OCIDistribution) listTags(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) getManifest(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) checkBlobExists(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) getBlob(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) startBlobUpload(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) uploadBlobChunk(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) completeBlobUpload(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) deleteBlob(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) putManifest(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) deleteManifest(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}

func (ocid *OCIDistribution) getCatalog(w http.ResponseWriter, r *http.Request) {
	ocid.log.Info("Hello")
}
