package models

import "github.com/fairway-corp/swagchat-api/utils"

type RoomUsers struct {
	Users []string `json:"users,omitempty"`
}

type RoomUser struct {
	Id                      uint64         `json:"-" db:"id"`
	RoomId                  string         `json:"roomId,omitempty" db:"room_id"`
	UserId                  string         `json:"userId,omitempty" db:"user_id"`
	UnreadCount             *int64         `json:"unreadCount,omitempty" db:"unread_count"`
	MetaData                utils.JSONText `json:"metaData,omitempty" db:"meta_data"`
	NotificationSubscribeId *string        `json:"-" db:"notification_subscribe_id"`
	Created                 int64          `json:"created,omitempty" db:"created"`
	Modified                int64          `json:"modified,omitempty" db:"modified"`
	Deleted                 int64          `json:"-" db:"deleted"`
}

type ErrorRoomUser struct {
	UserId string         `json:"userId,omitempty"`
	Error  *ProblemDetail `json:"error"`
}

type ResponseRoomUser struct {
	RoomUsers []RoomUser      `json:"roomUsers,omitempty"`
	Errors    []ErrorRoomUser `json:"errors,omitempty"`
}
