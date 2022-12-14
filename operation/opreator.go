package operation

import (
	"fmt"
	"syncData/global"
	"syncData/model"
)

func TransferData() {
	updatePersonalOfAsset()
	updateOrgOfAsset()
	updateOrgAsset()
}

/*
UPDATE `gocyber_yezhenghan`.`sys_asset` sa,
(SELECT asset_id,account_id FROM `gocyber_yezhenghan`.`sys_account_asset`) saa SET sa.creator_account_id = saa.account_id WHERE sa.id = saa.asset_id;
*/
// 更新asset中个人部分
func updatePersonalOfAsset() {
	//sysAccountAssaets := []*model.SysAccountAsset{}
	sysAssets := []*model.SysAsset{}

	global.DB.Find(&sysAssets)

	if err := global.DB.Model(&sysAssets).Update("creator_account_id",
		global.DB.Model(&model.SysAccountAsset{}).Select("account_id").Where("sys_asset.id = sys_account_asset.asset_id")).Error; err != nil {
		fmt.Println("first stage failed")
	}

}

/*
UPDATE `gocyber_yezhenghan`.`sys_asset` sa,
(SELECT `gocyber_yezhenghan`.`sys_org_asset`.asset_id, `gocyber_yezhenghan`.`sys_account`.id  FROM `gocyber_yezhenghan`.`sys_org_asset`
LEFT JOIN `gocyber_yezhenghan`.`sys_asset` ON `gocyber_yezhenghan`.`sys_org_asset`.asset_id =`gocyber_yezhenghan`.`sys_asset`.id
LEFT JOIN `gocyber_yezhenghan`.`sys_account` ON `gocyber_yezhenghan`.`sys_org_asset`.org_id =`gocyber_yezhenghan`.`sys_account`.org_id AND `gocyber_yezhenghan`.`sys_asset`.creator_user_id =`gocyber_yezhenghan`.`sys_account`.user_id) soa
SET sa.creator_account_id = soa.id WHERE sa.id = soa.asset_id
*/
// 更新asset中企业部分
func updateOrgOfAsset() {

	if err := global.DB.Exec(
		`UPDATE sys_asset sa,
		(SELECT sys_org_asset.asset_id, sys_account.id  FROM sys_org_asset
		LEFT JOIN sys_asset ON sys_org_asset.asset_id =sys_asset.id
		LEFT JOIN sys_account ON sys_org_asset.org_id = sys_account.org_id AND sys_asset.creator_user_id =sys_account.user_id) soa
		SET sa.creator_account_id = soa.id WHERE sa.id = soa.asset_id`).Error; err != nil {

		fmt.Println("second stage failed")
	}
}

/*
UPDATE `gocyber_yezhenghan`.`sys_org_asset` sa,
(SELECT id,creator_account_id FROM `gocyber_yezhenghan`.`sys_asset`) saa SET sa.account_id = saa.creator_account_id  WHERE sa.creator_account_id = saa.id;
*/
// 把asset表中org部分的accountid补充回org
func updateOrgAsset() {
	sysOrgAssets := []*model.SysOrgAsset{}

	global.DB.Find(&sysOrgAssets)

	if err := global.DB.Model(&sysOrgAssets).Update("account_id",
		global.DB.Model(&model.SysAsset{}).Select("creator_account_id").Where("sys_org_asset.asset_id = sys_asset.id")).Error; err != nil {
		fmt.Println("third stage failed")
	}
}
