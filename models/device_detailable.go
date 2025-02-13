package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceDetailable 
type DeviceDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrowser()(*string)
    GetBrowserId()(*string)
    GetDeviceId()(*string)
    GetDisplayName()(*string)
    GetIsCompliant()(*bool)
    GetIsManaged()(*bool)
    GetOperatingSystem()(*string)
    GetTrustType()(*string)
    SetBrowser(value *string)()
    SetBrowserId(value *string)()
    SetDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetIsCompliant(value *bool)()
    SetIsManaged(value *bool)()
    SetOperatingSystem(value *string)()
    SetTrustType(value *string)()
}
