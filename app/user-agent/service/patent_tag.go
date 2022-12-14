package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
)

type PatentTag struct {
	service.Service
}

// GetTagIdByPatentId 通过PatentId获得TagId
func (e *PatentTag) GetTagIdByPatentId(c *dto.PatentTagGetPageReq, list *[]models.PatentTag, count *int64) error {
	var err error
	var data models.PatentTag

	err = e.Orm.Model(&data).
		Where("Patent_Id = ?", c.GetPatentId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetTagPages 通过TagId获取Tag列表（TagName等）
func (e *PatentTag) GetTagPages(d *dto.TagsByIdsForRelationshipPatents, list *[]models.Tag, count *int64) error {

	var err error
	var ids []int = d.GetTagId()

	for i := 0; i < len(ids); i++ {

		if ids[i] != 0 {

			var data1 models.Tag

			err = e.Orm.Model(&data1).
				Where("Tag_Id = ? ", ids[i]).
				First(&data1).Limit(-1).Offset(-1).
				Count(count).Error

			*list = append(*list, data1)

			if err != nil {
				e.Log.Errorf("db error:%s", err)
				return err
			}
		}
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentIdByTagId 通过TagId获得PatentId
func (e *PatentTag) GetPatentIdByTagId(c *dto.PatentTagGetPageReq, list *[]models.PatentTag, count *int64) error {
	var err error
	var data models.PatentTag

	err = e.Orm.Model(&data).
		Where("Tag_Id = ?", c.GetTagId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertPatentTagRelationship 创建专利标签关系
func (e *PatentTag) InsertPatentTagRelationship(c *dto.PatentTagInsertReq) error {
	var err error
	var data models.PatentTag
	var i int64
	err = e.Orm.Model(&data).Where("Patent_Id = ? AND Tag_Id = ? ", c.PatentId, c.TagId).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GeneratePatentTag(&data)

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// RemoveRelationship 根据Patent_Id、 Tag_Id删除标签专利关系
func (e *Patent) RemoveRelationship(c *dto.PatentTagInsertReq) error {
	var err error
	var data models.PatentTag

	db := e.Orm.Where("Patent_Id = ? AND Tag_Id = ? ", c.PatentId, c.TagId).
		Delete(&data)

	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
