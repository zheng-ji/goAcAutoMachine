package goAcAutoMachine

type AcNode struct {
	fail      *AcNode
	isPattern bool
	next      map[rune]*AcNode
}

func newAcNode() *AcNode {
	return &AcNode{
		fail:      nil,
		isPattern: false,
		next:      map[rune]*AcNode{},
	}
}

type AcAutoMachine struct {
	root *AcNode
}

type Result struct {
	Key   string
	Start int
	End   int
}

func NewAcAutoMachine() *AcAutoMachine {
	return &AcAutoMachine{
		root: newAcNode(),
	}
}

func (ac *AcAutoMachine) AddPattern(pattern string) {
	chars := []rune(pattern)
	iter := ac.root
	for _, c := range chars {
		if _, ok := iter.next[c]; !ok {
			iter.next[c] = newAcNode()
		}
		iter = iter.next[c]
	}
	iter.isPattern = true
}

func (ac *AcAutoMachine) Build() {
	queue := []*AcNode{}
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		parent := queue[0]
		queue = queue[1:]

		for char, child := range parent.next {
			if parent == ac.root {
				child.fail = ac.root
			} else {
				failAcNode := parent.fail
				for failAcNode != nil {
					if _, ok := failAcNode.next[char]; ok {
						child.fail = failAcNode.next[char]
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil {
					child.fail = ac.root
				}
			}
			queue = append(queue, child)
		}
	}
}

func (ac *AcAutoMachine) Query(content string) (results []Result) {
	chars := []rune(content)
	iter := ac.root
	var start, end int
	for i, c := range chars {
		_, ok := iter.next[c]
		for !ok && iter != ac.root {
			iter = iter.fail
		}
		if _, ok = iter.next[c]; ok {
			if iter == ac.root { // this is the first match, record the start position
				start = i
			}
			iter = iter.next[c]
			if iter.isPattern {
				end = i // this is the end match, record one result
				result := Result{
					Key:   string([]rune(content)[start : end+1]),
					Start: start,
					End:   end + 1,
				}
				// results = append(results, string([]rune(content)[start:end+1]))
				results = append(results, result)
			}
		}
	}
	return
}

// 仅匹配最长的关键字
func (ac *AcAutoMachine) QueryLast(content string) (results []Result) {
	chars := []rune(content)
	iter := ac.root
	var start, lastStart, end int
	var lastResult Result // cant remove!!!
	var result Result
	for i, c := range chars {
		_, ok := iter.next[c]
		for !ok && iter != ac.root {
			iter = iter.fail
		}
		if _, ok = iter.next[c]; ok {
			if iter == ac.root { // this is the first match, record the start position
				start = i
				result = Result{}
			}
			iter = iter.next[c]
			if iter.isPattern {
				end = i // this is the end match, record one result
				result = Result{
					Key:   string([]rune(content)[start : end+1]),
					Start: start,
					End:   end + 1,
				}
				lastResult = result
			}
		}

		// 结束或者lastStart变了
		if i == len(chars)-1 || lastStart != start {
			if (lastResult != Result{}) {
				results = append(results, lastResult)
			}
		}

		lastStart = start
	}
	return
}
