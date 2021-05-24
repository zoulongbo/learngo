package queue

type List []interface{}			//interface{} 任意类型

func (q *List) Push(v ...interface{})  {
	*q = append(*q, v ...)
}

func (q *List) Pop() interface{}  {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *List) IsEmpty() bool  {
	return len(*q) == 0
}


