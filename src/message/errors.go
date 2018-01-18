package message

type ProviderError struct {
	error string
	Type int
}

const (
	ErrCritcal = iota
	ErrOverQuota
)

func (p ProviderError) Error() string {
	return p.error
}

func NewProviderError(s string) *ProviderError {
	return &ProviderError{
		error: s,
	}
}
