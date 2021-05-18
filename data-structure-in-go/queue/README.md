# Queue  
These are **FIFO** data structures _(first-in first-out)_, 
where the first element to be inserted will be the 
first to be removed, items are **added at the end** 
and **removed the beginning**.

```go

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

```


 
