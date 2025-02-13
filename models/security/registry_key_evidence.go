package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistryKeyEvidence 
type RegistryKeyEvidence struct {
    AlertEvidence
    // The registryHive property
    registryHive *string
    // The registryKey property
    registryKey *string
}
// NewRegistryKeyEvidence instantiates a new RegistryKeyEvidence and sets the default values.
func NewRegistryKeyEvidence()(*RegistryKeyEvidence) {
    m := &RegistryKeyEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateRegistryKeyEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegistryKeyEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegistryKeyEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistryKeyEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["registryHive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryHive(val)
        }
        return nil
    }
    res["registryKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryKey(val)
        }
        return nil
    }
    return res
}
// GetRegistryHive gets the registryHive property value. The registryHive property
func (m *RegistryKeyEvidence) GetRegistryHive()(*string) {
    if m == nil {
        return nil
    } else {
        return m.registryHive
    }
}
// GetRegistryKey gets the registryKey property value. The registryKey property
func (m *RegistryKeyEvidence) GetRegistryKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.registryKey
    }
}
// Serialize serializes information the current object
func (m *RegistryKeyEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("registryHive", m.GetRegistryHive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryKey", m.GetRegistryKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRegistryHive sets the registryHive property value. The registryHive property
func (m *RegistryKeyEvidence) SetRegistryHive(value *string)() {
    if m != nil {
        m.registryHive = value
    }
}
// SetRegistryKey sets the registryKey property value. The registryKey property
func (m *RegistryKeyEvidence) SetRegistryKey(value *string)() {
    if m != nil {
        m.registryKey = value
    }
}
