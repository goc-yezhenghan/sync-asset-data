package model

import (
	"time"

	"gorm.io/gorm"
)

type SysAsset struct {
	gorm.Model
	Name                    string `json:"name" gorm:"size:64;comment:资产名称"`
	ContentPath             string `json:"contentUrl" gorm:"size:255;comment:文件路径"`
	CreatorUserID           uint   `json:"creatorUserId" gorm:"comment:创建者用户ID"`
	CreatorAccountId        uint   `json:"creatorAccountId" gorm:"comment:创建者账号ID"`
	ThumbnailPath           string `json:"thumbnailPath" gorm:"size:255;comment:资产缩略图"`
	OriginalContentPath     string `json:"originalContentPath" gorm:"size:255;comment:文件原始路径,仅当文件为glb文件时发离线任务请求"`
	OriginalContentFileSize uint64 `json:"originalContentFileSize" gorm:"comment:原始文件大小,单位字节"`
	Type                    string `json:"type" gorm:"size:255;comment:资产类型"`
	ExtensionPropertyUrl    string `json:"extensionPropertyUrl" gorm:"size:255;comment:扩展属性 Mongo url"`
	PipelineStatus          string `json:"pipelineStatus" gorm:"size:64;comment:流水线处理状态"`
	PipelineStatusUpdateAt  time.Time
}

type SysOrgAsset struct {
	gorm.Model
	ClientAssetID string `json:"clientAssetId" gorm:"size:64;comment:第三方组织资产ID"`
	AccountId     uint   `json:"accountId" gorm:"comment:用户ID"`
	OrgID         uint   `json:"orgId" gorm:"uniqueIndex:uk_org_asset"`
	AssetID       uint   `json:"assetId" gorm:"uniqueIndex:uk_org_asset"`
}

type SysAccountAsset struct {
	gorm.Model
	AccountID uint `json:"accountId" gorm:"uniqueIndex:uk_account_asset"`
	AssetID   uint `json:"assetId" gorm:"uniqueIndex:uk_account_asset"`
}

type SysAccount struct {
	gorm.Model
	UserID         uint   `json:"userId" gorm:"uniqueIndex:uk_account"`
	OrgID          uint   `json:"orgId" gorm:"uniqueIndex:uk_account; comment:当个人账户时orgID为0"`
	Type           string `json:"type" gorm:"size:16;comment:personal/org_member/org_customer"`
	DisplayName    string `json:"displayName" gorm:"size:64;comment:账户昵称"`
	ProfilePicUrl  string `json:"profilePicUrl" gorm:"size:255;comment:账户头像"`
	AvatarModelUrl string `json:"avatarModelUrl" gorm:"size:255;comment:形象模型"`
	Origin         string `json:"origin" gorm:"size:64;comment:账号来源"`
}
