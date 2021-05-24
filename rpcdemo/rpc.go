package rpcdemo

import "errors"

type DemoService struct {}

type Args struct {
	A, B int
}
//go rpc 要求参数必须有两个
func (DemoService) Div (args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
