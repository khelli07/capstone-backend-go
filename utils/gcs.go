package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

var storageClient *storage.Client
var storageContext context.Context

func InitGCS() {
	storageContext = context.Background()
	var err error
	storageClient, err = storage.NewClient(storageContext, option.WithCredentialsFile(os.Getenv("GCS_CREDENTIALS")))
	if err != nil {
		log.Println(err)
	}
	log.Println("GCS initialized")
}

func UploadFile(f multipart.File, uploadedFile *multipart.FileHeader) (string, error) {
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), uploadedFile.Filename)
	fmt.Println(os.Getenv("GCS_BUCKET"), fileName, storageContext)

	sw := storageClient.Bucket(os.Getenv("GCS_BUCKET")).Object(fileName).NewWriter(storageContext)
	if _, err := io.Copy(sw, f); err != nil {
		return "", errors.Wrap(err, "Failed to upload file")
	}

	if err := sw.Close(); err != nil {
		return "", errors.Wrap(err, "Failed to upload file")
	}

	u, err := url.Parse("/" + os.Getenv("GCS_BUCKET") + "/" + sw.Attrs().Name)
	if err != nil {
		return "", errors.Wrap(err, "Failed to upload file")
	}

	return u.String(), nil
}
