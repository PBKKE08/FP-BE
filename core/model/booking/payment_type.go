package booking

type PaymentType string

const (
	BCA     PaymentType = "BCA"
	Mandiri PaymentType = "Mandiri"
)

func IsPaymentTypeValid(t string) bool {
	return t == "BCA" || t == "Mandiri"
}

func NewPaymentType(mode string) PaymentType {
	return PaymentType(mode)
}
