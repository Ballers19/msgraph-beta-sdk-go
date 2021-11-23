package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PasswordSingleSignOnField 
type PasswordSingleSignOnField struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Title/label override for customization.
    customizedLabel *string;
    // Label that would be used if no customizedLabel is provided. Read only.
    defaultLabel *string;
    // Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
    fieldId *string;
    // Type of the credential. The values can be text, password.
    type_escaped *string;
}
// NewPasswordSingleSignOnField instantiates a new passwordSingleSignOnField and sets the default values.
func NewPasswordSingleSignOnField()(*PasswordSingleSignOnField) {
    m := &PasswordSingleSignOnField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnField) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCustomizedLabel gets the customizedLabel property value. Title/label override for customization.
func (m *PasswordSingleSignOnField) GetCustomizedLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedLabel
    }
}
// GetDefaultLabel gets the defaultLabel property value. Label that would be used if no customizedLabel is provided. Read only.
func (m *PasswordSingleSignOnField) GetDefaultLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLabel
    }
}
// GetFieldId gets the fieldId property value. Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
func (m *PasswordSingleSignOnField) GetFieldId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fieldId
    }
}
// GetType_escaped gets the type_escaped property value. Type of the credential. The values can be text, password.
func (m *PasswordSingleSignOnField) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordSingleSignOnField) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["customizedLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedLabel(val)
        }
        return nil
    }
    res["defaultLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLabel(val)
        }
        return nil
    }
    res["fieldId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFieldId(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *PasswordSingleSignOnField) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PasswordSingleSignOnField) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customizedLabel", m.GetCustomizedLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultLabel", m.GetDefaultLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fieldId", m.GetFieldId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnField) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCustomizedLabel sets the customizedLabel property value. Title/label override for customization.
func (m *PasswordSingleSignOnField) SetCustomizedLabel(value *string)() {
    m.customizedLabel = value
}
// SetDefaultLabel sets the defaultLabel property value. Label that would be used if no customizedLabel is provided. Read only.
func (m *PasswordSingleSignOnField) SetDefaultLabel(value *string)() {
    m.defaultLabel = value
}
// SetFieldId sets the fieldId property value. Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
func (m *PasswordSingleSignOnField) SetFieldId(value *string)() {
    m.fieldId = value
}
// SetType_escaped sets the type_escaped property value. Type of the credential. The values can be text, password.
func (m *PasswordSingleSignOnField) SetType_escaped(value *string)() {
    m.type_escaped = value
}
