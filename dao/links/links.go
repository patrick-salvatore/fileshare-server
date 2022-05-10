package links

import "go.mongodb.org/mongo-driver/bson/primitive"

type LinksJson struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Url         string             `bson:"url" json:"url"`
	ExpiresDate string             `bson:"expiresDate" json:"expiresDate"`
	OpenedDate  string             `bson:"openedDate" json:"openedDate"`
}
