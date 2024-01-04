package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
)

func (c *Firebase) UploadFile(ctx context.Context, file multipart.File, name, contentType string) (string, error) {
	bucket, err := c.Storage.DefaultBucket()
	if err != nil {
		return "", err
	}

	oh := bucket.Object(name)
	writer := oh.NewWriter(ctx)
	writer.ObjectAttrs = storage.ObjectAttrs{
		Name:        name,
		ContentType: contentType,
	}

	if _, err := io.Copy(writer, file); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	metaURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s", oh.BucketName(), url.QueryEscape(oh.ObjectName()))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, metaURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var meta map[string]any
	if err := json.Unmarshal(b, &meta); err != nil {
		return "", err
	}

	downloadURL := fmt.Sprintf("%s?alt=media&token=%s", metaURL, meta["downloadTokens"].(string))

	return downloadURL, nil
}
