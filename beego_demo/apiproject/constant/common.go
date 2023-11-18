package constant

// MainWgAddNum 服务启动控制数
const MainWgAddNum = 1

// MakeArrLenDefault 设置数组默认长度
const MakeArrLenDefault = 0

// AdapterConsole 日志类型【控制台】
const AdapterConsole = "console"

// AdapterFile 日志类型【文件】
const AdapterFile = "file"

// AdapterConsoleFile 日志类型【文件+控制台】
const AdapterConsoleFile = "console_file"

// SwitchStatus 默认开关状态
const SwitchStatus = 1

// ParamsDefault 请求参数默认值
const ParamsDefault = -1

// ParamsMin 参数最小值
const ParamsMin = 0

// BaseTypeMain 主基站
const BaseTypeMain = 1

// BaseTypeFollow 从基站
const BaseTypeFollow = 0

// BaseTypeBle 蓝牙基站
const BaseTypeBle = 2

// CarCountDefault 台车统计默认值
const CarCountDefault = 1

// BaseTypeFrom 从基站
const BaseTypeFrom = 0

// AreaDataTypeTof 小区定位类型【1: tof 2: tdoa】
const AreaDataTypeTof = 1

// AreaDataType 小区定位类型【1: tof 2: tdoa】
const AreaDataType = 2

// AreaUpdateType 小区更新类型
const AreaUpdateType = 1

// AreaDimensionOne 小区维度类型
const AreaDimensionOne = 1

// AreaDimensionTwo 小区维度类型
const AreaDimensionTwo = 2

// AreaSignal 信号强度门限
const AreaSignal = -110

// AreaDist 距离门限
const AreaDist = 20

// AreaAcross 跨区域门限
const AreaAcross = 50

// AreaMonitor 误差监测
const AreaMonitor = 1

// AreaMethods 算法
const AreaMethods = "nearSearch"

// TdoAAreaDist tdoA距离门限
const TdoAAreaDist = 1

// AreaTdoAMinBaseNum 最小定位基站数
const AreaTdoAMinBaseNum = 4

// AreaTdoAOneMinBaseNum 最小定位基站数【一维】
const AreaTdoAOneMinBaseNum = 2

// AreaTofMinBaseNum 最小定位基站数
const AreaTofMinBaseNum = 3

// AreaMinBaseNum 最小定位基站数
const AreaMinBaseNum = 1

// AreaDirectionDefault 一维默认方向
const AreaDirectionDefault = 1

// RedisListLenMax 默认最大队列数
const RedisListLenMax = 500

// TofCarStatus 台车状态标识
const TofCarStatus = 1

// WarnTypeBat 低电量
const WarnTypeBat = 0x01

// WarnTypeCharging 充电
const WarnTypeCharging = 0x02

// WarnTypeCut 剪断
const WarnTypeCut = 0x04

// WarnTypeLeave 撤离
const WarnTypeLeave = 0x08

// WarnTypeHelp 求救
const WarnTypeHelp = 0x10

// WarnTypeDown 下行区域报警
const WarnTypeDown = 0x20

// WarnTagBatMax 低电量报警最大值
const WarnTagBatMax = 10

// WarnTagBatOverMax 低电量报警最大值
const WarnTagBatOverMax = 25

// TrackDayDefault 历史轨迹默认保存天数
const TrackDayDefault = -7

// CarAlarmDist 台车测距状态
const CarAlarmDist = 0

// CarAlarmStatus 台车报警状态
const CarAlarmStatus = 1

// CarAlarmEnd 台车报警结束
const CarAlarmEnd = 2

// AreaTagTypeLeft 小区类型【左洞】
const AreaTagTypeLeft = -1

// AreaTagTypeRight 小区类型【右洞】
const AreaTagTypeRight = 1

// BaseStatusTrue 基站状态
const BaseStatusTrue = 2

// BaseStatusFalse 基站状态
const BaseStatusFalse = 0

// CmdTypeUwb 命令类型
const CmdTypeUwb = "uwb"

// CmdTypeAp 命令类型
const CmdTypeAp = "ap"

// CmdTypeAgent 命令类型
const CmdTypeAgent = "agent"

// CmdTypeBle 命令类型
const CmdTypeBle = "ble"

// LocationCoordinate 坐标系超时时间
const LocationCoordinate = 2 * 1000

// BaseCarDistTime 台车移动间隔时间
const BaseCarDistTime = 5 * 1000

// BaseCarDistLen 台车单位时间移动的最大距离
const BaseCarDistLen = 10

// CarTypeMining 开挖台车
const CarTypeMining = 1

// CarTypeWaterproof 防水台车
const CarTypeWaterproof = 2

// CarTypeInches 二寸台车
const CarTypeInches = 3

// CarTypeRaise 养护台车
const CarTypeRaise = 4

// CmdTypeConfig 下发命令类型 配置
const CmdTypeConfig = "config"

// ToolDataLenMax tool缓存数据最大长度
const ToolDataLenMax = 200

// ToolDataRssiOutTime tool信号强度缓存最大时间
const ToolDataRssiOutTime = 1000 * 60 * 2

// LoginTokenOverDay 登录默认过期天数
const LoginTokenOverDay = 3

// UserLoginStatus 登录禁用状态
const UserLoginStatus = 1

// ProjectMapDefault 地图和项目默认状态
const ProjectMapDefault = 1

// WarnStatusRun 正在报警
const WarnStatusRun = 0

// WarnStatusOver 停止报警
const WarnStatusOver = 1

// WarnStatusIgnore 忽略报警
const WarnStatusIgnore = 2

// RolePowerSuper 权限等级超级
const RolePowerSuper = 1

// RolePowerOrdinary 权限等级普通
const RolePowerOrdinary = 2

// PositionMinFilter TOF一维数据定位L滤波参数
const PositionMinFilter = 0.5

// PositionMaxFilter TOF一维数据定位L滤波参数
const PositionMaxFilter = 0.5

// DataTypeTdoA tdoA数据类型
const DataTypeTdoA = "TDOA"

// DataTypeTdoATof TDOA_TOF数据类型
const DataTypeTdoATof = "TDOA_TOF"

// DataTypeTof TOF数据类型
const DataTypeTof = "TOF"

// DataTypeSubG SubG数据类型
const DataTypeSubG = "SubG"

// RawDataTypeName raw数据类型
const RawDataTypeName = "raw"

// DataTypeBle ble数据类型
const DataTypeBle = "BLE"

// DataRawTypeBle raw数据类型
const DataRawTypeBle = "RAW_BLE"

// DataRawTypeTdoa raw数据类型
const DataRawTypeTdoa = "RAW_TDOA"

// DataTypeLow TOF低功耗数据类型
const DataTypeLow = "LOW"

// DataTypeTag 标签数据类型
const DataTypeTag = "Tag"

// BaseBatMax 设备最大电量
const BaseBatMax = 100

// TagOfflineTime 标签默认定位离线时间【秒】
const TagOfflineTime = 60 * 14.5

// TagOfflineWarnOutTime 标签缓存报警超时时间
const TagOfflineWarnOutTime = 1000 * 5

// UserImportLen 用户导入目录数量
const UserImportLen = 8

// FenceTypeInto 围栏类型入界
const FenceTypeInto = 1

// FenceTypeOut 围栏类型出界
const FenceTypeOut = 2

// FenceTypeOver 围栏类型超员
const FenceTypeOver = 3

// FenceTypeUnder 围栏类型缺员
const FenceTypeUnder = 4

// FenceTypeStop 围栏类型静止
const FenceTypeStop = 5

// FenceTypeStay 围栏类型滞留
const FenceTypeStay = 7

// FenceTypeOverSpeed 围栏类型超速
const FenceTypeOverSpeed = 8

// FenceTypeOverSpeed 围栏类型超速
const FenceTypeRollCall = 9

// WarnTypeInto 报警类型入界
const WarnTypeInto = -1

// WarnTypeOut 报警类型出界
const WarnTypeOut = -2

// WarnTypeOver 报警类型超员
const WarnTypeOver = -3

// WarnTypeUnder 报警类型缺员
const WarnTypeUnder = -4

// WarnTypeStop 报警类型静止
const WarnTypeStop = -5

// WarnTypeVisitorOut 报警类型访客超时
const WarnTypeVisitorOut = -6

// WarnTypeStay 报警类型滞留
const WarnTypeStay = -7

// WarnTypeOverSpeed 报警类型超速
const WarnTypeOverSpeed = -8

// BaseTypeBeacon 蓝牙Beacon
const BaseTypeBeacon = 2

// WorkTimeLenMin 工时统计数据最小长度
const WorkTimeLenMin = 1

// DiscTypeTake 主动测距类型
const DiscTypeTake = 1

// CarLocationOfflineTime 台车定位离线时间
const CarLocationOfflineTime = 1000 * 60 * 30

// CarLocationOfflineStatus 台车离线状态
const CarLocationOfflineStatus = 1

// TofDirectionLeft 开挖方向【左 -> 右】
const TofDirectionLeft = -1

// TofDirectionRight 开挖方向【右 -> 左】
const TofDirectionRight = 1

// TofDirectionTop 开挖方向【上 -> 下】
const TofDirectionTop = -2

// TofDirectionDown 开挖方向【下 -> 上】
const TofDirectionDown = 2

// WorkHourOutTime 工时统计超时时间
const WorkHourOutTime = 1 * 60

// BaseDualMacPrefix 双通道基站Mac前缀
const BaseDualMacPrefix = "10"

// MedianSliceMinLen 数组切片最小数量
const MedianSliceMinLen = 3

// MedianSliceNormalLen 数组切片正常数量
const MedianSliceNormalLen = 5

// LocationCarOutTime 台车位置缓存超时时间
const LocationCarOutTime = 10 * 1000

// KalmanN 滤波次数
const KalmanN = 2.0

// KalmanQ 滤波观测噪声
const KalmanQ = 0.3

// KalmanR 滤波频率
const KalmanR = 1

// KalmanT 滤波超时时间
const KalmanT = 5

// PositionBaseStatus 基站标点状态
const PositionBaseStatus = 1

// CoordinatesParamsNum 坐标参数数量
const CoordinatesParamsNum = 3

// LocationPushStatus 定位推送开关
const LocationPushStatus = 1

// LittleEndianLongOne 端序解析
const LittleEndianLongOne = 1

// LittleEndianLongTwo 端序解析
const LittleEndianLongTwo = 2

// LittleEndianLongFour 端序解析
const LittleEndianLongFour = 4

// LittleEndianLongEight 端序解析
const LittleEndianLongEight = 8

// BaseFwRssiTitle 子基站信号质量标题
const BaseFwRssiTitle = "rssi"

// BaseFwTimeTitle 子基站更新时间标题
const BaseFwTimeTitle = "time"

// BaseFwStateOutTime 子基站状态判断过期时间
const BaseFwStateOutTime = 5 * 60

// BasePositionStatus 基站点位状态
const BasePositionStatus = 1

// PushTypeTunnel 推送隧道类型
const PushTypeTunnel = "tunnel"

// BaseOfflineOutDay 设备离线天数
const BaseOfflineOutDay = -30

// BaseSpecialStatus 基站特殊点位标识
const BaseSpecialStatus = 1

// BaseSpecialRun 点位巡查进行中
const BaseSpecialRun = 0

// BaseSpecialOver 点位巡查结束
const BaseSpecialOver = 1

// BaseSpecialFail 点位巡查结束
const BaseSpecialFail = 2

// FenceNumMin 围栏报警最低总数
const FenceNumMin = 1

// WarnNumMin 报警最低总数
const WarnNumMin = 1

// TrackDistMin 轨迹距离最小值
const TrackDistMin = 3.0

// BaseTypeSREG 设备类型
const BaseTypeSREG = "SREG"

// HelpWarnStatusRun 报警状态
const HelpWarnStatusRun = 1

// TagMotionStop 标签静止状态
const TagMotionStop = 0

// TagMotionRun 标签运动状态
const TagMotionRun = 1

// UserTypeVisitor 访客
const UserTypeVisitor = 1

// UserTypeUser 员工
const UserTypeUser = 0

// UserTypeCar 车辆
const UserTypeCar = 3

// UserVisitorTimeLen 访客时间长度
const UserVisitorTimeLen = 2

// UserImportExcelLen 用户导入长度
const UserImportExcelLen = 6

// CensusTypeFence 围栏统计类型
const CensusTypeFence = "fence"

// CensusTypeRegion 区域统计类型
const CensusTypeRegion = "region"

// LocationChargingTime 标签充电超时时间
const LocationChargingTime = 10 * 1000

// LocationChargingNum 正在充电
const LocationChargingNum = 1

// ProjectTunnelType 隧道类型项目
const ProjectTunnelType = 1

// ProjectConfigMapMin 项目地图配置长度
const ProjectConfigMapMin = 1

// FenceTagTypeExclude 围栏标签类型【排除】
const FenceTagTypeExclude = 1

// FenceTagTypeInclude 围栏标签类型【包括】
const FenceTagTypeInclude = 0

// FenceCacheTypeInto 围栏缓存【进入】
const FenceCacheTypeInto = "into"

// FenceCacheTypeOut 围栏缓存【出去】
const FenceCacheTypeOut = "out"

// FenceTypeJob 作业围栏
const FenceTypeJob = 1
