package firebaseUpdate

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func UpdateRecord(collection string, docID string, field string, result bool) error {
	//firebase init
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	collectRef := client.Collection(collection)
	docRef := collectRef.Doc(docID)

	updateObject := firestore.Update{}
	if field == "email" {
		updateObject.Path = "subEmailVerified"
		updateObject.Value = result
	} else if field == "phone" {
		updateObject.Path = "subEmailVerified"
		updateObject.Value = result
	} else if field == "IsProcessed" {
		updateObject.Path = "IsProcessed"
		updateObject.Value = result
	} else {
		return errors.New("field is not valid")
	}

	_, err = docRef.Update(ctx, []firestore.Update{updateObject})
	if err != nil {
		log.Fatalln("firestore error: ", err)
	}
	// log.Println("write result: ", wr)

	return nil
}
