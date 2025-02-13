package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosAzureAdSingleSignOnExtension 
type IosAzureAdSingleSignOnExtension struct {
    IosSingleSignOnExtension
    // An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
    bundleIdAccessControlList []string
    // Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
    configurations []KeyTypedValuePairable
    // Enables or disables shared device mode.
    enableSharedDeviceMode *bool
}
// NewIosAzureAdSingleSignOnExtension instantiates a new IosAzureAdSingleSignOnExtension and sets the default values.
func NewIosAzureAdSingleSignOnExtension()(*IosAzureAdSingleSignOnExtension) {
    m := &IosAzureAdSingleSignOnExtension{
        IosSingleSignOnExtension: *NewIosSingleSignOnExtension(),
    }
    return m
}
// CreateIosAzureAdSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosAzureAdSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosAzureAdSingleSignOnExtension(), nil
}
// GetBundleIdAccessControlList gets the bundleIdAccessControlList property value. An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
func (m *IosAzureAdSingleSignOnExtension) GetBundleIdAccessControlList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bundleIdAccessControlList
    }
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *IosAzureAdSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.configurations
    }
}
// GetEnableSharedDeviceMode gets the enableSharedDeviceMode property value. Enables or disables shared device mode.
func (m *IosAzureAdSingleSignOnExtension) GetEnableSharedDeviceMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSharedDeviceMode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosAzureAdSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosSingleSignOnExtension.GetFieldDeserializers()
    res["bundleIdAccessControlList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBundleIdAccessControlList(res)
        }
        return nil
    }
    res["configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyTypedValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyTypedValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyTypedValuePairable)
            }
            m.SetConfigurations(res)
        }
        return nil
    }
    res["enableSharedDeviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSharedDeviceMode(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IosAzureAdSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBundleIdAccessControlList() != nil {
        err = writer.WriteCollectionOfStringValues("bundleIdAccessControlList", m.GetBundleIdAccessControlList())
        if err != nil {
            return err
        }
    }
    if m.GetConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurations()))
        for i, v := range m.GetConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSharedDeviceMode", m.GetEnableSharedDeviceMode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundleIdAccessControlList sets the bundleIdAccessControlList property value. An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
func (m *IosAzureAdSingleSignOnExtension) SetBundleIdAccessControlList(value []string)() {
    if m != nil {
        m.bundleIdAccessControlList = value
    }
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *IosAzureAdSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    if m != nil {
        m.configurations = value
    }
}
// SetEnableSharedDeviceMode sets the enableSharedDeviceMode property value. Enables or disables shared device mode.
func (m *IosAzureAdSingleSignOnExtension) SetEnableSharedDeviceMode(value *bool)() {
    if m != nil {
        m.enableSharedDeviceMode = value
    }
}
