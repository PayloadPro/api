package entities

import "github.com/PayloadPro/api/responses"

type Bin struct {
}

type Bins []Bin

func (bins Bins) Data() []responses.Data {
	data := make([]responses.Data, 0)
	data = append(data, responses.Data{
		Type: "bin",
	})

	return data
}

func (bin Bin) Data() responses.Data {
	data := responses.Data{
		Type: "bin",
	}

	return data
}
