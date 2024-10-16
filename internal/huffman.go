package internal

import (
	"container/heap"
)

type HuffmanNode struct {
	Character string
	Frequency uint64
	Left      *HuffmanNode
	Right     *HuffmanNode
}

type HuffmanHeap []*HuffmanNode

func (h HuffmanHeap) Len() int {
	return len(h)
}

func (h HuffmanHeap) Less(i, j int) bool {
	return h[i].Frequency < h[j].Frequency
}

func (h HuffmanHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanNode))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	x := old[n-1]

	*h = old[0 : n-1]
	return x
}

func BuildHuffmanTree(frequencies map[string]uint64) *HuffmanNode {
	h := &HuffmanHeap{}
	heap.Init(h)

	for char, freq := range frequencies {
		heap.Push(h, &HuffmanNode{Character: char, Frequency: freq})
	}

	for h.Len() > 1 {
		left := heap.Pop(h).(*HuffmanNode)
		right := heap.Pop(h).(*HuffmanNode)
		freqSum := left.Frequency + right.Frequency

		heap.Push(h, &HuffmanNode{
			Frequency: freqSum,
			Left:      left,
			Right:     right,
		})
	}

	return heap.Pop(h).(*HuffmanNode)
}

func GenerateHuffmanCodes(node *HuffmanNode, code string, codes map[string]string) {
	if node == nil {
		return
	}

	if node.Right == nil && node.Left == nil {
		codes[node.Character] = code
	}

	GenerateHuffmanCodes(node.Right, code+"1", codes)
	GenerateHuffmanCodes(node.Left, code+"0", codes)
}
