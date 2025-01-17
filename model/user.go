package model

import (
	"time"

	"github.com/benpate/data/journal"
	"github.com/benpate/null"
	"github.com/benpate/schema"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a person or machine account that can own pages and sections.
type User struct {
	UserID      primitive.ObjectID   `path:"userId"      json:"userId"      bson:"_id"`         // Unique identifier for this user.
	GroupIDs    []primitive.ObjectID `path:"groupIds"    json:"groupIds"    bson:"groupIds"`    // Slice of IDs for the groups that this user belongs to.
	Identities  []string             `path:"identities"  json:"identities"  bson:"identities"`  // Slice of globally unique identities for contacting this user.
	DisplayName string               `path:"displayName" json:"displayName" bson:"displayName"` // Name to be displayed for this user
	Description string               `path:"description" json:"description" bson:"description"` // Status summary for this user (used by ActivityPub)
	Username    string               `path:"username"    json:"username"    bson:"username"`    // This is the primary public identifier for the user.
	Password    string               `path:"password"    json:"password"    bson:"password"`    // This password should be encrypted with BCrypt.
	IsOwner     bool                 `path:"isOwner"     json:"isOwner"     bson:"isOwner"`     // If TRUE, then this user is a website owner with FULL privileges.
	AvatarURL   string               `path:"avatarUrl"   json:"avatarUrl"   bson:"avatarUrl"`   // Avatar image of this user.
	InboxID     primitive.ObjectID   `path:"inboxId"     json:"inboxId"     bson:"inboxId"`     // ID of the parent stream for storing this user's social inbox.
	OutboxID    primitive.ObjectID   `path:"outboxId"    json:"outboxId"    bson:"outboxId"`    // ID of the parent stream for storing this user's social outbox.

	journal.Journal `json:"journal" bson:"journal"`
}

func NewUser() User {
	return User{
		UserID:     primitive.NewObjectID(),
		GroupIDs:   make([]primitive.ObjectID, 0),
		Identities: make([]string, 0),
	}
}

// ID returns the primary key for this record
func (user *User) ID() string {
	return user.UserID.Hex()
}

// Copy returns a duplicate copy of this User
// NOTE: This must NOT be a pointer receiver, so that a true COPY
// of this record is returned.
func (user User) Copy() User {
	return user
}

func (user *User) Schema() schema.Schema {
	return schema.Schema{
		ID: "whisper.model.user",
		Element: schema.Object{
			Properties: map[string]schema.Element{
				"userId":      schema.String{Format: "objectId"},
				"groupIds":    schema.Array{Items: schema.String{Format: "objectId"}},
				"displayName": schema.String{MaxLength: null.NewInt(50)},
				"username":    schema.String{MaxLength: null.NewInt(50)},
				"avatarUrl":   schema.String{MaxLength: null.NewInt(100)},
			},
		},
	}
}

/******************************
 Steranko Interfaces
*******************************/

// GetUsername returns the username for this User.  A part of the "steranko.User" interface.
func (user *User) GetUsername() string {
	return user.Username
}

// GetPassword returns the (encrypted) passsword for this User.  A part of the "steranko.User" interface.
func (user *User) GetPassword() string {
	return user.Password
}

// SetUsername updates the username for this User.  A part of the "steranko.User" interface.
func (user *User) SetUsername(username string) {
	user.Username = username
}

// SetPassword updates the password for this User.  A part of the "steranko.User" interface.
func (user *User) SetPassword(password string) {
	user.Password = password
}

// Claims returns all access privileges given to this user.  A part of the "steranko.User" interface.
func (user *User) Claims() jwt.Claims {

	result := Authorization{
		UserID:      user.UserID,
		GroupIDs:    user.GroupIDs,
		DomainOwner: user.IsOwner,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),                   // Current create date.  (Used by Steranko to refresh tokens)
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(10, 0, 0)), // Expires ten years from now (but re-validated sooner by Steranko)
		},
	}

	return result
}
