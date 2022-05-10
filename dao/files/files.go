package files

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/patrick-salvatore/fileshare-service/db"
)

type FileJson struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	FileEx       string             `bson:"fileEx" json:"fileEx"` // 'pdf' | 'jpg' | 'doc' | 'any'
	Description  string             `bson:"description" json:"description"`
	UploadedDate string             `bson:"uploadedDate" json:"uploadedDate"`
	Links        string             `bson:"links" json:"links"`
}

type Files []FileJson

var ErrorMethodNotAllowed = "method Not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

var (
	ErrorInvalidFileData = "invalid file data"
)

func (f *FileJson) GoString() {
	log.WithFields(log.Fields{
		"Id":           f.Id,
		"Name":         f.Name,
		"FileEx":       f.FileEx,
		"Description":  f.Description,
		"UploadedDate": f.UploadedDate,
	}).Info("File JSON")
}

const (
	COLLECTION = "files"
)

func GetFiles() (Files, error) {
	filingCabinet := db.DatabaseInstance.GetCollection(COLLECTION)
	var results []FileJson
	cur, err := filingCabinet.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return []FileJson{}, err
	}
	cur.All(context.TODO(), &results)

	fmt.Printf("Found multiple documents: %+v\n", results)

	return results, nil
}

func GetFile(req *http.Request) (FileJson, error) {
	vars := mux.Vars(req)
	var fileId = vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var searchId primitive.ObjectID
	searchId, _ = primitive.ObjectIDFromHex(fileId)

	var result FileJson
	err := db.DatabaseInstance.GetCollection(COLLECTION).FindOne(ctx, bson.D{{"_id", searchId}}).Decode(&result)
	if err != nil {
		log.Info(err)

		return FileJson{}, err
	}
	fmt.Printf("found document %v", result)

	return result, nil
}

func PostFiles(req *http.Request) (FileJson, error) {
	var f = FileJson{UploadedDate: time.Now().Format(time.RFC3339), Id: primitive.NewObjectID()}

	if decode_err := json.NewDecoder(req.Body).Decode(&f); decode_err != nil {
		log.Fatal(decode_err)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		fileResult, err := db.DatabaseInstance.GetCollection(COLLECTION).InsertOne(ctx, f)
		if err != nil {
			log.Fatal(err)

		} else {
			fmt.Printf("Inserted %v documents into episode collection!\n", fileResult)
		}
	}

	return f, nil
}
