package stablediffusion


type Text2ImgParams struct {
	EnableHr                          bool          `json:"enable_hr"`
	DenoisingStrength                 int64         `json:"denoising_strength"`
	FirstphaseWidth                   int64         `json:"firstphase_width"`
	FirstphaseHeight                  int64         `json:"firstphase_height"`
	HrScale                           int64         `json:"hr_scale"`
	HrUpscaler                        string        `json:"hr_upscaler"`
	HrSecondPassSteps                 int64         `json:"hr_second_pass_steps"`
	HrResizeX                         int64         `json:"hr_resize_x"`
	HrResizeY                         int64         `json:"hr_resize_y"`
	Prompt                            string        `json:"prompt"`
	Styles                            []string      `json:"styles"`
	Seed                              int64         `json:"seed"`
	Subseed                           int64         `json:"subseed"`
	SubseedStrength                   int64         `json:"subseed_strength"`
	SeedResizeFromH                   int64         `json:"seed_resize_from_h"`
	SeedResizeFromW                   int64         `json:"seed_resize_from_w"`
	SamplerBame                       string        `json:"sampler_name"`
	BatchSize                         int64         `json:"batch_size"`
	NIter                             int64         `json:"n_iter"`
	Steps                             int64         `json:"steps"`
	CfgScale                          int64         `json:"cfg_scale"`
	Width                             uint64        `json:"width"`
	Height                            uint64        `json:"height"`
	RestoreFaces                      bool          `json:"restore_faces"`
	Tiling                            bool          `json:"tiling"`
	DoNotSaveSamples                  bool          `json:"do_not_save_samples"`
	DoNotSaveGrid                     bool          `json:"do_not_save_grid"`
	NegativePrompt                    string        `json:"negative_prompt"`
	Eta                               int64         `json:"eta"`
	SChurn                            int64         `json:"s_churn"`
	STmax                             int64         `json:"s_tmax"`
	STmin                             int64         `json:"s_tmin"`
	Snoise                            int64         `json:"s_noise"`
	OverrideSettings                  interface{}   `json:"override_settings"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards"`
	ScriptArgs                        []interface{} `json:"script_args"`
	SamplerIndex                      string        `json:"sampler_index"`
	ScriptName                        string        `json:"script_name"`
	SendImages                        bool          `json:"send_images"`
	SaveImages                        bool          `json:"save_images"`
	AlwaysonScripts                   interface{}   `json:"alwayson_scripts"`
}
