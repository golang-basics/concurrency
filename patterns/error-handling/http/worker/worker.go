package worker

type Worker struct {
}

func Init() (*Worker, error) {
	worker := &Worker{}
	return worker, nil
}

func (w Worker) Start() error {
	//maxRetries := 5
	//logger := logging.Logger()
	//logger.Info("starting worker")
	//retries := 0
	//for retries <= maxRetries {
	//	retries++
	//}
	return nil
}

func (w Worker) Stop() error {
	return nil
}
