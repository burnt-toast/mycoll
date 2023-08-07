package actions

import (
	"log"
	"mycoll/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

/*
Get collections for current user
*/
func GetCollections(c buffalo.Context) error {
	return c.Redirect(302, "/")
}

func AddCollection(c buffalo.Context) error {
	params := c.Params()
	collection := &models.Collection{}
	rName := params.Get("name")
	rDesc := params.Get("description")

	if rName == "" || rDesc == "" {
		return errors.New("Required inputs not found")
	}

	log.Println("About to get current user")
	u := c.Session().Get("current_user_id")
	log.Println("Current user: ", u)
	collection.Name = rName
	collection.Description = rDesc
	//collection.Owner = u.Name
	log.Println("About to create new collection:", collection)
	tx := c.Value("tx").(*pop.Connection)
	err := tx.Save(collection)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(302, "/")
}
