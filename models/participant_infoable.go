package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantInfoable 
type ParticipantInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountryCode()(*string)
    GetEndpointType()(*EndpointType)
    GetIdentity()(IdentitySetable)
    GetLanguageId()(*string)
    GetNonAnonymizedIdentity()(IdentitySetable)
    GetParticipantId()(*string)
    GetPlatformId()(*string)
    GetRegion()(*string)
    SetCountryCode(value *string)()
    SetEndpointType(value *EndpointType)()
    SetIdentity(value IdentitySetable)()
    SetLanguageId(value *string)()
    SetNonAnonymizedIdentity(value IdentitySetable)()
    SetParticipantId(value *string)()
    SetPlatformId(value *string)()
    SetRegion(value *string)()
}
