package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AutomaticRepliesSettingable 
type AutomaticRepliesSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalAudience()(*ExternalAudienceScope)
    GetExternalReplyMessage()(*string)
    GetInternalReplyMessage()(*string)
    GetScheduledEndDateTime()(DateTimeTimeZoneable)
    GetScheduledStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*AutomaticRepliesStatus)
    SetExternalAudience(value *ExternalAudienceScope)()
    SetExternalReplyMessage(value *string)()
    SetInternalReplyMessage(value *string)()
    SetScheduledEndDateTime(value DateTimeTimeZoneable)()
    SetScheduledStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *AutomaticRepliesStatus)()
}
