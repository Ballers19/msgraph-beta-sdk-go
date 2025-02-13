package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonNamePronounciationable 
type PersonNamePronounciationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetFirst()(*string)
    GetLast()(*string)
    GetMaiden()(*string)
    GetMiddle()(*string)
    SetDisplayName(value *string)()
    SetFirst(value *string)()
    SetLast(value *string)()
    SetMaiden(value *string)()
    SetMiddle(value *string)()
}
