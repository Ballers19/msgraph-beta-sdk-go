package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantContactInformationable 
type TenantContactInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetName()(*string)
    GetNotes()(*string)
    GetPhone()(*string)
    GetTitle()(*string)
    SetEmail(value *string)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetPhone(value *string)()
    SetTitle(value *string)()
}
