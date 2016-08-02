package iterator

type Iterable interface {
	Iterator() chan Item
}

//Could be any specific type
type Item interface{}

type Container struct {
	//could be anything, here, just a slice
	data []Item
	//any other fields...
}

func (c *Container) Size() int {
	return len(c.data)
}

func (c *Container) Iterator() chan Item {
	channel := make(chan Item);
	go func () {
		for i := 0; i < c.Size(); i++ {
			channel <- c.data[i]
		}
	} ();
	return channel
}
//Missing more methods

func (c *Container) Add(i Item) {
	if c.data == nil {
		c.data = make([]Item, 0)
	}
	c.data = append(c.data, i)
}


