package data

type Response struct{
	Code   int  `json:"code"`
	Status int  `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type PositionResponse struct{
    RoomId    string  `json:"roomId"`
	MetaId    string  `json:"metaId"`
	X_position string `json:"xPosition"`
	Y_position string `json:"yposition"`
}