package operations

import (
	"io"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetZipArchiveRequest struct {
	// Type of archive to make
	ArchiveKind components.ArchiveKind `pathParam:"style=simple,explode=false,name=archiveKind"`
	// commit ID of the archive to get as a ZIP file
	CommitID string `pathParam:"style=simple,explode=false,name=commitId"`
}

func (o *GetZipArchiveRequest) GetArchiveKind() components.ArchiveKind {
	if o == nil {
		return components.ArchiveKind("")
	}
	return o.ArchiveKind
}

func (o *GetZipArchiveRequest) GetCommitID() string {
	if o == nil {
		return ""
	}
	return o.CommitID
}

type GetZipArchiveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Stream io.ReadCloser
}

func (o *GetZipArchiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetZipArchiveResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}
