package entity

type Postback struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e *Postback) GetId() int {
	return e.ID
}

func (e *Postback) GetName() string {
	return e.Name
}

func (e *Postback) GetValue() string {
	return e.Value
}
