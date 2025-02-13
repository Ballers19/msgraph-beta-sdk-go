package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivateLinkDetailsable 
type PrivateLinkDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPolicyTenantId()(*string)
    GetResourceId()(*string)
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPolicyTenantId(value *string)()
    SetResourceId(value *string)()
}
