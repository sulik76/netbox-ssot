// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapNetApp = map[string]DeviceData{
    "E2824": {
        Manufacturer: "NetApp",
        Model: "E2824",
        Slug: "netapp-e2824",
        UHeight: 2,
        PartNumber: "E2824",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 27.44,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "con0", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU1", Label: "", Position: "1" },
            { Name: "PSU2", Label: "", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "e0a", Label: "", Type: "16gfc-sfpp", MgmtOnly: false },
            { Name: "e0b", Label: "", Type: "16gfc-sfpp", MgmtOnly: false },
            { Name: "e0c", Label: "", Type: "16gfc-sfpp", MgmtOnly: false },
            { Name: "e0d", Label: "", Type: "16gfc-sfpp", MgmtOnly: false },
        },
    },
    "FAS2700 Controller 10GBASE-T": {
        Manufacturer: "NetApp",
        Model: "FAS2700 Controller 10GBASE-T",
        Slug: "netapp-fas2700-controller-10gbase-t",
        UHeight: 0,
        PartNumber: "X3219A",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "con0", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU3", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU4", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "e0a", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0b", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0c", Label: "", Type: "10gbase-t", MgmtOnly: false },
            { Name: "e0d", Label: "", Type: "10gbase-t", MgmtOnly: false },
            { Name: "e0e", Label: "", Type: "10gbase-t", MgmtOnly: false },
            { Name: "e0f", Label: "", Type: "10gbase-t", MgmtOnly: false },
            { Name: "e0m", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "FAS2700 Controller CNA": {
        Manufacturer: "NetApp",
        Model: "FAS2700 Controller CNA",
        Slug: "netapp-fas2700-controller-cna",
        UHeight: 0,
        PartNumber: "X3218A",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "con0", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "e0a", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0b", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0c", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0d", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0e", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0f", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0m", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
    "FAS2750": {
        Manufacturer: "NetApp",
        Model: "FAS2750",
        Slug: "netapp-fas2750",
        UHeight: 2,
        PartNumber: "FAS2750",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "con0", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "A", Label: "", Position: "" },
            { Name: "B", Label: "", Position: "" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "FAS8200": {
        Manufacturer: "NetApp",
        Model: "FAS8200",
        Slug: "netapp-fas8200",
        UHeight: 3,
        PartNumber: "FAS8200A",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "A", Label: "", Position: "" },
            { Name: "B", Label: "", Position: "" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "FAS8200 Controller": {
        Manufacturer: "NetApp",
        Model: "FAS8200 Controller",
        Slug: "netapp-fas8200-controller",
        UHeight: 0,
        PartNumber: "X3172A",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "con0", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "e0a", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0b", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0c", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "e0d", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "e0e", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0f", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0g", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0h", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "e0m", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
