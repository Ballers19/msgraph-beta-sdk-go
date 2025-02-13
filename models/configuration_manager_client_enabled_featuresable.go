package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerClientEnabledFeaturesable 
type ConfigurationManagerClientEnabledFeaturesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliancePolicy()(*bool)
    GetDeviceConfiguration()(*bool)
    GetEndpointProtection()(*bool)
    GetInventory()(*bool)
    GetModernApps()(*bool)
    GetOfficeApps()(*bool)
    GetResourceAccess()(*bool)
    GetWindowsUpdateForBusiness()(*bool)
    SetCompliancePolicy(value *bool)()
    SetDeviceConfiguration(value *bool)()
    SetEndpointProtection(value *bool)()
    SetInventory(value *bool)()
    SetModernApps(value *bool)()
    SetOfficeApps(value *bool)()
    SetResourceAccess(value *bool)()
    SetWindowsUpdateForBusiness(value *bool)()
}
