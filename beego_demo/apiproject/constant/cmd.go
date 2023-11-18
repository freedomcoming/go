package constant

// CmdBaseIdAdd 添加关联基站命令前缀
const CmdBaseIdAdd = "matlist -a "

// CmdBaseIdDel 删除关联基站命令前缀
const CmdBaseIdDel = "matlist -d "

// CmdBaseIdUpdate 修改关联基站命令前缀
const CmdBaseIdUpdate = "matlist -m "

// DownsideSuffix 下行命令通用后缀
const DownsideSuffix = "\r\n"

// CmdBaseTypeAdd 添加
const CmdBaseTypeAdd = "add"

// CmdBaseTypeDel 删除
const CmdBaseTypeDel = "del"

// CmdBaseTypeUpdate 修改
const CmdBaseTypeUpdate = "update"

// CmdMRangName 基站测距命令
const CmdMRangName = "mrang -c80 -d"

// WarnTagRun 标签报警命令
const WarnTagRun = "MSGMGT -wto10 -wset"

// WarnTagStop 停止报警命令
const WarnTagStop = "MSGMGT -wcleanset"

const WarnTagKeyName = "tag_warn_"

// WarnTagTimeMin 报警超时其实时间
const WarnTagTimeMin = 10 * 1000

// WarnTagTimeMax 报警超时截止时间
const WarnTagTimeMax = 60 * 1000

// DiscDataTimeOut 测距信息超时时间
const DiscDataTimeOut = 10 * 60 * 1000

// WorksTimeMin 工时统计时间门限
const WorksTimeMin = 5 * 60

// CarSafeL 台车安全步距门限
const CarSafeL = 200

// CarSafeStart 台车安全步距起始阶段
const CarSafeStart = 0

// CarSafeEnd 台车安全步距截止阶段
const CarSafeEnd = 20000

// CarSafeLevel 台车安全步距等级
const CarSafeLevel = 1
