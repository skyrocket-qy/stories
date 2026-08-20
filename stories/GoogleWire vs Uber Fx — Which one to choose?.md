TL;DR: “Wire”

Wire and Fx are also popular DI tool in Go, but they implement with different method.

Wire: Compile-time DI tool

Fx: Runtime DI framework

Learn about Medium’s values
The main difference is: What time did they do the computation?

Dependency Graph
Life cycle management
Error handling
Reflection
While Wire do anything on compile-time(also without reflection), it has no extra overhead on the runtime, also get the shorter server restart time.

Here is the simple benchmark:

type Service interface {
 Do()
}

type MyService struct{}

func (s *MyService) Do() {}

func NewService() Service {
 return &MyService{}
}
Fx: 152545 ns/op

Wire: 0.3146 ns/op

Wire is roughly 484,885 times faster!!!

Code[!https://github.com/skyrocket-qy/fx-vs-wire]

