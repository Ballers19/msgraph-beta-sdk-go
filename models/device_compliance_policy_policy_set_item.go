package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyPolicySetItem 
type DeviceCompliancePolicyPolicySetItem struct {
    PolicySetItem
}
// NewDeviceCompliancePolicyPolicySetItem instantiates a new DeviceCompliancePolicyPolicySetItem and sets the default values.
func NewDeviceCompliancePolicyPolicySetItem()(*DeviceCompliancePolicyPolicySetItem) {
    m := &DeviceCompliancePolicyPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    return m
}
// CreateDeviceCompliancePolicyPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicyPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
