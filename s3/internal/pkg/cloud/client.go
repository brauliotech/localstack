package cloud

import (
	"context"
	"io"
)

type BucketClient interface {
	Create(ctx context.Context, bucket string) error
	UploadObject(ctx context.Context, bucket, fileName string, body io.Reader) (string, error)
	DownloadObject(ctx context.Context, bucket, fileName string, body io.WriterAt) error
	DeleteObject(ctx context.Context, bucket, fileName string) error
	ListObjects(ctx context.Context, bucket string) ([]*Object, error)
	FetchObject(ctx context.Context, bucket, fileName string) (io.ReadCloser, error)
}
