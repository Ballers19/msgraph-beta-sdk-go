package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleEligibilitySchedule 
type UnifiedRoleEligibilitySchedule struct {
    UnifiedRoleScheduleBase
    // Membership type of the eligible assignment. It can either be Inherited, Direct, or Group. Supports $filter (eq).
    memberType *string
    // The schedule object of the eligible role assignment request.
    scheduleInfo RequestScheduleable
}
// NewUnifiedRoleEligibilitySchedule instantiates a new unifiedRoleEligibilitySchedule and sets the default values.
func NewUnifiedRoleEligibilitySchedule()(*UnifiedRoleEligibilitySchedule) {
    m := &UnifiedRoleEligibilitySchedule{
        UnifiedRoleScheduleBase: *NewUnifiedRoleScheduleBase(),
    }
    return m
}
// CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleEligibilitySchedule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleEligibilitySchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleBase.GetFieldDeserializers()
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val)
        }
        return nil
    }
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(RequestScheduleable))
        }
        return nil
    }
    return res
}
// GetMemberType gets the memberType property value. Membership type of the eligible assignment. It can either be Inherited, Direct, or Group. Supports $filter (eq).
func (m *UnifiedRoleEligibilitySchedule) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// GetScheduleInfo gets the scheduleInfo property value. The schedule object of the eligible role assignment request.
func (m *UnifiedRoleEligibilitySchedule) GetScheduleInfo()(RequestScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleEligibilitySchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMemberType sets the memberType property value. Membership type of the eligible assignment. It can either be Inherited, Direct, or Group. Supports $filter (eq).
func (m *UnifiedRoleEligibilitySchedule) SetMemberType(value *string)() {
    if m != nil {
        m.memberType = value
    }
}
// SetScheduleInfo sets the scheduleInfo property value. The schedule object of the eligible role assignment request.
func (m *UnifiedRoleEligibilitySchedule) SetScheduleInfo(value RequestScheduleable)() {
    if m != nil {
        m.scheduleInfo = value
    }
}
