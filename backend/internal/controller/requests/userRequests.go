package requests

// 在這個檔案定義所有 users 路徑下的 APIs 可能會有的 request 型別
// 這邊的 request 型別是指從 client 端送來的 request body

// UserSignInRequest 是使用者登入的 request body
type UserSignInAndSignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
