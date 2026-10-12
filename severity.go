package ledger

// define ledger severity levels
type Severity int

const (
	Info  Severity = iota //0
	Warn                  //1
	Error                 //2
	Fatal                 //3
)

// checks if severity is within the allowed bounds
func (severity Severity) IsValid() bool {
	return severity >= Info && severity <= Fatal
}

// string conversion of severity
func (severity Severity) String() string {
	switch severity {
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}

	return ""
}
