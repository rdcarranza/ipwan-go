package login

type Request struct {
	Username      string `xml:"Username"`
	Password      string `xml:"Password"`
	Password_type int    `xml:"password_type"`
}
