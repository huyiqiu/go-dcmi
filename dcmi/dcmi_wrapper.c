#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>
#include "dcmi_wrapper.h"


void *dcmiHandle;

dcmiReturn_t (*dcmiInitFunc)();
dcmiReturn_t dcmi_init() {
    if (dcmiInitFunc == NULL) {
        return 1;
    }
    return dcmiInitFunc();
}

dcmiReturn_t (*dcmiGetDcmiVersionFunc)(char *dcmi_ver, unsigned int len);
dcmiReturn_t dcmi_get_dcmi_version(char *dcmi_ver, unsigned int len) {
    if (dcmiGetDcmiVersionFunc == NULL) {
        return 1;
    }
    return dcmiGetDcmiVersionFunc(dcmi_ver, len);
}

dcmiReturn_t (*dcmiGetCardListFunc)(int *card_num, int *card_list, int list_len);
dcmiReturn_t dcmi_get_card_list(int *card_num, int *card_list, int list_len) {
    if (dcmiGetCardListFunc == NULL) {
        return 1;
    }
    return dcmiGetCardListFunc(card_num, card_list, list_len);
}

dcmiReturn_t (*dcmiGetDeviceNumInCardFunc)(int card_id, int *device_num);
dcmiReturn_t dcmi_get_device_num_in_card(int card_id, int *device_num) {
    if (dcmiGetDeviceNumInCardFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceNumInCardFunc(card_id, device_num);
}

dcmiReturn_t (*dcmiGetDeviceTypeFunc)(int card_id, int device_id, enum dcmi_unit_type *device_type);
dcmiReturn_t dcmi_get_device_type(int card_id, int device_id, enum dcmi_unit_type *device_type) {
    if (dcmiGetDeviceTypeFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceTypeFunc(card_id, device_id, device_type);
}

dcmiReturn_t (*dcmiGetDeviceChipInfoFunc) (int card_id, int device_id, struct dcmi_chip_info *chip_info);
dcmiReturn_t dcmi_get_device_chip_info(int card_id, int device_id, struct dcmi_chip_info *chip_info) {
    if (dcmiGetDeviceChipInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceChipInfoFunc(card_id, device_id, chip_info);
}

dcmiReturn_t (*dcmiGetDevicePcieInfoFunc) (int card_id, int device_id, struct dcmi_pcie_info *pcie_info);
dcmiReturn_t dcmi_get_device_pcie_info(int card_id, int device_id, struct dcmi_pcie_info *pcie_info) {
    if (dcmiGetDevicePcieInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDevicePcieInfoFunc(card_id, device_id, pcie_info);
}

dcmiReturn_t (*dcmiGetDeviceElabelInfoFunc) (int card_id, int device_id, struct dcmi_elabel_info *elabel_info);
dcmiReturn_t dcmi_get_device_elabel_info(int card_id, int device_id, struct dcmi_elabel_info *elabel_info) {
    if (dcmiGetDeviceElabelInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceElabelInfoFunc(card_id, device_id, elabel_info);
}

dcmiReturn_t (*dcmiGetDeviceHealthFunc) (int card_id, int device_id, unsigned int *health);
dcmiReturn_t dcmi_get_device_health(int card_id, int device_id, unsigned int *health) {
    if (dcmiGetDeviceHealthFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceHealthFunc(card_id, device_id, health);
}

dcmiReturn_t (*dcmiGetDeviceMemoryInfoFunc) (int card_id, int device_id, struct dcmi_get_memory_info_stru *memory_info);
dcmiReturn_t dcmi_get_device_memory_info_v3(int card_id, int device_id, struct dcmi_get_memory_info_stru *memory_info) {
    if (dcmiGetDeviceMemoryInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceMemoryInfoFunc(card_id, device_id, memory_info);
}

dcmiReturn_t (*dcmiGetDeviceUtilizationRateFunc) (int card_id, int device_id, int input_type, unsigned int *utilization_rate);
dcmiReturn_t dcmi_get_device_utilization_rate(int card_id, int device_id, int input_type, unsigned int *utilization_rate) {
    if (dcmiGetDeviceUtilizationRateFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceUtilizationRateFunc(card_id, device_id, input_type, utilization_rate);
}

dcmiReturn_t (*dcmiGetDeviceAicoreInfoFunc) (int card_id, int device_id, struct dcmi_aicore_info *aicore_info);
dcmiReturn_t dcmi_get_device_aicore_info(int card_id, int device_id, struct dcmi_aicore_info *aicore_info) {
    if (dcmiGetDeviceAicoreInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceAicoreInfoFunc(card_id, device_id, aicore_info);
}

dcmiReturn_t (*dcmiGetDeviceResourceInfoFunc) (int card_id, int device_id, struct dcmi_proc_mem_info *proc_info, int *proc_num);
dcmiReturn_t dcmi_get_device_resource_info(int card_id, int device_id, struct dcmi_proc_mem_info *proc_info, int *proc_num) {
    if (dcmiGetDeviceResourceInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceResourceInfoFunc(card_id, device_id, proc_info, proc_num);
}

dcmiReturn_t (*dcmiGetDevicePowerInfoFunc) (int card_id, int device_id, int *power);
dcmiReturn_t dcmi_get_device_power_info(int card_id, int device_id, int *power) {
    if (dcmiGetDevicePowerInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDevicePowerInfoFunc(card_id, device_id, power);
}

dcmiReturn_t (*dcmiGetDeviceTemperatureFunc) (int card_id, int device_id, int *temperature);
dcmiReturn_t dcmi_get_device_temperature(int card_id, int device_id, int *temperature) {
    if (dcmiGetDeviceTemperatureFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceTemperatureFunc(card_id, device_id, temperature);
}

dcmiReturn_t (*dcmiGetDeviceHbmInfoFunc) (int card_id, int device_id, struct dcmi_hbm_info *device_hbm_info);
dcmiReturn_t dcmi_get_device_hbm_info(int card_id, int device_id, struct dcmi_hbm_info *device_hbm_info) {
    if (dcmiGetDeviceHbmInfoFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceHbmInfoFunc(card_id, device_id, device_hbm_info);
}

dcmiReturn_t (*dcmiGetDeviceShareEnableFunc) (int card_id, int device_id, int *enable_flag);
dcmiReturn_t dcmi_get_device_share_enable(int card_id, int device_id, int *enable_flag) {
    if (dcmiGetDeviceShareEnableFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceShareEnableFunc(card_id, device_id, enable_flag);
}

dcmiReturn_t (*dcmiGetDeviceChipSlotFunc) (int card_id, int device_id, int *chip_pos_id);
dcmiReturn_t dcmi_get_device_chip_slot(int card_id, int device_id, int *chip_pos_id) {
    if (dcmiGetDeviceChipSlotFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceChipSlotFunc(card_id, device_id, chip_pos_id);
}

dcmiReturn_t (*dcmiGetDeviceLogicIdFunc) (int *device_logic_id, int card_id, int device_id);
dcmiReturn_t dcmi_get_device_logic_id(int *device_logic_id, int card_id, int device_id) {
    if (dcmiGetDeviceLogicIdFunc == NULL) {
        return 1;
    }
    return dcmiGetDeviceLogicIdFunc(device_logic_id, card_id, device_id);
}

dcmiReturn_t (*dcmiGetDevicePhyicIdFromLogicIdFunc) (unsigned int logicid, unsigned int *phyid);
dcmiReturn_t dcmi_get_device_phyid_from_logicid(unsigned int logicid, unsigned int *phyid) {
    if (dcmiGetDevicePhyicIdFromLogicIdFunc == NULL) {
        return 1;
    }
    return dcmiGetDevicePhyicIdFromLogicIdFunc(logicid, phyid);
}

dcmiReturn_t (*dcmiGetDeviceDieFunc) (int card_id, int device_id, struct dcmi_soc_die_stru *device_die);
dcmiReturn_t dcmi_get_device_die(int card_id, int device_id, struct dcmi_soc_die_stru *device_die) {
    if (dcmiGetDeviceDieFunc == NULL) {
        return 1;
    } 
    return dcmiGetDeviceDieFunc(card_id, device_id, device_die);
}

dcmiReturn_t (*dcmiSetDeviceShareEnableFunc) (int card_id, int device_id, int enable_flag);
dcmiReturn_t dcmi_set_device_share_enable(int card_id, int device_id, int enable_flag) {
    if (dcmiSetDeviceShareEnableFunc == NULL) {
        return 1;
    }
    return dcmiSetDeviceShareEnableFunc(card_id, device_id, enable_flag);
}

dcmiReturn_t dl_init() {
    dcmiHandle = dlopen(DCMI_LIBRARY_PATH, RTLD_LAZY);
    if (!dcmiHandle) {
        return DCMI_ERROR_LIBRARY_NOT_FOUND;
    }
    dcmiInitFunc = dlsym(dcmiHandle, "dcmi_init");
    if (!dcmiInitFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }

    dcmiGetDcmiVersionFunc = dlsym(dcmiHandle, "dcmi_get_dcmi_version");
    if (!dcmiGetDcmiVersionFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetCardListFunc = dlsym(dcmiHandle, "dcmi_get_card_list");
    if (!dcmiGetCardListFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceNumInCardFunc = dlsym(dcmiHandle, "dcmi_get_device_num_in_card");
    if (!dcmiGetDeviceNumInCardFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceTypeFunc = dlsym(dcmiHandle, "dcmi_get_device_type");
    if (!dcmiGetDeviceTypeFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceChipInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_chip_info");
    if (!dcmiGetDeviceChipInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDevicePcieInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_pcie_info");
    if (!dcmiGetDeviceChipInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceElabelInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_elabel_info");
    if (!dcmiGetDeviceElabelInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceHealthFunc = dlsym(dcmiHandle, "dcmi_get_device_health");
    if (!dcmiGetDeviceHealthFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceMemoryInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_memory_info_v3");
    if (!dcmiGetDeviceMemoryInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceUtilizationRateFunc = dlsym(dcmiHandle, "dcmi_get_device_utilization_rate");
    if (!dcmiGetDeviceUtilizationRateFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceAicoreInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_aicore_info");
    if (!dcmiGetDeviceAicoreInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceResourceInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_resource_info");
    if (!dcmiGetDeviceResourceInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDevicePowerInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_power_info");
    if (!dcmiGetDevicePowerInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceTemperatureFunc = dlsym(dcmiHandle, "dcmi_get_device_temperature");
    if (!dcmiGetDeviceTemperatureFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceHbmInfoFunc = dlsym(dcmiHandle, "dcmi_get_device_hbm_info");
    if (!dcmiGetDeviceHbmInfoFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceShareEnableFunc = dlsym(dcmiHandle, "dcmi_get_device_share_enable");
    if (!dcmiGetDeviceShareEnableFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceChipSlotFunc = dlsym(dcmiHandle, "dcmi_get_device_chip_slot");
    if (!dcmiGetDeviceChipSlotFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceLogicIdFunc = dlsym(dcmiHandle, "dcmi_get_device_logic_id");
    if (!dcmiGetDeviceLogicIdFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDevicePhyicIdFromLogicIdFunc = dlsym(dcmiHandle, "dcmi_get_device_phyid_from_logicid");
    if (!dcmiGetDevicePhyicIdFromLogicIdFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    dcmiGetDeviceDieFunc = dlsym(dcmiHandle, "dcmi_get_device_die");
    if (!dcmiGetDeviceDieFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }


    dcmiSetDeviceShareEnableFunc = dlsym(dcmiHandle, "dcmi_set_device_share_enable");
    if (!dcmiSetDeviceShareEnableFunc) {
        return DCMI_ERROR_SYMBOL_NOT_FOUND;
    }
    return 0;
}


