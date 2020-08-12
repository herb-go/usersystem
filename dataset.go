package usersystem

import (
	"sync"
)

type Datatype string

type Dataset interface {
	InitType(datatype Datatype)
	Set(datatype Datatype, id string, v interface{})
	Get(datatype Datatype, id string) (interface{}, bool)
	Delete(datatype Datatype, id string)
	Flush(id string)
}

type PlainDataset struct {
	Dataset map[Datatype]*sync.Map
}

func (d *PlainDataset) InitType(datatype Datatype) {
	d.Dataset[datatype] = &sync.Map{}
}
func (d *PlainDataset) Set(datatype Datatype, id string, v interface{}) {
	d.Dataset[datatype].Store(id, v)
}
func (d *PlainDataset) Get(datatype Datatype, id string) (interface{}, bool) {
	return d.Dataset[datatype].Load(id)
}
func (d *PlainDataset) Delete(datatype Datatype, id string) {
	d.Dataset[datatype].Delete(id)
}
func (d *PlainDataset) Flush(id string) {
	for k := range d.Dataset {
		d.Dataset[k].Delete(id)
	}
}

func NewPlainDataset() *PlainDataset {
	return &PlainDataset{
		Dataset: map[Datatype]*sync.Map{},
	}
}
