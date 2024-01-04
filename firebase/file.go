package firebase

import (
	"errors"
	"mime/multipart"
	"path"
	"strings"
)

func ContentType(fh *multipart.FileHeader) string {
	return fh.Header.Get("Content-Type")
}

func Extension(fh *multipart.FileHeader) (string, error) {
	cd := fh.Header.Get("Content-Disposition")
	cd = strings.ReplaceAll(cd, "\"", "")

	splitted := strings.Split(cd, ";")
	if len(splitted) >= 3 {
		fileName := splitted[2]
		ext := path.Ext(fileName)

		return ext, nil
	}

	return "", errors.New("can not extract file extension")
}
