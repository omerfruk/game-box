package context

type ResponseModel struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	HataVarMi bool        `json:"hataVarMi"`
}
