package updatewindowsdeviceaccount

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type UpdateWindowsDeviceAccountRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    updateWindowsDeviceAccountActionParameter *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter;
}
// Instantiates a new updateWindowsDeviceAccountRequestBody and sets the default values.
func NewUpdateWindowsDeviceAccountRequestBody()(*UpdateWindowsDeviceAccountRequestBody) {
    m := &UpdateWindowsDeviceAccountRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindowsDeviceAccountRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the updateWindowsDeviceAccountActionParameter property value. 
func (m *UpdateWindowsDeviceAccountRequestBody) GetUpdateWindowsDeviceAccountActionParameter()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter) {
    if m == nil {
        return nil
    } else {
        return m.updateWindowsDeviceAccountActionParameter
    }
}
// The deserialization information for the current model
func (m *UpdateWindowsDeviceAccountRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["updateWindowsDeviceAccountActionParameter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUpdateWindowsDeviceAccountActionParameter() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowsDeviceAccountActionParameter(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter))
        }
        return nil
    }
    return res
}
func (m *UpdateWindowsDeviceAccountRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UpdateWindowsDeviceAccountRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateWindowsDeviceAccountActionParameter", m.GetUpdateWindowsDeviceAccountActionParameter())
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
func (m *UpdateWindowsDeviceAccountRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the updateWindowsDeviceAccountActionParameter property value. 
// Parameters:
//  - value : Value to set for the updateWindowsDeviceAccountActionParameter property.
func (m *UpdateWindowsDeviceAccountRequestBody) SetUpdateWindowsDeviceAccountActionParameter(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter)() {
    m.updateWindowsDeviceAccountActionParameter = value
}
