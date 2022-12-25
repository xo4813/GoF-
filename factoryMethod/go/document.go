package factoryMethod

import "fmt"

type iDocument interface {
	setName(name string)
	getName() string
	approve() bool
}

type document struct {
	name    string
	docType int
	data    string
}

func (d *document) setName(name string) {
	d.name = name
}

func (d *document) getName() string {
	return d.name
}

func (d *document) approve() bool {
	fmt.Printf("Approving %s", d.name)

	return true
}
