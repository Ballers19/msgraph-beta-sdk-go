package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutgoingCallOptions 
type OutgoingCallOptions struct {
    CallOptions
}
// NewOutgoingCallOptions instantiates a new OutgoingCallOptions and sets the default values.
func NewOutgoingCallOptions()(*OutgoingCallOptions) {
    m := &OutgoingCallOptions{
        CallOptions: *NewCallOptions(),
    }
    return m
}
// CreateOutgoingCallOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutgoingCallOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutgoingCallOptions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutgoingCallOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CallOptions.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OutgoingCallOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CallOptions.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
