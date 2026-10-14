package ledger

/*
Severity represents the importance of a recorded message.
*/
type Severity int

/*
Predefined severity levels, ordered from least to most severe.
*/
const (
	// Info represents routine program events.
	Info Severity = iota

	// Warn represents unusual conditions without significant issues.
	Warn

	// Error represents a problem that affects normal operation.
	Error

	// Fatal represents a serious condition that prevents the program
	// from functioning as intended.
	Fatal
)

/*
IsValid reports whether severity is within Ledger's defined severity levels.
*/
func (severity Severity) IsValid() bool {
	return severity >= Info && severity <= Fatal
}

/*
String returns the uppercase name of a severity.
It returns an empty string if the severity is invalid.
*/
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
