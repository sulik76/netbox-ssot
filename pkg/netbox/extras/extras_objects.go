package extras

import (
	"fmt"

	"github.com/bl4ko/netbox-ssot/pkg/netbox/common"
)

// CustomFieldTypes are predefined netbox's types for CustomFields
type CustomFieldType struct {
	common.Choice
}

// Predefined netbox's types for CustomFields
// https://github.com/netbox-community/netbox/blob/35be4f05ef376e28d9af4d7245ba10cc286bb62a/netbox/extras/choices.py#L10
var (
	CustomFieldTypeText     = CustomFieldType{common.Choice{Value: "text", Label: "Text"}}
	CustomFieldTypeLongText = CustomFieldType{common.Choice{Value: "longtext", Label: "Text (long)"}}
	CustomFieldTypeInteger  = CustomFieldType{common.Choice{Value: "integer", Label: "Integer"}}
	CustomFieldTypeDecimal  = CustomFieldType{common.Choice{Value: "decimal", Label: "Decimal"}}
	CustomFieldTypeBoolean  = CustomFieldType{common.Choice{Value: "boolean", Label: "Boolean (true/false)"}}
	CustomFieldTypeDate     = CustomFieldType{common.Choice{Value: "date", Label: "Date"}}
)

type FilterLogic struct {
	common.Choice
}

var (
	FilterLogicLoose = FilterLogic{common.Choice{Value: "loose", Label: "Loose"}}
)

type UIVisibility struct {
	common.Choice
}

var (
	UIVisibilityReadWrite = UIVisibility{common.Choice{Value: "read-write", Label: "Read/write"}}
)

type CustomField struct {
	ID int `json:"id,omitempty"`
	// Name of the custom field (e.g. host_cpu_cores). This field is required.
	Name string `json:"name,omitempty"`
	// Label represents name of the field as displayed to users (e.g. Physical CPU cores). If not provided, the name will be used instead.
	Label string `json:"label,omitempty"`
	// Type is the type of the custom field. Valid choices are: text, integer, boolean, date, url, select, multiselect. This field is required.
	Type CustomFieldType `json:"type,omitempty"`
	// Type of the related object (for object/multi-object fields only) (e.g. dcim.device). This field is required.
	ContentTypes []string `json:"content_types,omitempty"`
	// Description is a description of the field. This field is optional.
	Description string `json:"description,omitempty"`
	// Weighting for search. Lower values are considered more important. Deafult (1000)
	SearchWeight int `json:"search_weight,omitempty"`
	// Filter logic. This field is required. (Default loose)
	FilterLogic FilterLogic `json:"filter_logic,omitempty"`
	// UI visibility. This field is required. (Default read-write)
	UIVisibility UIVisibility `json:"ui_visibility,omitempty"`
	// Display Weight. Fields with higher weights appear lower in a form.
	// default 100
	DisplayWeight int `json:"weight,omitempty"`
}

func (cf CustomField) String() string {
	return fmt.Sprintf("CustomField{ID: %d, Name: %s, Label: %s, ...}", cf.ID, cf.Name, cf.Label)
}
