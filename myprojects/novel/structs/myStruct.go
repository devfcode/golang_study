package structs

import "gopkg.in/mgo.v2/bson"

//部分操作完成后的返回结构
type Tip struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

//返回数据的一般格式
type ResultList struct {
	Code int `json:"code"`
	Msg string `json:"msg"`

	Content Content `json:"content"`
}
type Content struct {
	Header Header `json:"header"`
	Other Other `json:"other"`
	Data []interface{}  `json:"data"`
}
type Header struct {
	Title string `json:"title"`
	Token string `json:"token"`
}
type Other struct {
	Total int `json:"total"`
	PageIndex int64 `json:"pageIndex"`
	PageSize int `json:"pageSize"`
}


//后台账号 添加数据的结构
type BackAccount struct {
	Account string `bson:"account" json:"account"`  //账户
	Password string `bson:"password" json:"password"` //密码
	Remarks string `bson:"remarks" json:"remarks"`  //备注
	AccountManagerHomepage bool `bson:"account_manager_homepage" json:"account_manager_homepage"` //用户管理 首页数据

	AccountManagerBlackInfoSee bool `bson:"account_manager_black_info_see" json:"account_manager_black_info_see"` //用户管理 黑名单信息管理 查看
	AccountManagerBlackUnseal bool `bson:"account_manager_black_unseal" json:"account_manager_black_unseal"` //用户管理 用户管理 解封
	AccountManagerUserlist bool `bson:"account_manager_userlist" json:"account_manager_userlist"` //用户管理 玩家列表
	AccountManagerUserinfoSee bool `bson:"account_manager_userinfo_see" json:"account_manager_userinfo_see"` //用户管理 用户信息 查看
	AccountManagerUserinfoOffline bool `bson:"account_manager_userinfo_offline" json:"account_manager_userinfo_offline"` //用户管理 用户信息 踢下线
	AccountManagerUserinfoSeal bool `bson:"account_manager_userinfo_seal" json:"account_manager_userinfo_seal"` //用户管理 用户信息 封号
	AccountManagerUserinfoLastloginip bool `bson:"account_manager_userinfo_lastloginip" json:"account_manager_userinfo_lastloginip"` //用户管理 用户信息 最后登陆ip
	AccountManagerUserinfoLastloginmapaddr bool `bson:"account_manager_userinfo_lastloginmapaddr" json:"account_manager_userinfo_lastloginmapaddr"`//用户管理 用户信息 最后登录MAC地址
	AccountManagerUserinfoMacaddr bool `bson:"account_manager_userinfo_macaddr" json:"account_manager_userinfo_macaddr"` //用户管理 用户信息 MAC地址
	AccountManagerUserinfoNickname bool `bson:"account_manager_userinfo_nickname" json:"account_manager_userinfo_nickname"`//用户管理 用户信息 玩家呢称
	AccountManagerReaddetail bool `bson:"account_manager_readdetail" json:"account_manager_readdetail"` //用户管理 阅卷明细
	AccountManagerCardgamerecord bool `bson:"account_manager_cardgamerecord" json:"account_manager_cardgamerecord"` //用户管理 牌局记录
	AccountManagerRechargerecord bool `bson:"account_manager_rechargerecord" json:"account_manager_rechargerecord"`//用户管理 充值记录

	RechargeBlockPlayerRecordQuery bool `bson:"recharge_block_player_record_query" json:"recharge_block_player_record_query"`//充值模块 玩家订单查询
	RechargeBlockRechargeDataAnalysis bool `bson:"recharge_block_recharge_data_analysis" json:"recharge_block_recharge_data_analysis"`//充值模块 充值数据分析
	RechargeBlockMonthRechargeCount bool `bson:"recharge_block_month_recharge_count" json:"recharge_block_month_recharge_count"`//充值模块 月充值统计
	RechargeBlockPlayerRechargeRangeSee bool `bson:"recharge_block_player_recharge_range_see" json:"recharge_block_player_recharge_range_see"`//充值模块 玩家充值排行 查看
	RechargeBlockPlayerRechargeRangeOutexcel bool `bson:"recharge_block_player_recharge_range_outexcel" json:"recharge_block_player_recharge_range_outexcel"`//充值模块 玩家充值排行 Excel导出
	RechargeBlockBigPlayerRechargeRange bool `bson:"recharge_block_big_player_recharge_range" json:"recharge_block_big_player_recharge_range"`//充值模块 大玩家充值排行
	RechargeBlockChannelManagerSee bool `bson:"recharge_block_channel_manager_see" json:"recharge_block_channel_manager_see"` //充值模块 充值渠道管理 查看
	RechargeBlockChannelManagerAdd bool `bson:"recharge_block_channel_manager_add" json:"recharge_block_channel_manager_add"` //充值模块 充值渠道管理 新增
	RechargeBlockChannelManagerDelete bool `bson:"recharge_block_channel_manager_delete" json:"recharge_block_channel_manager_delete"`//充值模块 充值渠道管理 删除
	RechargeBlockChannelManagerUpdate bool `bson:"recharge_block_channel_manager_update" json:"recharge_block_channel_manager_update"`//充值模块 充值渠道管理 修改
	RechargeBlockChannelManagerNotshow bool `bson:"recharge_block_channel_manager_notshow" json:"recharge_block_channel_manager_notshow"`//充值模块 充值渠道管理 不显示

	BookBlockManageAdd bool `bson:"book_block_manage_add" json:"book_block_manage_add"`//书籍管理 查看
	BookBlockManageDelete bool `bson:"book_block_manage_delete" json:"book_block_manage_delete"`//书籍管理 删除
	BookBlockManageNotshow bool `bson:"book_block_manage_notshow" json:"book_block_manage_notshow"`//书籍管理 不显示


	ServiceBlockPlayerAdviseSee bool `bson:"service_block_player_advise_see" json:"service_block_player_advise_see"`//客服模块 玩家反馈建议 查看
	ServiceBlockPlayerAdviseReply bool `bson:"service_block_player_advise_reply" json:"service_block_player_advise_reply"` //客服模块 玩家反馈建议 回复
	ServiceBlockPlayerAdviseMarkhandled bool `bson:"service_block_player_advise_markhandled" json:"service_block_player_advise_markhandled"`//客服模块 玩家反馈建议 修改为已经处理
	ServiceBlockPlayerAdviseDelete bool `bson:"service_block_player_advise_delete" json:"service_block_player_advise_delete"` //客服模块 玩家反馈建议 删除
	ServiceBlockPlayerAdviseManyDelete bool `bson:"service_block_player_advise_many_delete" json:"service_block_player_advise_many_delete"`//客服模块 玩家反馈建议 批量删除
	ServiceBlockPlayerAdviseManyReply bool `bson:"service_block_player_advise_many_reply" json:"service_block_player_advise_many_reply"`//客服模块 玩家反馈建议 批量回复
	ServiceBlockMessageList bool `bson:"service_block_message_list" json:"service_block_message_list"`//客服模块 消息列表
	ServiceBlockSendMessageSee bool `bson:"service_block_send_message_see" json:"service_block_send_message_see"`//客服模块 发送短消息 查看
	ServiceBlockSendMessageOne bool `bson:"service_block_send_message_one" json:"service_block_send_message_one"`//客服模块 发送短消息 单次
	ServiceBlockSendMessageEveryHour bool `bson:"service_block_send_message_every_hour" json:"service_block_send_message_every_hour"`//客服模块 发送短消息 每小时
	ServiceBlockSendMessageEveryDay bool `bson:"service_block_send_message_every_day" json:"service_block_send_message_every_day"`//客服模块 发送短消息 每天
	ServiceBlockSendMessageEmailQuerySee bool `bson:"service_block_send_message_email_query_see" json:"service_block_send_message_email_query_see"`//客服模块 发件箱查询 查看
	ServiceBlockSendMessageEmailQueryManyDelete bool `bson:"service_block_send_message_email_query_many_delete" json:"service_block_send_message_email_query_many_delete"`//客服模块 发件箱查询 批量删除
	ServiceBlockSysMessageListSee bool `bson:"service_block_sys_message_list_see" json:"service_block_sys_message_list_see"` //客服模块 系统消息列表 查看
	ServiceBlockSysMessageListManyDelete bool `bson:"service_block_sys_message_list_many_delete" json:"service_block_sys_message_list_many_delete"`//客服模块 系统消息列表 批量删除
	ServiceBlockSysMessageListPageDelete bool `bson:"service_block_sys_message_list_page_delete" json:"service_block_sys_message_list_page_delete"`//客服模块 系统消息列表 删除本页

	DataStatisticsAll bool `bson:"data_statistics_all" json:"data_statistics_all"` //数据统计 运营数据总表
	DataStatisticsMonth bool `bson:"data_statistics_month" json:"data_statistics_month"` //数据统计 运营数据总表(月)
	DataStatisticsGamedata bool `bson:"data_statistics_gamedata" json:"data_statistics_gamedata"`//数据统计 游戏数据报表
	DataStatisticsAllConnection bool `bson:"data_statistics_all_connection" json:"data_statistics_all_connection"`//数据统计 历史链接人数

	VersionManageIOSServerManageSee bool `bson:"version_manage_ios_server_manage_see" json:"version_manage_ios_server_manage_see"`//版本更新管理 IOS升级服务器管理 查看
	VersionManageIOSServerManageHandle bool `bson:"version_manage_ios_server_manage_handle" json:"version_manage_ios_server_manage_handle"`//版本更新管理 IOS升级服务器管理 操作
	VersionManageAllServerManageSee bool `bson:"version_manage_all_server_manage_see" json:"version_manage_all_server_manage_see"` //版本更新管理 整包升级服务器管理 查看
	VersionManageAllServerManageHandle bool `bson:"version_manage_all_server_manage_handle" json:"version_manage_all_server_manage_handle"`//版本更新管理 整包升级服务器管理 操作
	VersionManageBlockServerManageSee bool `bson:"version_manage_block_server_manage_see" json:"version_manage_block_server_manage_see"`//版本更新管理 模块升级服务器管理 查看
	VersionManageBlockServerManageHandle bool `bson:"version_manage_block_server_manage_handle" json:"version_manage_block_server_manage_handle"`//版本更新管理 模块升级服务器管理 操作

	GameConfigurationSysNoticeSee bool `bson:"game_configuration_sys_notice_see" json:"game_configuration_sys_notice_see"`//游戏配置 系统公告 查看
	GameConfigurationSysNoticeAdd bool `bson:"game_configuration_sys_notice_add" json:"game_configuration_sys_notice_add"`//游戏配置 系统公告 增加系统公告
	GameConfigurationSysNoticeShow bool `bson:"game_configuration_sys_notice_show" json:"game_configuration_sys_notice_show"`//游戏配置 系统公告 显示
	GameConfigurationSysNoticeUpdate bool `bson:"game_configuration_sys_notice_update" json:"game_configuration_sys_notice_update"` //游戏配置 系统公告 修改
	GameConfigurationSysNoticeDelete bool `bson:"game_configuration_sys_notice_delete" json:"game_configuration_sys_notice_delete"` //游戏配置 系统公告 删除
	GameConfigurationPayModeSwithSee bool  `bson:"game_configuration_pay_mode_swith_see" json:"game_configuration_pay_mode_swith_see"`//游戏配置 支付模式切换开关 查看
	GameConfigurationPayModeSwithMaintain bool `bson:"game_configuration_pay_mode_swith_maintain" json:"game_configuration_pay_mode_swith_maintain"`//游戏配置 支付模式切换开关 游戏维护
	GameConfigurationPayModeSwithPayMode bool `bson:"game_configuration_pay_mode_swith_pay_mode" json:"game_configuration_pay_mode_swith_pay_mode"`//游戏配置 支付模式切换开关 支付模式
	GameConfigurationPlayerActionSee bool `bson:"game_configuration_player_action_see" json:"game_configuration_player_action_see"` //游戏配置 玩家行为配置 查看
	GameConfigurationPlayerActionAdd bool `bson:"game_configuration_player_action_add" json:"game_configuration_player_action_add"` //游戏配置 玩家行为配置 新增
	GameConfigurationPlayerActionUpdate bool `bson:"game_configuration_player_action_update" json:"game_configuration_player_action_update"` //游戏配置 玩家行为配置 修改
	GameConfigurationPlayerActionDelete bool `bson:"game_configuration_player_action_delete" json:"game_configuration_player_action_delete"`//游戏配置 玩家行为配置 删除
	GameConfigurationAdvSwithSee bool `bson:"game_configuration_adv_swith_see" json:"game_configuration_adv_swith_see"`//游戏配置 新增广告墙开关 查看
	GameConfigurationAdvSwithAdd bool `bson:"game_configuration_adv_swith_add" json:"game_configuration_adv_swith_add"`//游戏配置 新增广告墙开关 新增

	SysSettingBackgroundPowerSee bool `bson:"sys_setting_background_power_see" json:"sys_setting_background_power_see"` //系统设置 后台权限设定 查看
	SysSettingBackgroundPowerHandle bool `bson:"sys_setting_background_power_handle" json:"sys_setting_background_power_handle"`//系统设置 后台权限设定 操作
	SysSettingBackgroundRecord bool `bson:"sys_setting_background_record" json:"sys_setting_background_record"`//系统设置 后台操作记录
	SysSettingServiceManageSys bool `bson:"sys_setting_service_manage_sys" json:"sys_setting_service_manage_sys"`//系统设置 客服管理系统

	CreateTime string `bson:"create_time" json:"create_time"` //账号创建时间
}


//后台访问记录
type VistRecord struct {
	Time string `bson:"time" json:"time"`
	IP string `bson:"ip" json:"ip"`
	User string `bson:"user" json:"user"`
	Url string `bson:"url" json:"url"`
	Param string `bson:"param" json:"param"`
}


//公告管理
type Notice struct {
	Type string `bson:"type" json:"type"`
	Priority string `bson:"priority" json:"priority"`
	Status string `bson:"status" json:"status"`
	Platform string `bson:"platform" json:"platform"`
	StartTime string `bson:"start_time" json:"start_time"`
	EndTime string `bson:"end_time" json:"end_time"`
	Substance string `bson:"substance" json:"substance"`
}
//公告管理 修改内容
type NoticeEdit struct {
	Id string `json:"_id"`
	Type string `json:"type"`
	Priority string `json:"priority"`
	Status string `json:"status"`
	Platform string `json:"platform"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Substance string `json:"substance"`
}


//用户反馈	客服管理 --- 玩家反馈  回复界面
type ChatRecord struct {
	Num int `bson:"num" json:"num"`
	UserId string `bson:"user_id" json:"user_id"`
	ReplyTime string `bson:"reply_time" json:"reply_time"`
	Text string `bson:"text" json:"text"`
	//IsLookOver bool `bson:"is_look_over" json:"is_look_over"`
}

//玩家反馈表 (仅供测试用)
type Service struct {
	Num int `bson:"num" json:"num"`
	UserId int `bson:"user_id" json:"user_id"`
	Account string `bson:"account" json:"account"`
	Mark string `bson:"mark" json:"mark"`
	Problem string `bson:"problem" json:"problem"`
	Status string `bson:"status" json:"status"`
	IpAddr string `bson:"ip_address" json:"ip_addr"`
	CreateTime string `bson:"create_time" json:"create_time"`

}
//UserFeedback 用户反馈问题记录表
type UserFeedbackRecord struct {
	Num int `bson:"num" json:"num"` //问题编号
	UserId string `bson:"user_id" json:"user_id"`
	Account string `bson:"account" json:"account"` //用户账号
	Mark string `bson:"mark" json:"mark"`//客服标记问题
	Problem string `bson:"problem" json:"problem"`//用户提的问题
	Status string `bson:"status" json:"status"`//状态
	IpAddr string `bson:"ip_address" json:"ip_addr"`//ip地址
	CreateTime string `bson:"create_time" json:"create_time"`//提问时间
	ReplyTime string `bson:"reply_time" json:"reply_time"`//回复时间
	Text string `bson:"text" json:"text"` //客服回复内容
}

//玩家反馈列表 (仅供测试用)
type ServiceList struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Content []Service `json:"content"`
}


//书籍管理  获取书籍列表
type BookList struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Category string `bson:"category" json:"category"`
	SubCategory string `bson:"sub_category" json:"sub_category"`
	Name string	`bson:"name" json:"name"`
	Author string `bson:"author" json:"author"`
	WordsCount int64 `bson:"words_count" json:"words_count"`
	ChapterCount int `bson:"chapter_count" json:"chapter_count"`
	Free bool `bson:"free" json:"free"`
	ChapterPrice float64 `bson:"chapter_price" json:"chapter_price"`
	FreeChapterCount int `bson:"free_chapter_count" json:"free_chapter_count"`
	Hide bool `bson:"hide" json:"hide"`	//书籍信息是否隐藏
	CreateAt string `bson:"create_at" json:"create_at"`
	ChapterList []ChapterList `bson:"chapter_list" json:"chapter_list"`
}
type ChapterList struct {
	Name string `bson:"name" json:"name"` //章节名
	Crawled bool `bson:"crawled" json:"crawled"`//是否爬取
	Free bool `bson:"free" json:"free"` //这一章是否免费
	Words string `bson:"words" json:"words"` //这一章单词数
}

//书籍管理 	修改书籍信息
type BookInfo struct {
	Id string `json:"id"`	//书籍ID
	Category string `json:"category"`	//一级分类
	SubCategory string `json:"sub_category"` //二级分类
	ChapterPrice float64 `json:"chapter_price"`	//章节单价
	FreeChapterCount int `json:"free_chapter_count"`	//免费章节数量
	Free bool `json:"free"`	//本书是否免费
	Hide bool `json:"hide"`	//书籍信息是否隐藏
}
