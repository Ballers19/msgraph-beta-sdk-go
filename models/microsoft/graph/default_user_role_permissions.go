package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DefaultUserRolePermissions 
type DefaultUserRolePermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the default user role can create applications.
    allowedToCreateApps *bool;
    // Indicates whether the default user role can create security groups.
    allowedToCreateSecurityGroups *bool;
    // Indicates whether the default user role can read other users.
    allowedToReadOtherUsers *bool;
}
// NewDefaultUserRolePermissions instantiates a new defaultUserRolePermissions and sets the default values.
func NewDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    m := &DefaultUserRolePermissions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedToCreateApps gets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) GetAllowedToCreateApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateApps
    }
}
// GetAllowedToCreateSecurityGroups gets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) GetAllowedToCreateSecurityGroups()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateSecurityGroups
    }
}
// GetAllowedToReadOtherUsers gets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) GetAllowedToReadOtherUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToReadOtherUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRolePermissions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedToCreateApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateApps(val)
        }
        return nil
    }
    res["allowedToCreateSecurityGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateSecurityGroups(val)
        }
        return nil
    }
    res["allowedToReadOtherUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToReadOtherUsers(val)
        }
        return nil
    }
    return res
}
func (m *DefaultUserRolePermissions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DefaultUserRolePermissions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowedToCreateApps", m.GetAllowedToCreateApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToCreateSecurityGroups", m.GetAllowedToCreateSecurityGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToReadOtherUsers", m.GetAllowedToReadOtherUsers())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedToCreateApps sets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) SetAllowedToCreateApps(value *bool)() {
    m.allowedToCreateApps = value
}
// SetAllowedToCreateSecurityGroups sets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(value *bool)() {
    m.allowedToCreateSecurityGroups = value
}
// SetAllowedToReadOtherUsers sets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) SetAllowedToReadOtherUsers(value *bool)() {
    m.allowedToReadOtherUsers = value
}
