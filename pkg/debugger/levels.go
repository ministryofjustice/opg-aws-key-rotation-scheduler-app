package debugger

var LEVEL int = DETAILED

const (
	ERR      int = 0
	INFO     int = 10
	DETAILED int = 20
	VERBOSE  int = 30
	ALL      int = 900
)

func SetLevel(lvl int) {
	LEVEL = lvl
}
