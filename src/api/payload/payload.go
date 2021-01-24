package payload

type Login struct{
	User string `json:"user"`
	Password string `json:"password"`
}