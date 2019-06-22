package main

const defaultChannelCapacity = 100

type job func(in, out chan interface{})

func (j job) do(in chan interface{}) (out chan interface{}) {
    out = make(chan interface{}, defaultChannelCapacity)
    j(in, out)
    return out
}

func Pipe(funcs ...job) {
    in := make(chan interface{}, defaultChannelCapacity)
    for _, f := range funcs {
        in = f.do(in)
        close(in)
    }
}