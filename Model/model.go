package Model

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
}

type Participant struct {
	ID        int    `json:"id"`
	AccountID int    `json:"id_account"`
	Username  string `json:"username"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RoomDetail struct {
	Room         Room          `json:"room"`
	Participants []Participant `json:"participants"`
}
