package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcUserRoleScopeTagInfo 
type CloudPcUserRoleScopeTagInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Scope tag display name.
    displayName *string
    // Scope tag ID.
    roleScopeTagId *string
}
// NewCloudPcUserRoleScopeTagInfo instantiates a new cloudPcUserRoleScopeTagInfo and sets the default values.
func NewCloudPcUserRoleScopeTagInfo()(*CloudPcUserRoleScopeTagInfo) {
    m := &CloudPcUserRoleScopeTagInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcUserRoleScopeTagInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcUserRoleScopeTagInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Scope tag display name.
func (m *CloudPcUserRoleScopeTagInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcUserRoleScopeTagInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["roleScopeTagId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleScopeTagId(val)
        }
        return nil
    }
    return res
}
// GetRoleScopeTagId gets the roleScopeTagId property value. Scope tag ID.
func (m *CloudPcUserRoleScopeTagInfo) GetRoleScopeTagId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagId
    }
}
// Serialize serializes information the current object
func (m *CloudPcUserRoleScopeTagInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roleScopeTagId", m.GetRoleScopeTagId())
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
func (m *CloudPcUserRoleScopeTagInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Scope tag display name.
func (m *CloudPcUserRoleScopeTagInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetRoleScopeTagId sets the roleScopeTagId property value. Scope tag ID.
func (m *CloudPcUserRoleScopeTagInfo) SetRoleScopeTagId(value *string)() {
    if m != nil {
        m.roleScopeTagId = value
    }
}
