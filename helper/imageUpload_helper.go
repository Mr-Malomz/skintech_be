package helper

import (
	"context"
	"log"
	"mime/multipart"
	"time"

	"github.com/Mr-Malomz/skintech_be/config"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(file *multipart.FileHeader) (string, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	
	uploadURL, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: file.Filename, Folder: "skintech_be"})
	defer cancel()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return uploadURL.SecureURL, err
}
