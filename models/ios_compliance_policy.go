package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCompliancePolicy 
type IosCompliancePolicy struct {
    DeviceCompliancePolicy
    // Device threat protection levels for the Device Threat Protection API.
    advancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Require that devices have enabled device threat protection .
    deviceThreatProtectionEnabled *bool
    // Device threat protection levels for the Device Threat Protection API.
    deviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Indicates whether or not to require a managed email profile.
    managedEmailProfileRequired *bool
    // Maximum IOS build version.
    osMaximumBuildVersion *string
    // Maximum IOS version.
    osMaximumVersion *string
    // Minimum IOS build version.
    osMinimumBuildVersion *string
    // Minimum IOS version.
    osMinimumVersion *string
    // Indicates whether or not to block simple passcodes.
    passcodeBlockSimple *bool
    // Number of days before the passcode expires. Valid values 1 to 65535
    passcodeExpirationDays *int32
    // The number of character sets required in the password.
    passcodeMinimumCharacterSetCount *int32
    // Minimum length of passcode. Valid values 4 to 14
    passcodeMinimumLength *int32
    // Minutes of inactivity before a passcode is required.
    passcodeMinutesOfInactivityBeforeLock *int32
    // Minutes of inactivity before the screen times out.
    passcodeMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous passcodes to block. Valid values 1 to 24
    passcodePreviousPasscodeBlockCount *int32
    // Indicates whether or not to require a passcode.
    passcodeRequired *bool
    // Possible values of required passwords.
    passcodeRequiredType *RequiredPasswordType
    // Require the device to not have the specified apps installed. This collection can contain a maximum of 100 elements.
    restrictedApps []AppListItemable
    // Devices must not be jailbroken or rooted.
    securityBlockJailbrokenDevices *bool
}
// NewIosCompliancePolicy instantiates a new IosCompliancePolicy and sets the default values.
func NewIosCompliancePolicy()(*IosCompliancePolicy) {
    m := &IosCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    return m
}
// CreateIosCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosCompliancePolicy(), nil
}
// GetAdvancedThreatProtectionRequiredSecurityLevel gets the advancedThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *IosCompliancePolicy) GetAdvancedThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    if m == nil {
        return nil
    } else {
        return m.advancedThreatProtectionRequiredSecurityLevel
    }
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection .
func (m *IosCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceThreatProtectionEnabled
    }
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *IosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    if m == nil {
        return nil
    } else {
        return m.deviceThreatProtectionRequiredSecurityLevel
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["advancedThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["deviceThreatProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionEnabled(val)
        }
        return nil
    }
    res["deviceThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["managedEmailProfileRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedEmailProfileRequired(val)
        }
        return nil
    }
    res["osMaximumBuildVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumBuildVersion(val)
        }
        return nil
    }
    res["osMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumVersion(val)
        }
        return nil
    }
    res["osMinimumBuildVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumBuildVersion(val)
        }
        return nil
    }
    res["osMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumVersion(val)
        }
        return nil
    }
    res["passcodeBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockSimple(val)
        }
        return nil
    }
    res["passcodeExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeExpirationDays(val)
        }
        return nil
    }
    res["passcodeMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumCharacterSetCount(val)
        }
        return nil
    }
    res["passcodeMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumLength(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passcodePreviousPasscodeBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodePreviousPasscodeBlockCount(val)
        }
        return nil
    }
    res["passcodeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequired(val)
        }
        return nil
    }
    res["passcodeRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["restrictedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetRestrictedApps(res)
        }
        return nil
    }
    res["securityBlockJailbrokenDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityBlockJailbrokenDevices(val)
        }
        return nil
    }
    return res
}
// GetManagedEmailProfileRequired gets the managedEmailProfileRequired property value. Indicates whether or not to require a managed email profile.
func (m *IosCompliancePolicy) GetManagedEmailProfileRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.managedEmailProfileRequired
    }
}
// GetOsMaximumBuildVersion gets the osMaximumBuildVersion property value. Maximum IOS build version.
func (m *IosCompliancePolicy) GetOsMaximumBuildVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMaximumBuildVersion
    }
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum IOS version.
func (m *IosCompliancePolicy) GetOsMaximumVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMaximumVersion
    }
}
// GetOsMinimumBuildVersion gets the osMinimumBuildVersion property value. Minimum IOS build version.
func (m *IosCompliancePolicy) GetOsMinimumBuildVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMinimumBuildVersion
    }
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum IOS version.
func (m *IosCompliancePolicy) GetOsMinimumVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMinimumVersion
    }
}
// GetPasscodeBlockSimple gets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosCompliancePolicy) GetPasscodeBlockSimple()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeBlockSimple
    }
}
// GetPasscodeExpirationDays gets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosCompliancePolicy) GetPasscodeExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeExpirationDays
    }
}
// GetPasscodeMinimumCharacterSetCount gets the passcodeMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *IosCompliancePolicy) GetPasscodeMinimumCharacterSetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinimumCharacterSetCount
    }
}
// GetPasscodeMinimumLength gets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosCompliancePolicy) GetPasscodeMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinimumLength
    }
}
// GetPasscodeMinutesOfInactivityBeforeLock gets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinutesOfInactivityBeforeLock
    }
}
// GetPasscodeMinutesOfInactivityBeforeScreenTimeout gets the passcodeMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinutesOfInactivityBeforeScreenTimeout
    }
}
// GetPasscodePreviousPasscodeBlockCount gets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosCompliancePolicy) GetPasscodePreviousPasscodeBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodePreviousPasscodeBlockCount
    }
}
// GetPasscodeRequired gets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosCompliancePolicy) GetPasscodeRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeRequired
    }
}
// GetPasscodeRequiredType gets the passcodeRequiredType property value. Possible values of required passwords.
func (m *IosCompliancePolicy) GetPasscodeRequiredType()(*RequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.passcodeRequiredType
    }
}
// GetRestrictedApps gets the restrictedApps property value. Require the device to not have the specified apps installed. This collection can contain a maximum of 100 elements.
func (m *IosCompliancePolicy) GetRestrictedApps()([]AppListItemable) {
    if m == nil {
        return nil
    } else {
        return m.restrictedApps
    }
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *IosCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityBlockJailbrokenDevices
    }
}
// Serialize serializes information the current object
func (m *IosCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetAdvancedThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("advancedThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceThreatProtectionEnabled", m.GetDeviceThreatProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetDeviceThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("deviceThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedEmailProfileRequired", m.GetManagedEmailProfileRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumBuildVersion", m.GetOsMaximumBuildVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumBuildVersion", m.GetOsMinimumBuildVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockSimple", m.GetPasscodeBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeExpirationDays", m.GetPasscodeExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumCharacterSetCount", m.GetPasscodeMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumLength", m.GetPasscodeMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeLock", m.GetPasscodeMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeScreenTimeout", m.GetPasscodeMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodePreviousPasscodeBlockCount", m.GetPasscodePreviousPasscodeBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeRequired", m.GetPasscodeRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasscodeRequiredType() != nil {
        cast := (*m.GetPasscodeRequiredType()).String()
        err = writer.WriteStringValue("passcodeRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRestrictedApps()))
        for i, v := range m.GetRestrictedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("restrictedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityBlockJailbrokenDevices", m.GetSecurityBlockJailbrokenDevices())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionRequiredSecurityLevel sets the advancedThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *IosCompliancePolicy) SetAdvancedThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    if m != nil {
        m.advancedThreatProtectionRequiredSecurityLevel = value
    }
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection .
func (m *IosCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    if m != nil {
        m.deviceThreatProtectionEnabled = value
    }
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *IosCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    if m != nil {
        m.deviceThreatProtectionRequiredSecurityLevel = value
    }
}
// SetManagedEmailProfileRequired sets the managedEmailProfileRequired property value. Indicates whether or not to require a managed email profile.
func (m *IosCompliancePolicy) SetManagedEmailProfileRequired(value *bool)() {
    if m != nil {
        m.managedEmailProfileRequired = value
    }
}
// SetOsMaximumBuildVersion sets the osMaximumBuildVersion property value. Maximum IOS build version.
func (m *IosCompliancePolicy) SetOsMaximumBuildVersion(value *string)() {
    if m != nil {
        m.osMaximumBuildVersion = value
    }
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum IOS version.
func (m *IosCompliancePolicy) SetOsMaximumVersion(value *string)() {
    if m != nil {
        m.osMaximumVersion = value
    }
}
// SetOsMinimumBuildVersion sets the osMinimumBuildVersion property value. Minimum IOS build version.
func (m *IosCompliancePolicy) SetOsMinimumBuildVersion(value *string)() {
    if m != nil {
        m.osMinimumBuildVersion = value
    }
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum IOS version.
func (m *IosCompliancePolicy) SetOsMinimumVersion(value *string)() {
    if m != nil {
        m.osMinimumVersion = value
    }
}
// SetPasscodeBlockSimple sets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosCompliancePolicy) SetPasscodeBlockSimple(value *bool)() {
    if m != nil {
        m.passcodeBlockSimple = value
    }
}
// SetPasscodeExpirationDays sets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosCompliancePolicy) SetPasscodeExpirationDays(value *int32)() {
    if m != nil {
        m.passcodeExpirationDays = value
    }
}
// SetPasscodeMinimumCharacterSetCount sets the passcodeMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *IosCompliancePolicy) SetPasscodeMinimumCharacterSetCount(value *int32)() {
    if m != nil {
        m.passcodeMinimumCharacterSetCount = value
    }
}
// SetPasscodeMinimumLength sets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosCompliancePolicy) SetPasscodeMinimumLength(value *int32)() {
    if m != nil {
        m.passcodeMinimumLength = value
    }
}
// SetPasscodeMinutesOfInactivityBeforeLock sets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLock(value *int32)() {
    if m != nil {
        m.passcodeMinutesOfInactivityBeforeLock = value
    }
}
// SetPasscodeMinutesOfInactivityBeforeScreenTimeout sets the passcodeMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    if m != nil {
        m.passcodeMinutesOfInactivityBeforeScreenTimeout = value
    }
}
// SetPasscodePreviousPasscodeBlockCount sets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosCompliancePolicy) SetPasscodePreviousPasscodeBlockCount(value *int32)() {
    if m != nil {
        m.passcodePreviousPasscodeBlockCount = value
    }
}
// SetPasscodeRequired sets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosCompliancePolicy) SetPasscodeRequired(value *bool)() {
    if m != nil {
        m.passcodeRequired = value
    }
}
// SetPasscodeRequiredType sets the passcodeRequiredType property value. Possible values of required passwords.
func (m *IosCompliancePolicy) SetPasscodeRequiredType(value *RequiredPasswordType)() {
    if m != nil {
        m.passcodeRequiredType = value
    }
}
// SetRestrictedApps sets the restrictedApps property value. Require the device to not have the specified apps installed. This collection can contain a maximum of 100 elements.
func (m *IosCompliancePolicy) SetRestrictedApps(value []AppListItemable)() {
    if m != nil {
        m.restrictedApps = value
    }
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *IosCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    if m != nil {
        m.securityBlockJailbrokenDevices = value
    }
}
