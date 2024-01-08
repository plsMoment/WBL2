package dev07

// or merge done-channels and returns result channel which send signal and close
// after first receiving from any channel
func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	anyFinished := make(chan interface{})

	for i := range channels {
		go func(ch <-chan interface{}) {
			select {
			case sig := <-ch:
				close(anyFinished)
				out <- sig
				close(out)
			case <-anyFinished:
				return
			}
		}(channels[i])
	}

	return out
}
