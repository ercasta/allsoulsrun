package engine

func SaveComponent(history *[]ComponentHistory, c chan ComponentHistory, doneChan chan bool) {
	for i := range c {
		*history = append(*history, i)
	}
	doneChan <- true
}
