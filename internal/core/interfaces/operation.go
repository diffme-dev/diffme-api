package interfaces

const (
	OperationAdd     = "add"
	OperationReplace = "replace"
	OperationRemove  = "remove"
	OperationMove    = "move"
	OperationCopy    = "copy"
	OperationTest    = "test"
)

type StringPointer string

type DiffOperation struct {
	Type     string      `json:"op"`
	From     StringPointer     `json:"from,omitempty"`
	Path     StringPointer     `json:"path"`
	OldValue interface{} `json:"old_value,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}