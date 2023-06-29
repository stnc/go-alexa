package alexaapi

//
//
// Interface: Dialog

const (
	DirectiveTypeDialogDelegate              DirectiveType = "Dialog.Delegate"
	DirectiveTypeDialogElicitSlot            DirectiveType = "Dialog.ElicitSlot"
	DirectiveTypeDialogConfirmSlot           DirectiveType = "Dialog.ConfirmSlot"
	DirectiveTypeDialogConfirmIntent         DirectiveType = "Dialog.ConfirmIntent"
	DirectiveTypeDialogUpdateDynamicEntities DirectiveType = "Dialog.UpdateDynamicEntities"
)

//
//
// Directive: Dialog.Delegate

type DirectiveDialogDelegate struct {
	Type          DirectiveType `json:"type"`
	UpdatedIntent *Intent       `json:"updatedIntent,omitempty"`
}

func CreateDirectiveDialogDelegate(updatedIntent *Intent) *DirectiveDialogDelegate {
	return &DirectiveDialogDelegate{
		Type:          DirectiveTypeDialogDelegate,
		UpdatedIntent: updatedIntent,
	}
}

//
//
// Directive: Dialog.ElicitSlot

type DirectiveDialogElicitSlot struct {
	Type          DirectiveType `json:"type"`
	UpdatedIntent *Intent       `json:"updatedIntent,omitempty"`
	SlotToElicit  string        `json:"slotToElicit,omitempty"`
}

func CreateDirectiveDialogElicitSlot(updatedIntent *Intent, slotToElicit string) *DirectiveDialogElicitSlot {
	return &DirectiveDialogElicitSlot{
		Type:          DirectiveTypeDialogElicitSlot,
		UpdatedIntent: updatedIntent,
		SlotToElicit:  slotToElicit,
	}
}

//
//
// Directive: Dialog.ConfirmSlot

type DirectiveDialogConfirmSlot struct {
	Type          DirectiveType `json:"type"`
	SlotToConfirm string        `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent       `json:"updatedIntent,omitempty"`
}

func CreateDirectiveDialogConfirmSlot(updatedIntent *Intent, slotToConfirm string) *DirectiveDialogConfirmSlot {
	return &DirectiveDialogConfirmSlot{
		Type:          DirectiveTypeDialogConfirmSlot,
		UpdatedIntent: updatedIntent,
		SlotToConfirm: slotToConfirm,
	}
}

//
//
// Directive: Dialog.ConfirmIntent

type DirectiveDialogConfirmIntent struct {
	Type          DirectiveType `json:"type"`
	UpdatedIntent *Intent       `json:"updatedIntent,omitempty"`
}

func CreateDirectiveDialogConfirmIntent(updatedIntent *Intent) *DirectiveDialogConfirmIntent {
	return &DirectiveDialogConfirmIntent{
		Type:          DirectiveTypeDialogConfirmIntent,
		UpdatedIntent: updatedIntent,
	}
}

//
//
// Directive: Dialog.UpdateDynamicEntities

type DirectiveDialogUpdateDynamicEntities struct {
	Type           DirectiveType                               `json:"type"`
	UpdateBehavior UpdateDynamicEntitiesUpdateBehavior         `json:"updateBehavior,omitempty"`
	Types          []*DirectiveDialogUpdateDynamicEntitiesType `json:"types,omitempty"`
}

type DirectiveDialogUpdateDynamicEntitiesType struct {
	Name   string                                           `json:"name"`
	Values []*DirectiveDialogUpdateDynamicEntitiesTypeValue `json:"values"`
}

type DirectiveDialogUpdateDynamicEntitiesTypeValue struct {
	Id   string                                            `json:"id"`
	Name DirectiveDialogUpdateDynamicEntitiesTypeValueName `json:"name"`
}

type DirectiveDialogUpdateDynamicEntitiesTypeValueName struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

func CreateDirectiveDialogUpdateDynamicEntities(updateBehavior UpdateDynamicEntitiesUpdateBehavior, types []*DirectiveDialogUpdateDynamicEntitiesType) *DirectiveDialogUpdateDynamicEntities {
	return &DirectiveDialogUpdateDynamicEntities{
		Type:           DirectiveTypeDialogUpdateDynamicEntities,
		UpdateBehavior: updateBehavior,
		Types:          types,
	}
}

type UpdateDynamicEntitiesUpdateBehavior string

const (
	UpdateDynamicEntitiesUpdateBehaviorUnspecified UpdateDynamicEntitiesUpdateBehavior = ""
	UpdateDynamicEntitiesUpdateBehaviorReplace     UpdateDynamicEntitiesUpdateBehavior = "REPLACE"
	UpdateDynamicEntitiesUpdateBehaviorClear       UpdateDynamicEntitiesUpdateBehavior = "CLEAR"
)
