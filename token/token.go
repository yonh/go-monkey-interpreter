package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 未知的词法单元或字符
	EOF     = "EOF"     // 文件结尾

	// 标识符 字面量
	IDENT = "IDENT" // add, foo, x, y ...
	INT   = "INT"   // 1, 2, 3

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// 关键字列表
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// 方法会根据关键字表赖查找字符串是否为关键字，如果是则返回关键字的TokenType,否则为IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
