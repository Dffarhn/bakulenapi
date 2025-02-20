package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var FirestoreClient *firestore.Client
var StorageClient *storage.Client
var AuthClient       *auth.Client

// InitFirebase initializes both Firestore and Storage clients
func InitFirebase() {
	ctx := context.Background()
	// Initialize Firestore client
	firestoreOpt := option.WithCredentialsFile("bakulendatabase-firebase-adminsdk-fbsvc-c5d1d48f7d.json")
	app, err := firebase.NewApp(ctx, nil, firestoreOpt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Firestore: %v", err)
	}
	FirebaseApp = app
	log.Println("Firebase Firestore initialized successfully")

	// Initialize Firestore client
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	AuthClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth client: %v", err)
	}
	log.Println("Firebase Auth initialized successfully")

	// Initialize Firebase Storage client
	storageOpt := option.WithCredentialsFile("bekaspakaistorage-firebase-adminsdk-hedsy-d41a469e13.json")
	storageClient, err := storage.NewClient(ctx, storageOpt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Storage client: %v", err)
	}
	StorageClient = storageClient
	log.Println("Firebase Storage initialized successfully")
}

// GetFirestoreClient returns the Firestore client instance
func GetFirestoreClient() *firestore.Client {
	return FirestoreClient
}

// GetFirebaseStorageClient returns the Firebase Storage client instance
func GetFirebaseStorageClient() *storage.Client {
	return StorageClient
}

func GetFirebaseAuthClient() *auth.Client{
	return AuthClient
}
