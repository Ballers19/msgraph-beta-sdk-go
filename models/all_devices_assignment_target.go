package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AllDevicesAssignmentTarget 
type AllDevicesAssignmentTarget struct {
    DeviceAndAppManagementAssignmentTarget
}
// NewAllDevicesAssignmentTarget instantiates a new AllDevicesAssignmentTarget and sets the default values.
func NewAllDevicesAssignmentTarget()(*AllDevicesAssignmentTarget) {
    m := &AllDevicesAssignmentTarget{
        DeviceAndAppManagementAssignmentTarget: *NewDeviceAndAppManagementAssignmentTarget(),
    }
    return m
}
// CreateAllDevicesAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAllDevicesAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllDevicesAssignmentTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AllDevicesAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentTarget.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AllDevicesAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
