package piscine

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return true
	case string, rune:
		return false
	}
	return false
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), comp func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if comp(it) {
			f(it)
		}
		it = it.Next
	}
}
