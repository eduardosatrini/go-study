package queue

type Queue struct {
	values []string
}

func (q *Queue) Insert(name string) {

	q.values = append(q.values, name)

}

func (q *Queue) Remove() {

	q.values = q.values[1:]

}
