package service

import (
	"github.com/benpate/data"
	"github.com/benpate/data/option"
	"github.com/benpate/derp"
	"github.com/benpate/exp"
	"github.com/benpate/ghost/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Group manages all interactions with the Group collection
type Group struct {
	collection data.Collection
}

// NewGroup returns a fully populated Group service
func NewGroup(collection data.Collection) Group {
	return Group{
		collection: collection,
	}
}

/*******************************************
 * COMMON DATA FUNCTIONS
 *******************************************/

// List returns an iterator containing all of the Groups who match the provided criteria
func (service *Group) List(criteria exp.Expression, options ...option.Option) (data.Iterator, error) {
	return service.collection.List(notDeleted(criteria), options...)
}

// Load retrieves an Group from the database
func (service *Group) Load(criteria exp.Expression, result *model.Group) error {
	if err := service.collection.Load(notDeleted(criteria), result); err != nil {
		return derp.Wrap(err, "service.Group", "Error loading Group", criteria)
	}

	return nil
}

// Save adds/updates an Group in the database
func (service *Group) Save(user *model.Group, note string) error {

	if err := service.collection.Save(user, note); err != nil {
		return derp.Wrap(err, "service.Group", "Error saving Group", user, note)
	}

	return nil
}

// Delete removes an Group from the database (virtual delete)
func (service *Group) Delete(user *model.Group, note string) error {

	if err := service.collection.Delete(user, note); err != nil {
		return derp.Wrap(err, "service.Group", "Error deleting Group", user, note)
	}

	return nil
}

/*******************************************
 * GENERIC DATA FUNCTIONS
 *******************************************/

// New returns a fully initialized model.Stream as a data.Object.
func (service *Group) ObjectNew() data.Object {
	result := model.NewGroup()
	return &result
}

func (service *Group) ObjectList(criteria exp.Expression, options ...option.Option) (data.Iterator, error) {
	return service.List(criteria, options...)
}

func (service *Group) ObjectLoad(criteria exp.Expression) (data.Object, error) {
	result := model.NewGroup()
	err := service.Load(criteria, &result)
	return &result, err
}

func (service *Group) ObjectSave(object data.Object, comment string) error {
	return service.Save(object.(*model.Group), comment)
}

func (service *Group) ObjectDelete(object data.Object, comment string) error {
	return service.Delete(object.(*model.Group), comment)
}

/*******************************************
 * CUSTOM QUERIES
 *******************************************/

// LoadByID loads a single model.Group object that matches the provided userID
func (service *Group) LoadByID(userID primitive.ObjectID, result *model.Group) error {
	criteria := exp.Equal("_id", userID)
	return service.Load(criteria, result)
}

// LoadByGroupname loads a single model.Group object that matches the provided token
func (service *Group) LoadByToken(token string, result *model.Group) error {

	// If the token *looks* like an ObjectID then try that first.  If it works, then return in triumph
	if userID, err := primitive.ObjectIDFromHex(token); err == nil {
		if err := service.LoadByID(userID, result); err == nil {
			return nil
		}
	}

	// Otherwise, use the token as a username
	criteria := exp.Equal("token", token)
	return service.Load(criteria, result)
}

// ListByGroup returns all users that match a provided group name
func (service *Group) ListByGroup(group string) (data.Iterator, error) {
	return service.List(exp.Equal("groupId", group))
}