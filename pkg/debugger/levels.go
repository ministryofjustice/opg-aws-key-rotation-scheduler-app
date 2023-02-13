package debugger

var LEVEL int = DETAILED

const (
	ERR      int = 0
	INFO     int = 10
	DETAILED int = 20
	VERBOSE  int = 30
	ALL      int = 900
)

var levelToString map[int]string = map[int]string{
	ERR:      "ERR",
	INFO:     "INFO",
	DETAILED: "DETAILED",
	VERBOSE:  "VERBOSE",
	ALL:      "ALL",
}
