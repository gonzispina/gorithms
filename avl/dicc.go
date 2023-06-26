package avl

type Dicc struct {
	avl *Node
}

func NewDicc() *Dicc {
	return &Dicc{avl: nil}
}

func (d *Dicc) Get(key string) interface{} {
	if d.avl == nil {
		return nil
	}

	return search(d.avl, key)
}

func (d *Dicc) Define(key string, value interface{}) {
	if d.avl == nil {
		d.avl = newNode(key, value)
		return
	}

	insert(d.avl, key, value)
}
