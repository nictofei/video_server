package defs

//Request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"password"`
}

type VideoInfo struct{
  Id string
  AuthorId int
  Name string
  DisplayCtime string
}
