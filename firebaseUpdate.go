package firebaseUpdate

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	app *firebase.App
	err error
	ctx context.Context
)

func init() {
	//firebase init
	ctx = context.Background()
	opt := option.WithCredentialsFile("./serviceAccount.json")
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
}

func DeleteRecord(collection string, docID string) error {
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	collectRef := client.Collection(collection)
	docRef := collectRef.Doc(docID)
	_, err = docRef.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetRecord(collection string, docID string) (map[string]interface{}, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	collectRef := client.Collection(collection)
	docRef := collectRef.Doc(docID)

	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc.Data(), nil
}

func GetRecords(collection string) ([]map[string]interface{}, error) {
	res := []map[string]interface{}{}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	collectRef := client.Collection(collection)

	iter := collectRef.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln("error getting document: ", err)
		}
		res = append(res, doc.Data())
	}

	return res, nil
}

func UpdateRecord(collection string, docID string, field string, result bool) error {

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
		return err
	}
	return nil
}
