// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

/*
UFileBucketSet - DescribeBucket

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type UFileBucketSet struct {

	// Bucket所属业务, general或vod或udb general: 普通业务； vod: 视频云业务; udb: 云数据库业务
	Biz string

	// Bucket的ID
	BucketId string

	// Bucket名称
	BucketName string

	// 与Bucket关联的CND加速域名的ID列表
	CdnDomainId []string

	// Bucket的创建时间
	CreateTime int

	// Bucket的域名集合 参数见 UFileDomainSet
	Domain UFileDomainSet

	// 是否存在自定义域名。0不存在，1存在，2错误
	HasUserDomain int

	// Bucket的修改时间
	ModifyTime int

	// Bucket所属地域
	Region string

	// 所属业务组
	Tag string

	// Bucket访问类型
	Type string
}