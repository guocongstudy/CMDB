package host

const (
	ProvateIDC Vendor = iota
	Tencent
	Aliyun
	HuaWei
)

//用int做枚举
type Vendor int

type Host struct {
	*Base
	*Resource
	*Describe
}

type Base struct {
	Id           string `json:id`              //全局唯一Id
	SyncAt       int64  `json:"sync_at"`       //同步时间
	Vendor       Vendor `json:"vendor"`        //厂商
	Region       string `json:"region"`        //地域
	Zone         string `json:"zone"`          //区域
	CreateAt     int64  `json:"create_at"`     //创建时间
	InstanceId   string `json:"instance_id"`   //实例ID
	ResourceId   string `json:"resource_id"`   //基础数据Hash
	DescribeHash string `json:"describe_hash"` //描述数据Hash
}

type Resource struct {
	ExpireAt    int64             `json:"expire_at"`    //过期时间
	Category    string            `json:"category"`     //种类
	Type        string            `json:"type"`         //规格
	Name        string            `json:"name"`         //名称
	Description string            `json:"description"`  //描述
	Status      string            `json:"status"`       //服务商中的状态
	Tags        map[string]string `json:"tags"`         //标签
	UpdateAt    int64             `json:"update_at"`    //更新时间
	SyncAccount string            `json:"sync_account"` //同步账号
	PublicIP    string            `json:"public_ip"`    //公网IP
	PrivateIP   string            `json:"private_ip"`   //私网IP
	PayType     string            `json:"pay_type"`     //实例付费方式
}

type Describe struct {
	ResourceId string `json:"resource_id"`
	CPU int `json:"cpu"`
	Memory int `json:"memory"`
	GPUAmount int `json:"gpu_amount"`
	GPUSpec string `json:"gpu_spec"`
	OSType string `json:"os_type"`
	OSName string `json:"os_name"`
	SerialNumber string `json:"serial_number"`
	ImageID string `json:"image_id"`
	InternetMaxBandwidthOut int `json:"internet_max_bandwidth_out"`
	InternetMaxBandwidthIn int `json:"internet_max_bandwidth_in"`
	KeyPairName string `json:"key_pair_name"`
	SecurityGroups string `json:"security_groups"`
}