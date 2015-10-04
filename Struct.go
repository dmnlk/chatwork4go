package chatwork4go

type Status struct {
	AccountID        int    `json:"account_id"`
	AvatarImageURL   string `json:"avatar_image_url"`
	ChatworkID       string `json:"chatwork_id"`
	Department       string `json:"department"`
	Name             string `json:"name"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	RoomID           int    `json:"room_id"`
}
