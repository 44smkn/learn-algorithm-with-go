package rangesumbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	queue := make(chan int, 10)

	go func() {
		search(root, L, R, queue)
		close(queue)
	}()

	for n := range queue {
		sum += n
	}
	return sum
}

func search(node *TreeNode, L, R int, ch chan<- int) {
	if node == nil {
		return
	}
	v := node.Val
	switch {
	case v < L:
		search(node.Right, L, R, ch)
	case v > R:
		search(node.Left, L, R, ch)
	default:
		ch <- v
		search(node.Left, L, R, ch)
		search(node.Right, L, R, ch)
	}
}
