package usersystem

import (
	"sync"
)

type DataType string

type Dataset interface {
	InitType(datatype DataType)
	Set(datatype DataType, id string, v interface{})
	Get(datatype DataType, id string) (interface{}, bool)
	Delete(datatype DataType, id string)
	Flush(id string)
}

type PlainDataset struct {
	Dataset map[DataType]*sync.Map
}

func (d *PlainDataset) InitType(datatype DataType) {
	d.Dataset[datatype] = &sync.Map{}
}
func (d *PlainDataset) Set(datatype DataType, id string, v interface{}) {
	d.Dataset[datatype].Store(id, v)
}
func (d *PlainDataset) Get(datatype DataType, id string) (interface{}, bool) {
	return d.Dataset[datatype].Load(id)
}
func (d *PlainDataset) Delete(datatype DataType, id string) {
	d.Dataset[datatype].Delete(id)
}
func (d *PlainDataset) Flush(id string) {
	for k := range d.Dataset {
		d.Dataset[k].Delete(id)
	}
}

func NewPlainDataset() *PlainDataset {
	return &PlainDataset{
		Dataset: map[DataType]*sync.Map{},
	}
}
