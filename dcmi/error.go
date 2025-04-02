package dcmi

import "fmt"

/*----------------------------------------------*
 #define DCMI_OK 0
 #define DCMI_ERROR_CODE_BASE (-8000)
 #define DCMI_ERR_CODE_INVALID_PARAMETER             (DCMI_ERROR_CODE_BASE - 1)
 #define DCMI_ERR_CODE_OPER_NOT_PERMITTED            (DCMI_ERROR_CODE_BASE - 2)
 #define DCMI_ERR_CODE_MEM_OPERATE_FAIL              (DCMI_ERROR_CODE_BASE - 3)
 #define DCMI_ERR_CODE_SECURE_FUN_FAIL               (DCMI_ERROR_CODE_BASE - 4)
 #define DCMI_ERR_CODE_INNER_ERR                     (DCMI_ERROR_CODE_BASE - 5)
 #define DCMI_ERR_CODE_TIME_OUT                      (DCMI_ERROR_CODE_BASE - 6)
 #define DCMI_ERR_CODE_INVALID_DEVICE_ID             (DCMI_ERROR_CODE_BASE - 7)
 #define DCMI_ERR_CODE_DEVICE_NOT_EXIST              (DCMI_ERROR_CODE_BASE - 8)
 #define DCMI_ERR_CODE_IOCTL_FAIL                    (DCMI_ERROR_CODE_BASE - 9)
 #define DCMI_ERR_CODE_SEND_MSG_FAIL                 (DCMI_ERROR_CODE_BASE - 10)
 #define DCMI_ERR_CODE_RECV_MSG_FAIL                 (DCMI_ERROR_CODE_BASE - 11)
 #define DCMI_ERR_CODE_NOT_REDAY                     (DCMI_ERROR_CODE_BASE - 12)
 #define DCMI_ERR_CODE_NOT_SUPPORT_IN_CONTAINER      (DCMI_ERROR_CODE_BASE - 13)
 #define DCMI_ERR_CODE_FILE_OPERATE_FAIL             (DCMI_ERROR_CODE_BASE - 14)
 #define DCMI_ERR_CODE_RESET_FAIL                    (DCMI_ERROR_CODE_BASE - 15)
 #define DCMI_ERR_CODE_ABORT_OPERATE                 (DCMI_ERROR_CODE_BASE - 16)
 #define DCMI_ERR_CODE_IS_UPGRADING                  (DCMI_ERROR_CODE_BASE - 17)
 #define DCMI_ERR_CODE_RESOURCE_OCCUPIED             (DCMI_ERROR_CODE_BASE - 20)
 #define DCMI_ERR_CODE_PARTITION_NOT_RIGHT           (DCMI_ERROR_CODE_BASE - 22)
 #define DCMI_ERR_CODE_NOT_SUPPORT                   (DCMI_ERROR_CODE_BASE - 255)
*----------------------------------------------*/

var errorMessages = map[int]string{
    -8001: "Invalid parameter",
    -8002: "Operation not permitted",
    -8003: "Memory operate fail",
    -8004: "Secure function fail",
    -8005: "Inner error",
    -8006: "Timeout",
    -8007: "Invalid device ID",
    -8008: "Device not exist",
    -8009: "IOCTL fail",
    -8010: "Send message fail",
    -8011: "Receive message fail",
    -8012: "Not ready",
    -8013: "Not support in container",
    -8014: "File operate fail",
    -8015: "Reset fail",
    -8016: "Abort operate",
    -8017: "Is upgrading",
    -8020: "Resource occupied",
    -8022: "Partition not right",
    -8255: "Not support",
}


func dcmiError(ret int) error {
    if msg, ok := errorMessages[ret]; ok {
        return fmt.Errorf("%s", msg)
    }
    return fmt.Errorf("unknown error: %d", ret)
}