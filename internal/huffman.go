package internal

import (
	"bufio"
	"container/heap"
	"fmt"
	"unicode/utf8"
)

type HuffmanNode struct {
	Character rune
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

func BuildHuffmanTree(frequencies map[rune]uint64) *HuffmanNode {
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

func GenerateHuffmanCodes(node *HuffmanNode, code string, codes map[rune]string) {
	if node == nil {
		return
	}

	if node.Right == nil && node.Left == nil {
		codes[node.Character] = code
	}

	GenerateHuffmanCodes(node.Left, code+"0", codes)
	GenerateHuffmanCodes(node.Right, code+"1", codes)
}

func SerializeHuffmanTree(node *HuffmanNode, builder *[]byte) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*builder = append(*builder, '0')
		runeBytes := make([]byte, utf8.RuneLen(rune(node.Character)))
		utf8.EncodeRune(runeBytes, rune(node.Character))

		// Append the encoded rune bytes
		*builder = append(*builder, runeBytes...)
	} else {
		*builder = append(*builder, '1')
	}

	SerializeHuffmanTree(node.Left, builder)
	SerializeHuffmanTree(node.Right, builder)
}

func DeserializeHuffmanTree(reader *bufio.Reader) (*HuffmanNode, error) {
	char, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	if char == '0' {
		leafChar, _, err := reader.ReadRune()
		if err != nil {
			return nil, err
		}
		return &HuffmanNode{Character: leafChar}, nil
	}

	if char == '1' {
		left, err := DeserializeHuffmanTree(reader)
		if err != nil {
			return nil, err
		}

		right, err := DeserializeHuffmanTree(reader)
		if err != nil {
			return nil, err
		}

		return &HuffmanNode{Left: left, Right: right}, nil
	}

	return nil, fmt.Errorf("invalid character encountered during deserialization: %c", char)
}
