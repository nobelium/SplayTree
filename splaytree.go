package splaytree

type SplayNode struct {
	left  *SplayNode
	right *SplayNode
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

func (t *SplayTree) Insert(key Item, replace bool) Item {
	if key == nil {
		panic("Inserting nil item into SplayTree")
	}
	cur := t.GetRoot()

	if cur == nil {
		t.size = 1
		t.SetRoot(newSplayNode(key))
	}

	for cur != nil {
		switch {
		case cur.Less(key):
			if cur.right == nil {
				t.size++
				cur.right = newSplayNode(key)
			}
			cur = cur.right
		case key.Less(cur):
			if cur.left == nil {
				t.size++
				cur.left = newSplayNode(key)
			}
			cur = cur.left
		default:
			// If keys are equal replace based on argument
			if replace {
				return cur.replaceItem(key)
			}
		}
	}
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

// TODO
func (t *SplayTree) Splay() {

}
