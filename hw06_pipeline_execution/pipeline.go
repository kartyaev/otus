package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}
	out := wrapWithDone(in, done)
	for _, s := range stages {
		out = s(wrapWithDone(out, done))
	}
	return out
}

func wrapWithDone(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case c, ok := <-in:
				if !ok {
					return
				}
				out <- c
			}
		}
	}()

	return out
}
