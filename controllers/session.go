package controllers

//func Signup(resp http.ResponseWriter, req *http.Request) {
//	//take mail, password
//	//if mail exist un db
//	//	Error
//	//crypt password
//	//create account
//	user := &models.User{}
//
//	err := json.NewDecoder(req.Body).Decode(user)
//	if err != nil {
//		rest.Send(resp, rest.Message(false, "Error - Invalid data"))
//		return
//	}
//
//	err := user.Create()
//	if err != nil {
//		rest.Send(resp, rest.Message(false, err.Error()))
//		return
//	}
//
//	rest.Respond(resp, rest.Message(true, "Success - User created"))
//}
//
//func LogIn(w http.ResponseWriter, r *http.Request) {
//	//take mail, password
//	//if password != db_password
//	//	error
//	//assign jwttoken
//}
//
//func LogOut(w http.ResponseWriter, r *http.Request) {
//	//empty jwttoken
//}
