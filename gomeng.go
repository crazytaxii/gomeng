package gomeng

import ()

const PLATFORM_IOS = "ios"
const PLATFORM_ANDROID = "android"
const DOMAIN = "https://msgapi.umeng.com"

type Config struct {
	ProductMode     bool
	Platform        string
	AppKey          string
	AppMasterSecret string
}

var configMap = make(map[string]*Config)

/**
 * Android平台配置初始化
 */
func Init2(mode bool, appKey string, appMasterSecret string) {
	configMap["android"] = &Config{
		ProductMode:     mode,
		Platform:        PLATFORM_ANDROID,
		AppKey:          appKey,
		AppMasterSecret: appMasterSecret,
	}
} // InitAndroid()

/**
 * iOS平台配置初始化
 */
func Init1(mode bool, appKey string, appMasterSecret string) {
	configMap["ios"] = &Config{
		ProductMode:     mode,
		Platform:        PLATFORM_IOS,
		AppKey:          appKey,
		AppMasterSecret: appMasterSecret,
	}
} // InitiOS()
