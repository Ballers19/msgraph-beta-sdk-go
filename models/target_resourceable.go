package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetResourceable 
type TargetResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetGroupType()(*GroupType)
    GetId()(*string)
    GetModifiedProperties()([]ModifiedPropertyable)
    GetType()(*string)
    GetUserPrincipalName()(*string)
    SetDisplayName(value *string)()
    SetGroupType(value *GroupType)()
    SetId(value *string)()
    SetModifiedProperties(value []ModifiedPropertyable)()
    SetType(value *string)()
    SetUserPrincipalName(value *string)()
}
