package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SynchronizationError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    code *string;
    // 
    message *string;
    // 
    tenantActionable *bool;
}
// Instantiates a new synchronizationError and sets the default values.
func NewSynchronizationError()(*SynchronizationError) {
    m := &SynchronizationError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the code property value. 
func (m *SynchronizationError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the message property value. 
func (m *SynchronizationError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the tenantActionable property value. 
func (m *SynchronizationError) GetTenantActionable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.tenantActionable
    }
}
// The deserialization information for the current model
func (m *SynchronizationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["tenantActionable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTenantActionable(val)
        return nil
    }
    return res
}
func (m *SynchronizationError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SynchronizationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("tenantActionable", m.GetTenantActionable())
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
func (m *SynchronizationError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the code property value. 
// Parameters:
//  - value : Value to set for the code property.
func (m *SynchronizationError) SetCode(value *string)() {
    m.code = value
}
// Sets the message property value. 
// Parameters:
//  - value : Value to set for the message property.
func (m *SynchronizationError) SetMessage(value *string)() {
    m.message = value
}
// Sets the tenantActionable property value. 
// Parameters:
//  - value : Value to set for the tenantActionable property.
func (m *SynchronizationError) SetTenantActionable(value *bool)() {
    m.tenantActionable = value
}
