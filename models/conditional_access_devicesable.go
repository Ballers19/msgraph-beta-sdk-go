package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessDevicesable 
type ConditionalAccessDevicesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceFilter()(ConditionalAccessFilterable)
    GetExcludeDevices()([]string)
    GetExcludeDeviceStates()([]string)
    GetIncludeDevices()([]string)
    GetIncludeDeviceStates()([]string)
    SetDeviceFilter(value ConditionalAccessFilterable)()
    SetExcludeDevices(value []string)()
    SetExcludeDeviceStates(value []string)()
    SetIncludeDevices(value []string)()
    SetIncludeDeviceStates(value []string)()
}
