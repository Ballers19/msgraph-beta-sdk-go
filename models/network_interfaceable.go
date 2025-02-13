package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkInterfaceable 
type NetworkInterfaceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetIpV4Address()(*string)
    GetIpV6Address()(*string)
    GetLocalIpV6Address()(*string)
    GetMacAddress()(*string)
    SetDescription(value *string)()
    SetIpV4Address(value *string)()
    SetIpV6Address(value *string)()
    SetLocalIpV6Address(value *string)()
    SetMacAddress(value *string)()
}
