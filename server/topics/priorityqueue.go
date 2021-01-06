package topics

import (
	"container/heap"
)

type PriorityQueue struct {
	itemHeap *itemHeap
	lookup   map[int]*item
}

func New() PriorityQueue {
	return PriorityQueue{
		itemHeap: &itemHeap{},
		lookup:   make(map[int]*item),
	}
}

func (p *PriorityQueue) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue) Insert(topic *Topic, priority int) {
	_, ok := p.lookup[topic.ID]
	if ok {
		return
	}

	newItem := &item{
		topic:    topic,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
	p.lookup[topic.ID] = newItem
}

func (p *PriorityQueue) Pop() *Topic {
	if len(*p.itemHeap) == 0 {
		return nil
	}

	item := heap.Pop(p.itemHeap).(*item)
	delete(p.lookup, item.topic.ID)
	return item.topic
}

func (p *PriorityQueue) UpdatePriority(topicID int, newPriority int) {
	item, ok := p.lookup[topicID]
	if !ok {
		return
	}

	item.priority = newPriority
	heap.Fix(p.itemHeap, item.index)
}

type itemHeap []*item

type item struct {
	topic    *Topic
	priority int
	index    int
}

func (ih *itemHeap) Len() int {
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].priority > (*ih)[j].priority
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap) Push(x interface{}) {
	it := x.(*item)
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap) Pop() interface{} {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
