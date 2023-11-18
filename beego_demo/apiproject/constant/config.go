package constant

// RegionPointAlias 匹配小区边界别名
const RegionPointAlias = "matchRegionPoint"

// RegionPointTitle 匹配小区边界标题
const RegionPointTitle = "是否匹配小区边界上的最近点位[0:不匹配; 1:匹配]"

// PointMinDistAlias 匹配小区边界上最近点位的最小距离别名
const PointMinDistAlias = "matchPointMinDist"

// PointMinDistTitle 匹配小区边界上最近点位的最小距离标题
const PointMinDistTitle = "匹配小区边界上最近点位的最小距离[2m]"

// CollectTidAlias 数据采集标签ID
const CollectTidAlias = "collectTid"

// CollectTidTitle 数据采集标签ID
const CollectTidTitle = "数据采集标签ID,英文逗号隔开"

// TrackDayAlias 轨迹数据默认保存时间
const TrackDayAlias = "trackDay"

// TrackDayTitle 历史轨迹默认保存天数
const TrackDayTitle = "历史轨迹默认保存天数"

// MdnsServiceIpAlias mdns服务IP地址
const MdnsServiceIpAlias = "MdnsIp"

// MdnsServiceIpTitle 基站自动上线服务IP地址
const MdnsServiceIpTitle = "基站自动上线服务IP地址"

// TcpPushAddrAlias tcp定位推送服务地址
const TcpPushAddrAlias = "TcpAddr"

// TcpPushAddrTitle tcp定位推送服务地址
const TcpPushAddrTitle = "tcp数据推送服务地址"

// LocationPushStatusAlias 定位推送开关
const LocationPushStatusAlias = "PushStatus"

// LocationPushStatusTitle 定位推送开关[1: 开启]
const LocationPushStatusTitle = "定位推送开关[1: 开启]"

// BeaconRssiAlias beacon信号质量门限
const BeaconRssiAlias = "BeaconRssi"

// BeaconRssiTitle beacon信号质量门限
const BeaconRssiTitle = "beacon信号质量门限"

// BaseDiscDistSAlias 基站自动测距标准差
const BaseDiscDistSAlias = "baseDistS"

// BaseDiscDistSTitle 基站自动测距标准差
const BaseDiscDistSTitle = "基站自动测距标准差"

// HeightParamsAlias 定位人员高度
const HeightParamsAlias = "heightParams"

// HeightParamsTitle 定位人员高度
const HeightParamsTitle = "高度【status:动态人员高度, height:人员高度】"

// SpecialParamsAlias 巡检点位参数
const SpecialParamsAlias = "specialParams"

// SpecialParamsTitle 巡检点位参数
const SpecialParamsTitle = "巡检点位参数【rss:信号门限, time:延时统计时间[s], num: 最低时长[s]】"

// LocationPushTidAlias 数据推送标签ID
const LocationPushTidAlias = "pushTid"

// LocationPushTidTitle 数据推送标签ID
const LocationPushTidTitle = "数据推送标签ID"

// DepartTypeAlias 部门分类
const DepartTypeAlias = "departType"

// DepartTypeTitle 部门分类
const DepartTypeTitle = "部门分类"

// DevicePointAlias 点位同步配置
const DevicePointAlias = "devicePoint"

// DevicePointTitle 点位同步配置
const DevicePointTitle = "点位同步配置"

// LocationSubGAlignAlias subG数据缓存时间
const LocationSubGAlignAlias = "subGAlignTime"

// LocationSubGAlignTitle subG数据缓存时间
const LocationSubGAlignTitle = "subG数据缓存时间【ms】"

// KalmanParamsAlias 卡尔曼滤波参数
const KalmanParamsAlias = "kalmanParams"

// KalmanParamsTitle 卡尔曼滤波参数
const KalmanParamsTitle = "滤波参数[Q:噪声,R:频率,N:次数,T:时间]"

// SubParamsAlias sub定位参数
const SubParamsAlias = "SubParams"

// SubParamsTitle sub定位参数
const SubParamsTitle = "sub参数[d:最小跨层距离,num:切换数量,val:气压冗余,r:信号门限]"

// LocationSearchParamsAlias 定位搜索参数
const LocationSearchParamsAlias = "searchParams"

// LocationSearchParamsTitle 定位搜索参数
const LocationSearchParamsTitle = "搜索参数[step:步长,first:首次距离,dist:搜索范围]"

// SubUwbBleAlias sub融合定位参数
const SubUwbBleAlias = "subUwbBle"

// SubUwbBleTitle sub融合定位参数
const SubUwbBleTitle = "sub融合定位参数"

// VoltageAlias 电压范围参数
const VoltageAlias = "voltage"

// VoltageTitle 电压范围参数
const VoltageTitle = "电压范围参数"

// PushDataTimeAlias 数据推送间隔
const PushDataTimeAlias = "pushTime"

// PushDataTimeTitle 数据推送间隔
const PushDataTimeTitle = "数据推送间隔[ms]"

// SearchTdoAMaxSearchDist  最大搜索距离
const SearchTdoAMaxSearchDist = 10.0

// SearchTofMaxSearchDist Tof最大搜索距离
const SearchTofMaxSearchDist = 10.0

// SearchStep 搜索步长
const SearchStep = 0.1

// SearchFirstMaxDist 首次搜索最大距离
const SearchFirstMaxDist = 1000.0

// BorderOutCountAlias 超出小区边界计数
const BorderOutCountAlias = "borderOutCount"

// BorderOutCountTitle 超出小区边界计数
const BorderOutCountTitle = "超出小区边界计数[次数]"

// SubFloorSwitchNumMax 蓝牙楼层切换最大数量
const SubFloorSwitchNumMax = 5

// SubFloorSwitchDismMin 蓝牙楼层切换最小距离
const SubFloorSwitchDismMin = -1

// BleSubGParamsNum subG参数
const BleSubGParamsNum = 1.5

const (
	SubRssiThreshold1 = -65 //信号强度门限
	SubRssiThreshold2 = -80 //信号强度门限
	SubRssiThreshold3 = -95 //信号强度门限
)

// LocationHeightDefault 默认定位人员高度
const LocationHeightDefault = 1.2

// MdnsStatusOff mdns开启状态
const MdnsStatusOff = 0

// SpecialRssiMin 点位巡检最低信号门限
const SpecialRssiMin = -60

// SpecialTimeMax 点位巡检最大延时时间
const SpecialTimeMax = 60

// SpecialNumMin 点位巡检最小时长
const SpecialNumMin = 60

// DevicePointIp 设备点位地址
const DevicePointIp = "139.9.244.124:30002"

// SubMainUwb sub主定位
const SubMainUwb = "uwb"

// SubOutTime sub超时时间
const SubOutTime = 3 * 1000

// LocationDistMax 定位最大距离
const LocationDistMax = 50

// VoltageMinDefault 电压默认最小值
const VoltageMinDefault = 6.0

// VoltageMaxDefault 电压默认最大值
const VoltageMaxDefault = 8.4

// BorderOutCount 超出小区边界计数
const BorderOutCount = 3
