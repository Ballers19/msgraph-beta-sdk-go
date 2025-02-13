package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkHardwareDetailable 
type TeamworkHardwareDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMacAddresses()([]string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetSerialNumber()(*string)
    GetUniqueId()(*string)
    SetMacAddresses(value []string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetSerialNumber(value *string)()
    SetUniqueId(value *string)()
}
