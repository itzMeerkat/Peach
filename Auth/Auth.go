package Auth

func Login(usrname string, pw string) int {
	if usrname == "AG3" && pw == "123456" {
		return 0
	}
	return -1
}
