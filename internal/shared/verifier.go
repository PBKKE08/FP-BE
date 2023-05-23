package shared

type Verifier interface {
	Verify() error
}
