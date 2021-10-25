package addcopy

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AddCopyRequestBody struct {
    additionalData map[string]interface{};
    contentType *string;
}
func NewAddCopyRequestBody()(*AddCopyRequestBody) {
    m := &AddCopyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AddCopyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AddCopyRequestBody) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *AddCopyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    return res
}
func (m *AddCopyRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *AddCopyRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
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
func (m *AddCopyRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AddCopyRequestBody) SetContentType(value *string)() {
    m.contentType = value
}
