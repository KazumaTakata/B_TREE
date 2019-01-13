package Graph

type Queue struct {
	queue []ID
}

func (q *Queue) Push(id ID) {
	q.queue = append(q.queue, id)
}

func (q *Queue) Pop() ID {
	var x ID
	x, q.queue = q.queue[0], q.queue[1:]
	return x
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
