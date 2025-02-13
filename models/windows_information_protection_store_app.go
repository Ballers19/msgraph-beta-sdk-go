package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionStoreApp 
type WindowsInformationProtectionStoreApp struct {
    WindowsInformationProtectionApp
}
// NewWindowsInformationProtectionStoreApp instantiates a new WindowsInformationProtectionStoreApp and sets the default values.
func NewWindowsInformationProtectionStoreApp()(*WindowsInformationProtectionStoreApp) {
    m := &WindowsInformationProtectionStoreApp{
        WindowsInformationProtectionApp: *NewWindowsInformationProtectionApp(),
    }
    return m
}
// CreateWindowsInformationProtectionStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionStoreApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsInformationProtectionApp.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsInformationProtectionApp.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
