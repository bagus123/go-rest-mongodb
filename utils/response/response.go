package response

// Response ...
type Response struct {
	Data interface{} `json:"data"`
}

// Error ...
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
