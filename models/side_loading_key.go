package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SideLoadingKey sideLoadingKey entity is required for Windows 8 and 8.1 devices to intall Line Of Business Apps for a tenant.
type SideLoadingKey struct {
    Entity
    // Side Loading Key description displayed to the ITPro Admins..
    description *string
    // Side Loading Key Name displayed to the ITPro Admins.
    displayName *string
    // Side Loading Key Last Updated Date displayed to the ITPro Admins.
    lastUpdatedDateTime *string
    // Side Loading Key Total Activation displayed to the ITPro Admins.
    totalActivation *int32
    // Side Loading Key Value, it is 5x5 value, seperated by hiphens.
    value *string
}
// NewSideLoadingKey instantiates a new sideLoadingKey and sets the default values.
func NewSideLoadingKey()(*SideLoadingKey) {
    m := &SideLoadingKey{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSideLoadingKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSideLoadingKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSideLoadingKey(), nil
}
// GetDescription gets the description property value. Side Loading Key description displayed to the ITPro Admins..
func (m *SideLoadingKey) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Side Loading Key Name displayed to the ITPro Admins.
func (m *SideLoadingKey) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SideLoadingKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["totalActivation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActivation(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Side Loading Key Last Updated Date displayed to the ITPro Admins.
func (m *SideLoadingKey) GetLastUpdatedDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetTotalActivation gets the totalActivation property value. Side Loading Key Total Activation displayed to the ITPro Admins.
func (m *SideLoadingKey) GetTotalActivation()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalActivation
    }
}
// GetValue gets the value property value. Side Loading Key Value, it is 5x5 value, seperated by hiphens.
func (m *SideLoadingKey) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *SideLoadingKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalActivation", m.GetTotalActivation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Side Loading Key description displayed to the ITPro Admins..
func (m *SideLoadingKey) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Side Loading Key Name displayed to the ITPro Admins.
func (m *SideLoadingKey) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Side Loading Key Last Updated Date displayed to the ITPro Admins.
func (m *SideLoadingKey) SetLastUpdatedDateTime(value *string)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetTotalActivation sets the totalActivation property value. Side Loading Key Total Activation displayed to the ITPro Admins.
func (m *SideLoadingKey) SetTotalActivation(value *int32)() {
    if m != nil {
        m.totalActivation = value
    }
}
// SetValue sets the value property value. Side Loading Key Value, it is 5x5 value, seperated by hiphens.
func (m *SideLoadingKey) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
