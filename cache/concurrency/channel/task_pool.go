package channel

import "context"

type Task func()

type TaskPool struct {
	tasks chan Task
}

func NewTaskPool(numG int, cap int) *TaskPool {

	pool := &TaskPool{
		tasks: make(chan Task, cap),
	}

	for i := 0; i < numG; i++ {
		go func() {
			for t := range pool.tasks {
				t()
			}
		}()
	}

	return pool
}
func (p *TaskPool) Do(t Task) {

}

func (p *TaskPool) Submit(ctx context.Context, t Task) error {

	select {
	case <-ctx.Done():
		return ctx.Err()
	case p.tasks <- t:
	}

	return nil
}

func (p *TaskPool) Close() {
	close(p.tasks)
}
