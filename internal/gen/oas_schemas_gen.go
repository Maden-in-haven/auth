// Code generated by ogen, DO NOT EDIT.

package gen

type AuthLoginPostInternalServerError struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *AuthLoginPostInternalServerError) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *AuthLoginPostInternalServerError) SetMessage(val OptString) {
	s.Message = val
}

func (*AuthLoginPostInternalServerError) authLoginPostRes() {}

type AuthLoginPostOK struct {
	// Токен доступа JWT.
	AccessToken OptString `json:"access_token"`
	// Токен обновления JWT.
	RefreshToken OptString `json:"refresh_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *AuthLoginPostOK) GetAccessToken() OptString {
	return s.AccessToken
}

// GetRefreshToken returns the value of RefreshToken.
func (s *AuthLoginPostOK) GetRefreshToken() OptString {
	return s.RefreshToken
}

// SetAccessToken sets the value of AccessToken.
func (s *AuthLoginPostOK) SetAccessToken(val OptString) {
	s.AccessToken = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *AuthLoginPostOK) SetRefreshToken(val OptString) {
	s.RefreshToken = val
}

func (*AuthLoginPostOK) authLoginPostRes() {}

type AuthLoginPostReq struct {
	// Логин пользователя.
	Username OptString `json:"username"`
	// Пароль пользователя.
	Password OptString `json:"password"`
}

// GetUsername returns the value of Username.
func (s *AuthLoginPostReq) GetUsername() OptString {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *AuthLoginPostReq) GetPassword() OptString {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *AuthLoginPostReq) SetUsername(val OptString) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *AuthLoginPostReq) SetPassword(val OptString) {
	s.Password = val
}

type AuthLoginPostUnauthorized struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *AuthLoginPostUnauthorized) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *AuthLoginPostUnauthorized) SetMessage(val OptString) {
	s.Message = val
}

func (*AuthLoginPostUnauthorized) authLoginPostRes() {}

type AuthRefreshPostInternalServerError struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *AuthRefreshPostInternalServerError) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *AuthRefreshPostInternalServerError) SetMessage(val OptString) {
	s.Message = val
}

func (*AuthRefreshPostInternalServerError) authRefreshPostRes() {}

type AuthRefreshPostOK struct {
	// Новый токен доступа JWT.
	AccessToken OptString `json:"access_token"`
	// Новый токен обновления JWT.
	RefreshToken OptString `json:"refresh_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *AuthRefreshPostOK) GetAccessToken() OptString {
	return s.AccessToken
}

// GetRefreshToken returns the value of RefreshToken.
func (s *AuthRefreshPostOK) GetRefreshToken() OptString {
	return s.RefreshToken
}

// SetAccessToken sets the value of AccessToken.
func (s *AuthRefreshPostOK) SetAccessToken(val OptString) {
	s.AccessToken = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *AuthRefreshPostOK) SetRefreshToken(val OptString) {
	s.RefreshToken = val
}

func (*AuthRefreshPostOK) authRefreshPostRes() {}

type AuthRefreshPostReq struct {
	// Токен обновления JWT.
	RefreshToken OptString `json:"refresh_token"`
}

// GetRefreshToken returns the value of RefreshToken.
func (s *AuthRefreshPostReq) GetRefreshToken() OptString {
	return s.RefreshToken
}

// SetRefreshToken sets the value of RefreshToken.
func (s *AuthRefreshPostReq) SetRefreshToken(val OptString) {
	s.RefreshToken = val
}

type AuthRefreshPostUnauthorized struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *AuthRefreshPostUnauthorized) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *AuthRefreshPostUnauthorized) SetMessage(val OptString) {
	s.Message = val
}

func (*AuthRefreshPostUnauthorized) authRefreshPostRes() {}

type AuthVerifyPostInternalServerError struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *AuthVerifyPostInternalServerError) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *AuthVerifyPostInternalServerError) SetMessage(val OptString) {
	s.Message = val
}

func (*AuthVerifyPostInternalServerError) authVerifyPostRes() {}

type AuthVerifyPostOK struct {
	Valid OptBool `json:"valid"`
}

// GetValid returns the value of Valid.
func (s *AuthVerifyPostOK) GetValid() OptBool {
	return s.Valid
}

// SetValid sets the value of Valid.
func (s *AuthVerifyPostOK) SetValid(val OptBool) {
	s.Valid = val
}

func (*AuthVerifyPostOK) authVerifyPostRes() {}

type AuthVerifyPostReq struct {
	// Токен доступа JWT.
	AccessToken OptString `json:"access_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *AuthVerifyPostReq) GetAccessToken() OptString {
	return s.AccessToken
}

// SetAccessToken sets the value of AccessToken.
func (s *AuthVerifyPostReq) SetAccessToken(val OptString) {
	s.AccessToken = val
}

type AuthVerifyPostUnauthorized struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *AuthVerifyPostUnauthorized) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *AuthVerifyPostUnauthorized) SetMessage(val OptString) {
	s.Message = val
}

func (*AuthVerifyPostUnauthorized) authVerifyPostRes() {}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
