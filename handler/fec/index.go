package fec

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
)



type TraceInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    Uuid string `binding:"required" form:"uuid" json:"uuid" bson:"uuid"`
    WebsiteId string `binding:"required" form:"website_id" json:"website_id" bson:"website_id"` 
    Devide string `form:"devide" json:"devide" bson:"devide"`
    UserAgent string `form:"user_agent" json:"user_agent" bson:"user_agent"`
    BrowserName string `form:"browser_name" json:"browser_name" bson:"browser_name"`
    BrowserVersion string `form:"browser_version" json:"browser_version" bson:"browser_version"`
    BrowserDate string `form:"browser_date" json:"browser_date" bson:"browser_date"`
    BrowserLang string `form:"browser_lang" json:"browser_lang" bson:"browser_lang"`
    Operate string `form:"operate" json:"operate" bson:"operate"`
    OperateRelase string `form:"operate_relase" json:"operate_relase" bson:"operate_relase"`
    Url string `form:"url" json:"url" bson:"url"`
    Domain string `form:"domain" json:"domain" bson:"domain"`
    Title string `form:"title" json:"title" bson:"title"`
    ReferUrl string `form:"refer_url" json:"refer_url" bson:"refer_url"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    FirstReferrerUrl string `form:"first_referrer_url" json:"first_referrer_url" bson:"first_referrer_url"`
    ClActivity string `form:"cl_activity" json:"cl_activity" bson:"cl_activity"`
    ClActivity_child string `form:"icl_activity_childd" json:"cl_activity_child" bson:"cl_activity_child"`
    IsReturn string `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage string `form:"first_page" json:"first_page" bson:"first_page"`
    DevicePixelRatio string `form:"device_pixel_ratio" json:"device_pixel_ratio" bson:"device_pixel_ratio"`
    Resolution string `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth string `form:"color_depth" json:"color_depth" bson:"color_depth"`
    
}
func (traceInfo TraceInfo) TableName() string {
    return "trace_info"
}

func SaveJsData(c *gin.Context){
    var traceInfo TraceInfo
    err := c.ShouldBindQuery(&traceInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    err = mongodb.MC(traceInfo.TableName(), func(coll *mgo.Collection) error {
        // c.Find(bson.M{"_id": id}).One(traceInfo)
        traceInfo.Id_ = bson.NewObjectId()
        err := coll.Insert(traceInfo)
        return err
    })
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "traceInfo": traceInfo,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}