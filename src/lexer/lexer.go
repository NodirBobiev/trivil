package lexer

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"trivil/env"
)

var _ = fmt.Printf

type Lexer struct {
	source *env.Source
	src    []byte

	// состояние
	ch       rune // текущий unicode символ
	offset   int  // смещение (начало) текущего символа
	rdOffset int  // позиция чтения
	uOffset  int  // смещение в unicode символах
}

func (s *Lexer) Init(source *env.Source) {
	s.source = source
	s.src = source.Bytes

	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0
	s.uOffset = 0

	// пропустить BOM
	if len(s.src) >= 3 && s.src[0] == 0xEF && s.src[1] == 0xBB && s.src[2] == 0xBF {
		s.offset = 3
	}

	s.source.AddLine(s.uOffset)
	s.next()
}

func (s *Lexer) error(ofs int, id string, args ...interface{}) {
	env.AddError(s.source.MakePos(ofs), id, args...)
}

// Read the next Unicode char into s.ch.
// s.ch < 0 means end-of-file.
func (s *Lexer) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		r := rune(s.src[s.rdOffset])
		w := 1
		switch {
		case r == 0:
			s.error(s.offset, "ЛЕК-ОШ-СИМ", fmt.Sprintf("%#U", rune(0)))
			s.uOffset++
		case r == '\n':
			s.uOffset++
			s.source.AddLine(s.uOffset)
		case r == '\r':
			r = '\n'
			s.uOffset++
			if s.rdOffset+1 < len(s.src) && rune(s.src[s.rdOffset+1]) == '\n' {
				s.uOffset++
				w = 2
			}
			s.source.AddLine(s.uOffset)
		case r >= utf8.RuneSelf:
			// not ASCII
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				s.error(s.offset, "ЛЕК-UTF8")
			}
			s.uOffset++
		default:
			s.uOffset++ // ascii
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		s.source.AddLine(s.uOffset)
		s.ch = -1 // eof
	}
}

func (s *Lexer) peek() byte {
	if s.rdOffset < len(s.src) {
		return s.src[s.rdOffset]
	}
	return 0
}

//====

func lower(ch rune) rune     { return ('a' - 'A') | ch }
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isHex(ch rune) bool     { return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f' }

func isLetter(ch rune) bool {
	return ch >= utf8.RuneSelf && unicode.IsLetter(ch) ||
		'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch == '№'

}

func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func (s *Lexer) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' {
		s.next()
	}
}

func (s *Lexer) scanModifier() string {
	ofs := s.offset - 1
	if !isLetter(s.ch) {
		s.error(s.offset, "ЛЕК-МОДИФИКАТОР")
		return ""
	}
	for isLetter(s.ch) {
		s.next()
	}
	return string(s.src[ofs:s.offset])
}

func (s *Lexer) scanLineComment() int {
	// Первый '/' уже взят
	ofs := s.offset - 1
	s.next()
	for s.ch != '\n' && s.ch >= 0 {
		s.next()
	}
	if s.ch == '\n' {
		s.next() // перешли на след. символ после комментария
	}
	return ofs
}

func (s *Lexer) scanBlockComment() int {
	// '/' уже взят
	ofs := s.offset - 1
	s.next() // '*'
	for true {
		if s.ch < 0 {
			s.error(ofs, "ЛЕК-НЕТ-*/")
			break
		}
		ch := s.ch
		s.next()
		if ch == '*' && s.ch == '/' {
			s.next()
			break
		} else if ch == '/' && s.ch == '*' {
			s.scanBlockComment()
		}
	}
	return ofs
}

func (s *Lexer) scanIdentifier() (Token, string) {
	var ofs = s.offset
	var wordStart = s.offset
	// состояние после слова
	var we_ch rune
	var we_offset, we_rdOffest, we_uOffset int

	for {
		for {
			s.next()
			if !isLetter(s.ch) && !isDigit(s.ch) {
				break
			}
		}
		//fmt.Printf("word '%s' %d:%d\n", string(s.src[wordStart:s.offset]), wordStart, s.offset)
		if s.ch == '-' {
			_ch, _offset, _rdOffest, _uOffset := s.ch, s.offset, s.rdOffset, s.uOffset
			s.next()
			if !isLetter(s.ch) {
				s.ch, s.offset, s.rdOffset, s.uOffset = _ch, _offset, _rdOffest, _uOffset
				break
			}
		} else if s.ch == '!' || s.ch == '?' {
			s.next()
			break
		} else {
			// это пробел или не часть идентификатора, проверяю последнее слово на keyword
			lit := string(s.src[wordStart:s.offset])
			//fmt.Printf("проверка keyword '%s' %d:%d rdO=%d\n", lit, wordStart, s.offset, s.rdOffset)
			tok := Lookup(lit)
			if tok != IDENT {
				if wordStart == ofs {
					// первое слово - ключевое
					return tok, lit
				} else {
					// второе слово - ключевое, возвращаю предыдущую часть
					// ключевое слово будет взято следующим вызовом
					//fmt.Printf("берем предудущую часть %d:%d\n", ofs, wordEnd)
					lit := string(s.src[ofs:we_offset])
					//fmt.Printf("not keyword '%s' %d:%d next rdO=%d\n", lit, ofs, wordEnd, wordEnd)
					s.ch, s.offset, s.rdOffset, s.uOffset = we_ch, we_offset, we_rdOffest, we_uOffset

					return IDENT, lit
				}
			}

			if s.ch != ' ' {
				break
			}

			we_ch, we_offset, we_rdOffest, we_uOffset = s.ch, s.offset, s.rdOffset, s.uOffset

			s.next()
			if !isLetter(s.ch) {
				s.ch, s.offset, s.rdOffset, s.uOffset = we_ch, we_offset, we_rdOffest, we_uOffset

				//fmt.Printf("выход на пробеле ofs=%d rdO=%d\n", s.offset, s.rdOffset)

				break
			}
			wordStart = s.offset
		}
	}

	lit := string(s.src[ofs:s.offset])

	return IDENT, lit
}

func (s *Lexer) scanString(opening rune) string {
	// первая кавычка уже взята
	ofs := s.offset

	for {
		ch := s.ch
		if ch == '\n' || ch < 0 {
			s.error(ofs, "ЛЕК-СТРОКА-ЗАВЕРШЕНИЕ")
			break
		}
		s.next()
		if ch == opening {
			break
		}
		if ch < ' ' {
			s.error(ofs, "ЛЕК-ОШ-СИМ", fmt.Sprintf("%#U", ch))
		} else if ch == '\\' {
			s.scanEscape(opening)
		}
	}

	if ofs >= s.offset {
		return "!" // ошибка уже выдана
	}

	return string(s.src[ofs : s.offset-1])
}

func (s *Lexer) scanSymbol(opening rune) string {
	// первая кавычка уже взята
	var ofs = s.offset

	var valid = true
	var n = 0

	for {
		ch := s.ch
		if ch == '\n' || ch < 0 {
			if valid {
				s.error(ofs, "ЛЕК-СТРОКА-ЗАВЕРШЕНИЕ")
				valid = false
			}
			break
		}
		s.next()
		if ch == opening {
			break
		}
		n++
		if ch < ' ' {
			s.error(ofs, "ЛЕК-ОШ-СИМ", fmt.Sprintf("%#U", ch))
			valid = false
		} else if ch == '\\' {
			if !s.scanEscape(opening) {
				valid = false
			}
		}
	}

	if valid && n != 1 {
		s.error(ofs, "ЛЕК-ОШ-ДЛИНА-СИМВОЛА")
	}

	if !valid && ofs >= s.offset {
		return "!" // ошибка уже выдана
	}

	return string(s.src[ofs : s.offset-1])
}

// Сканирует escape sequence. В случае ошибки возвращает false
// Не проверяет корректность
func (s *Lexer) scanEscape(quote rune) bool {
	ofs := s.offset

	var n int
	switch s.ch {
	case 'n', 'r', 't', '\'', '"', '\\':
		s.next()
		return true
	case 'u': // \uABCD
		n = 4
		s.next()
	default:
		if s.ch < 0 {
			s.error(ofs, "ЛЕК-ESCAPE-ЗАВЕРШЕНИЕ")
		} else {
			s.error(ofs, "ЛЕК-ОШ-ESCAPE")
		}
		return false
	}

	for n > 0 {
		d := uint32(digitVal(s.ch))
		if d >= 16 {
			if s.ch < 0 {
				s.error(s.offset, "ЛЕК-ESCAPE-ЗАВЕРШЕНИЕ")
			} else {
				s.error(s.offset, "ЛЕК-ОШ-СИМ", fmt.Sprintf("%#U", s.ch))
				return false
			}
		}
		s.next()
		n--
	}

	return true
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= lower(ch) && lower(ch) <= 'f':
		return int(lower(ch) - 'a' + 10)
	}
	return 16 // larger than any legal digit val
}

//==== number

func (s *Lexer) scanDigits(base int) {
	if base == 10 {
		for isDecimal(s.ch) {
			s.next()
		}
	} else if base == 16 {
		for isHex(s.ch) {
			s.next()
		}
	} else {
		panic("! wrong base")
	}
}

func (s *Lexer) scanNumber() (Token, string) {
	ofs := s.offset

	base := 10
	prefix := rune(0) // 0 - нет префикса, 'x' - 0x

	// целое
	if s.ch == '0' {
		s.next()
		if s.ch == 'x' {
			s.next()
			base = 16
			prefix = 'x'
		}

	}

	s.scanDigits(base)

	if s.ch != '.' {
		return INT, string(s.src[ofs:s.offset])
	}

	// дробная часть
	if prefix != rune(0) {
		s.error(s.offset, "ЛЕК-ВЕЩ-БАЗА")
	}

	s.next()
	s.scanDigits(10)

	return FLOAT, string(s.src[ofs:s.offset])
}

//====

func (s *Lexer) checkEqu(tok0, tok1 Token) Token {
	if s.ch == '=' {
		s.next()
		return tok1
	}
	return tok0
}

func (s *Lexer) checkNext(tok0 Token, next rune, tok1 Token) Token {
	if s.ch == next {
		s.next()
		return tok1
	}
	return tok0
}

func (s *Lexer) Scan() (pos int, tok Token, lit string) {

	s.skipWhitespace()

	// начало лексемы
	pos = s.source.MakePos(s.uOffset)
	ch := s.ch

	switch {
	case isLetter(ch):
		tok, lit = s.scanIdentifier()
	case isDecimal(ch):
		tok, lit = s.scanNumber()
	default:
		s.next() // всегда двигаемся
		switch ch {
		case -1:
			tok = EOF
		case '\n':
			// пропускаем все
			for s.ch == '\n' {
				s.next()
			}
			tok = NL
		case '"':
			tok = STRING
			lit = s.scanString('"')
		case '\'':
			tok = SYMBOL
			lit = s.scanSymbol('\'')
		case '@':
			tok = MODIFIER
			lit = s.scanModifier()
		case ':':
			switch s.ch {
			case '=':
				tok = ASSIGN
				s.next()
			case '&':
				tok = BITAND
				s.next()
			case '|':
				tok = BITOR
				s.next()
			case '\\':
				tok = BITXOR
				s.next()
			case '~':
				tok = BITNOT
				s.next()
			default:
				tok = COLON
			}

		case '.':
			tok = DOT
			if s.ch == '.' && s.peek() == '.' {
				s.next()
				s.next() // consume last '.'
				tok = ELLIPSIS
			}
		case ',':
			tok = COMMA
		case ';':
			tok = SEMI
			lit = ";"
		case '(':
			tok = s.checkNext(LPAR, ':', LCONV)
		case ')':
			tok = RPAR
		case '[':
			tok = LBRACK
		case ']':
			tok = RBRACK
		case '{':
			tok = LBRACE
		case '}':
			tok = RBRACE
		case '+':
			tok = s.checkNext(ADD, '+', INC)
		case '-':
			tok = s.checkNext(SUB, '-', DEC)
		case '*':
			tok = MUL
		case '/':
			if s.ch == '/' {
				tok = LINE_COMMENT
				ofs := s.scanLineComment()
				lit = string(s.src[ofs:s.offset])
			} else if s.ch == '*' {
				tok = BLOCK_COMMENT
				ofs := s.scanBlockComment()
				lit = string(s.src[ofs:s.offset])
			} else {
				tok = QUO
			}

		case '%':
			tok = REM
		case '^':
			tok = NOTNIL
		case '<':
			switch s.ch {
			case '=':
				tok = LEQ
				s.next()
			case '<':
				tok = SHL
				s.next()
			default:
				tok = LSS
			}
		case '>':
			switch s.ch {
			case '=':
				tok = GEQ
				s.next()
			case '>':
				tok = SHR
				s.next()
			default:
				tok = GTR
			}
		case '=':
			tok = EQ
		case '#':
			tok = NEQ
		case '~':
			tok = NOT
		case '&':
			tok = AND
		case '|':
			tok = OR
		default:
			s.error(s.offset, "ЛЕК-ОШ-СИМ", fmt.Sprintf("%#U", ch))
			tok = Invalid
			lit = string(ch)
		}
	}

	return
}

func (s *Lexer) WhitespaceBefore(c rune) bool {

	ofs := s.offset - 2
	if ofs < 0 {
		return false
	}
	if s.src[ofs+1] != byte(c) {
		return false
	}

	ch := s.src[ofs]
	return ch == ' ' || ch == '\t'
}
