package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceSettings 
type ComplianceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The eDiscoveryResultTemplates property
    eDiscoveryResultTemplates []DisplayTemplateable
}
// NewComplianceSettings instantiates a new complianceSettings and sets the default values.
func NewComplianceSettings()(*ComplianceSettings) {
    m := &ComplianceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateComplianceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEDiscoveryResultTemplates gets the eDiscoveryResultTemplates property value. The eDiscoveryResultTemplates property
func (m *ComplianceSettings) GetEDiscoveryResultTemplates()([]DisplayTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.eDiscoveryResultTemplates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eDiscoveryResultTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDisplayTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DisplayTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DisplayTemplateable)
            }
            m.SetEDiscoveryResultTemplates(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ComplianceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEDiscoveryResultTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEDiscoveryResultTemplates()))
        for i, v := range m.GetEDiscoveryResultTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("eDiscoveryResultTemplates", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEDiscoveryResultTemplates sets the eDiscoveryResultTemplates property value. The eDiscoveryResultTemplates property
func (m *ComplianceSettings) SetEDiscoveryResultTemplates(value []DisplayTemplateable)() {
    if m != nil {
        m.eDiscoveryResultTemplates = value
    }
}
