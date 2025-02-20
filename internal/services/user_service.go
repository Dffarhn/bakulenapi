package service

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Dffarhn/bakulenapi/config"
	"github.com/Dffarhn/bakulenapi/internal/models"
)

type UserService struct {
	FirestoreClient *firestore.Client
}

func NewUserService() *UserService {
	return &UserService{
		FirestoreClient: config.GetFirestoreClient(),
	}
}

func (s *UserService) GetUser(id string) (*models.User, error) {
    userRef := s.FirestoreClient.Collection("users").Doc(id)
    userDoc, err := userRef.Get(context.Background())
    if err != nil {
        return nil, err
    }

    // Map Firestore document fields to the User struct
    var user models.User
    err = userDoc.DataTo(&user)
    if err != nil {
        return nil, err
    }

    // Set the user ID manually since Firestore does not include it in the document fields
    user.ID = userDoc.Ref.ID

    return &user, nil
}

// UpdateUser updates the user's fields in Firestore (e.g., username, profile_picture)
// UpdateUser updates the user's fields dynamically
func (s *UserService) UpdateUser(id string, data map[string]interface{}) error {
	// Reference to the user's document in Firestore
	userRef := s.FirestoreClient.Collection("users").Doc(id)

	// Prepare the updates dynamically
	var updates []firestore.Update

	// Check if username is provided and add it to the updates
	if name, ok := data["name"]; ok {
		updates = append(updates, firestore.Update{
			Path:  "name",
			Value: name,
		})
	}

	// Check if profile_picture is provided and add it to the updates
	if profilePicture, ok := data["profile_picture"]; ok {
		updates = append(updates, firestore.Update{
			Path:  "profile_picture",
			Value: profilePicture,
		})
	}

	// Only proceed if there are updates to apply
	if len(updates) > 0 {
		_, err := userRef.Update(context.Background(), updates)
		if err != nil {
			log.Printf("Error updating user: %v", err)
			return fmt.Errorf("failed to update user: %v", err)
		}
	}

	return nil
}