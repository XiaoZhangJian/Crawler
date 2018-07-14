package scheduler

import "github.com/XiaoZhangJian/Crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) WokerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var wokerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(wokerQ) > 0 {
				activeWorker = wokerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)

			case w := <-q.workerChan:
				wokerQ = append(wokerQ, w)
			case activeWorker <- activeRequest:
				wokerQ = wokerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
