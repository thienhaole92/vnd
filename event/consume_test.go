package event

import (
	"context"
	"errors"
	"testing"

	"github.com/thienhaole92/vnd/logger"
)

func TestConsume(t *testing.T) {
	type BlockList struct {
		API             string `json:"api"`
		ClientRequestID string `json:"clientRequestId"`
		RequestID       string `json:"requestId"`
		ETag            string `json:"eTag"`
		ContentType     string `json:"contentType"`
		ContentLength   int    `json:"contentLength"`
		BlobType        string `json:"blobType"`
		URL             string `json:"url"`
		Sequencer       string `json:"sequencer"`
	}

	testcases := []struct {
		name      string
		event     EventString
		delegator func(*logger.Logger, context.Context, *EventSchema[BlockList]) error
		expect    error
	}{
		{
			"can consume valid event",
			`
			{
				"specversion": "1.0",
				"type": "Microsoft.Storage.BlobCreated",
				"source": "/subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.Storage/storageAccounts/{storage-account}",
				"id": "9aeb0fdf-c01e-0131-0922-9eb54906e209",
				"time": "2019-11-18T15:13:39.4589254Z",
				"subject": "blobServices/default/containers/{storage-container}/blobs/{new-file}",
				"data": {
					"api": "PutBlockList",
					"clientRequestId": "4c5dd7fb-2c48-4a27-bb30-5361b5de920a",
					"requestId": "9aeb0fdf-c01e-0131-0922-9eb549000000",
					"eTag": "0x8D76C39E4407333",
					"contentType": "image/png",
					"contentLength": 30699,
					"blobType": "BlockBlob",
					"url": "https://gridtesting.blob.core.windows.net/testcontainer/{new-file}",
					"sequencer": "000000000000000000000000000099240000000000c41c18"
				}
			}
			`,
			func(*logger.Logger, context.Context, *EventSchema[BlockList]) error {
				return nil
			},
			nil,
		},
		{
			"should return error while consume invalid event",
			`
			{
				"specversion": "1.0",
				"type": "Microsoft.Storage.BlobCreated",
				"source": "/subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.Storage/storageAccounts/{storage-account}",
				"id": "9aeb0fdf-c01e-0131-0922-9eb54906e209",
				"time": "2019-11-18T15:13:39.4589254Z",
				"subject": "blobServices/default/containers/{storage-container}/blobs/{new-file}",
				"data": {
					"api": "PutBlockList",
				}
			}
			`,
			func(*logger.Logger, context.Context, *EventSchema[BlockList]) error {
				return nil
			},
			errors.New("invalid character '}' looking for beginning of object key string"),
		},
		{
			"should return error if delegator throws an error",
			`
			{}
			`,
			func(*logger.Logger, context.Context, *EventSchema[BlockList]) error {
				return errors.New("delegator error")
			},
			errors.New("delegator error"),
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			out := Consume(context.TODO(), tt.event, "Testing", tt.delegator)
			if out != nil {
				if out.Error() != tt.expect.Error() {
					t.Errorf("want %v, got %v", tt.expect, out)
				}
			} else {
				if out != tt.expect {
					t.Errorf("want %v, got %v", tt.expect, out)
				}
			}
		})
	}
}
