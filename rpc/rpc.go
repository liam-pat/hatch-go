package rpcDemo

import "errors"

type DemoService struct {
}

type Args struct {
	A int
	B int
}

func (DemoService) Div(args Args, result *float64) error {

	if args.B == 0 {
		return errors.New("second params cannot be 0")
	}

	*result = float64(args.A) / float64(args.B)

	return nil
}
