package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	value       rune
	weight      int
	left, right *TreeNode
}

func NewTreeNode(value rune, weight int, left *TreeNode, right *TreeNode) TreeNode {
	return TreeNode{value: value, weight: weight, left: left, right: right}
}

func (n TreeNode) makeCodeForValue(value rune, parentCode string) string {
	if n.value == value {
		return parentCode
	} else {
		if n.left != nil {
			code := n.left.makeCodeForValue(value, parentCode+"0")
			if code != "" {
				return code
			}
		}

		if n.right != nil {
			code := n.right.makeCodeForValue(value, parentCode+"1")
			if code != "" {
				return code
			}
		}
	}

	return ""
}

func mainHuffmanEncode() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputString := sc.Text()
	charsFreq := countCharacters(inputString)

	nodesSlice := makeNodesSlice(charsFreq)

	tree := makeHuffmanTree(nodesSlice)

	codes := makeHuffmanCodes(tree, charsFreq)

	encodedString := huffmanEncode(inputString, codes)

	fmt.Println(len(charsFreq), len(encodedString))
	for char, freq := range codes {
		fmt.Printf("%c: %s\n", char, freq)
	}
	fmt.Println(encodedString)
}

func mainHuffmanDecode() {
	//var lettersCount, encodedStringSize int
	codes := make(map[string]string)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nums := strings.Split(sc.Text(), " ")
	lettersCount, _ := strconv.Atoi(nums[0])
	_, _ = strconv.Atoi(nums[1])

	for i := 0; i < lettersCount; i++ {
		sc.Scan()
		values := strings.Split(sc.Text(), ": ")

		char := values[0]
		freq := values[1]
		codes[freq] = char
	}

	sc.Scan()
	encodedString := sc.Text()
	decodedString := huffmanDecodeString(encodedString, codes)
	fmt.Println(decodedString)
}

func countCharacters(str string) map[rune]int {
	result := make(map[rune]int)
	for _, char := range str {
		result[char]++
	}

	return result
}

func nodesSort(nodesSlice []TreeNode) {

	sort.Slice(nodesSlice, func(i, j int) bool {
		return nodesSlice[i].weight > nodesSlice[j].weight
	})
}

func makeNodesSlice(charsFreq map[rune]int) []TreeNode {
	var nodes []TreeNode
	for char, freq := range charsFreq {
		nodes = append(nodes, NewTreeNode(char, freq, nil, nil))
	}

	return nodes
}

func makeHuffmanTree(nodesSlice []TreeNode) TreeNode {

	for len(nodesSlice) > 1 {
		var leftNode TreeNode
		var rightNode TreeNode
		var parentNode TreeNode

		nodesSort(nodesSlice)
		leftNode, nodesSlice = nodesSlice[len(nodesSlice)-1], nodesSlice[:len(nodesSlice)-1]
		rightNode, nodesSlice = nodesSlice[len(nodesSlice)-1], nodesSlice[:len(nodesSlice)-1]

		parentNode = NewTreeNode(0, leftNode.weight+rightNode.weight, &leftNode, &rightNode)

		nodesSlice = append(nodesSlice, parentNode)
	}

	return nodesSlice[0]
}

func makeHuffmanCodes(tree TreeNode, charsFreq map[rune]int) map[rune]string {
	codes := make(map[rune]string)
	for char, _ := range charsFreq {
		codes[char] = tree.makeCodeForValue(char, "")

		if len(charsFreq) == 1 {
			codes[char] = "0"
		}
	}

	return codes
}

func huffmanEncode(inputString string, codes map[rune]string) string {
	var result []string
	for i := 0; i < len(inputString); i++ {
		result = append(result, codes[rune(inputString[i])])
	}
	return strings.Join(result, "")
}

func huffmanDecode(encoded string, tree TreeNode) string {
	var result []string
	var node = tree

	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 48 {
			node = *node.left
		} else {
			node = *node.right
		}
		if node.value != 0 {
			result = append(result, string(node.value))
			node = tree
		}
	}
	return strings.Join(result, "")
}

func huffmanDecodeString(encoded string, codes map[string]string) string {
	var result []string
	var code string

	for i := 0; i < len(encoded); i++ {
		code += string(encoded[i])
		if char, ok := codes[code]; ok {
			result = append(result, char)
			code = ""
		} else {
			code += codes[code]
		}
	}
	return strings.Join(result, "")
}
