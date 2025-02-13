package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkGmailEasConfiguration 
type AndroidForWorkGmailEasConfiguration struct {
    AndroidForWorkEasEmailProfileBase
}
// NewAndroidForWorkGmailEasConfiguration instantiates a new AndroidForWorkGmailEasConfiguration and sets the default values.
func NewAndroidForWorkGmailEasConfiguration()(*AndroidForWorkGmailEasConfiguration) {
    m := &AndroidForWorkGmailEasConfiguration{
        AndroidForWorkEasEmailProfileBase: *NewAndroidForWorkEasEmailProfileBase(),
    }
    return m
}
// CreateAndroidForWorkGmailEasConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkGmailEasConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkGmailEasConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkGmailEasConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidForWorkEasEmailProfileBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AndroidForWorkGmailEasConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidForWorkEasEmailProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
