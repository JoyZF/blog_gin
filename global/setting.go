package global

import (
	"github.com/JoyZF/blog_gin/pkg/logger"
	"github.com/JoyZF/blog_gin/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettings
	EmailSetting    *setting.EmailSettingS
)
