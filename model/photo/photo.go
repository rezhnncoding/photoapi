package photo

import "time"

type Photo struct {
	Id            string    `bson:"_id,omitempty"`
	Title         string    `bson:"Title,omitempty"`
	Description   string    `bson:"Description,omitempty"`
	ImageName     string    `bson:"ImageName,omitempty"`
	CreateDate    time.Time `bson:"CreateDate,omitempty"`
	CreatorUserId string    `bson:"CreatorUserId,omitempty"`
	VisitCount    int       `bson:"VisitCount,omitempty"`
	LikeCount     int       `bson:"LikeCount,omitempty"`
}
