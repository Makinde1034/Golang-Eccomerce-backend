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
	Ok bool `json:"ok"`  
}