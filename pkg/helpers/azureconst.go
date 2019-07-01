// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

// AUTOGENERATED FILE

// GetAzureLocations provides all azure regions in prod.
// Related powershell to refresh this list:
//   Get-AzureRmLocation | Select-Object -Property Location
func GetAzureLocations() []string {
	return []string{
		"australiacentral",
		"australiacentral2",
		"australiaeast",
		"australiasoutheast",
		"brazilsouth",
		"canadacentral",
		"canadaeast",
		"centralindia",
		"centralus",
		"centraluseuap",
		"chinaeast",
		"chinaeast2",
		"chinanorth",
		"chinanorth2",
		"eastasia",
		"eastus",
		"eastus2",
		"eastus2euap",
		"francecentral",
		"francesouth",
		"japaneast",
		"japanwest",
		"koreacentral",
		"koreasouth",
		"northcentralus",
		"northeurope",
		"southafricanorth",
		"southafricawest",
		"southcentralus",
		"southeastasia",
		"southindia",
		"uksouth",
		"ukwest",
		"usdodcentral",
		"usdodeast",
		"westcentralus",
		"westeurope",
		"westindia",
		"westus",
		"westus2",
		"chinaeast",
		"chinanorth",
		"chinanorth2",
		"chinaeast2",
		"germanycentral",
		"germanynortheast",
		"usgovvirginia",
		"usgoviowa",
		"usgovarizona",
		"usgovtexas",
		"francecentral",
		"uaenorth",
		"uaecentral",
	}
}

// GetDCOSMasterAllowedSizes returns the master allowed sizes
func GetDCOSMasterAllowedSizes() string {
	return `      "allowedValues": [
        "Standard_A10",
        "Standard_A11",
        "Standard_A2",
        "Standard_A2_v2",
        "Standard_A2m_v2",
        "Standard_A3",
        "Standard_A4",
        "Standard_A4_v2",
        "Standard_A4m_v2",
        "Standard_A5",
        "Standard_A6",
        "Standard_A7",
        "Standard_A8",
        "Standard_A8_v2",
        "Standard_A8m_v2",
        "Standard_A9",
        "Standard_B2ms",
        "Standard_B4ms",
        "Standard_B8ms",
        "Standard_D11",
        "Standard_D11_v2",
        "Standard_D12",
        "Standard_D12_v2",
        "Standard_D13",
        "Standard_D13_v2",
        "Standard_D14",
        "Standard_D14_v2",
        "Standard_D15_v2",
        "Standard_D16_v3",
        "Standard_D16s_v3",
        "Standard_D2",
        "Standard_D2_v2",
        "Standard_D2_v3",
        "Standard_D2s_v3",
        "Standard_D3",
        "Standard_D32_v3",
        "Standard_D32s_v3",
        "Standard_D3_v2",
        "Standard_D4",
        "Standard_D4_v2",
        "Standard_D4_v3",
        "Standard_D4s_v3",
        "Standard_D5_v2",
        "Standard_D64_v3",
        "Standard_D64s_v3",
        "Standard_D8_v3",
        "Standard_D8s_v3",
        "Standard_DC2s",
        "Standard_DC4s",
        "Standard_DS11",
        "Standard_DS11-1_v2",
        "Standard_DS11_v2",
        "Standard_DS12",
        "Standard_DS12-1_v2",
        "Standard_DS12-2_v2",
        "Standard_DS12_v2",
        "Standard_DS13",
        "Standard_DS13-2_v2",
        "Standard_DS13-4_v2",
        "Standard_DS13_v2",
        "Standard_DS14",
        "Standard_DS14-4_v2",
        "Standard_DS14-8_v2",
        "Standard_DS14_v2",
        "Standard_DS15_v2",
        "Standard_DS3",
        "Standard_DS3_v2",
        "Standard_DS4",
        "Standard_DS4_v2",
        "Standard_DS5_v2",
        "Standard_E16-4s_v3",
        "Standard_E16-8s_v3",
        "Standard_E16_v3",
        "Standard_E16s_v3",
        "Standard_E20_v3",
        "Standard_E20s_v3",
        "Standard_E2_v3",
        "Standard_E2s_v3",
        "Standard_E32-16s_v3",
        "Standard_E32-8s_v3",
        "Standard_E32_v3",
        "Standard_E32s_v3",
        "Standard_E4-2s_v3",
        "Standard_E4_v3",
        "Standard_E4s_v3",
        "Standard_E64-16s_v3",
        "Standard_E64-32s_v3",
        "Standard_E64_v3",
        "Standard_E64i_v3",
        "Standard_E64is_v3",
        "Standard_E64s_v3",
        "Standard_E8-2s_v3",
        "Standard_E8-4s_v3",
        "Standard_E8_v3",
        "Standard_E8s_v3",
        "Standard_F16",
        "Standard_F16s",
        "Standard_F16s_v2",
        "Standard_F2",
        "Standard_F2s_v2",
        "Standard_F32s_v2",
        "Standard_F4",
        "Standard_F4s",
        "Standard_F4s_v2",
        "Standard_F64s_v2",
        "Standard_F72s_v2",
        "Standard_F8",
        "Standard_F8s",
        "Standard_F8s_v2",
        "Standard_G1",
        "Standard_G2",
        "Standard_G3",
        "Standard_G4",
        "Standard_G5",
        "Standard_GS1",
        "Standard_GS2",
        "Standard_GS3",
        "Standard_GS4",
        "Standard_GS4-4",
        "Standard_GS4-8",
        "Standard_GS5",
        "Standard_GS5-16",
        "Standard_GS5-8",
        "Standard_H16",
        "Standard_H16m",
        "Standard_H16mr",
        "Standard_H16r",
        "Standard_H8",
        "Standard_H8m",
        "Standard_HB60rs",
        "Standard_HC44rs",
        "Standard_L16s",
        "Standard_L16s_v2",
        "Standard_L32s",
        "Standard_L32s_v2",
        "Standard_L4s",
        "Standard_L64s_v2",
        "Standard_L80s_v2",
        "Standard_L8s",
        "Standard_L8s_v2",
        "Standard_M128",
        "Standard_M128-32ms",
        "Standard_M128-64ms",
        "Standard_M128m",
        "Standard_M128ms",
        "Standard_M128s",
        "Standard_M16-4ms",
        "Standard_M16-8ms",
        "Standard_M16ms",
        "Standard_M32-16ms",
        "Standard_M32-8ms",
        "Standard_M32ls",
        "Standard_M32ms",
        "Standard_M32ts",
        "Standard_M64",
        "Standard_M64-16ms",
        "Standard_M64-32ms",
        "Standard_M64ls",
        "Standard_M64m",
        "Standard_M64ms",
        "Standard_M64s",
        "Standard_M8-2ms",
        "Standard_M8-4ms",
        "Standard_M8ms",
        "Standard_NC12",
        "Standard_NC12s_v2",
        "Standard_NC12s_v3",
        "Standard_NC24",
        "Standard_NC24r",
        "Standard_NC24rs_v2",
        "Standard_NC24rs_v3",
        "Standard_NC24s_v2",
        "Standard_NC24s_v3",
        "Standard_NC6",
        "Standard_NC6s_v2",
        "Standard_NC6s_v3",
        "Standard_ND12s",
        "Standard_ND24rs",
        "Standard_ND24s",
        "Standard_ND6s",
        "Standard_NV12",
        "Standard_NV12s_v2",
        "Standard_NV12s_v3",
        "Standard_NV24",
        "Standard_NV24s_v2",
        "Standard_NV24s_v3",
        "Standard_NV48s_v3",
        "Standard_NV6",
        "Standard_NV6s_v2",
        "Standard_PB12s",
        "Standard_PB24s",
        "Standard_PB6s"
    ],
`
}

// GetKubernetesAllowedVMSKUs returns the allowed sizes for Kubernetes agent
func GetKubernetesAllowedVMSKUs() string {
	return `      "allowedValues": [
        "Standard_A0",
        "Standard_A1",
        "Standard_A10",
        "Standard_A11",
        "Standard_A1_v2",
        "Standard_A2",
        "Standard_A2_v2",
        "Standard_A2m_v2",
        "Standard_A3",
        "Standard_A4",
        "Standard_A4_v2",
        "Standard_A4m_v2",
        "Standard_A5",
        "Standard_A6",
        "Standard_A7",
        "Standard_A8",
        "Standard_A8_v2",
        "Standard_A8m_v2",
        "Standard_A9",
        "Standard_B1ls",
        "Standard_B1ms",
        "Standard_B1s",
        "Standard_B2ms",
        "Standard_B2s",
        "Standard_B4ms",
        "Standard_B8ms",
        "Standard_D1",
        "Standard_D11",
        "Standard_D11_v2",
        "Standard_D12",
        "Standard_D12_v2",
        "Standard_D13",
        "Standard_D13_v2",
        "Standard_D14",
        "Standard_D14_v2",
        "Standard_D15_v2",
        "Standard_D16_v3",
        "Standard_D16s_v3",
        "Standard_D1_v2",
        "Standard_D2",
        "Standard_D2_v2",
        "Standard_D2_v3",
        "Standard_D2s_v3",
        "Standard_D3",
        "Standard_D32_v3",
        "Standard_D32s_v3",
        "Standard_D3_v2",
        "Standard_D4",
        "Standard_D4_v2",
        "Standard_D4_v3",
        "Standard_D4s_v3",
        "Standard_D5_v2",
        "Standard_D64_v3",
        "Standard_D64s_v3",
        "Standard_D8_v3",
        "Standard_D8s_v3",
        "Standard_DC2s",
        "Standard_DC4s",
        "Standard_DS1",
        "Standard_DS11",
        "Standard_DS11-1_v2",
        "Standard_DS11_v2",
        "Standard_DS12",
        "Standard_DS12-1_v2",
        "Standard_DS12-2_v2",
        "Standard_DS12_v2",
        "Standard_DS13",
        "Standard_DS13-2_v2",
        "Standard_DS13-4_v2",
        "Standard_DS13_v2",
        "Standard_DS14",
        "Standard_DS14-4_v2",
        "Standard_DS14-8_v2",
        "Standard_DS14_v2",
        "Standard_DS15_v2",
        "Standard_DS1_v2",
        "Standard_DS2",
        "Standard_DS2_v2",
        "Standard_DS3",
        "Standard_DS3_v2",
        "Standard_DS4",
        "Standard_DS4_v2",
        "Standard_DS5_v2",
        "Standard_E16-4s_v3",
        "Standard_E16-8s_v3",
        "Standard_E16_v3",
        "Standard_E16s_v3",
        "Standard_E20_v3",
        "Standard_E20s_v3",
        "Standard_E2_v3",
        "Standard_E2s_v3",
        "Standard_E32-16s_v3",
        "Standard_E32-8s_v3",
        "Standard_E32_v3",
        "Standard_E32s_v3",
        "Standard_E4-2s_v3",
        "Standard_E4_v3",
        "Standard_E4s_v3",
        "Standard_E64-16s_v3",
        "Standard_E64-32s_v3",
        "Standard_E64_v3",
        "Standard_E64i_v3",
        "Standard_E64is_v3",
        "Standard_E64s_v3",
        "Standard_E8-2s_v3",
        "Standard_E8-4s_v3",
        "Standard_E8_v3",
        "Standard_E8s_v3",
        "Standard_F1",
        "Standard_F16",
        "Standard_F16s",
        "Standard_F16s_v2",
        "Standard_F1s",
        "Standard_F2",
        "Standard_F2s",
        "Standard_F2s_v2",
        "Standard_F32s_v2",
        "Standard_F4",
        "Standard_F4s",
        "Standard_F4s_v2",
        "Standard_F64s_v2",
        "Standard_F72s_v2",
        "Standard_F8",
        "Standard_F8s",
        "Standard_F8s_v2",
        "Standard_G1",
        "Standard_G2",
        "Standard_G3",
        "Standard_G4",
        "Standard_G5",
        "Standard_GS1",
        "Standard_GS2",
        "Standard_GS3",
        "Standard_GS4",
        "Standard_GS4-4",
        "Standard_GS4-8",
        "Standard_GS5",
        "Standard_GS5-16",
        "Standard_GS5-8",
        "Standard_H16",
        "Standard_H16m",
        "Standard_H16mr",
        "Standard_H16r",
        "Standard_H8",
        "Standard_H8m",
        "Standard_HB60rs",
        "Standard_HC44rs",
        "Standard_L16s",
        "Standard_L16s_v2",
        "Standard_L32s",
        "Standard_L32s_v2",
        "Standard_L4s",
        "Standard_L64s_v2",
        "Standard_L80s_v2",
        "Standard_L8s",
        "Standard_L8s_v2",
        "Standard_M128",
        "Standard_M128-32ms",
        "Standard_M128-64ms",
        "Standard_M128m",
        "Standard_M128ms",
        "Standard_M128s",
        "Standard_M16-4ms",
        "Standard_M16-8ms",
        "Standard_M16ms",
        "Standard_M32-16ms",
        "Standard_M32-8ms",
        "Standard_M32ls",
        "Standard_M32ms",
        "Standard_M32ts",
        "Standard_M64",
        "Standard_M64-16ms",
        "Standard_M64-32ms",
        "Standard_M64ls",
        "Standard_M64m",
        "Standard_M64ms",
        "Standard_M64s",
        "Standard_M8-2ms",
        "Standard_M8-4ms",
        "Standard_M8ms",
        "Standard_NC12",
        "Standard_NC12s_v2",
        "Standard_NC12s_v3",
        "Standard_NC24",
        "Standard_NC24r",
        "Standard_NC24rs_v2",
        "Standard_NC24rs_v3",
        "Standard_NC24s_v2",
        "Standard_NC24s_v3",
        "Standard_NC6",
        "Standard_NC6s_v2",
        "Standard_NC6s_v3",
        "Standard_ND12s",
        "Standard_ND24rs",
        "Standard_ND24s",
        "Standard_ND6s",
        "Standard_NV12",
        "Standard_NV12s_v2",
        "Standard_NV12s_v3",
        "Standard_NV24",
        "Standard_NV24s_v2",
        "Standard_NV24s_v3",
        "Standard_NV48s_v3",
        "Standard_NV6",
        "Standard_NV6s_v2",
        "Standard_PB12s",
        "Standard_PB24s",
        "Standard_PB6s"
    ],
`
}

// GetSizeMap returns the size / storage map
func GetSizeMap() string {
	return `    "vmSizesMap": {
    "Standard_A0": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A10": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A11": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A1_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A2_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A2m_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A4_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A4m_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A5": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A6": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A7": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A8": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A8_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A8m_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A9": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_B1ls": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_B1ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_B1s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_B2ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_B2s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_B4ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_B8ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_D1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D11": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D11_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D12": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D12_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D13": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D13_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D14": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D14_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D15_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D16_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D16s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_D1_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_D3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D32_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D32s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_D3_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_D5_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D64_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D64s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_D8_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D8s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DC2s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DC4s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS1": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS11": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS11-1_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS11_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12-1_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12-2_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13-2_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13-4_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14-4_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14-8_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS15_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS1_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS2_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS3_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS4": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS4_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS5_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E16-4s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E16-8s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E16_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E16s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E20_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E20s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E2_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E2s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E32-16s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E32-8s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E32_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E32s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E4-2s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E4_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E4s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E64-16s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E64-32s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E64_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E64i_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E64is_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E64s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E8-2s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E8-4s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_E8_v3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_E8s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F16": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F16s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F16s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F1s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F2s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F2s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F32s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F4s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F4s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F64s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F72s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F8": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F8s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F8s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_G1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G5": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_GS1": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS4": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS4-4": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS4-8": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS5": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS5-16": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS5-8": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_H16": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H16m": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H16mr": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H16r": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H8": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H8m": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_HB60rs": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_HC44rs": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L16s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L16s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L32s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L32s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L4s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L64s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L80s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L8s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L8s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M128": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_M128-32ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M128-64ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M128m": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_M128ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M128s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M16-4ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M16-8ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M16ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M32-16ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M32-8ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M32ls": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M32ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M32ts": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M64": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_M64-16ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M64-32ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M64ls": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M64m": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_M64ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M64s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M8-2ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M8-4ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M8ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC12": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC12s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC12s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC24": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC24r": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC24rs_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC24rs_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC24s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC24s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC6": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC6s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC6s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_ND12s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_ND24rs": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_ND24s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_ND6s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NV12": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NV12s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NV12s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NV24": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NV24s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NV24s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NV48s_v3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NV6": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NV6s_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_PB12s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_PB24s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_PB6s": {
      "storageAccountType": "Premium_LRS"
    }
   }
`
}
