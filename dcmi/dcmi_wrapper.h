#include "dcmi_interface_api.h"
#define DCMI_LIBRARY_PATH "libdcmi.so"
#define DCMI_ERROR_LIBRARY_NOT_FOUND 1
#define DCMI_ERROR_SYMBOL_NOT_FOUND 2


typedef int dcmiReturn_t;

dcmiReturn_t dl_init();
dcmiReturn_t dcmi_init();
dcmiReturn_t dcmi_get_dcmi_version(char *dcmi_ver, unsigned int len);
dcmiReturn_t dcmi_get_card_list(int *card_num, int *card_list, int list_len);
dcmiReturn_t dcmi_get_device_num_in_card(int card_id, int *device_num);
dcmiReturn_t dcmi_get_device_type(int card_id, int device_id, enum dcmi_unit_type *device_type);
dcmiReturn_t dcmi_get_device_chip_info(int card_id, int device_id, struct dcmi_chip_info *chip_info);
dcmiReturn_t dcmi_get_device_pcie_info(int card_id, int device_id, struct dcmi_pcie_info *pcie_info);
dcmiReturn_t dcmi_get_device_health(int card_id, int device_id, unsigned int *health);
dcmiReturn_t dcmi_get_device_elabel_info(int card_id, int device_id, struct dcmi_elabel_info *elabel_info);
dcmiReturn_t dcmi_get_device_memory_info_v3(int card_id, int device_id, struct dcmi_get_memory_info_stru *memory_info);
dcmiReturn_t dcmi_get_device_utilization_rate(int card_id, int device_id, int input_type, unsigned int *utilization_rate);
dcmiReturn_t dcmi_get_device_aicore_info(int card_id, int device_id, struct dcmi_aicore_info *aicore_info);
dcmiReturn_t dcmi_get_device_resource_info(int card_id, int device_id, struct dcmi_proc_mem_info *proc_info, int *proc_num);
dcmiReturn_t dcmi_get_device_power_info(int card_id, int device_id, int *power);
dcmiReturn_t dcmi_get_device_temperature(int card_id, int device_id, int *temperature);
dcmiReturn_t dcmi_get_device_hbm_info(int card_id, int device_id, struct dcmi_hbm_info *device_hbm_info);
dcmiReturn_t dcmi_get_device_share_enable(int card_id, int device_id, int *enable_flag);
dcmiReturn_t dcmi_get_device_chip_slot(int card_id, int device_id, int *chip_pos_id);
dcmiReturn_t dcmi_get_device_logic_id(int *device_logic_id, int card_id, int device_id);
dcmiReturn_t dcmi_get_device_phyid_from_logicid(unsigned int logicid, unsigned int *phyid);
dcmiReturn_t dcmi_get_device_die(int card_id, int device_id, struct dcmi_soc_die_stru *device_die);

// set
dcmiReturn_t dcmi_set_device_share_enable(int card_id, int device_id, int enable_flag);