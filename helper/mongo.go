package helper

import (
   
   
)

// 得到customer数据
func GetCustomerDbName() string {
    return "trace_customer"
}
// 得到customer表的数据
func GetCustomerCollName(websiteId string) string {
    return "customer" + "_" + websiteId
}
func GetOutCustomerEmailCollName(websiteId string) string {
    return "customer_email" + "_" + websiteId
}

// 得到当前时间对应的 mongodb   的 dbName
// nowDateStr 格式：  2009-12-12
func GetDbName(dbName string, nowDateStr string) (string){
    if nowDateStr == "" {
        nowDateStr = DateUTCStr()
    }
    return dbName + "_" + nowDateStr
}
// 得到当前时间对应的 mongodb   的 collection Name
func GetCollName(collName string, websiteId string) string {
    return collName + "_" + websiteId
}


// 初始数据接收的数据库
func GetTraceDbName() (string){
    return GetDbName("trace", "")
}

// dateStr 格式：  2009-12-12 
// 通过时间，得到相应时间的库
func GetTraceDbNameByDate(dateStr string) string {
    return GetDbName("trace", dateStr)
}

func GetTraceDataCollName(websiteId string) (string){
    return GetCollName("trace_data", websiteId)
}

// 得到Browser统计后的数据输出的collection
func GetOutWholeBrowserCollName(websiteId string) (string){
    return GetCollName("trace_whole_browser_data", websiteId)
}

// 得到Browser统计后的数据输出的collection
func GetOutWholeAllCollName(websiteId string) (string){
    return GetCollName("trace_whole_all_data", websiteId)
}

// 得到Refer统计后的数据输出的collection
func GetOutWholeReferCollName(websiteId string) (string){
    return GetCollName("trace_whole_refer_data", websiteId)
}

// 得到 country 统计后的数据输出的collection
func GetOutWholeCountryCollName(websiteId string) (string){
    return GetCollName("trace_whole_country_data", websiteId)
}

// 得到 devide 统计后的数据输出的collection
func GetOutWholeDevideCollName(websiteId string) (string){
    return GetCollName("trace_whole_devide_data", websiteId)
}

// 得到 sku 统计后的数据输出的collection
func GetOutWholeSkuCollName(websiteId string) (string){
    return GetCollName("trace_whole_sku_data", websiteId)
}

// 得到 sku refer 统计后的数据输出的collection
func GetOutWholeSkuReferCollName(websiteId string) (string){
    return GetCollName("trace_whole_sku_refer_data", websiteId)
}

// 得到 search 统计后的数据输出的collection
func GetOutWholeSearchCollName(websiteId string) (string){
    return GetCollName("trace_whole_search_data", websiteId)
}

// 得到 search lang 统计后的数据输出的collection
func GetOutWholeSearchLangCollName(websiteId string) (string){
    return GetCollName("trace_whole_search_lang_data", websiteId)
}

// 得到 Url 统计后的数据输出的collection
func GetOutWholeUrlCollName(websiteId string) (string){
    return GetCollName("trace_whole_url_data", websiteId)
}
// 得到 First Url 统计后的数据输出的collection
func GetOutWholeFirstUrlCollName(websiteId string) (string){
    return GetCollName("trace_whole_first_url_data", websiteId)
}

// 得到 Category 统计后的数据输出的collection
func GetOutWholeCategoryCollName(websiteId string) (string){
    return GetCollName("trace_whole_category_data", websiteId)
}


// 得到 App 统计后的数据输出的collection
func GetOutWholeAppCollName(websiteId string) (string){
    return GetCollName("trace_whole_app_data", websiteId)
}

// 得到 Advertise Fid 统计后的数据输出的collection
func GetOutAdvertiseFidCollName(websiteId string) (string){
    return GetCollName("trace_advertise_fid_data", websiteId)
}

// 得到 Advertise Content 统计后的数据输出的collection
func GetOutAdvertiseContentCollName(websiteId string) (string){
    return GetCollName("trace_advertise_content_data", websiteId)
}

// 得到 Advertise MarketGroup 统计后的数据输出的collection
func GetOutAdvertiseMarketGroupCollName(websiteId string) (string){
    return GetCollName("trace_advertise_market_group_data", websiteId)
}

// 得到 Advertise Design 统计后的数据输出的 collection
func GetOutAdvertiseDesignCollName(websiteId string) (string){
    return GetCollName("trace_advertise_design_data", websiteId)
}

// 得到 Advertise Campaign 统计后的数据输出的 collection
func GetOutAdvertiseCampaignCollName(websiteId string) (string){
    return GetCollName("trace_advertise_campaign_data", websiteId)
}


// 得到 Advertise Medium 统计后的数据输出的 collection
func GetOutAdvertiseMediumCollName(websiteId string) (string){
    return GetCollName("trace_advertise_medium_data", websiteId)
}


// 得到 Advertise source 统计后的数据输出的 collection
func GetOutAdvertiseSourceCollName(websiteId string) (string){
    return GetCollName("trace_advertise_source_data", websiteId)
}


// 得到 Advertise edm 统计后的数据输出的 collection
func GetOutAdvertiseEdmCollName(websiteId string) (string){
    return GetCollName("trace_advertise_edm_data", websiteId)
}

// 得到 customer uuid 统计后的数据输出的 collection
func GetOutCustomerUuidCollName(websiteId string) (string){
    return GetCollName("trace_customer_uuid_data", websiteId)
}

// 得到 Advertise Eid 统计后的数据输出的collection
func GetOutAdvertiseEidCollName(websiteId string) (string){
    return GetCollName("trace_advertise_eid_data", websiteId)
}















