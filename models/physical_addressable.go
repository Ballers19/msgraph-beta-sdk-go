package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhysicalAddressable 
type PhysicalAddressable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetPostalCode()(*string)
    GetPostOfficeBox()(*string)
    GetState()(*string)
    GetStreet()(*string)
    GetType()(*PhysicalAddressType)
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetPostalCode(value *string)()
    SetPostOfficeBox(value *string)()
    SetState(value *string)()
    SetStreet(value *string)()
    SetType(value *PhysicalAddressType)()
}
