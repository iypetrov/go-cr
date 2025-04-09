package distribution

import (
	"fmt"
)

type Code string

const (
	BLOB_UNKNOWN          Code = "BLOB_UNKNOWN"
	BLOB_UPLOAD_INVALID   Code = "BLOB_UPLOAD_INVALID"
	BLOB_UPLOAD_UNKNOWN   Code = "BLOB_UPLOAD_UNKNOWN"
	DIGEST_INVALID        Code = "DIGEST_INVALID"
	MANIFEST_BLOB_UNKNOWN Code = "MANIFEST_BLOB_UNKNOWN"
	MANIFEST_INVALID      Code = "MANIFEST_INVALID"
	MANIFEST_UNKNOWN      Code = "MANIFEST_UNKNOWN"
	NAME_INVALID          Code = "NAME_INVALID"
	NAME_UNKNOWN          Code = "NAME_UNKNOWN"
	SIZE_INVALID          Code = "SIZE_INVALID"
	UNAUTHORIZED          Code = "UNAUTHORIZED"
	DENIED                Code = "DENIED"
	UNSUPPORTED           Code = "UNSUPPORTED"
	TOOMANYREQUESTS       Code = "TOOMANYREQUESTS"
)

type Error struct {
	StatusCode int
	Code       Code   `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"details"`
}

func (e Error) Error() string {
	return fmt.Sprintf("custom error: %s", e.Message)
}

type Errors struct {
	Errors []Error `json:"errors"`
}
