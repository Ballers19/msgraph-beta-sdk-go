package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssignmentFilterValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicator to valid or invalid rule.
    isValidRule *bool;
}
// Instantiates a new assignmentFilterValidationResult and sets the default values.
func NewAssignmentFilterValidationResult()(*AssignmentFilterValidationResult) {
    m := &AssignmentFilterValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isValidRule property value. Indicator to valid or invalid rule.
func (m *AssignmentFilterValidationResult) GetIsValidRule()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidRule
    }
}
// The deserialization information for the current model
func (m *AssignmentFilterValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isValidRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsValidRule(val)
        return nil
    }
    return res
}
func (m *AssignmentFilterValidationResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssignmentFilterValidationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isValidRule", m.GetIsValidRule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AssignmentFilterValidationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isValidRule property value. Indicator to valid or invalid rule.
// Parameters:
//  - value : Value to set for the isValidRule property.
func (m *AssignmentFilterValidationResult) SetIsValidRule(value *bool)() {
    m.isValidRule = value
}
