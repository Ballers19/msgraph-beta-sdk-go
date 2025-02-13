package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagement 
type DeviceAppManagement struct {
    Entity
    // Android managed app policies.
    androidManagedAppProtections []AndroidManagedAppProtectionable
    // Default managed app policies.
    defaultManagedAppProtections []DefaultManagedAppProtectionable
    // Device app management tasks.
    deviceAppManagementTasks []DeviceAppManagementTaskable
    // The Windows Enterprise Code Signing Certificate.
    enterpriseCodeSigningCertificates []EnterpriseCodeSigningCertificateable
    // The IOS Lob App Provisioning Configurations.
    iosLobAppProvisioningConfigurations []IosLobAppProvisioningConfigurationable
    // iOS managed app policies.
    iosManagedAppProtections []IosManagedAppProtectionable
    // Whether the account is enabled for syncing applications from the Microsoft Store for Business.
    isEnabledForMicrosoftStoreForBusiness *bool
    // Managed app policies.
    managedAppPolicies []ManagedAppPolicyable
    // The managed app registrations.
    managedAppRegistrations []ManagedAppRegistrationable
    // The managed app statuses.
    managedAppStatuses []ManagedAppStatusable
    // The mobile eBook categories.
    managedEBookCategories []ManagedEBookCategoryable
    // The Managed eBook.
    managedEBooks []ManagedEBookable
    // Windows information protection for apps running on devices which are MDM enrolled.
    mdmWindowsInformationProtectionPolicies []MdmWindowsInformationProtectionPolicyable
    // The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
    microsoftStoreForBusinessLanguage *string
    // The last time an application sync from the Microsoft Store for Business was completed.
    microsoftStoreForBusinessLastCompletedApplicationSyncTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time the apps from the Microsoft Store for Business were synced successfully for the account.
    microsoftStoreForBusinessLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
    microsoftStoreForBusinessPortalSelection *MicrosoftStoreForBusinessPortalSelectionOptions
    // The mobile app categories.
    mobileAppCategories []MobileAppCategoryable
    // The Managed Device Mobile Application Configurations.
    mobileAppConfigurations []ManagedDeviceMobileAppConfigurationable
    // The mobile apps.
    mobileApps []MobileAppable
    // The PolicySet of Policies and Applications
    policySets []PolicySetable
    // Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
    sideLoadingKeys []SideLoadingKeyable
    // The WinPhone Symantec Code Signing Certificate.
    symantecCodeSigningCertificate SymantecCodeSigningCertificateable
    // Targeted managed app configurations.
    targetedManagedAppConfigurations []TargetedManagedAppConfigurationable
    // List of Vpp tokens for this organization.
    vppTokens []VppTokenable
    // The collection of Windows Defender Application Control Supplemental Policies.
    wdacSupplementalPolicies []WindowsDefenderApplicationControlSupplementalPolicyable
    // Windows information protection device registrations that are not MDM enrolled.
    windowsInformationProtectionDeviceRegistrations []WindowsInformationProtectionDeviceRegistrationable
    // Windows information protection for apps running on devices which are not MDM enrolled.
    windowsInformationProtectionPolicies []WindowsInformationProtectionPolicyable
    // Windows information protection wipe actions.
    windowsInformationProtectionWipeActions []WindowsInformationProtectionWipeActionable
    // Windows managed app policies.
    windowsManagedAppProtections []WindowsManagedAppProtectionable
    // Windows management app.
    windowsManagementApp WindowsManagementAppable
}
// NewDeviceAppManagement instantiates a new DeviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceAppManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagement(), nil
}
// GetAndroidManagedAppProtections gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedAppProtections
    }
}
// GetDefaultManagedAppProtections gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable) {
    if m == nil {
        return nil
    } else {
        return m.defaultManagedAppProtections
    }
}
// GetDeviceAppManagementTasks gets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) GetDeviceAppManagementTasks()([]DeviceAppManagementTaskable) {
    if m == nil {
        return nil
    } else {
        return m.deviceAppManagementTasks
    }
}
// GetEnterpriseCodeSigningCertificates gets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) GetEnterpriseCodeSigningCertificates()([]EnterpriseCodeSigningCertificateable) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseCodeSigningCertificates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidManagedAppProtectionable)
            }
            m.SetAndroidManagedAppProtections(res)
        }
        return nil
    }
    res["defaultManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDefaultManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(DefaultManagedAppProtectionable)
            }
            m.SetDefaultManagedAppProtections(res)
        }
        return nil
    }
    res["deviceAppManagementTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAppManagementTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAppManagementTaskable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceAppManagementTaskable)
            }
            m.SetDeviceAppManagementTasks(res)
        }
        return nil
    }
    res["enterpriseCodeSigningCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnterpriseCodeSigningCertificateable, len(val))
            for i, v := range val {
                res[i] = v.(EnterpriseCodeSigningCertificateable)
            }
            m.SetEnterpriseCodeSigningCertificates(res)
        }
        return nil
    }
    res["iosLobAppProvisioningConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosLobAppProvisioningConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(IosLobAppProvisioningConfigurationable)
            }
            m.SetIosLobAppProvisioningConfigurations(res)
        }
        return nil
    }
    res["iosManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(IosManagedAppProtectionable)
            }
            m.SetIosManagedAppProtections(res)
        }
        return nil
    }
    res["isEnabledForMicrosoftStoreForBusiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForMicrosoftStoreForBusiness(val)
        }
        return nil
    }
    res["managedAppPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppPolicyable)
            }
            m.SetManagedAppPolicies(res)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistrationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppRegistrationable)
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedAppStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppStatusable)
            }
            m.SetManagedAppStatuses(res)
        }
        return nil
    }
    res["managedEBookCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedEBookCategoryable)
            }
            m.SetManagedEBookCategories(res)
        }
        return nil
    }
    res["managedEBooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedEBookable)
            }
            m.SetManagedEBooks(res)
        }
        return nil
    }
    res["mdmWindowsInformationProtectionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MdmWindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(MdmWindowsInformationProtectionPolicyable)
            }
            m.SetMdmWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["microsoftStoreForBusinessLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLanguage(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessPortalSelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftStoreForBusinessPortalSelectionOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessPortalSelection(val.(*MicrosoftStoreForBusinessPortalSelectionOptions))
        }
        return nil
    }
    res["mobileAppCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppCategoryable)
            }
            m.SetMobileAppCategories(res)
        }
        return nil
    }
    res["mobileAppConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationable)
            }
            m.SetMobileAppConfigurations(res)
        }
        return nil
    }
    res["mobileApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppable)
            }
            m.SetMobileApps(res)
        }
        return nil
    }
    res["policySets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetable, len(val))
            for i, v := range val {
                res[i] = v.(PolicySetable)
            }
            m.SetPolicySets(res)
        }
        return nil
    }
    res["sideLoadingKeys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSideLoadingKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SideLoadingKeyable, len(val))
            for i, v := range val {
                res[i] = v.(SideLoadingKeyable)
            }
            m.SetSideLoadingKeys(res)
        }
        return nil
    }
    res["symantecCodeSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSymantecCodeSigningCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSymantecCodeSigningCertificate(val.(SymantecCodeSigningCertificateable))
        }
        return nil
    }
    res["targetedManagedAppConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(TargetedManagedAppConfigurationable)
            }
            m.SetTargetedManagedAppConfigurations(res)
        }
        return nil
    }
    res["vppTokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVppTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VppTokenable, len(val))
            for i, v := range val {
                res[i] = v.(VppTokenable)
            }
            m.SetVppTokens(res)
        }
        return nil
    }
    res["wdacSupplementalPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDefenderApplicationControlSupplementalPolicyable)
            }
            m.SetWdacSupplementalPolicies(res)
        }
        return nil
    }
    res["windowsInformationProtectionDeviceRegistrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionDeviceRegistrationable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionDeviceRegistrationable)
            }
            m.SetWindowsInformationProtectionDeviceRegistrations(res)
        }
        return nil
    }
    res["windowsInformationProtectionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionPolicyable)
            }
            m.SetWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["windowsInformationProtectionWipeActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionWipeActionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionWipeActionable)
            }
            m.SetWindowsInformationProtectionWipeActions(res)
        }
        return nil
    }
    res["windowsManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsManagedAppProtectionable)
            }
            m.SetWindowsManagedAppProtections(res)
        }
        return nil
    }
    res["windowsManagementApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsManagementAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsManagementApp(val.(WindowsManagementAppable))
        }
        return nil
    }
    return res
}
// GetIosLobAppProvisioningConfigurations gets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) GetIosLobAppProvisioningConfigurations()([]IosLobAppProvisioningConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.iosLobAppProvisioningConfigurations
    }
}
// GetIosManagedAppProtections gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtectionable) {
    if m == nil {
        return nil
    } else {
        return m.iosManagedAppProtections
    }
}
// GetIsEnabledForMicrosoftStoreForBusiness gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledForMicrosoftStoreForBusiness
    }
}
// GetManagedAppPolicies gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.managedAppPolicies
    }
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistrationable) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
// GetManagedAppStatuses gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatusable) {
    if m == nil {
        return nil
    } else {
        return m.managedAppStatuses
    }
}
// GetManagedEBookCategories gets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) GetManagedEBookCategories()([]ManagedEBookCategoryable) {
    if m == nil {
        return nil
    } else {
        return m.managedEBookCategories
    }
}
// GetManagedEBooks gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBookable) {
    if m == nil {
        return nil
    } else {
        return m.managedEBooks
    }
}
// GetMdmWindowsInformationProtectionPolicies gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.mdmWindowsInformationProtectionPolicies
    }
}
// GetMicrosoftStoreForBusinessLanguage gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLanguage
    }
}
// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
    }
}
// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
    }
}
// GetMicrosoftStoreForBusinessPortalSelection gets the microsoftStoreForBusinessPortalSelection property value. Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessPortalSelection()(*MicrosoftStoreForBusinessPortalSelectionOptions) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessPortalSelection
    }
}
// GetMobileAppCategories gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategoryable) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppCategories
    }
}
// GetMobileAppConfigurations gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppConfigurations
    }
}
// GetMobileApps gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileAppable) {
    if m == nil {
        return nil
    } else {
        return m.mobileApps
    }
}
// GetPolicySets gets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) GetPolicySets()([]PolicySetable) {
    if m == nil {
        return nil
    } else {
        return m.policySets
    }
}
// GetSideLoadingKeys gets the sideLoadingKeys property value. Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
func (m *DeviceAppManagement) GetSideLoadingKeys()([]SideLoadingKeyable) {
    if m == nil {
        return nil
    } else {
        return m.sideLoadingKeys
    }
}
// GetSymantecCodeSigningCertificate gets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) GetSymantecCodeSigningCertificate()(SymantecCodeSigningCertificateable) {
    if m == nil {
        return nil
    } else {
        return m.symantecCodeSigningCertificate
    }
}
// GetTargetedManagedAppConfigurations gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.targetedManagedAppConfigurations
    }
}
// GetVppTokens gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppTokenable) {
    if m == nil {
        return nil
    } else {
        return m.vppTokens
    }
}
// GetWdacSupplementalPolicies gets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) GetWdacSupplementalPolicies()([]WindowsDefenderApplicationControlSupplementalPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.wdacSupplementalPolicies
    }
}
// GetWindowsInformationProtectionDeviceRegistrations gets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionDeviceRegistrations
    }
}
// GetWindowsInformationProtectionPolicies gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionPolicies
    }
}
// GetWindowsInformationProtectionWipeActions gets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) GetWindowsInformationProtectionWipeActions()([]WindowsInformationProtectionWipeActionable) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionWipeActions
    }
}
// GetWindowsManagedAppProtections gets the windowsManagedAppProtections property value. Windows managed app policies.
func (m *DeviceAppManagement) GetWindowsManagedAppProtections()([]WindowsManagedAppProtectionable) {
    if m == nil {
        return nil
    } else {
        return m.windowsManagedAppProtections
    }
}
// GetWindowsManagementApp gets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) GetWindowsManagementApp()(WindowsManagementAppable) {
    if m == nil {
        return nil
    } else {
        return m.windowsManagementApp
    }
}
// Serialize serializes information the current object
func (m *DeviceAppManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidManagedAppProtections()))
        for i, v := range m.GetAndroidManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultManagedAppProtections()))
        for i, v := range m.GetDefaultManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("defaultManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceAppManagementTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceAppManagementTasks()))
        for i, v := range m.GetDeviceAppManagementTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceAppManagementTasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseCodeSigningCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnterpriseCodeSigningCertificates()))
        for i, v := range m.GetEnterpriseCodeSigningCertificates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseCodeSigningCertificates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosLobAppProvisioningConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosLobAppProvisioningConfigurations()))
        for i, v := range m.GetIosLobAppProvisioningConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("iosLobAppProvisioningConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosManagedAppProtections()))
        for i, v := range m.GetIosManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("iosManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForMicrosoftStoreForBusiness", m.GetIsEnabledForMicrosoftStoreForBusiness())
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppPolicies()))
        for i, v := range m.GetManagedAppPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppStatuses()))
        for i, v := range m.GetManagedAppStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBookCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBookCategories()))
        for i, v := range m.GetManagedEBookCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedEBookCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBooks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBooks()))
        for i, v := range m.GetManagedEBooks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedEBooks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMdmWindowsInformationProtectionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMdmWindowsInformationProtectionPolicies()))
        for i, v := range m.GetMdmWindowsInformationProtectionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mdmWindowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftStoreForBusinessLanguage", m.GetMicrosoftStoreForBusinessLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastCompletedApplicationSyncTime", m.GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastSuccessfulSyncDateTime", m.GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftStoreForBusinessPortalSelection() != nil {
        cast := (*m.GetMicrosoftStoreForBusinessPortalSelection()).String()
        err = writer.WriteStringValue("microsoftStoreForBusinessPortalSelection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppCategories()))
        for i, v := range m.GetMobileAppCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppConfigurations()))
        for i, v := range m.GetMobileAppConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileApps()))
        for i, v := range m.GetMobileApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicySets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicySets()))
        for i, v := range m.GetPolicySets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("policySets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSideLoadingKeys() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSideLoadingKeys()))
        for i, v := range m.GetSideLoadingKeys() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sideLoadingKeys", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("symantecCodeSigningCertificate", m.GetSymantecCodeSigningCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedManagedAppConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedManagedAppConfigurations()))
        for i, v := range m.GetTargetedManagedAppConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targetedManagedAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVppTokens() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVppTokens()))
        for i, v := range m.GetVppTokens() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("vppTokens", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWdacSupplementalPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWdacSupplementalPolicies()))
        for i, v := range m.GetWdacSupplementalPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("wdacSupplementalPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionDeviceRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionDeviceRegistrations()))
        for i, v := range m.GetWindowsInformationProtectionDeviceRegistrations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionDeviceRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionPolicies()))
        for i, v := range m.GetWindowsInformationProtectionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionWipeActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionWipeActions()))
        for i, v := range m.GetWindowsInformationProtectionWipeActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionWipeActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsManagedAppProtections()))
        for i, v := range m.GetWindowsManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsManagementApp", m.GetWindowsManagementApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidManagedAppProtections sets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtectionable)() {
    if m != nil {
        m.androidManagedAppProtections = value
    }
}
// SetDefaultManagedAppProtections sets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)() {
    if m != nil {
        m.defaultManagedAppProtections = value
    }
}
// SetDeviceAppManagementTasks sets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) SetDeviceAppManagementTasks(value []DeviceAppManagementTaskable)() {
    if m != nil {
        m.deviceAppManagementTasks = value
    }
}
// SetEnterpriseCodeSigningCertificates sets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) SetEnterpriseCodeSigningCertificates(value []EnterpriseCodeSigningCertificateable)() {
    if m != nil {
        m.enterpriseCodeSigningCertificates = value
    }
}
// SetIosLobAppProvisioningConfigurations sets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) SetIosLobAppProvisioningConfigurations(value []IosLobAppProvisioningConfigurationable)() {
    if m != nil {
        m.iosLobAppProvisioningConfigurations = value
    }
}
// SetIosManagedAppProtections sets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtectionable)() {
    if m != nil {
        m.iosManagedAppProtections = value
    }
}
// SetIsEnabledForMicrosoftStoreForBusiness sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    if m != nil {
        m.isEnabledForMicrosoftStoreForBusiness = value
    }
}
// SetManagedAppPolicies sets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicyable)() {
    if m != nil {
        m.managedAppPolicies = value
    }
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistrationable)() {
    if m != nil {
        m.managedAppRegistrations = value
    }
}
// SetManagedAppStatuses sets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatusable)() {
    if m != nil {
        m.managedAppStatuses = value
    }
}
// SetManagedEBookCategories sets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) SetManagedEBookCategories(value []ManagedEBookCategoryable)() {
    if m != nil {
        m.managedEBookCategories = value
    }
}
// SetManagedEBooks sets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBookable)() {
    if m != nil {
        m.managedEBooks = value
    }
}
// SetMdmWindowsInformationProtectionPolicies sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)() {
    if m != nil {
        m.mdmWindowsInformationProtectionPolicies = value
    }
}
// SetMicrosoftStoreForBusinessLanguage sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    if m != nil {
        m.microsoftStoreForBusinessLanguage = value
    }
}
// SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
    }
}
// SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
    }
}
// SetMicrosoftStoreForBusinessPortalSelection sets the microsoftStoreForBusinessPortalSelection property value. Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessPortalSelection(value *MicrosoftStoreForBusinessPortalSelectionOptions)() {
    if m != nil {
        m.microsoftStoreForBusinessPortalSelection = value
    }
}
// SetMobileAppCategories sets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategoryable)() {
    if m != nil {
        m.mobileAppCategories = value
    }
}
// SetMobileAppConfigurations sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)() {
    if m != nil {
        m.mobileAppConfigurations = value
    }
}
// SetMobileApps sets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) SetMobileApps(value []MobileAppable)() {
    if m != nil {
        m.mobileApps = value
    }
}
// SetPolicySets sets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) SetPolicySets(value []PolicySetable)() {
    if m != nil {
        m.policySets = value
    }
}
// SetSideLoadingKeys sets the sideLoadingKeys property value. Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
func (m *DeviceAppManagement) SetSideLoadingKeys(value []SideLoadingKeyable)() {
    if m != nil {
        m.sideLoadingKeys = value
    }
}
// SetSymantecCodeSigningCertificate sets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) SetSymantecCodeSigningCertificate(value SymantecCodeSigningCertificateable)() {
    if m != nil {
        m.symantecCodeSigningCertificate = value
    }
}
// SetTargetedManagedAppConfigurations sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)() {
    if m != nil {
        m.targetedManagedAppConfigurations = value
    }
}
// SetVppTokens sets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) SetVppTokens(value []VppTokenable)() {
    if m != nil {
        m.vppTokens = value
    }
}
// SetWdacSupplementalPolicies sets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) SetWdacSupplementalPolicies(value []WindowsDefenderApplicationControlSupplementalPolicyable)() {
    if m != nil {
        m.wdacSupplementalPolicies = value
    }
}
// SetWindowsInformationProtectionDeviceRegistrations sets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)() {
    if m != nil {
        m.windowsInformationProtectionDeviceRegistrations = value
    }
}
// SetWindowsInformationProtectionPolicies sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)() {
    if m != nil {
        m.windowsInformationProtectionPolicies = value
    }
}
// SetWindowsInformationProtectionWipeActions sets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) SetWindowsInformationProtectionWipeActions(value []WindowsInformationProtectionWipeActionable)() {
    if m != nil {
        m.windowsInformationProtectionWipeActions = value
    }
}
// SetWindowsManagedAppProtections sets the windowsManagedAppProtections property value. Windows managed app policies.
func (m *DeviceAppManagement) SetWindowsManagedAppProtections(value []WindowsManagedAppProtectionable)() {
    if m != nil {
        m.windowsManagedAppProtections = value
    }
}
// SetWindowsManagementApp sets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) SetWindowsManagementApp(value WindowsManagementAppable)() {
    if m != nil {
        m.windowsManagementApp = value
    }
}
