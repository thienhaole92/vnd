package event

import (
	"encoding/json"

	"github.com/thienhaole92/vnd/logger"
)

/*
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
*/
type EventSchema[T any] struct {
	Specversion string `json:"specversion"`
	Type        string `json:"type"`
	Source      string `json:"source"`
	Id          string `json:"id"`
	Time        string `json:"time"`
	Subject     string `json:"subject"`
	Data        T      `json:"data"`
}

type EventString string

func (es *EventString) UnpackEvent(data interface{}) error {
	log := logger.GetLogger("UnpackEvent")
	defer log.Sync()

	if len(*es) == 0 {
		return nil
	}

	return json.Unmarshal([]byte(*es), data)
}
