// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapWatchGuard = map[string]DeviceData{
    "Firebox M370": {
        Manufacturer: "WatchGuard",
        Model: "Firebox M370",
        Slug: "watchguard-firebox-m370",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU", Label: "", Type: "iec-60320-c14", MaximumDraw: 220, AllocatedDraw: 220 },
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
            { Name: "eth0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth7", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "Firebox T80": {
        Manufacturer: "WatchGuard",
        Model: "Firebox T80",
        Slug: "watchguard-firebox-t80",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "left-to-right",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 2.6,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "mgmnt", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth7", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "Firebox T85-PoE": {
        Manufacturer: "WatchGuard",
        Model: "Firebox T85-PoE",
        Slug: "watchguard-firebox-t85-poe",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "left-to-right",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 2.6,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "mgmnt", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth7", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "M270": {
        Manufacturer: "WatchGuard",
        Model: "M270",
        Slug: "watchguard-firebox-m270",
        UHeight: 1,
        PartNumber: "M270",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 8.5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
            { Name: "usb1", Type: "usb-a", Label: "", Poe: false },
            { Name: "usb2", Type: "usb-a", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "iec-60320-c14", MaximumDraw: 75, AllocatedDraw: 0 },
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
            { Name: "0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "7", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
