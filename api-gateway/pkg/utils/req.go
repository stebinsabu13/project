package utils

type UserDetails struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	MobileNum string `json:"mobilenum" binding:"required,min=10,max=10"`
}

type MethodReq struct {
	Method   int `json:"method" binding:"required"`
	WaitTime int `json:"waitTime" binding:"required"`
}
