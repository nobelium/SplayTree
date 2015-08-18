package splaytree

type SplayNode struct {
	left   *SplayNode
	right  *SplayNode
	parent *SplayNode
	Item
}

func newSplayNode(item Item) *SplayNode {
	return &SplayNode{Item: item, left: nil, right: nil, parent: nil}
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
	if n == nil {
		return
	}
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
			if n == parent.left {
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
		panic("Anomaly in Tree structure makeLeftChildParent")
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
		panic("Anomaly in Tree structure makeRightChildParent")
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
		case cur.Item.Less(key):
			if cur.right == nil {
				t.size++
				cur.right = newNode
				cur.right.parent = cur
			}
			cur = cur.right
		case key.Less(cur.Item):
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

func (t *SplayTree) Find(key Item) Item {
	node := t.findNode(key)
	if node != nil {
		return node.Item
	}
	return nil
}

func (t *SplayTree) findNode(key Item) *SplayNode {
	cur := t.GetRoot()
	for cur != nil {
		switch {
		case cur.Item.Less(key):
			cur = cur.right
		case key.Less(cur.Item):
			cur = cur.left
		default:
			t.Splay(cur)
			return cur
		}
	}
	return nil
}

func (t *SplayTree) Remove(key Item) Item {
	node := t.findNode(key)
	if node != nil {
		t.remove(node)
		return node.Item
	}
	return nil
}

func (t *SplayTree) remove(n *SplayNode) {
	if n == nil {
		return
	}

	t.Splay(n)
	if n.left != nil && n.right != nil {
		// find inorder predecessor
		prev := n.left
		for prev.right != nil {
			prev = prev.right
		}
		prev.right = n.right
		n.right.parent = prev
		n.left.parent = nil
		t.SetRoot(n.left)
	} else if n.right != nil {
		n.right.parent = nil
		t.SetRoot(n.right)
	} else if n.left != nil {
		n.left.parent = nil
		t.SetRoot(n.left)
	} else {
		t.SetRoot(nil)
	}
	n.parent = nil
	n.right = nil
	n.left = nil
	n = nil
	t.size--
}

// TODO
func (t *SplayTree) Iterator() {

}
