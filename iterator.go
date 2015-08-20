package splaytree

type Iterator func(i Item) bool

// Ascend will call iterator once for each element greater or equal than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *SplayTree) Inorder(pivot Item, iterator Iterator) {
	t.inorder(t.GetRoot(), pivot, iterator)
}

func (t *SplayTree) inorder(x *SplayNode, pivot Item, iterator Iterator) bool {
	if x == nil {
		return true
	}

	if x.Item.Less(pivot) {
		if !t.inorder(x.left, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.inorder(x.right, pivot, iterator)
}
