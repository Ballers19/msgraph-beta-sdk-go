package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyB2BSetting 
type CrossTenantAccessPolicyB2BSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of applications targeted with your cross-tenant access policy.
    applications CrossTenantAccessPolicyTargetConfigurationable
    // The list of users and groups targeted with your cross-tenant access policy.
    usersAndGroups CrossTenantAccessPolicyTargetConfigurationable
}
// NewCrossTenantAccessPolicyB2BSetting instantiates a new crossTenantAccessPolicyB2BSetting and sets the default values.
func NewCrossTenantAccessPolicyB2BSetting()(*CrossTenantAccessPolicyB2BSetting) {
    m := &CrossTenantAccessPolicyB2BSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyB2BSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyB2BSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplications gets the applications property value. The list of applications targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) GetApplications()(CrossTenantAccessPolicyTargetConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyB2BSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyTargetConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplications(val.(CrossTenantAccessPolicyTargetConfigurationable))
        }
        return nil
    }
    res["usersAndGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyTargetConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersAndGroups(val.(CrossTenantAccessPolicyTargetConfigurationable))
        }
        return nil
    }
    return res
}
// GetUsersAndGroups gets the usersAndGroups property value. The list of users and groups targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) GetUsersAndGroups()(CrossTenantAccessPolicyTargetConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.usersAndGroups
    }
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyB2BSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applications", m.GetApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("usersAndGroups", m.GetUsersAndGroups())
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
func (m *CrossTenantAccessPolicyB2BSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplications sets the applications property value. The list of applications targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) SetApplications(value CrossTenantAccessPolicyTargetConfigurationable)() {
    if m != nil {
        m.applications = value
    }
}
// SetUsersAndGroups sets the usersAndGroups property value. The list of users and groups targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) SetUsersAndGroups(value CrossTenantAccessPolicyTargetConfigurationable)() {
    if m != nil {
        m.usersAndGroups = value
    }
}
