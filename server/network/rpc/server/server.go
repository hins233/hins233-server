package server

import (
	"errors"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Greeter struct {
}

func (t *Greeter) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Greeter) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

type Greet string
type Answer string

func (t *Greeter) SayHello(greet *Greet, answer *Answer) error {
	fmt.Println("server: hello my dear son", *greet)
	*answer = "deal by server"
	return nil
}
