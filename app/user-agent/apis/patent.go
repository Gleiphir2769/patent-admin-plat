package apis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"net/http"
)

type Patent struct {
	api.Api
}

// GetPatentById
// @Summary 通过专利id获取单个对象
// @Description 获取JSON,希望可以通过以下参数高级搜索，暂时只支持patentId
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/{patent_id} [get]
// @Security Bearer
func (e Patent) GetPatentById(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Patent
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// GetPatentLists
// @Summary 列表专利信息数据
// @Description 获取JSON
// @Tags 专利表
// @Router /api/v1/user-agent/patent [get]
// @Security Bearer
func (e Patent) GetPatentLists(c *gin.Context) { //gin框架里的上下文
	s := service.Patent{}         //service中查询或者返回的结果赋值给s变量
	req := dto.PatentGetPageReq{} //被绑定的数据
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]models.Patent, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// InsertPatent
// @Summary 创建专利
// @Description 不是必须要有主键PatentId值（自增），其他需要修改什么输入什么
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "专利表数据"
// @Router /api/v1/user-agent/patent [post]
// @Security Bearer
func (e Patent) InsertPatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetPatentId(), "创建成功")
}

// UpdatePatent
// @Summary 修改专利表数据
// @Description 在post的json数组必须要有主键PatentId值（默认0不可重复），其他需要修改什么输入什么
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentUpdateReq true "body"
// @Router /api/v1/user-agent/patent [put]
// @Security Bearer
func (e Patent) UpdatePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.UpdateLists(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetPatentId(), "更新成功")
}

// DeletePatentByPatentId
// @Summary 输入专利id删除专利表
// @Description  输入专利id删除专利表
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeletePatentByPatentId(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentById{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetPatentId(), "删除成功")
}

// ClaimPatent
// @Summary 认领专利
// @Description 认领专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/claim [post]
// @Security Bearer
func (e Patent) ClaimPatent(c *gin.Context) {

	pid, err := e.internalInsertIfAbsent(c)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.UserPatent{}
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentClaim(user.GetUserId(c), pid, user.GetUserId(c))

	if err = s.Insert(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "认领成功")
}

// FocusPatent
// @Summary 关注专利
// @Description 关注专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/focus [post]
// @Security Bearer
func (e Patent) FocusPatent(c *gin.Context) {

	pid, err := e.internalInsertIfAbsent(c)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.UserPatent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), pid, user.GetUserId(c))

	if err = s.Insert(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "关注成功")
}

func (e Patent) internalInsertIfAbsent(c *gin.Context) (int, error) {
	ps := service.Patent{}
	req := dto.PatentInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		return 0, err
	}
	return ps.InsertIfAbsent(&req)
}

// GetFocusPages
// @Summary 获取关注列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/focus [get]
// @Security Bearer
func (e Patent) GetFocusPages(c *gin.Context) {
	s := service.UserPatent{}         //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req.UserId = user.GetUserId(c)
	req1 := dto.PatentsByIdsForRelationshipUsers{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	list := make([]models.UserPatent, 0)
	list1 := make([]models.Patent, 0)
	var count int64
	err = s.GetClaimLists(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	var count2 int64
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors
	req1.PatentIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}
	err = s.GetPatentPagesByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")
}

// GetClaimPages
// @Summary 获取认领列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/claim [get]
// @Security Bearer
func (e Patent) GetClaimPages(c *gin.Context) {
	s := service.UserPatent{}         //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req1 := dto.PatentsByIdsForRelationshipUsers{}

	req.UserId = user.GetUserId(c)

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	list := make([]models.UserPatent, 0)
	list1 := make([]models.Patent, 0)

	var count int64

	err = s.GetCollectionLists(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors

	req1.PatentIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s.GetPatentPagesByIds(&req1, &list1, &count2)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list1, "查询成功")
}

//// InsertUserPatentRelationship
//// @Summary 创建用户专利关系
//// @Description Type和PatentId为必要输入，Type只能是 认领 或者 关注 关系
//// @Tags 用户专利关系表
//// @Accept  application/json
//// @Product application/json
//// @Param data body dto.UserPatentInsertReq true "Type和PatentId为必要输入"
//// @Router /api/v1/user-patent/ [post]
//// @Security Bearer
//func (e UserPatent) InsertUserPatentRelationship(c *gin.Context) {
//	s := service.UserPatent{}
//	req := dto.UserPatentInsertReq{}
//	req.UserId = user.GetUserId(c)
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, binding.JSON).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	// 设置创建人
//	req.SetCreateBy(user.GetUserId(c))
//
//	if req.Type == "认领" {
//		err = s.Insert(&req)
//	} else if req.Type == "关注" {
//		err = s.InsertCollectionRelationship(&req)
//	} else {
//		e.Logger.Error(err)
//		e.Error(404, err, fmt.Sprintf("invalid req.Type: %s, should be 认领/关注", req.Type))
//		return
//	}
//
//	if req.PatentId == 0 {
//		e.Logger.Error(err)
//		e.Error(404, err, "您输入的专利id不存在！")
//		return
//	}
//
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	e.OK(req, "创建成功")
//}

// DeleteFocus
// @Summary 取消关注
// @Description  取消关注
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/focus/{patent_id}  [delete]
// @Security Bearer
func (e Patent) DeleteFocus(c *gin.Context) {
	s := service.Patent{}
	req := dto.UserPatentObject{}
	req.UserId = user.GetUserId(c)

	req.SetUpdateBy(user.GetUserId(c))

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req). //在这一步传入request数据
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveFocus(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "删除成功")
}

// DeleteClaim
// @Summary 取消认领
// @Description  取消认领
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/claim/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeleteClaim(c *gin.Context) {
	s := service.Patent{}
	req := dto.UserPatentObject{}
	req.UserId = user.GetUserId(c)

	req.SetUpdateBy(user.GetUserId(c))

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req). //在这一步传入request数据
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveClaim(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "删除成功")
}

// UpdateUserPatentRelationship
// @Summary 修改用户专利关系
// @Description 需要输入专利id
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpDateUserPatentObject true "body"
// @Router /api/v1/user-agent/patent [put]
// @Security Bearer
func (e Patent) UpdateUserPatentRelationship(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.UpDateUserPatentObject{}
	req.UserId = user.GetUserId(c)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	if req.PatentId == 0 {
		e.Logger.Error(err)
		e.Error(404, err, "请输入专利id")
		return
	}

	err = s.UpdateUserPatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "更新成功")
}
