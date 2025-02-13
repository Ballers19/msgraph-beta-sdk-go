package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeHomeButtonOpensNewTab 
type EdgeHomeButtonOpensNewTab struct {
    EdgeHomeButtonConfiguration
}
// NewEdgeHomeButtonOpensNewTab instantiates a new EdgeHomeButtonOpensNewTab and sets the default values.
func NewEdgeHomeButtonOpensNewTab()(*EdgeHomeButtonOpensNewTab) {
    m := &EdgeHomeButtonOpensNewTab{
        EdgeHomeButtonConfiguration: *NewEdgeHomeButtonConfiguration(),
    }
    return m
}
// CreateEdgeHomeButtonOpensNewTabFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeHomeButtonOpensNewTabFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeHomeButtonOpensNewTab(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeHomeButtonOpensNewTab) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeHomeButtonConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *EdgeHomeButtonOpensNewTab) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeHomeButtonConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
