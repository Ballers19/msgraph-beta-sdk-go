package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditActorable 
type AuditActorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationDisplayName()(*string)
    GetApplicationId()(*string)
    GetIpAddress()(*string)
    GetRemoteTenantId()(*string)
    GetRemoteUserId()(*string)
    GetServicePrincipalName()(*string)
    GetType()(*string)
    GetUserId()(*string)
    GetUserPermissions()([]string)
    GetUserPrincipalName()(*string)
    GetUserRoleScopeTags()([]RoleScopeTagInfoable)
    SetApplicationDisplayName(value *string)()
    SetApplicationId(value *string)()
    SetIpAddress(value *string)()
    SetRemoteTenantId(value *string)()
    SetRemoteUserId(value *string)()
    SetServicePrincipalName(value *string)()
    SetType(value *string)()
    SetUserId(value *string)()
    SetUserPermissions(value []string)()
    SetUserPrincipalName(value *string)()
    SetUserRoleScopeTags(value []RoleScopeTagInfoable)()
}
