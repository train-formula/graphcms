package calls

type Call interface {
	Validation()
}

type DemoCall struct {
	Request  struct{}
	Response struct{}
}
