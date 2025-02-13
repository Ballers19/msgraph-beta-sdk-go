package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChromeOSDevicePropertyable 
type ChromeOSDevicePropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetUpdatable()(*bool)
    GetValue()(*string)
    GetValueType()(*string)
    SetName(value *string)()
    SetUpdatable(value *bool)()
    SetValue(value *string)()
    SetValueType(value *string)()
}
