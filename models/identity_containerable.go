package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityContainerable 
type IdentityContainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiConnectors()([]IdentityApiConnectorable)
    GetB2cUserFlows()([]B2cIdentityUserFlowable)
    GetB2xUserFlows()([]B2xIdentityUserFlowable)
    GetConditionalAccess()(ConditionalAccessRootable)
    GetContinuousAccessEvaluationPolicy()(ContinuousAccessEvaluationPolicyable)
    GetIdentityProviders()([]IdentityProviderBaseable)
    GetUserFlowAttributes()([]IdentityUserFlowAttributeable)
    GetUserFlows()([]IdentityUserFlowable)
    SetApiConnectors(value []IdentityApiConnectorable)()
    SetB2cUserFlows(value []B2cIdentityUserFlowable)()
    SetB2xUserFlows(value []B2xIdentityUserFlowable)()
    SetConditionalAccess(value ConditionalAccessRootable)()
    SetContinuousAccessEvaluationPolicy(value ContinuousAccessEvaluationPolicyable)()
    SetIdentityProviders(value []IdentityProviderBaseable)()
    SetUserFlowAttributes(value []IdentityUserFlowAttributeable)()
    SetUserFlows(value []IdentityUserFlowable)()
}
