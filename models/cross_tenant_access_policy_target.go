package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyTarget 
type CrossTenantAccessPolicyTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier of the user, group, or application; one of the following keywords: AllUsers and AllApplications; or for targets that are applications, you may use reserved values.
    target *string
    // The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
    targetType *CrossTenantAccessPolicyTargetType
}
// NewCrossTenantAccessPolicyTarget instantiates a new crossTenantAccessPolicyTarget and sets the default values.
func NewCrossTenantAccessPolicyTarget()(*CrossTenantAccessPolicyTarget) {
    m := &CrossTenantAccessPolicyTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCrossTenantAccessPolicyTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    res["targetType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCrossTenantAccessPolicyTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*CrossTenantAccessPolicyTargetType))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The unique identifier of the user, group, or application; one of the following keywords: AllUsers and AllApplications; or for targets that are applications, you may use reserved values.
func (m *CrossTenantAccessPolicyTarget) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetTargetType gets the targetType property value. The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
func (m *CrossTenantAccessPolicyTarget) GetTargetType()(*CrossTenantAccessPolicyTargetType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("targetType", &cast)
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
func (m *CrossTenantAccessPolicyTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTarget sets the target property value. The unique identifier of the user, group, or application; one of the following keywords: AllUsers and AllApplications; or for targets that are applications, you may use reserved values.
func (m *CrossTenantAccessPolicyTarget) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
// SetTargetType sets the targetType property value. The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
func (m *CrossTenantAccessPolicyTarget) SetTargetType(value *CrossTenantAccessPolicyTargetType)() {
    if m != nil {
        m.targetType = value
    }
}
