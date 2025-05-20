package interfaces

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type note struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt int64         `bson:"created_at" json:"created_at"`
	UpdatedAt int64         `bson:"updated_at" json:"updated_at"`
	IsDone    bool          `bson:"is_done" json:"is_done"`
	DeadLine  int64         `bson:"dead_line" json:"dead_line"`
}
