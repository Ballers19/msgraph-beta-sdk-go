package getmanagementconditionstatementsforplatformwithplatform

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GetManagementConditionStatementsForPlatformWithPlatform struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// Instantiates a new getManagementConditionStatementsForPlatformWithPlatform and sets the default values.
func NewGetManagementConditionStatementsForPlatformWithPlatform()(*GetManagementConditionStatementsForPlatformWithPlatform) {
    m := &GetManagementConditionStatementsForPlatformWithPlatform{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetManagementConditionStatementsForPlatformWithPlatform) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// The deserialization information for the current model
func (m *GetManagementConditionStatementsForPlatformWithPlatform) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *GetManagementConditionStatementsForPlatformWithPlatform) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetManagementConditionStatementsForPlatformWithPlatform) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetManagementConditionStatementsForPlatformWithPlatform) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
