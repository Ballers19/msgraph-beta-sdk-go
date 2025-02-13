package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CredentialUserRegistrationsSummary provides operations to manage the collection of accessReviewDecision entities.
type CredentialUserRegistrationsSummary struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of users that are capable of performing multi-factor authentication or self service password reset. Optional. Read-only.
    mfaAndSsprCapableUserCount *int32
    // The state of a conditional access policy that enforces multi-factor authentication. Optional. Read-only.
    mfaConditionalAccessPolicyState *string
    // The number of users in the multi-factor authentication exclusion security group (Microsoft 365 Lighthouse - MFA exclusions). Optional. Read-only.
    mfaExcludedUserCount *int32
    // The number of users registered for multi-factor authentication. Optional. Read-only.
    mfaRegisteredUserCount *int32
    // A flag indicating whether Identity Security Defaults is enabled. Optional. Read-only.
    securityDefaultsEnabled *bool
    // The number of users enabled for self service password reset. Optional. Read-only.
    ssprEnabledUserCount *int32
    // The number of users registered for self service password reset. Optional. Read-only.
    ssprRegisteredUserCount *int32
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string
    // The total number of users in the given managed tenant. Optional. Read-only.
    totalUserCount *int32
}
// NewCredentialUserRegistrationsSummary instantiates a new credentialUserRegistrationsSummary and sets the default values.
func NewCredentialUserRegistrationsSummary()(*CredentialUserRegistrationsSummary) {
    m := &CredentialUserRegistrationsSummary{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationsSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationsSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["mfaAndSsprCapableUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaAndSsprCapableUserCount(val)
        }
        return nil
    }
    res["mfaConditionalAccessPolicyState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaConditionalAccessPolicyState(val)
        }
        return nil
    }
    res["mfaExcludedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaExcludedUserCount(val)
        }
        return nil
    }
    res["mfaRegisteredUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaRegisteredUserCount(val)
        }
        return nil
    }
    res["securityDefaultsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDefaultsEnabled(val)
        }
        return nil
    }
    res["ssprEnabledUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsprEnabledUserCount(val)
        }
        return nil
    }
    res["ssprRegisteredUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsprRegisteredUserCount(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["totalUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserCount(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// GetMfaAndSsprCapableUserCount gets the mfaAndSsprCapableUserCount property value. The number of users that are capable of performing multi-factor authentication or self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaAndSsprCapableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mfaAndSsprCapableUserCount
    }
}
// GetMfaConditionalAccessPolicyState gets the mfaConditionalAccessPolicyState property value. The state of a conditional access policy that enforces multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaConditionalAccessPolicyState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mfaConditionalAccessPolicyState
    }
}
// GetMfaExcludedUserCount gets the mfaExcludedUserCount property value. The number of users in the multi-factor authentication exclusion security group (Microsoft 365 Lighthouse - MFA exclusions). Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaExcludedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mfaExcludedUserCount
    }
}
// GetMfaRegisteredUserCount gets the mfaRegisteredUserCount property value. The number of users registered for multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaRegisteredUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mfaRegisteredUserCount
    }
}
// GetSecurityDefaultsEnabled gets the securityDefaultsEnabled property value. A flag indicating whether Identity Security Defaults is enabled. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetSecurityDefaultsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityDefaultsEnabled
    }
}
// GetSsprEnabledUserCount gets the ssprEnabledUserCount property value. The number of users enabled for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetSsprEnabledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ssprEnabledUserCount
    }
}
// GetSsprRegisteredUserCount gets the ssprRegisteredUserCount property value. The number of users registered for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetSsprRegisteredUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ssprRegisteredUserCount
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetTotalUserCount gets the totalUserCount property value. The total number of users in the given managed tenant. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetTotalUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// Serialize serializes information the current object
func (m *CredentialUserRegistrationsSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaAndSsprCapableUserCount", m.GetMfaAndSsprCapableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mfaConditionalAccessPolicyState", m.GetMfaConditionalAccessPolicyState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaExcludedUserCount", m.GetMfaExcludedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaRegisteredUserCount", m.GetMfaRegisteredUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDefaultsEnabled", m.GetSecurityDefaultsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ssprEnabledUserCount", m.GetSsprEnabledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ssprRegisteredUserCount", m.GetSsprRegisteredUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshedDateTime = value
    }
}
// SetMfaAndSsprCapableUserCount sets the mfaAndSsprCapableUserCount property value. The number of users that are capable of performing multi-factor authentication or self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaAndSsprCapableUserCount(value *int32)() {
    if m != nil {
        m.mfaAndSsprCapableUserCount = value
    }
}
// SetMfaConditionalAccessPolicyState sets the mfaConditionalAccessPolicyState property value. The state of a conditional access policy that enforces multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaConditionalAccessPolicyState(value *string)() {
    if m != nil {
        m.mfaConditionalAccessPolicyState = value
    }
}
// SetMfaExcludedUserCount sets the mfaExcludedUserCount property value. The number of users in the multi-factor authentication exclusion security group (Microsoft 365 Lighthouse - MFA exclusions). Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaExcludedUserCount(value *int32)() {
    if m != nil {
        m.mfaExcludedUserCount = value
    }
}
// SetMfaRegisteredUserCount sets the mfaRegisteredUserCount property value. The number of users registered for multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaRegisteredUserCount(value *int32)() {
    if m != nil {
        m.mfaRegisteredUserCount = value
    }
}
// SetSecurityDefaultsEnabled sets the securityDefaultsEnabled property value. A flag indicating whether Identity Security Defaults is enabled. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetSecurityDefaultsEnabled(value *bool)() {
    if m != nil {
        m.securityDefaultsEnabled = value
    }
}
// SetSsprEnabledUserCount sets the ssprEnabledUserCount property value. The number of users enabled for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetSsprEnabledUserCount(value *int32)() {
    if m != nil {
        m.ssprEnabledUserCount = value
    }
}
// SetSsprRegisteredUserCount sets the ssprRegisteredUserCount property value. The number of users registered for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetSsprRegisteredUserCount(value *int32)() {
    if m != nil {
        m.ssprRegisteredUserCount = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetTotalUserCount sets the totalUserCount property value. The total number of users in the given managed tenant. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetTotalUserCount(value *int32)() {
    if m != nil {
        m.totalUserCount = value
    }
}
