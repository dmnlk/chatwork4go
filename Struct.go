package chatwork4go

type Status struct {
	MentionNum     int `json:"mention_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskNum      int `json:"mytask_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	UnreadRoomNum  int `json:"unread_room_num"`
}
