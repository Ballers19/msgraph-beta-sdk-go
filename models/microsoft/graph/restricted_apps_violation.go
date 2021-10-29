package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RestrictedAppsViolation struct {
    Entity
    // Device configuration profile unique identifier, must be Guid
    deviceConfigurationId *string;
    // Device configuration profile name
    deviceConfigurationName *string;
    // Device name
    deviceName *string;
    // Managed device unique identifier, must be Guid
    managedDeviceId *string;
    // Platform type. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
    platformType *PolicyPlatformType;
    // List of violated restricted apps
    restrictedApps []ManagedDeviceReportedApp;
    // Restricted apps state. Possible values are: prohibitedApps, notApprovedApps.
    restrictedAppsState *RestrictedAppsState;
    // User unique identifier, must be Guid
    userId *string;
    // User name
    userName *string;
}
// Instantiates a new restrictedAppsViolation and sets the default values.
func NewRestrictedAppsViolation()(*RestrictedAppsViolation) {
    m := &RestrictedAppsViolation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetDeviceConfigurationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationId
    }
}
// Gets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) GetDeviceConfigurationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationName
    }
}
// Gets the deviceName property value. Device name
func (m *RestrictedAppsViolation) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the platformType property value. Platform type. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
func (m *RestrictedAppsViolation) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// Gets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) GetRestrictedApps()([]ManagedDeviceReportedApp) {
    if m == nil {
        return nil
    } else {
        return m.restrictedApps
    }
}
// Gets the restrictedAppsState property value. Restricted apps state. Possible values are: prohibitedApps, notApprovedApps.
func (m *RestrictedAppsViolation) GetRestrictedAppsState()(*RestrictedAppsState) {
    if m == nil {
        return nil
    } else {
        return m.restrictedAppsState
    }
}
// Gets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userName property value. User name
func (m *RestrictedAppsViolation) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// The deserialization information for the current model
func (m *RestrictedAppsViolation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfigurationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationId(val)
        return nil
    }
    res["deviceConfigurationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationName(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        cast := val.(PolicyPlatformType)
        m.SetPlatformType(&cast)
        return nil
    }
    res["restrictedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceReportedApp() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceReportedApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceReportedApp))
        }
        m.SetRestrictedApps(res)
        return nil
    }
    res["restrictedAppsState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRestrictedAppsState)
        if err != nil {
            return err
        }
        cast := val.(RestrictedAppsState)
        m.SetRestrictedAppsState(&cast)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
        return nil
    }
    return res
}
func (m *RestrictedAppsViolation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RestrictedAppsViolation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceConfigurationId", m.GetDeviceConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceConfigurationName", m.GetDeviceConfigurationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := m.GetPlatformType().String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRestrictedApps()))
        for i, v := range m.GetRestrictedApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("restrictedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedAppsState() != nil {
        cast := m.GetRestrictedAppsState().String()
        err = writer.WriteStringValue("restrictedAppsState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
// Parameters:
//  - value : Value to set for the deviceConfigurationId property.
func (m *RestrictedAppsViolation) SetDeviceConfigurationId(value *string)() {
    m.deviceConfigurationId = value
}
// Sets the deviceConfigurationName property value. Device configuration profile name
// Parameters:
//  - value : Value to set for the deviceConfigurationName property.
func (m *RestrictedAppsViolation) SetDeviceConfigurationName(value *string)() {
    m.deviceConfigurationName = value
}
// Sets the deviceName property value. Device name
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *RestrictedAppsViolation) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the managedDeviceId property value. Managed device unique identifier, must be Guid
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *RestrictedAppsViolation) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the platformType property value. Platform type. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
// Parameters:
//  - value : Value to set for the platformType property.
func (m *RestrictedAppsViolation) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// Sets the restrictedApps property value. List of violated restricted apps
// Parameters:
//  - value : Value to set for the restrictedApps property.
func (m *RestrictedAppsViolation) SetRestrictedApps(value []ManagedDeviceReportedApp)() {
    m.restrictedApps = value
}
// Sets the restrictedAppsState property value. Restricted apps state. Possible values are: prohibitedApps, notApprovedApps.
// Parameters:
//  - value : Value to set for the restrictedAppsState property.
func (m *RestrictedAppsViolation) SetRestrictedAppsState(value *RestrictedAppsState)() {
    m.restrictedAppsState = value
}
// Sets the userId property value. User unique identifier, must be Guid
// Parameters:
//  - value : Value to set for the userId property.
func (m *RestrictedAppsViolation) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userName property value. User name
// Parameters:
//  - value : Value to set for the userName property.
func (m *RestrictedAppsViolation) SetUserName(value *string)() {
    m.userName = value
}
