package response

type RegisterResponse struct {
	Firstname string
	Lastname  string
	Email     string
	Token     string
	ID        interface{}
}

type Error struct {
	Msg string `json:"msg"`
	Ok string `json:"ok"`  
}