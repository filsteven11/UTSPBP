package Controller

import (
	"PBPUTS/Model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	// Query untuk mendapatkan semua ruangan
	query := "SELECT id, room_name FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data ruangan")
		return
	}
	defer rows.Close()

	var rooms []Model.Room
	for rows.Next() {
		var room Model.Room
		if err := rows.Scan(&room.ID, &room.RoomName); err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, "Gagal membaca data room")
			return
		}
		rooms = append(rooms, room)
	}

	response := struct {
		Status string       `json:"status"`
		Data   []Model.Room `json:"data"`
	}{
		Status: "success",
		Data:   rooms,
	}
	SendJSONResponse(w, http.StatusOK, response)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	response := Model.Response{Status: "error", Message: message}
	SendJSONResponse(w, code, response)
}

func SendJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	roomIdStr := params.Get("id")
	if roomIdStr == "" {
		SendErrorResponse(w, http.StatusBadRequest, "ID ruangan tidak valid")
		return
	}

	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "ID ruangan tidak valid")
		return
	}

	room, err := Model.RoomDetail(roomId)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil detail ruangan")
		return
	}

	participants, err := Model.GetParticipantsRoom(roomId)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data partisipan ruangan")
		return
	}

	// Construct the response
	response := struct {
		Status string           `json:"status"`
		Data   Model.RoomDetail `json:"data"`
	}{
		Status: "success",
		Data: Model.RoomDetail{
			Room:         room,
			Participants: participants,
		},
	}

	SendJSONResponse(w, http.StatusOK, response)
}
