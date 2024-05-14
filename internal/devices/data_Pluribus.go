// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapPluribus = map[string]DeviceData{
    "Freedom E28Q-L": {
        Manufacturer: "Pluribus",
        Model: "Freedom E28Q-L",
        Slug: "pluribus-freedom-e28q-l",
        UHeight: 2,
        PartNumber: "E28Q-L",
        IsFullDepth: true,
        Airflow: "rear-to-front",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 25,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "CON", Poe: false },
            { Name: "Console USB", Type: "usb-a", Label: "USB", Poe: false },
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
            { Name: "PSU-1", Label: "1", Position: "PSU-1" },
            { Name: "PSU-2", Label: "2", Position: "PSU-2" },
            { Name: "PCIe slot 1 - FH", Label: "PCIe 1", Position: "PCIe-1" },
            { Name: "PCIe slot 2 - FH", Label: "PCIe 2", Position: "PCIe-2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "mgmt", Label: "MGMT", Type: "1000base-t", MgmtOnly: true },
            { Name: "1", Label: "1", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "2", Label: "2", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "3", Label: "3", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "4", Label: "4", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "5", Label: "5..8", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "5-10GbE", Label: "5..8", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "6-10GbE", Label: "5..8", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "7-10GbE", Label: "5..8", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "8-10GbE", Label: "5..8", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "9", Label: "9..12", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "9-10GbE", Label: "9..12", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "10-10GbE", Label: "9..12", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "11-10GbE", Label: "9..12", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "12-10GbE", Label: "9..12", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "13", Label: "13..16", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "13-10GbE", Label: "13..16", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "14-10GbE", Label: "13..16", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "15-10GbE", Label: "13..16", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "16-10GbE", Label: "13..16", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "17", Label: "17..20", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "17-10GbE", Label: "17..20", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "18-10GbE", Label: "17..20", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "19-10GbE", Label: "17..20", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "20-10GbE", Label: "17..20", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "21", Label: "21..24", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "21-10GbE", Label: "21..24", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "22-10GbE", Label: "21..24", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "23-10GbE", Label: "21..24", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "24-10GbE", Label: "21..24", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "25", Label: "25..28", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "25-10GbE", Label: "25..28", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "26-10GbE", Label: "25..28", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "27-10GbE", Label: "25..28", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "28-10GbE", Label: "25..28", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "29", Label: "29..32", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "29-10GbE", Label: "29..32", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "30-10GbE", Label: "29..32", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "31-10GbE", Label: "29..32", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "32-10GbE", Label: "29..32", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "33", Label: "33..36", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "33-10GbE", Label: "33..36", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "34-10GbE", Label: "33..36", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "35-10GbE", Label: "33..36", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "36-10GbE", Label: "33..36", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "37", Label: "37..40", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "37-10GbE", Label: "37..40", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "38-10GbE", Label: "37..40", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "39-10GbE", Label: "37..40", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "40-10GbE", Label: "37..40", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "41", Label: "41..44", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "41-10GbE", Label: "41..44", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "42-10GbE", Label: "41..44", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "43-10GbE", Label: "41..44", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "44-10GbE", Label: "41..44", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "45", Label: "45..48", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "45-10GbE", Label: "45..48", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "46-10GbE", Label: "45..48", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "47-10GbE", Label: "45..48", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "48-10GbE", Label: "45..48", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "49", Label: "49..52", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "49-10GbE", Label: "49..52", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "50-10GbE", Label: "49..52", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "51-10GbE", Label: "49..52", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "52-10GbE", Label: "49..52", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "53", Label: "53..56", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "53-10GbE", Label: "53..56", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "54-10GbE", Label: "53..56", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "55-10GbE", Label: "53..56", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "56-10GbE", Label: "53..56", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "57", Label: "57..60", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "57-10GbE", Label: "57..60", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "58-10GbE", Label: "57..60", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "59-10GbE", Label: "57..60", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "60-10GbE", Label: "57..60", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "61", Label: "61..64", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "61-10GbE", Label: "61..64", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "62-10GbE", Label: "61..64", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "63-10GbE", Label: "61..64", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "64-10GbE", Label: "61..64", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "65", Label: "65..68", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "65-10GbE", Label: "65..68", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "66-10GbE", Label: "65..68", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "67-10GbE", Label: "65..68", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "68-10GbE", Label: "65..68", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "69", Label: "69..72", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "69-10GbE", Label: "69..72", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "70-10GbE", Label: "69..72", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "71-10GbE", Label: "69..72", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "72-10GbE", Label: "69..72", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "73", Label: "73..76", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "73-10GbE", Label: "73..76", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "74-10GbE", Label: "73..76", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "75-10GbE", Label: "73..76", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "76-10GbE", Label: "73..76", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "77", Label: "77..80", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "77-10GbE", Label: "77..80", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "78-10GbE", Label: "77..80", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "79-10GbE", Label: "77..80", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "80-10GbE", Label: "77..80", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "81", Label: "81..84", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "81-10GbE", Label: "81..84", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "82-10GbE", Label: "81..84", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "83-10GbE", Label: "81..84", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "84-10GbE", Label: "81..84", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "85", Label: "85", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "86", Label: "86", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "87", Label: "87", Type: "40gbase-x-qsfpp", MgmtOnly: false },
            { Name: "88", Label: "88", Type: "40gbase-x-qsfpp", MgmtOnly: false },
        },
    },
}
