package splaytree

type SplayNode struct {
	left   *SplayNode
	right  *SplayNode
	parent *SplayNode
	Item
}

func newSplayNode(item Item) *SplayNode {
	return &SplayNode{Item: item}
}

func (n *SplayNode) replaceItem(item Item) Item {
	curItem := n.Item
	n.Item = item
	return curItem
}

type Item interface {
	Less(than Item) bool
}

type SplayTree struct {
	root *SplayNode
	size int
}

func NewSplayTree() *SplayTree {
	return &SplayTree{}
}

func (t *SplayTree) SetRoot(r *SplayNode) {
	t.root = r
}

func (t *SplayTree) GetRoot() *SplayNode {
	return t.root
}

func (t *SplayTree) Size() int {
	return t.size
}

func (t *SplayTree) Splay(n *SplayNode) {
	for n.parent != nil {
		parent := n.parent
		grandParent := parent.parent
		if grandParent == nil {
			if n == parent.left {
				t.makeLeftChildParent(n, parent)
			} else {
				t.makeRightChildParent(n, parent)
			}
		} else {
			if n == parent.right {
				if parent == grandParent.left {
					t.makeLeftChildParent(parent, grandParent)
					t.makeLeftChildParent(n, parent)
				} else {
					t.makeLeftChildParent(n, parent) // After this n.parent would be updated
					t.makeRightChildParent(n, n.parent)
				}
			} else {
				if parent == grandParent.left {
					t.makeRightChildParent(n, parent) // After this n.parent would be updated
					t.makeLeftChildParent(n, n.parent)
				} else {
					t.makeRightChildParent(parent, grandParent)
					t.makeRightChildParent(n, parent)
				}
			}
		}
	}
	t.SetRoot(n)
}

func (t *SplayTree) makeLeftChildParent(x, y *SplayNode) {
	if x == nil || y == nil || x.parent != y || y.left != x {
		panic("Anomaly in Tree structure")
	}

	if y.parent != nil {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent
	y.parent = x
	y.left = x.right
	x.right = y
}

func (t *SplayTree) makeRightChildParent(x, y *SplayNode) {
	if x == nil || y == nil || x.parent != y || y.right != x {
		panic("Anomaly in Tree structure")
	}

	if y.parent != nil {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}

	if x.left != nil {
		x.left.parent = y
	}

	x.parent = y.parent
	y.parent = x
	y.right = x.left
	x.left = y
}

func (t *SplayTree) Insert(key Item, replace bool) Item {
	if key == nil {
		panic("Inserting nil item into SplayTree")
	}
	cur := t.GetRoot()

	if cur == nil {
		t.size = 1
		t.SetRoot(newSplayNode(key))
		return nil
	}

	newNode := newSplayNode(key)
	for cur != nil {
		switch {
		case cur.Less(key):
			if cur.right == nil {
				t.size++
				cur.right = newNode
				cur.right.parent = cur
			}
			cur = cur.right
		case key.Less(cur):
			if cur.left == nil {
				t.size++
				cur.left = newNode
				cur.left.parent = cur
			}
			cur = cur.left
		default:
			// If keys are equal replace based on argument
			if replace {
				return cur.replaceItem(key)
			}
		}
	}
	t.Splay(newNode)
	return nil
}

// TODO Incomplete
func (t *SplayTree) Find(key Item) Item {
	cur := t.GetRoot()
	for cur != nil {
		switch {
		case cur.Less(key):
			cur = cur.right
		case key.Less(cur):
			cur = cur.left
		default:
			return cur.Item
		}
	}
	return nil
}

// TODO
func (t *SplayTree) Remove(key Item) {

}

// TODO
func (t *SplayTree) Iterator() {

}
