package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsEncryptedIdentity 
type CommunicationsEncryptedIdentity struct {
    Identity
}
// NewCommunicationsEncryptedIdentity instantiates a new CommunicationsEncryptedIdentity and sets the default values.
func NewCommunicationsEncryptedIdentity()(*CommunicationsEncryptedIdentity) {
    m := &CommunicationsEncryptedIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateCommunicationsEncryptedIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsEncryptedIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsEncryptedIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsEncryptedIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CommunicationsEncryptedIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
