package channel

/**
 * 1. Funnel : multiple input -> single output
 * 2. Fan out : single input -> multiple output
 * 3. Turn out : multiple input -> multiple output
 */

func FanOut[T any](in <-chan T, firstOut, secondOut chan<- T) {
	// receive data from single channel(in)
	for data := range in {
		// send data to each channel(out) which is available
		select {
		case firstOut <- data:
		case secondOut <- data:
		}
	}
}

func TurnOut[T any](quit <-chan struct{}, firstIn, secondIn chan T, firstOut, secondOut chan<- T) {
	var data T

	for {
		select {
		case data = <-firstIn:
		case data = <-secondIn:
			// quit channel only has a single purpose of closing this channel when the sending operation finishes from the data sender(go-routine)..
		case <-quit:
			/**
			 * closing receive channel is actually a anti pattern and compiler won't let you close receive channel
			 * however, in this scenario it's fine because sender will close 'quit' channel.
			 */
			close(firstIn)
			close(secondIn)

			FanOut(firstIn, firstIn, secondIn)
			FanOut(secondIn, firstIn, secondIn)
			return // escape from the for loop
		}

		select {
		case firstOut <- data:
		case secondOut <- data:
		}
	}
}
