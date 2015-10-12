package chatwork4go

type Status struct {
	MentionNum     int `json:"mention_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskNum      int `json:"mytask_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	UnreadRoomNum  int `json:"unread_room_num"`
}

type Task struct {
	AssignedByAccount struct {
		AccountID      int    `json:"account_id"`
		AvatarImageURL string `json:"avatar_image_url"`
		Name           string `json:"name"`
	} `json:"assigned_by_account"`
	Body      string `json:"body"`
	LimitTime int    `json:"limit_time"`
	MessageID int    `json:"message_id"`
	Room      struct {
		IconPath string `json:"icon_path"`
		Name     string `json:"name"`
		RoomID   int    `json:"room_id"`
	} `json:"room"`
	Status string `json:"status"`
	TaskID int    `json:"task_id"`
}
