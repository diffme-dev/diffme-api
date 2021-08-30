package interfaces

type Operation string

const (
	OperationAdd     Operation = "add"
	OperationReplace Operation = "replace"
	OperationRemove  Operation = "remove"
	OperationMove    Operation = "move"
	OperationCopy    Operation = "copy"
	OperationTest    Operation = "test"
)

type StringPointer string

type DiffOperation struct {
	Type     string        `json:"op"`
	From     StringPointer `json:"from,omitempty"`
	Path     StringPointer `json:"path"`
	OldValue interface{}   `json:"old_value,omitempty"`
	Value    interface{}   `json:"value,omitempty"`
}
