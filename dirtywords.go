package dirtywords

type TrieTree struct {
	Root *trieNode `json:"root"`
	Skip []rune    `json:"skip"`
}

type trieNode struct {
	Children map[rune]*trieNode `json:"children"`
	Done     bool               `json:"done"`
}

func BuildTree(words [][]rune, skip []rune) (tree *TrieTree) {
	if skip == nil {
		skip = []rune{}
	}
	tree = &TrieTree{
		Root: &trieNode{
			Children: make(map[rune]*trieNode),
		},
		Skip: skip,
	}

	for _, word := range words {
		tree.Root.insertWord(word, tree.Skip)
	}
	return
}

func (tree *TrieTree) Replace(targetStr string, mask rune) string {
	target := []rune(targetStr)
	cur := 0
	length := len(target)
	for cur < length {

		node := tree.Root
		length2 := 0
		for j := cur; j < length; j++ {
			tmp, found := node.Children[target[j]]
			if found {
				node = tmp
				if node.Done {
					length2 = j - cur + 1
				}
			} else {
				break
			}
		}

		if length2 > 0 {
			for i := 0; i < length2; i++ {
				target[cur+i] = mask
			}
			cur = cur + length2
		} else {
			cur++
		}
	}
	return string(target)
}

func (tree *TrieTree) Check(targetStr string) bool {
	target := []rune(targetStr)
	cur := 0
	length := len(target)
	for cur < length {
		node := tree.Root
		for j := cur; j < length; j++ {
			tmp, found := node.Children[target[j]]
			if found {
				node = tmp
				if node.Done {
					return true
				}
			} else {
				break
			}
		}
		cur++
	}
	return false
}

func (node *trieNode) insertWord(word []rune, skip []rune) {
	cur := node
	for _, char := range word {
		cur = cur.findOrCreate(char, skip)
	}
	cur.Done = true
}

func (node *trieNode) find(char rune) (child *trieNode) {
	child, _ = node.Children[char]
	return
}

func (node *trieNode) findOrCreate(char rune, skips []rune) (child *trieNode) {
	child, found := node.Children[char]
	if !found {
		child = &trieNode{
			Children: make(map[rune]*trieNode, 4),
		}
		node.Children[char] = child
		for _, skip := range skips {
			child.Children[skip] = child
		}
	}
	return
}
