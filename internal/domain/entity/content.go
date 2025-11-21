package entity

type Content struct {
	ID         int `json:"id"`
	ServiceID  int `json:"service_id"`
	Service    Service
	Name       string `json:"name"`
	OriginAddr string `json:"origin_addr"`
	Value      string `json:"value"`
}

func (e *Content) GetId() int {
	return e.ID
}

func (e *Content) GetName() string {
	return e.Name
}

func (e *Content) GetOriginAddr() string {
	return e.OriginAddr
}

func (e *Content) GetValue() string {
	return e.Value
}
