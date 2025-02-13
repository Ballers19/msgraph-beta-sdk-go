package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceEndpoint 
type ServiceEndpoint struct {
    Endpoint
}
// NewServiceEndpoint instantiates a new ServiceEndpoint and sets the default values.
func NewServiceEndpoint()(*ServiceEndpoint) {
    m := &ServiceEndpoint{
        Endpoint: *NewEndpoint(),
    }
    return m
}
// CreateServiceEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceEndpoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Endpoint.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ServiceEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Endpoint.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
