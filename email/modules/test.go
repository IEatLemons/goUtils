package EmailModules

type Test struct {
	User    string `json:"user"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
	Number  string `json:"number"`
}

func (M *Test) GetTo() string {
	return M.User
}

func (M *Test) GetSubject() string {
	return M.Subject
}

func (M *Test) GetBody() string {
	body := "name : " + M.Name + "; number : " + M.Number
	return body
}

func (M *Test) GetMailType() string {
	return "text"
}