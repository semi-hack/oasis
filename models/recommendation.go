package models

import (
	"time"

   "go.mongodb.org/mongo-driver/bson/primitive"
)

type Recommendation struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string             `json:"title" bson:"title"`
	Description       string             `json:"description" bson:"description"`
	CategoryId    primitive.ObjectID             `json:"categoryId" bson:"catgoryId"`
	CoverImage string `json:"coverImage" bson:"coverImage"`
	CreatorId  primitive.ObjectID `json:"creatorId" bson:"creatorId"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}