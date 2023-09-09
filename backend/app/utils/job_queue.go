// https://www.opsdash.com/blog/job-queues-in-go.html
// https://myungsworld.tistory.com/5

package utils

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type (
	// Queue
	//
	// 작업 큐에 사용되는 여러 채널 포함
	Queue struct {
		jobChannel chan Job // 들어온 요청을 임시로 저장할 작업 큐 버퍼
	}

	// Job
	//
	// 사용자가 북마크 저장할 때 스크래핑 큐에도 저장
	Job struct {
		UserID uint   // 북마크를 저장한 유저 아이디
		Link   string `validate:"required,http_url"` // 스크래핑 할 웹사이트 링크
	}
)

var (
	asyncQueue *Queue
	WG         sync.WaitGroup
)

func InitQueue() {
	asyncQueue = &Queue{
		jobChannel: make(chan Job, 1000),
	}
	log.Debugw("[func Init]", "AsyncQueue", asyncQueue, "len", len(asyncQueue.jobChannel), "capacity", cap(asyncQueue.jobChannel))
}

func StopQueue() {
	close(asyncQueue.jobChannel)
	log.Info("[func StopQueue] Queue stopped")
}

func process(job Job) {
	log.Debugw("[func process] Start", "job", job)
	time.Sleep(time.Second * 30)
	log.Debugw("[func process] End", "job", job)
}

func RunWorker() {
	defer WG.Done()

	for job := range asyncQueue.jobChannel {
		process(job)
	}
}

func Enqueue(job *Job) {
	asyncQueue.jobChannel <- *job
	log.Debugw("[func Enqueue]", "job", *job)
}
