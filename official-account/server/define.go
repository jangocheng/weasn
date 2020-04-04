package server

// 消息类型
const (
	MsgTypeText       = "text"                      // 文本消息
	MsgTypeImage      = "image"                     // 图片消息
	MsgTypeVoice      = "voice"                     // 语音消息
	MsgTypeVideo      = "video"                     // 视频消息
	MsgTypeShortVideo = "shortvideo"                // 短视频消息[限接收]
	MsgTypeLocation   = "location"                  // 坐标消息[限接收]
	MsgTypeLink       = "link"                      // 链接消息[限接收]
	MsgTypeMusic      = "music"                     // 音乐消息[限回复]
	MsgTypeNews       = "news"                      // 图文消息[限回复]
	MsgTypeTransfer   = "transfer_customer_service" // 消息消息转发到客服
	MsgTypeEvent      = "event"                     // 事件推送消息
)

// 事件类型
const (
	EventSubscribe             = "subscribe"             // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
	EventUnsubscribe           = "unsubscribe"           // 取消关注事件
	EventScan                  = "SCAN"                  // 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventLocation              = "LOCATION"              // 上报地理位置事件
	EventClick                 = "CLICK"                 // 点击菜单拉取消息时的事件推送
	EventView                  = "VIEW"                  // 点击菜单跳转链接时的事件推送
	EventScancodePush          = "scancode_push"         // 扫码推事件的事件推送
	EventScancodeWaitmsg       = "scancode_waitmsg"      // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventPicSysphoto           = "pic_sysphoto"          // 弹出系统拍照发图的事件推送
	EventPicPhotoOrAlbum       = "pic_photo_or_album"    // 弹出拍照或者相册发图的事件推送
	EventPicWeixin             = "pic_weixin"            // 弹出微信相册发图器的事件推送
	EventLocationSelect        = "location_select"       // 弹出地理位置选择器的事件推送
	EventTemplateSendJobFinish = "TEMPLATESENDJOBFINISH" // EventTemplateSendJobFinish = "TEMPLATESENDJOBFINISH"/

)

// 函数映射
var objectMapper = map[string]string{
	"Text":                    MsgTypeText,
	"Image":                   MsgTypeImage,
	"Voice":                   MsgTypeVoice,
	"Video":                   MsgTypeVideo,
	"Music":                   MsgTypeMusic,
	"News":                    MsgTypeNews,
	"Link":                    MsgTypeLink,
	"Event":                   MsgTypeEvent,
	"Transfer":                MsgTypeTransfer,
	"ShortVideo":              MsgTypeShortVideo,
	"Location":                MsgTypeLocation,
	"E_Scan":                  EventScan,
	"E_Click":                 EventClick,
	"E_View":                  EventView,
	"E_Location":              EventLocation,
	"E_Subscribe":             EventSubscribe,
	"E_Unsubscribe":           EventUnsubscribe,
	"E_ScancodePush":          EventScancodePush,
	"E_ScancodeWaitmsg":       EventScancodeWaitmsg,
	"E_PicSysphoto":           EventPicSysphoto,
	"E_PicPhotoOrAlbum":       EventPicPhotoOrAlbum,
	"E_PicWeixin":             EventPicWeixin,
	"E_LocationSelect":        EventLocationSelect,
	"E_TemplateSendJobFinish": EventTemplateSendJobFinish,
}
