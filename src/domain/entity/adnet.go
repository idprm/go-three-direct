package entity

type Adnet struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e *Adnet) GetId() int {
	return e.ID
}

func (e *Adnet) GetName() string {
	return e.Name
}

func (e *Adnet) GetValue() string {
	return e.Value
}
