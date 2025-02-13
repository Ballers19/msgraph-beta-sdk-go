package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSIncludedApp contains properties of an included .app in a MacOS app.
type MacOSIncludedApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The CFBundleIdentifier.
    bundleId *string
    // The CFBundleVersion.
    bundleVersion *string
}
// NewMacOSIncludedApp instantiates a new macOSIncludedApp and sets the default values.
func NewMacOSIncludedApp()(*MacOSIncludedApp) {
    m := &MacOSIncludedApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSIncludedAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSIncludedAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSIncludedApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSIncludedApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBundleId gets the bundleId property value. The CFBundleIdentifier.
func (m *MacOSIncludedApp) GetBundleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bundleId
    }
}
// GetBundleVersion gets the bundleVersion property value. The CFBundleVersion.
func (m *MacOSIncludedApp) GetBundleVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bundleVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSIncludedApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    res["bundleVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleVersion(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MacOSIncludedApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bundleVersion", m.GetBundleVersion())
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
func (m *MacOSIncludedApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBundleId sets the bundleId property value. The CFBundleIdentifier.
func (m *MacOSIncludedApp) SetBundleId(value *string)() {
    if m != nil {
        m.bundleId = value
    }
}
// SetBundleVersion sets the bundleVersion property value. The CFBundleVersion.
func (m *MacOSIncludedApp) SetBundleVersion(value *string)() {
    if m != nil {
        m.bundleVersion = value
    }
}
