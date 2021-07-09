package queue

type Queue struct {
	First  *Node
	Last   *Node
	Length int
}

func (q *Queue) Push(node *Node) int {
	if q.First == nil {
		q.First = node
	}

	if q.Last != nil {
		q.Last.Next = node
	}

	q.Last = node
	q.Length++

	return q.Length
}

func (q *Queue) Pop() *Node {
	if q.Length == 0 {
		return nil
	}

	first := q.First
	q.First = first.Next

	if q.Length == 1 {
		q.Last = nil
	}

	q.Length--

	return first
}
