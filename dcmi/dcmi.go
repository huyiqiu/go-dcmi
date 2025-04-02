package dcmi

/*
#cgo LDFLAGS: -ldl
#include "dcmi_wrapper.h"
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)


const (
	MAX_CHIP_NAME_LEN = 32
	DCMI_SUCCESS      = 0
	MAX_CARD_NUM      = 64
	MAX_LEN           = 256
)

const (
	DCMI_UTILIZATION_RATE_DDR           = 1
	DCMI_UTILIZATION_RATE_AICORE        = 2
	DCMI_UTILIZATION_RATE_AICPU         = 3
	DCMI_UTILIZATION_RATE_CTRLCPU       = 4
	DCMI_UTILIZATION_RATE_DDR_BANDWIDTH = 5
	DCMI_UTILIZATION_RATE_HBM           = 6
	DCMI_UTILIZATION_RATE_HBM_BANDWIDTH = 10
	DCMI_UTILIZATION_RATE_VECTORCORE    = 12
)

type DcmiUnitType int

const (
	NPU_TYPE     DcmiUnitType = 0
	MCU_TYPE     DcmiUnitType = 1
	CPU_TYPE     DcmiUnitType = 2
	INVALID_TYPE DcmiUnitType = 0xFF
)

func DcmiInit() error {
	ret := C.dl_init()
	if ret != 0 {
		return fmt.Errorf("failed to fetch libdcmi")
	}
	ret = C.dcmi_init()
	if ret != DCMI_SUCCESS {
		return fmt.Errorf("failed to initialize dcmi, %v", dcmiError(int(ret)))
	}
	return nil
}

func GetDcmiVersion() (string, error) {
	ver := make([]byte, 64)
	ret := C.dcmi_get_dcmi_version((*C.char)(unsafe.Pointer(&ver[0])), 64)
	if ret != DCMI_SUCCESS {
		return "", fmt.Errorf("failed to get dcmi version, %v", dcmiError(int(ret)))
	}
	return string(ver[:clen(ver)]), nil
}

func GetCardList() (int, []int, error) {
	var cardNum C.int
	cardList := make([]C.int, MAX_CARD_NUM)

	ret := C.dcmi_get_card_list(&cardNum, &cardList[0], C.int(MAX_CARD_NUM))
	if ret != DCMI_SUCCESS {
		return 0, nil, fmt.Errorf("failed to get card list, %v", dcmiError(int(ret)))
	}

	goCardList := make([]int, int(cardNum))
	for i := 0; i < int(cardNum); i++ {
		goCardList[i] = int(cardList[i])
	}

	return int(cardNum), goCardList, nil
}

func GetDeviceNumInCard(cardId int) (int, error) {
	var deviceNum C.int
	ret := C.dcmi_get_device_num_in_card(C.int(cardId), &deviceNum)
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device num in card, %v", dcmiError(int(ret)))
	}
	return int(deviceNum), nil
}

type dcmiDevice struct {
	cardId   int
	deviceId int
}

func GetDeviceByCardIdAndDeviceId(cardId, deviceId int) Device {
	return &dcmiDevice{
		cardId:   cardId,
		deviceId: deviceId,
	}
}

func (d *dcmiDevice) GetDeviceType() (string, error) {
	var deviceType C.enum_dcmi_unit_type
	ret := C.dcmi_get_device_type(C.int(d.cardId), C.int(d.deviceId), &deviceType)
	if ret != DCMI_SUCCESS {
		return "", fmt.Errorf("failed to get device type, %v", dcmiError(int(ret)))
	}
	var dtype string
	switch DcmiUnitType(deviceType) {
	case NPU_TYPE:
		dtype = "NPU"
	case MCU_TYPE:
		dtype = "MCU"
	case CPU_TYPE:
		dtype = "CPU"
	default:
		dtype = "INVALID"
	}
	return dtype, nil
}

type cDcmiChipInfo struct {
	ChipType  [MAX_CHIP_NAME_LEN]C.uchar
	ChipName  [MAX_CHIP_NAME_LEN]C.uchar
	ChipVer   [MAX_CHIP_NAME_LEN]C.uchar
	AicoreCnt C.uint
}

func (d *dcmiDevice) GetDeviceChipInfo() (DcmiChipInfo, error) {
	var cInfo cDcmiChipInfo
	ret := C.dcmi_get_device_chip_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_chip_info)(unsafe.Pointer(&cInfo)))
	if ret != DCMI_SUCCESS {
		return DcmiChipInfo{}, fmt.Errorf("failed get card%d-device%d chip info, %v", d.cardId, d.deviceId, dcmiError(int(ret)))
	}
	goInfo := DcmiChipInfo{
		ChipType:  C.GoString((*C.char)(unsafe.Pointer(&cInfo.ChipType[0]))),
		ChipName:  C.GoString((*C.char)(unsafe.Pointer(&cInfo.ChipName[0]))),
		ChipVer:   C.GoString((*C.char)(unsafe.Pointer(&cInfo.ChipVer[0]))),
		AicoreCnt: uint32(cInfo.AicoreCnt),
	}
	return goInfo, nil
}

func (d *dcmiDevice) GetDevicePcieInfo() (DcmiPcieInfo, error) {
	var pcieInfo DcmiPcieInfo
	ret := C.dcmi_get_device_pcie_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_pcie_info)(unsafe.Pointer(&pcieInfo)))
	if ret != DCMI_SUCCESS {
		return DcmiPcieInfo{}, fmt.Errorf("failed get card%d-device%d pcie info, %v", d.cardId, d.deviceId, dcmiError(int(ret)))
	}
	return pcieInfo, nil
}

type cDcmiElabelInfo struct {
	ProductName      [MAX_LEN]C.char
	Model            [MAX_LEN]C.char
	Manufacturer     [MAX_LEN]C.char
	ManufacturerDate [MAX_LEN]C.char
	SerialNumber     [MAX_LEN]C.char
}

// 310P not support
func (d *dcmiDevice) GetDeviceElabelInfo() (DcmiElabelInfo, error) {
	var cInfo cDcmiElabelInfo
	ret := C.dcmi_get_device_elabel_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_elabel_info)(unsafe.Pointer(&cInfo)))
	if ret != DCMI_SUCCESS {
		return DcmiElabelInfo{}, fmt.Errorf("failed to get card%d-device%d elabel info, %v", d.cardId, d.deviceId, dcmiError(int(ret)))
	}
	goInfo := DcmiElabelInfo{
		ProductName:      C.GoString((*C.char)(unsafe.Pointer(&cInfo.ProductName[0]))),
		Model:            C.GoString((*C.char)(unsafe.Pointer(&cInfo.Model[0]))),
		Manufacturer:     C.GoString((*C.char)(unsafe.Pointer(&cInfo.Manufacturer[0]))),
		ManufacturerDate: C.GoString((*C.char)(unsafe.Pointer(&cInfo.ManufacturerDate[0]))),
		SerialNumber:     C.GoString((*C.char)(unsafe.Pointer(&cInfo.SerialNumber[0]))),
	}
	return goInfo, nil
}

func (d *dcmiDevice) GetDeviceHealth() (int, error) {
	var health C.uint
	ret := C.dcmi_get_device_health(C.int(d.cardId), C.int(d.deviceId), &health)
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device health status, %v", dcmiError(int(ret)))
	}
	return int(ret), nil
}

func (d *dcmiDevice) GetDeviceMemoryInfoV3() (DcmiMemoryInfoV3, error) {
	var memInfo DcmiMemoryInfoV3
	ret := C.dcmi_get_device_memory_info_v3(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_get_memory_info_stru)(unsafe.Pointer(&memInfo)))
	if ret != DCMI_SUCCESS {
		return memInfo, fmt.Errorf("failed to get device memory info, %v", dcmiError(int(ret)))
	}
	return memInfo, nil
}

func (d *dcmiDevice) GetDeviceUtilizationRate() (int, error) {
	var utilization C.uint
	ret := C.dcmi_get_device_utilization_rate(C.int(d.cardId), C.int(d.deviceId), DCMI_UTILIZATION_RATE_AICORE, &utilization)
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device utilization, %v", dcmiError(int(ret)))
	}
	return int(utilization), nil
}

func (d *dcmiDevice) GetDeviceAicoreInfo() (DcmiAicoreInfo, error) {
	var aicoreInfo DcmiAicoreInfo
	ret := C.dcmi_get_device_aicore_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_aicore_info)(unsafe.Pointer(&aicoreInfo)))
	if ret != DCMI_SUCCESS {
		return aicoreInfo, fmt.Errorf("failed to get device aicore info, %v", dcmiError(int(ret)))
	}
	return aicoreInfo, nil
}

func (d *dcmiDevice) GetDeviceResourceInfo() ([]DcmiProcMemInfo, error) {
	var procNum C.int
	ret := C.dcmi_get_device_resource_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_proc_mem_info)(unsafe.Pointer(&DcmiProcMemInfo{})), &procNum)
	if ret != DCMI_SUCCESS {
		return []DcmiProcMemInfo{}, fmt.Errorf("failed to get proc numbers, %v", dcmiError(int(ret)))
	}
	resourceInfo := make([]DcmiProcMemInfo, procNum)
	if procNum == 0 {
		return resourceInfo, nil
	}
	ret = C.dcmi_get_device_resource_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_proc_mem_info)(unsafe.Pointer(&resourceInfo[0])), &procNum)
	if ret != DCMI_SUCCESS {
		return []DcmiProcMemInfo{}, fmt.Errorf("failed to get device resource info, %v", dcmiError(int(ret)))
	}
	return resourceInfo, nil
}

func (d *dcmiDevice) GetPowerInfo() (string, error) {
	var power C.int
	ret := C.dcmi_get_device_power_info(C.int(d.cardId), C.int(d.deviceId), &power)
	if ret != DCMI_SUCCESS {
		return "N/A", fmt.Errorf("failed to get power info, %v", dcmiError(int(ret)))
	}
	powerString := fmt.Sprintf("%.1f", float64(power) / 10)
	return powerString, nil
}

func (d *dcmiDevice) GetDeviceTemperature() (int, error) {
	var temp C.int
	ret := C.dcmi_get_device_temperature(C.int(d.cardId), C.int(d.deviceId), &temp)
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device temperature, %v", dcmiError(int(ret)))
	}
	return int(temp), nil
}

func (d *dcmiDevice) GetDeviceHbmInfo() (DcmiHbmInfo, error) {
	var dcmiHbmInfo DcmiHbmInfo
	ret := C.dcmi_get_device_hbm_info(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_hbm_info)(unsafe.Pointer(&dcmiHbmInfo)))
	if ret != DCMI_SUCCESS {
		return dcmiHbmInfo, fmt.Errorf("failed to get device hbm info, %v", dcmiError(int(ret)))
	}
	return dcmiHbmInfo, nil
}

func (d *dcmiDevice) GetDeviceShareEnable() (bool, error) {
	var enable C.int
	ret := C.dcmi_get_device_share_enable(C.int(d.cardId), C.int(d.deviceId), &enable)
	if ret != DCMI_SUCCESS {
		return false, fmt.Errorf("failed to get device share enable, %v", dcmiError(int(ret)))
	}
	return enable == 1, nil
}

func (d *dcmiDevice) GetDeviceChipSlot() (int, error) {
	var chip_pod_id C.int
	ret := C.dcmi_get_device_chip_slot(C.int(d.cardId), C.int(d.deviceId), &chip_pod_id);
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device chip slot, %v", dcmiError(int(ret)))
	}
	return int(chip_pod_id), nil
}

func (d *dcmiDevice) GetDeviceLogicId() (int, error) {
	var logic_id C.int
	ret := C.dcmi_get_device_logic_id(&logic_id, C.int(d.cardId), C.int(d.deviceId));
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device logicId, %v", dcmiError(int(ret)))
	}
	return int(logic_id), nil
}

func (d *dcmiDevice) GetDevicePhyId() (int, error) {
	var phy_id C.uint
	var logic_id C.int
	ret := C.dcmi_get_device_logic_id(&logic_id, C.int(d.cardId), C.int(d.deviceId));
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device logicId, %v", dcmiError(int(ret)))
	}
	ret = C.dcmi_get_device_phyid_from_logicid(C.uint(logic_id), &phy_id);
	if ret != DCMI_SUCCESS {
		return 0, fmt.Errorf("failed to get device phyId by logicId, %v", dcmiError(int(ret)))
	}
	return int(phy_id), nil
}

type cDcmiSocDieStru struct {
	SocDie [5]C.uint
}

func (d *dcmiDevice) GetDeviceDie() (string, error) {
	var die cDcmiSocDieStru
	ret := C.dcmi_get_device_die(C.int(d.cardId), C.int(d.deviceId), (*C.struct_dcmi_soc_die_stru)(unsafe.Pointer(&die)))
	if ret != DCMI_SUCCESS {
		return "", fmt.Errorf("failed to get device die, %v", dcmiError(int(ret)))
	}
	var hexValues []string
	for _, item := range die.SocDie {
		hexValues = append(hexValues, fmt.Sprintf("%08X", item))
	}

	return strings.Join(hexValues, " "), nil
}


func (d *dcmiDevice) SetDeviceShareEnable(enable bool) error {
	var enable_flag C.int
	if enable {
		enable_flag = 1
	} else {
		enable_flag = 0
	}
	ret := C.dcmi_set_device_share_enable(C.int(d.cardId), C.int(d.deviceId), enable_flag)
	if ret != DCMI_SUCCESS {
		return fmt.Errorf("failed to set device share enable to %v, %v", enable, dcmiError(int(ret)))
	}
	return nil
}


type DcmiChipInfo struct {
	ChipType  string
	ChipName  string
	ChipVer   string
	AicoreCnt uint32
}

type DcmiElabelInfo struct {
	ProductName      string
	Model            string
	Manufacturer     string
	ManufacturerDate string
	SerialNumber     string
}

type DcmiMemoryInfoV3 struct {
	MemorySize      uint64
	MemoryAvailable uint64
	Freq            uint32
	HugePageSize    uint64
	HugePagesTotal  uint64
	HugePagesFree   uint64
	Utilization     uint32
	Reserve         [60]byte
}

type DcmiHbmInfo struct {
    MemorySize       uint64
    Freq             uint32
    MemoryUsage      uint64
    Temp             int32
    BandwidthUtilRate uint32
}

type DcmiAicoreInfo struct {
	Freq    uint32
	CurFreq uint32
}

type DcmiProcMemInfo struct {
	ProcId       int32
	ProcMemUsage uint64
}

type DcmiPcieInfo struct {
    DeviceID    uint32
    VenderID    uint32
    SubVenderID uint32
    SubDeviceID uint32
    BdfDeviceID uint32
    BdfBusID    uint32
    BdfFuncID   uint32
}


type Device interface {
	GetDeviceType() (string, error)
	GetDeviceChipInfo() (DcmiChipInfo, error)
	GetDeviceElabelInfo() (DcmiElabelInfo, error)
	GetDeviceHealth() (int, error)
	GetDeviceMemoryInfoV3() (DcmiMemoryInfoV3, error)
	GetDeviceUtilizationRate() (int, error)
	GetDeviceAicoreInfo() (DcmiAicoreInfo, error)
	GetDeviceResourceInfo() ([]DcmiProcMemInfo, error)
	GetPowerInfo() (string, error)
	GetDeviceTemperature() (int, error)
	GetDevicePcieInfo() (DcmiPcieInfo, error)
	GetDeviceHbmInfo() (DcmiHbmInfo, error)
	GetDeviceShareEnable() (bool, error)
	GetDeviceChipSlot() (int, error)
	GetDeviceLogicId() (int, error)
	GetDevicePhyId() (int, error)
	GetDeviceDie() (string, error)

	SetDeviceShareEnable(enable bool) error
}
