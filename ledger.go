package ledger

// define ledger severity levels
type Severity int

const (
	INFO  Severity = iota //0
	WARN                  //1
	ERROR                 //2
	FATAL                 //3
)

// struct containing configuration and state
type Ledger struct{}

// returns an explicitly constructed object
// log := ledger.New()
func New() *Ledger {
	return &Ledger{}
}
