package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningServicePrincipal 
type ProvisioningServicePrincipal struct {
    Identity
}
// NewProvisioningServicePrincipal instantiates a new ProvisioningServicePrincipal and sets the default values.
func NewProvisioningServicePrincipal()(*ProvisioningServicePrincipal) {
    m := &ProvisioningServicePrincipal{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateProvisioningServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningServicePrincipal(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ProvisioningServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
