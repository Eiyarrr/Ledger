package ledger

// struct containing configuration and state
type Ledger struct{}

// returns an explicitly constructed object
// log := ledger.New()
func New() *Ledger {
	return &Ledger{}
}
