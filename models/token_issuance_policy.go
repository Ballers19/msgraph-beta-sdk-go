package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TokenIssuancePolicy 
type TokenIssuancePolicy struct {
    StsPolicy
}
// NewTokenIssuancePolicy instantiates a new TokenIssuancePolicy and sets the default values.
func NewTokenIssuancePolicy()(*TokenIssuancePolicy) {
    m := &TokenIssuancePolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// CreateTokenIssuancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTokenIssuancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTokenIssuancePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TokenIssuancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TokenIssuancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
