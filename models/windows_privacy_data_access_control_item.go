package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPrivacyDataAccessControlItem 
type WindowsPrivacyDataAccessControlItem struct {
    Entity
    // This indicates an access level for the privacy data category to which the specified application will be given to. Possible values are: notConfigured, forceAllow, forceDeny, userInControl.
    accessLevel *WindowsPrivacyDataAccessLevel
    // The Package Family Name of a Windows app. When set, the access level applies to the specified application.
    appDisplayName *string
    // The Package Family Name of a Windows app. When set, the access level applies to the specified application.
    appPackageFamilyName *string
    // This indicates a privacy data category to which the specific access control will apply. Possible values are: notConfigured, accountInfo, appsRunInBackground, calendar, callHistory, camera, contacts, diagnosticsInfo, email, location, messaging, microphone, motion, notifications, phone, radios, tasks, syncWithDevices, trustedDevices.
    dataCategory *WindowsPrivacyDataCategory
}
// NewWindowsPrivacyDataAccessControlItem instantiates a new windowsPrivacyDataAccessControlItem and sets the default values.
func NewWindowsPrivacyDataAccessControlItem()(*WindowsPrivacyDataAccessControlItem) {
    m := &WindowsPrivacyDataAccessControlItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPrivacyDataAccessControlItem(), nil
}
// GetAccessLevel gets the accessLevel property value. This indicates an access level for the privacy data category to which the specified application will be given to. Possible values are: notConfigured, forceAllow, forceDeny, userInControl.
func (m *WindowsPrivacyDataAccessControlItem) GetAccessLevel()(*WindowsPrivacyDataAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.accessLevel
    }
}
// GetAppDisplayName gets the appDisplayName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppPackageFamilyName gets the appPackageFamilyName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) GetAppPackageFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPackageFamilyName
    }
}
// GetDataCategory gets the dataCategory property value. This indicates a privacy data category to which the specific access control will apply. Possible values are: notConfigured, accountInfo, appsRunInBackground, calendar, callHistory, camera, contacts, diagnosticsInfo, email, location, messaging, microphone, motion, notifications, phone, radios, tasks, syncWithDevices, trustedDevices.
func (m *WindowsPrivacyDataAccessControlItem) GetDataCategory()(*WindowsPrivacyDataCategory) {
    if m == nil {
        return nil
    } else {
        return m.dataCategory
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPrivacyDataAccessControlItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsPrivacyDataAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*WindowsPrivacyDataAccessLevel))
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appPackageFamilyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPackageFamilyName(val)
        }
        return nil
    }
    res["dataCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsPrivacyDataCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataCategory(val.(*WindowsPrivacyDataCategory))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsPrivacyDataAccessControlItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
        err = writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPackageFamilyName", m.GetAppPackageFamilyName())
        if err != nil {
            return err
        }
    }
    if m.GetDataCategory() != nil {
        cast := (*m.GetDataCategory()).String()
        err = writer.WriteStringValue("dataCategory", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessLevel sets the accessLevel property value. This indicates an access level for the privacy data category to which the specified application will be given to. Possible values are: notConfigured, forceAllow, forceDeny, userInControl.
func (m *WindowsPrivacyDataAccessControlItem) SetAccessLevel(value *WindowsPrivacyDataAccessLevel)() {
    if m != nil {
        m.accessLevel = value
    }
}
// SetAppDisplayName sets the appDisplayName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppPackageFamilyName sets the appPackageFamilyName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) SetAppPackageFamilyName(value *string)() {
    if m != nil {
        m.appPackageFamilyName = value
    }
}
// SetDataCategory sets the dataCategory property value. This indicates a privacy data category to which the specific access control will apply. Possible values are: notConfigured, accountInfo, appsRunInBackground, calendar, callHistory, camera, contacts, diagnosticsInfo, email, location, messaging, microphone, motion, notifications, phone, radios, tasks, syncWithDevices, trustedDevices.
func (m *WindowsPrivacyDataAccessControlItem) SetDataCategory(value *WindowsPrivacyDataCategory)() {
    if m != nil {
        m.dataCategory = value
    }
}
