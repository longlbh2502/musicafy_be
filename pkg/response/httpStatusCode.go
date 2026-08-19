package response

const (
	ErrCodeSusscess     = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid

)

var msg = map[int]string{
	ErrCodeSusscess:     "success",
	ErrCodeParamInvalid: "email is invalid",
}
