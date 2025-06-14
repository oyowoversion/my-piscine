package piscine

type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		return newNode
	}

	current := root
	for {
		if data < current.Data {
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				break
			}
			current = current.Right
		}
	}
	return root
}
