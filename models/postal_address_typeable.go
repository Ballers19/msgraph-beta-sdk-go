package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostalAddressTypeable 
type PostalAddressTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryLetterCode()(*string)
    GetPostalCode()(*string)
    GetState()(*string)
    GetStreet()(*string)
    SetCity(value *string)()
    SetCountryLetterCode(value *string)()
    SetPostalCode(value *string)()
    SetState(value *string)()
    SetStreet(value *string)()
}
