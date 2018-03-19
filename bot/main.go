package bot

import "time"

// shitpost every day
const interval = 24 * time.Hour

func Run() {
	for {
		// TODO: compose tweet
		
		time.Sleep(interval)
	}
}