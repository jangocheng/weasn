package media

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/service"
)

type Client struct {
	baseClient *service.BaseClient
}

func New(opts context.Options) *Client {
	return &Client{baseClient: service.NewBaseClient(opts)}
}

// UploadImage 上传图片
func (this *Client) UploadImage(path string) (*MediaUpload, error) {
	return this.upload("image", path)
}

// uploadVoice 上传声音
func (this *Client) UploadVoice(path string) (*MediaUpload, error) {
	return this.upload("voice", path)
}

// uploadVideo 上传视频
func (this *Client) UploadVideo(path string) (*MediaUpload, error) {
	return this.upload("video", path)
}

// UploadThumb 上传缩略图
func (this *Client) UploadThumb(path string) (*MediaUpload, error) {
	return this.upload("thumb", path)
}

// UploadVideoForBroadcasting 上传群发视频
func (this *Client) UploadVideoForBroadcasting(path, title, description string) (*MediaUpload, error) {
	response, err := this.UploadVideo(path)
	if err != nil {
		return nil, err
	}
	return this.CreateVideoForBroadcasting(response.MediaId, title, description)
}

// CreateVideoForBroadcasting 创建群发消息
func (this *Client) CreateVideoForBroadcasting(mediaId, title, description string) (*MediaUpload, error) {
	params := map[string]string{
		"media_id":    mediaId,
		"title":       title,
		"description": description,
	}

	var response MediaUpload
	return &response, this.baseClient.Post("cgi-bin/media/uploadvideo", params, &response)
}

func (this *Client) upload(mediaType, path string) (*MediaUpload, error) {
	var response MediaUpload

	// 上传文件
	files := map[string]string{
		"media": path,
	}

	// 请求参数
	params := map[string]string{
		"type": mediaType,
	}
	return &response, this.baseClient.Upload("cgi-bin/media/upload", files, params, &response)
}
