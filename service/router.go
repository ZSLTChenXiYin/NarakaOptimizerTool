package service

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func routerInit() {
	r.LoadHTMLGlob("./templates/*")
	r.StaticFS("/static", gin.Dir("./static", false))

	r.GET("/", getIndex)

	r.GET("/physic", getPhysic)
	r.POST("/physic", postPhysic)

	r.GET("/initialization", getInitialization)
	r.POST("/initialization", postInitialization)

	r.GET("/information", getInformation)
}

func getIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func getPhysic(c *gin.Context) {
	c.HTML(200, "physic.html", nil)
}

type QualitySettingsData struct {
	Preset                 int `json:"preset"`
	L22GraphicQualityLevel struct {
		MModelQualityLevel         int `json:"m_modelQualityLevel"`
		MTessellationQualityLevel  int `json:"m_tessellationQualityLevel"`
		MVisualEffectsQualityLevel int `json:"m_visualEffectsQualityLevel"`
		MTextureQualityLevel       int `json:"m_textureQualityLevel"`
		MShadowQualityLevel        int `json:"m_shadowQualityLevel"`
		MVolumetricLightLevel      int `json:"m_volumetricLightLevel"`
		MCloudQualityLevel         int `json:"m_cloudQualityLevel"`
		MAOLevel                   int `json:"m_aoLevel"`
		MSSRLevel                  int `json:"m_SSRLevel"`
		MAALevel                   int `json:"m_AALevel"`
		MPostProcessingLevel       int `json:"m_PostProcessingLevel"`
		MLightingQualityLevel      int `json:"m_LightingQualityLevel"`
	} `json:"l22GraphicQualityLevel"`
	L22SystemQualitySetting struct {
		RenderScale                  float64 `json:"renderScale"`
		RenderScaleStep              float64 `json:"renderScaleStep"`
		AAMode                       int     `json:"aaMode"`
		CheckboardRendering          int     `json:"checkboardRendering"`
		UpSamplingType               int     `json:"upSamplingType"`
		EnableDlssDx12               bool    `json:"enableDlssDx12"`
		EnableDlssG                  bool    `json:"enableDlssG"`
		EnableDlssRR                 bool    `json:"enableDlssRR"`
		RandomDiscardFactor          float64 `json:"randomDiscardFactor"`
		DlssMode                     int     `json:"dlssMode"`
		XessMode                     int     `json:"xessMode"`
		DlssSharpness                float64 `json:"dlssSharpness"`
		LXFsr2Mode                   int     `json:"lxFsr2Mode"`
		Fsr2Sharpness                float64 `json:"fsr2Sharpness"`
		UpsamplingMode               int     `json:"upsamplingMode"`
		LXFsr3Mode                   int     `json:"lxFsr3Mode"`
		EnbaleFSR3FrameInterpolation bool    `json:"enbaleFSR3FrameInterpolation"`
		UpsamplingSharpStops         float64 `json:"upsamplingSharpStops"`
		NisQuality                   int     `json:"nisQuality"`
		FullScreenMode               int     `json:"fullScreenMode"`
		ResolutionWidth              int     `json:"resolutionWidth"`
		ResolutionHeight             int     `json:"resolutionHeight"`
		FrameRateLimit               int     `json:"frameRateLimit"`
		VSyncCount                   int     `json:"vSyncCount"`
		Gamma                        float64 `json:"gamma"`
		MaxLuma                      float64 `json:"maxLuma"`
		MinLuma                      float64 `json:"minLuma"`
		PaperWhite                   float64 `json:"paperWhite"`
		MHDRMode                     int     `json:"mHDRMode"`
		ReflexMode                   int     `json:"reflexMode"`
		VrsMode                      int     `json:"vrsMode"`
		ColorBlindMode               int     `json:"colorBlindMode"`
		ColorBlindStrength           float64 `json:"colorBlindStrength"`
		ColorBlindQualityMode        bool    `json:"colorBlindQualityMode"`
		NvHighlightsEnabled          bool    `json:"nvHighlightsEnabled"`
		StyleMode                    int     `json:"styleMode"`
		MotionBlurEnabled            bool    `json:"motionBlurEnabled"`
		RaytracingEnabled            bool    `json:"raytracingEnabled"`
		RaytracingAO                 bool    `json:"raytracingAO"`
		RaytracingGI                 bool    `json:"raytracingGI"`
		RaytracingReflection         bool    `json:"raytracingReflection"`
		RaytracingShadow             bool    `json:"raytracingShadow"`
		RaytracingBVHAllInOne        bool    `json:"raytracingBVHAllInOne"`
		RaytracingBVHActorCountMax   int     `json:"raytracingBVHActorCountMax"`
		RtgiResolution               int     `json:"rtgiResolution"`
		CharacterAdditionalPhysics1  bool    `json:"characterAdditionalPhysics1"`
		XboxQualityOption            int     `json:"xboxQualityOption"`
	} `json:"l22SystemQualitySetting"`
}

func postPhysic(c *gin.Context) {
	physic := c.PostForm("physic")
	if developerMode() {
		DebugLogger.Printf(NewLog("Physic: %s", physic))
	}

	quality_settings_data_file_path := "/NarakaBladepoint_Data/QualitySettingsData.txt"

	var open_file_path string

	switch installationPlatform() {
	case "official":
		open_file_path = fmt.Sprintf("%s%s%s", narakaInstallPath(), "/program", quality_settings_data_file_path)
	case "steam", "epic":
		open_file_path = fmt.Sprintf("%s%s", narakaInstallPath(), quality_settings_data_file_path)
	default:
		open_file_path = fmt.Sprintf("%s%s", narakaInstallPath(), quality_settings_data_file_path)
	}

	file, err := os.OpenFile(open_file_path, os.O_RDWR, 0666)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to open file: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}
	defer file.Close()

	json_data, err := io.ReadAll(file)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to read file: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}

	quality_settings_data := &QualitySettingsData{}

	err = json.Unmarshal(json_data, &quality_settings_data)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to unmarshal json: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}

	if physic == "on" {
		quality_settings_data.L22SystemQualitySetting.CharacterAdditionalPhysics1 = true
	} else {
		quality_settings_data.L22SystemQualitySetting.CharacterAdditionalPhysics1 = false
	}

	json_data, err = json.Marshal(quality_settings_data)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to marshal json: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}

	err = file.Truncate(0)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to truncate file: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to seek file: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}

	_, err = file.Write(json_data)
	if err != nil {
		if developerMode() {
			DebugLogger.Printf(NewLog("Failed to write file: %v", err))
		}
		c.HTML(200, "result.html", gin.H{
			"msg": "物理效果设置失败。",
			"err": err.Error(),
		})
		return
	}

	c.HTML(200, "result.html", gin.H{
		"msg": "物理效果设置成功。",
	})
}

func getInitialization(c *gin.Context) {
	c.HTML(200, "initialization.html", nil)
}

func postInitialization(c *gin.Context) {
	naraka_install_path := c.DefaultPostForm("naraka_install_path", "")
	if developerMode() {
		DebugLogger.Printf(NewLog("Naraka install path: %s", naraka_install_path))
	}
	if naraka_install_path == "" {
		c.HTML(200, "result.html", gin.H{
			"msg": "初始化失败。",
			"err": "永劫无间安装路径不能为空。",
		})
		return
	}

	installation_platform := c.PostForm("installation_platform")
	if developerMode() {
		DebugLogger.Printf(NewLog("Installation platform: %s", installation_platform))
	}

	developer_mode := c.PostForm("developer_mode")
	if developerMode() {
		DebugLogger.Printf(NewLog("Developer mode: %s", developer_mode))
	}

	setNarakaInstallPath(naraka_install_path)
	if developerMode() {
		DebugLogger.Printf(NewLog("Naraka install path in map: %s", narakaInstallPath()))
	}
	setInstallationPlatform(installation_platform)
	if developerMode() {
		DebugLogger.Printf(NewLog("Installation platform in map: %s", installationPlatform()))
	}
	setDeveloperMode(developer_mode == "on")
	if developerMode() {
		DebugLogger.Printf(NewLog("Developer mode in map: %t", developerMode()))
	}

	saveConfig()

	c.HTML(200, "result.html", gin.H{
		"msg": "初始化成功。",
	})
}

func getInformation(c *gin.Context) {
	var developer_mode string
	if developerMode() {
		developer_mode = "开启"
	} else {
		developer_mode = "关闭"
	}

	var installation_platform string
	switch installationPlatform() {
	case "official":
		installation_platform = "官方"
	case "steam":
		installation_platform = "Steam"
	case "epic":
		installation_platform = "Epic"
	case "other":
		installation_platform = "其他"
	default:
		installation_platform = "未知"
	}

	naraka_install_path := narakaInstallPath()
	if naraka_install_path == "" {
		naraka_install_path = "未知"
	}

	c.HTML(200, "information.html", gin.H{
		"developer_mode":        developer_mode,
		"installation_platform": installation_platform,
		"naraka_install_path":   naraka_install_path,
	})
}
