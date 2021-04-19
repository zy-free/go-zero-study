package httpx

const (
	ApplicationJson = "application/json"
	ApplicationCsv  = "application/csv"
	ContentEncoding = "Content-Encoding"
	ContentSecurity = "X-Content-Security"
	ContentType     = "Content-Type"
	KeyField        = "key"
	SecretField     = "secret"
	TypeField       = "type"
	CryptionType    = 1
)

const (
	CodeSignaturePass = iota
	CodeSignatureInvalidHeader
	CodeSignatureWrongTime
	CodeSignatureInvalidToken
)
