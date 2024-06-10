package error

import "fmt"

// 此處定義一個 ValidationError 的 struct，用來表示 reqeust 參數驗證錯誤
type ValidationError struct {
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(e.Message)
}
