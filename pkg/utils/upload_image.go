package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Dffarhn/bakulenapi/config"
)

// UploadImage uploads the WebP image to Firebase Storage

func UploadImage(filename string, fileContent []byte) (string, error) {
	// Get Firebase Storage bucket
	bucket := config.StorageClient.Bucket("bekaspakaistorage.appspot.com")
	folder := fmt.Sprintf("images/bakulen/%s", filename)

	// Create a writer for the file object
	object := bucket.Object(folder)
	writer := object.NewWriter(context.Background())
	writer.ContentType = "image/webp" // Set correct content type
	writer.ChunkSize = 0              // Write the entire file in one go

	// Write the WebP file content to Firebase Storage
	if _, err := writer.Write(fileContent); err != nil {
		log.Println("Error uploading image to Firebase Storage:", err)
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Close the writer
	if err := writer.Close(); err != nil {
		log.Println("Error closing Firebase Storage writer:", err)
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Generate a signed URL
	signedURL, err := getSignedURL(folder)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %v", err)
	}

	return signedURL, nil
}

// getSignedURL generates a signed URL for accessing the file
func getSignedURL(filename string) (string, error) {

    // Define signing options
    opts := &storage.SignedURLOptions{
        GoogleAccessID: "firebase-adminsdk-hedsy@bekaspakaistorage.iam.gserviceaccount.com",
        Method:         "GET",
        Expires:        time.Now().Add(24 * time.Hour),
    }

    url, err := config.GetFirebaseStorageClient().Bucket("bekaspakaistorage.appspot.com").SignedURL(filename, opts)
    if err != nil {
        return "", err
    }

    return url, nil
}

func GenerateUniqueFilename(path string) string {

	return fmt.Sprintf("%s/%d", path, time.Now().UnixNano())

}
