package common

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item[V any] struct {
	Value    V   // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[V any] []*Item[V]

func (pq PriorityQueue[V]) Len() int { return len(pq) }

func (pq PriorityQueue[V]) Less(i, j int) bool {
	// Lowest priority score actually means highest priority
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[V]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[V]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[V])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[V]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[V]) Update(item *Item[V], value V, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
