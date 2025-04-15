package lexer

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (after reading char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

/*************  ✨ Windsurf Command ⭐  *************/
// readChar reads the next character from the input and updates the position and
// readPosition pointers.
/*******  3420a72d-c324-47c2-a816-60863b9c1737  *******/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
