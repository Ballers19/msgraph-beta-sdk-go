package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type Setting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The display name for the setting. Required. Read-only.
    displayName *string;
    // The value for the setting serialized as string of JSON. Required. Read-only.
    jsonValue *string;
    // A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
    overwriteAllowed *bool;
    // The data type for the setting. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
    valueType *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType;
}
// Instantiates a new setting and sets the default values.
func NewSetting()(*Setting) {
    m := &Setting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Setting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. The display name for the setting. Required. Read-only.
func (m *Setting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
func (m *Setting) GetJsonValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonValue
    }
}
// Gets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
func (m *Setting) GetOverwriteAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteAllowed
    }
}
// Gets the valueType property value. The data type for the setting. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
func (m *Setting) GetValueType()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// The deserialization information for the current model
func (m *Setting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["jsonValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJsonValue(val)
        return nil
    }
    res["overwriteAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOverwriteAllowed(val)
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)
        m.SetValueType(&cast)
        return nil
    }
    return res
}
func (m *Setting) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Setting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonValue", m.GetJsonValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteAllowed", m.GetOverwriteAllowed())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := m.GetValueType().String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *Setting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. The display name for the setting. Required. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Setting) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
// Parameters:
//  - value : Value to set for the jsonValue property.
func (m *Setting) SetJsonValue(value *string)() {
    m.jsonValue = value
}
// Sets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
// Parameters:
//  - value : Value to set for the overwriteAllowed property.
func (m *Setting) SetOverwriteAllowed(value *bool)() {
    m.overwriteAllowed = value
}
// Sets the valueType property value. The data type for the setting. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
// Parameters:
//  - value : Value to set for the valueType property.
func (m *Setting) SetValueType(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)() {
    m.valueType = value
}
