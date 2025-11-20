package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	large_scale_app "github.com/blueprint-uservices/blueprint/examples/large_scale_app/workflow/large_scale_app"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with http, and uses mongodb as NoSQL database backends (generated)",
	Build:       makeDockerSpec,
}

func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string
	var allServices []string

	//----- service171 (terminal) ----------
	service171_db := mongodb.Container(spec, "service171_db")
	allServices = append(allServices, service171_db)
	service171 := workflow.Service[large_scale_app.Service171](spec, "service171", service171_db)
	service171_ctr := applyDockerDefaults(spec, service171, "service171_proc", "service171_container")
	containers = append(containers, service171_ctr)
	allServices = append(allServices, service171)

	//----- service128 (terminal) ----------
	service128_db := mongodb.Container(spec, "service128_db")
	allServices = append(allServices, service128_db)
	service128 := workflow.Service[large_scale_app.Service128](spec, "service128", service128_db)
	service128_ctr := applyDockerDefaults(spec, service128, "service128_proc", "service128_container")
	containers = append(containers, service128_ctr)
	allServices = append(allServices, service128)

	//----- service341 (terminal) ----------
	service341_db := mongodb.Container(spec, "service341_db")
	allServices = append(allServices, service341_db)
	service341 := workflow.Service[large_scale_app.Service341](spec, "service341", service341_db)
	service341_ctr := applyDockerDefaults(spec, service341, "service341_proc", "service341_container")
	containers = append(containers, service341_ctr)
	allServices = append(allServices, service341)

	//----- service340 (terminal) ----------
	service340_db := mongodb.Container(spec, "service340_db")
	allServices = append(allServices, service340_db)
	service340 := workflow.Service[large_scale_app.Service340](spec, "service340", service340_db)
	service340_ctr := applyDockerDefaults(spec, service340, "service340_proc", "service340_container")
	containers = append(containers, service340_ctr)
	allServices = append(allServices, service340)

	//----- service339 (terminal) ----------
	service339_db := mongodb.Container(spec, "service339_db")
	allServices = append(allServices, service339_db)
	service339 := workflow.Service[large_scale_app.Service339](spec, "service339", service339_db)
	service339_ctr := applyDockerDefaults(spec, service339, "service339_proc", "service339_container")
	containers = append(containers, service339_ctr)
	allServices = append(allServices, service339)

	//----- service338 (terminal) ----------
	service338_db := mongodb.Container(spec, "service338_db")
	allServices = append(allServices, service338_db)
	service338 := workflow.Service[large_scale_app.Service338](spec, "service338", service338_db)
	service338_ctr := applyDockerDefaults(spec, service338, "service338_proc", "service338_container")
	containers = append(containers, service338_ctr)
	allServices = append(allServices, service338)

	//----- service337 (terminal) ----------
	service337_db := mongodb.Container(spec, "service337_db")
	allServices = append(allServices, service337_db)
	service337 := workflow.Service[large_scale_app.Service337](spec, "service337", service337_db)
	service337_ctr := applyDockerDefaults(spec, service337, "service337_proc", "service337_container")
	containers = append(containers, service337_ctr)
	allServices = append(allServices, service337)

	//----- service336 (terminal) ----------
	service336_db := mongodb.Container(spec, "service336_db")
	allServices = append(allServices, service336_db)
	service336 := workflow.Service[large_scale_app.Service336](spec, "service336", service336_db)
	service336_ctr := applyDockerDefaults(spec, service336, "service336_proc", "service336_container")
	containers = append(containers, service336_ctr)
	allServices = append(allServices, service336)

	//----- service335 (terminal) ----------
	service335_db := mongodb.Container(spec, "service335_db")
	allServices = append(allServices, service335_db)
	service335 := workflow.Service[large_scale_app.Service335](spec, "service335", service335_db)
	service335_ctr := applyDockerDefaults(spec, service335, "service335_proc", "service335_container")
	containers = append(containers, service335_ctr)
	allServices = append(allServices, service335)

	//----- service334 (terminal) ----------
	service334_db := mongodb.Container(spec, "service334_db")
	allServices = append(allServices, service334_db)
	service334 := workflow.Service[large_scale_app.Service334](spec, "service334", service334_db)
	service334_ctr := applyDockerDefaults(spec, service334, "service334_proc", "service334_container")
	containers = append(containers, service334_ctr)
	allServices = append(allServices, service334)

	//----- service333 (terminal) ----------
	service333_db := mongodb.Container(spec, "service333_db")
	allServices = append(allServices, service333_db)
	service333 := workflow.Service[large_scale_app.Service333](spec, "service333", service333_db)
	service333_ctr := applyDockerDefaults(spec, service333, "service333_proc", "service333_container")
	containers = append(containers, service333_ctr)
	allServices = append(allServices, service333)

	//----- service332 (terminal) ----------
	service332_db := mongodb.Container(spec, "service332_db")
	allServices = append(allServices, service332_db)
	service332 := workflow.Service[large_scale_app.Service332](spec, "service332", service332_db)
	service332_ctr := applyDockerDefaults(spec, service332, "service332_proc", "service332_container")
	containers = append(containers, service332_ctr)
	allServices = append(allServices, service332)

	//----- service331 (terminal) ----------
	service331_db := mongodb.Container(spec, "service331_db")
	allServices = append(allServices, service331_db)
	service331 := workflow.Service[large_scale_app.Service331](spec, "service331", service331_db)
	service331_ctr := applyDockerDefaults(spec, service331, "service331_proc", "service331_container")
	containers = append(containers, service331_ctr)
	allServices = append(allServices, service331)

	//----- service330 (terminal) ----------
	service330_db := mongodb.Container(spec, "service330_db")
	allServices = append(allServices, service330_db)
	service330 := workflow.Service[large_scale_app.Service330](spec, "service330", service330_db)
	service330_ctr := applyDockerDefaults(spec, service330, "service330_proc", "service330_container")
	containers = append(containers, service330_ctr)
	allServices = append(allServices, service330)

	//----- service329 (terminal) ----------
	service329_db := mongodb.Container(spec, "service329_db")
	allServices = append(allServices, service329_db)
	service329 := workflow.Service[large_scale_app.Service329](spec, "service329", service329_db)
	service329_ctr := applyDockerDefaults(spec, service329, "service329_proc", "service329_container")
	containers = append(containers, service329_ctr)
	allServices = append(allServices, service329)

	//----- service328 (terminal) ----------
	service328_db := mongodb.Container(spec, "service328_db")
	allServices = append(allServices, service328_db)
	service328 := workflow.Service[large_scale_app.Service328](spec, "service328", service328_db)
	service328_ctr := applyDockerDefaults(spec, service328, "service328_proc", "service328_container")
	containers = append(containers, service328_ctr)
	allServices = append(allServices, service328)

	//----- service327 (terminal) ----------
	service327_db := mongodb.Container(spec, "service327_db")
	allServices = append(allServices, service327_db)
	service327 := workflow.Service[large_scale_app.Service327](spec, "service327", service327_db)
	service327_ctr := applyDockerDefaults(spec, service327, "service327_proc", "service327_container")
	containers = append(containers, service327_ctr)
	allServices = append(allServices, service327)

	//----- service326 (terminal) ----------
	service326_db := mongodb.Container(spec, "service326_db")
	allServices = append(allServices, service326_db)
	service326 := workflow.Service[large_scale_app.Service326](spec, "service326", service326_db)
	service326_ctr := applyDockerDefaults(spec, service326, "service326_proc", "service326_container")
	containers = append(containers, service326_ctr)
	allServices = append(allServices, service326)

	//----- service325 (terminal) ----------
	service325_db := mongodb.Container(spec, "service325_db")
	allServices = append(allServices, service325_db)
	service325 := workflow.Service[large_scale_app.Service325](spec, "service325", service325_db)
	service325_ctr := applyDockerDefaults(spec, service325, "service325_proc", "service325_container")
	containers = append(containers, service325_ctr)
	allServices = append(allServices, service325)

	//----- service324 (terminal) ----------
	service324_db := mongodb.Container(spec, "service324_db")
	allServices = append(allServices, service324_db)
	service324 := workflow.Service[large_scale_app.Service324](spec, "service324", service324_db)
	service324_ctr := applyDockerDefaults(spec, service324, "service324_proc", "service324_container")
	containers = append(containers, service324_ctr)
	allServices = append(allServices, service324)

	//----- service323 (terminal) ----------
	service323_db := mongodb.Container(spec, "service323_db")
	allServices = append(allServices, service323_db)
	service323 := workflow.Service[large_scale_app.Service323](spec, "service323", service323_db)
	service323_ctr := applyDockerDefaults(spec, service323, "service323_proc", "service323_container")
	containers = append(containers, service323_ctr)
	allServices = append(allServices, service323)

	//----- service322 (terminal) ----------
	service322_db := mongodb.Container(spec, "service322_db")
	allServices = append(allServices, service322_db)
	service322 := workflow.Service[large_scale_app.Service322](spec, "service322", service322_db)
	service322_ctr := applyDockerDefaults(spec, service322, "service322_proc", "service322_container")
	containers = append(containers, service322_ctr)
	allServices = append(allServices, service322)

	//----- service321 (terminal) ----------
	service321_db := mongodb.Container(spec, "service321_db")
	allServices = append(allServices, service321_db)
	service321 := workflow.Service[large_scale_app.Service321](spec, "service321", service321_db)
	service321_ctr := applyDockerDefaults(spec, service321, "service321_proc", "service321_container")
	containers = append(containers, service321_ctr)
	allServices = append(allServices, service321)

	//----- service320 (terminal) ----------
	service320_db := mongodb.Container(spec, "service320_db")
	allServices = append(allServices, service320_db)
	service320 := workflow.Service[large_scale_app.Service320](spec, "service320", service320_db)
	service320_ctr := applyDockerDefaults(spec, service320, "service320_proc", "service320_container")
	containers = append(containers, service320_ctr)
	allServices = append(allServices, service320)

	//----- service319 (terminal) ----------
	service319_db := mongodb.Container(spec, "service319_db")
	allServices = append(allServices, service319_db)
	service319 := workflow.Service[large_scale_app.Service319](spec, "service319", service319_db)
	service319_ctr := applyDockerDefaults(spec, service319, "service319_proc", "service319_container")
	containers = append(containers, service319_ctr)
	allServices = append(allServices, service319)

	//----- service318 (terminal) ----------
	service318_db := mongodb.Container(spec, "service318_db")
	allServices = append(allServices, service318_db)
	service318 := workflow.Service[large_scale_app.Service318](spec, "service318", service318_db)
	service318_ctr := applyDockerDefaults(spec, service318, "service318_proc", "service318_container")
	containers = append(containers, service318_ctr)
	allServices = append(allServices, service318)

	//----- service317 (terminal) ----------
	service317_db := mongodb.Container(spec, "service317_db")
	allServices = append(allServices, service317_db)
	service317 := workflow.Service[large_scale_app.Service317](spec, "service317", service317_db)
	service317_ctr := applyDockerDefaults(spec, service317, "service317_proc", "service317_container")
	containers = append(containers, service317_ctr)
	allServices = append(allServices, service317)

	//----- service316 (terminal) ----------
	service316_db := mongodb.Container(spec, "service316_db")
	allServices = append(allServices, service316_db)
	service316 := workflow.Service[large_scale_app.Service316](spec, "service316", service316_db)
	service316_ctr := applyDockerDefaults(spec, service316, "service316_proc", "service316_container")
	containers = append(containers, service316_ctr)
	allServices = append(allServices, service316)

	//----- service315 (terminal) ----------
	service315_db := mongodb.Container(spec, "service315_db")
	allServices = append(allServices, service315_db)
	service315 := workflow.Service[large_scale_app.Service315](spec, "service315", service315_db)
	service315_ctr := applyDockerDefaults(spec, service315, "service315_proc", "service315_container")
	containers = append(containers, service315_ctr)
	allServices = append(allServices, service315)

	//----- service314 (terminal) ----------
	service314_db := mongodb.Container(spec, "service314_db")
	allServices = append(allServices, service314_db)
	service314 := workflow.Service[large_scale_app.Service314](spec, "service314", service314_db)
	service314_ctr := applyDockerDefaults(spec, service314, "service314_proc", "service314_container")
	containers = append(containers, service314_ctr)
	allServices = append(allServices, service314)

	//----- service313 (terminal) ----------
	service313_db := mongodb.Container(spec, "service313_db")
	allServices = append(allServices, service313_db)
	service313 := workflow.Service[large_scale_app.Service313](spec, "service313", service313_db)
	service313_ctr := applyDockerDefaults(spec, service313, "service313_proc", "service313_container")
	containers = append(containers, service313_ctr)
	allServices = append(allServices, service313)

	//----- service312 (terminal) ----------
	service312_db := mongodb.Container(spec, "service312_db")
	allServices = append(allServices, service312_db)
	service312 := workflow.Service[large_scale_app.Service312](spec, "service312", service312_db)
	service312_ctr := applyDockerDefaults(spec, service312, "service312_proc", "service312_container")
	containers = append(containers, service312_ctr)
	allServices = append(allServices, service312)

	//----- service311 (terminal) ----------
	service311_db := mongodb.Container(spec, "service311_db")
	allServices = append(allServices, service311_db)
	service311 := workflow.Service[large_scale_app.Service311](spec, "service311", service311_db)
	service311_ctr := applyDockerDefaults(spec, service311, "service311_proc", "service311_container")
	containers = append(containers, service311_ctr)
	allServices = append(allServices, service311)

	//----- service310 (terminal) ----------
	service310_db := mongodb.Container(spec, "service310_db")
	allServices = append(allServices, service310_db)
	service310 := workflow.Service[large_scale_app.Service310](spec, "service310", service310_db)
	service310_ctr := applyDockerDefaults(spec, service310, "service310_proc", "service310_container")
	containers = append(containers, service310_ctr)
	allServices = append(allServices, service310)

	//----- service309 (terminal) ----------
	service309_db := mongodb.Container(spec, "service309_db")
	allServices = append(allServices, service309_db)
	service309 := workflow.Service[large_scale_app.Service309](spec, "service309", service309_db)
	service309_ctr := applyDockerDefaults(spec, service309, "service309_proc", "service309_container")
	containers = append(containers, service309_ctr)
	allServices = append(allServices, service309)

	//----- service308 (terminal) ----------
	service308_db := mongodb.Container(spec, "service308_db")
	allServices = append(allServices, service308_db)
	service308 := workflow.Service[large_scale_app.Service308](spec, "service308", service308_db)
	service308_ctr := applyDockerDefaults(spec, service308, "service308_proc", "service308_container")
	containers = append(containers, service308_ctr)
	allServices = append(allServices, service308)

	//----- service307 (terminal) ----------
	service307_db := mongodb.Container(spec, "service307_db")
	allServices = append(allServices, service307_db)
	service307 := workflow.Service[large_scale_app.Service307](spec, "service307", service307_db)
	service307_ctr := applyDockerDefaults(spec, service307, "service307_proc", "service307_container")
	containers = append(containers, service307_ctr)
	allServices = append(allServices, service307)

	//----- service306 (terminal) ----------
	service306_db := mongodb.Container(spec, "service306_db")
	allServices = append(allServices, service306_db)
	service306 := workflow.Service[large_scale_app.Service306](spec, "service306", service306_db)
	service306_ctr := applyDockerDefaults(spec, service306, "service306_proc", "service306_container")
	containers = append(containers, service306_ctr)
	allServices = append(allServices, service306)

	//----- service305 (terminal) ----------
	service305_db := mongodb.Container(spec, "service305_db")
	allServices = append(allServices, service305_db)
	service305 := workflow.Service[large_scale_app.Service305](spec, "service305", service305_db)
	service305_ctr := applyDockerDefaults(spec, service305, "service305_proc", "service305_container")
	containers = append(containers, service305_ctr)
	allServices = append(allServices, service305)

	//----- service304 (terminal) ----------
	service304_db := mongodb.Container(spec, "service304_db")
	allServices = append(allServices, service304_db)
	service304 := workflow.Service[large_scale_app.Service304](spec, "service304", service304_db)
	service304_ctr := applyDockerDefaults(spec, service304, "service304_proc", "service304_container")
	containers = append(containers, service304_ctr)
	allServices = append(allServices, service304)

	//----- service303 (terminal) ----------
	service303_db := mongodb.Container(spec, "service303_db")
	allServices = append(allServices, service303_db)
	service303 := workflow.Service[large_scale_app.Service303](spec, "service303", service303_db)
	service303_ctr := applyDockerDefaults(spec, service303, "service303_proc", "service303_container")
	containers = append(containers, service303_ctr)
	allServices = append(allServices, service303)

	//----- service302 (terminal) ----------
	service302_db := mongodb.Container(spec, "service302_db")
	allServices = append(allServices, service302_db)
	service302 := workflow.Service[large_scale_app.Service302](spec, "service302", service302_db)
	service302_ctr := applyDockerDefaults(spec, service302, "service302_proc", "service302_container")
	containers = append(containers, service302_ctr)
	allServices = append(allServices, service302)

	//----- service301 (terminal) ----------
	service301_db := mongodb.Container(spec, "service301_db")
	allServices = append(allServices, service301_db)
	service301 := workflow.Service[large_scale_app.Service301](spec, "service301", service301_db)
	service301_ctr := applyDockerDefaults(spec, service301, "service301_proc", "service301_container")
	containers = append(containers, service301_ctr)
	allServices = append(allServices, service301)

	//----- service300 (terminal) ----------
	service300_db := mongodb.Container(spec, "service300_db")
	allServices = append(allServices, service300_db)
	service300 := workflow.Service[large_scale_app.Service300](spec, "service300", service300_db)
	service300_ctr := applyDockerDefaults(spec, service300, "service300_proc", "service300_container")
	containers = append(containers, service300_ctr)
	allServices = append(allServices, service300)

	//----- service299 (terminal) ----------
	service299_db := mongodb.Container(spec, "service299_db")
	allServices = append(allServices, service299_db)
	service299 := workflow.Service[large_scale_app.Service299](spec, "service299", service299_db)
	service299_ctr := applyDockerDefaults(spec, service299, "service299_proc", "service299_container")
	containers = append(containers, service299_ctr)
	allServices = append(allServices, service299)

	//----- service298 (terminal) ----------
	service298_db := mongodb.Container(spec, "service298_db")
	allServices = append(allServices, service298_db)
	service298 := workflow.Service[large_scale_app.Service298](spec, "service298", service298_db)
	service298_ctr := applyDockerDefaults(spec, service298, "service298_proc", "service298_container")
	containers = append(containers, service298_ctr)
	allServices = append(allServices, service298)

	//----- service297 (terminal) ----------
	service297_db := mongodb.Container(spec, "service297_db")
	allServices = append(allServices, service297_db)
	service297 := workflow.Service[large_scale_app.Service297](spec, "service297", service297_db)
	service297_ctr := applyDockerDefaults(spec, service297, "service297_proc", "service297_container")
	containers = append(containers, service297_ctr)
	allServices = append(allServices, service297)

	//----- service296 (terminal) ----------
	service296_db := mongodb.Container(spec, "service296_db")
	allServices = append(allServices, service296_db)
	service296 := workflow.Service[large_scale_app.Service296](spec, "service296", service296_db)
	service296_ctr := applyDockerDefaults(spec, service296, "service296_proc", "service296_container")
	containers = append(containers, service296_ctr)
	allServices = append(allServices, service296)

	//----- service295 (terminal) ----------
	service295_db := mongodb.Container(spec, "service295_db")
	allServices = append(allServices, service295_db)
	service295 := workflow.Service[large_scale_app.Service295](spec, "service295", service295_db)
	service295_ctr := applyDockerDefaults(spec, service295, "service295_proc", "service295_container")
	containers = append(containers, service295_ctr)
	allServices = append(allServices, service295)

	//----- service294 (terminal) ----------
	service294_db := mongodb.Container(spec, "service294_db")
	allServices = append(allServices, service294_db)
	service294 := workflow.Service[large_scale_app.Service294](spec, "service294", service294_db)
	service294_ctr := applyDockerDefaults(spec, service294, "service294_proc", "service294_container")
	containers = append(containers, service294_ctr)
	allServices = append(allServices, service294)

	//----- service293 (terminal) ----------
	service293_db := mongodb.Container(spec, "service293_db")
	allServices = append(allServices, service293_db)
	service293 := workflow.Service[large_scale_app.Service293](spec, "service293", service293_db)
	service293_ctr := applyDockerDefaults(spec, service293, "service293_proc", "service293_container")
	containers = append(containers, service293_ctr)
	allServices = append(allServices, service293)

	//----- service292 (terminal) ----------
	service292_db := mongodb.Container(spec, "service292_db")
	allServices = append(allServices, service292_db)
	service292 := workflow.Service[large_scale_app.Service292](spec, "service292", service292_db)
	service292_ctr := applyDockerDefaults(spec, service292, "service292_proc", "service292_container")
	containers = append(containers, service292_ctr)
	allServices = append(allServices, service292)

	//----- service291 (terminal) ----------
	service291_db := mongodb.Container(spec, "service291_db")
	allServices = append(allServices, service291_db)
	service291 := workflow.Service[large_scale_app.Service291](spec, "service291", service291_db)
	service291_ctr := applyDockerDefaults(spec, service291, "service291_proc", "service291_container")
	containers = append(containers, service291_ctr)
	allServices = append(allServices, service291)

	//----- service290 (terminal) ----------
	service290_db := mongodb.Container(spec, "service290_db")
	allServices = append(allServices, service290_db)
	service290 := workflow.Service[large_scale_app.Service290](spec, "service290", service290_db)
	service290_ctr := applyDockerDefaults(spec, service290, "service290_proc", "service290_container")
	containers = append(containers, service290_ctr)
	allServices = append(allServices, service290)

	//----- service289 (terminal) ----------
	service289_db := mongodb.Container(spec, "service289_db")
	allServices = append(allServices, service289_db)
	service289 := workflow.Service[large_scale_app.Service289](spec, "service289", service289_db)
	service289_ctr := applyDockerDefaults(spec, service289, "service289_proc", "service289_container")
	containers = append(containers, service289_ctr)
	allServices = append(allServices, service289)

	//----- service288 (terminal) ----------
	service288_db := mongodb.Container(spec, "service288_db")
	allServices = append(allServices, service288_db)
	service288 := workflow.Service[large_scale_app.Service288](spec, "service288", service288_db)
	service288_ctr := applyDockerDefaults(spec, service288, "service288_proc", "service288_container")
	containers = append(containers, service288_ctr)
	allServices = append(allServices, service288)

	//----- service287 (terminal) ----------
	service287_db := mongodb.Container(spec, "service287_db")
	allServices = append(allServices, service287_db)
	service287 := workflow.Service[large_scale_app.Service287](spec, "service287", service287_db)
	service287_ctr := applyDockerDefaults(spec, service287, "service287_proc", "service287_container")
	containers = append(containers, service287_ctr)
	allServices = append(allServices, service287)

	//----- service286 (terminal) ----------
	service286_db := mongodb.Container(spec, "service286_db")
	allServices = append(allServices, service286_db)
	service286 := workflow.Service[large_scale_app.Service286](spec, "service286", service286_db)
	service286_ctr := applyDockerDefaults(spec, service286, "service286_proc", "service286_container")
	containers = append(containers, service286_ctr)
	allServices = append(allServices, service286)

	//----- service285 (terminal) ----------
	service285_db := mongodb.Container(spec, "service285_db")
	allServices = append(allServices, service285_db)
	service285 := workflow.Service[large_scale_app.Service285](spec, "service285", service285_db)
	service285_ctr := applyDockerDefaults(spec, service285, "service285_proc", "service285_container")
	containers = append(containers, service285_ctr)
	allServices = append(allServices, service285)

	//----- service284 (terminal) ----------
	service284_db := mongodb.Container(spec, "service284_db")
	allServices = append(allServices, service284_db)
	service284 := workflow.Service[large_scale_app.Service284](spec, "service284", service284_db)
	service284_ctr := applyDockerDefaults(spec, service284, "service284_proc", "service284_container")
	containers = append(containers, service284_ctr)
	allServices = append(allServices, service284)

	//----- service283 (terminal) ----------
	service283_db := mongodb.Container(spec, "service283_db")
	allServices = append(allServices, service283_db)
	service283 := workflow.Service[large_scale_app.Service283](spec, "service283", service283_db)
	service283_ctr := applyDockerDefaults(spec, service283, "service283_proc", "service283_container")
	containers = append(containers, service283_ctr)
	allServices = append(allServices, service283)

	//----- service282 (terminal) ----------
	service282_db := mongodb.Container(spec, "service282_db")
	allServices = append(allServices, service282_db)
	service282 := workflow.Service[large_scale_app.Service282](spec, "service282", service282_db)
	service282_ctr := applyDockerDefaults(spec, service282, "service282_proc", "service282_container")
	containers = append(containers, service282_ctr)
	allServices = append(allServices, service282)

	//----- service281 (terminal) ----------
	service281_db := mongodb.Container(spec, "service281_db")
	allServices = append(allServices, service281_db)
	service281 := workflow.Service[large_scale_app.Service281](spec, "service281", service281_db)
	service281_ctr := applyDockerDefaults(spec, service281, "service281_proc", "service281_container")
	containers = append(containers, service281_ctr)
	allServices = append(allServices, service281)

	//----- service280 (terminal) ----------
	service280_db := mongodb.Container(spec, "service280_db")
	allServices = append(allServices, service280_db)
	service280 := workflow.Service[large_scale_app.Service280](spec, "service280", service280_db)
	service280_ctr := applyDockerDefaults(spec, service280, "service280_proc", "service280_container")
	containers = append(containers, service280_ctr)
	allServices = append(allServices, service280)

	//----- service279 (terminal) ----------
	service279_db := mongodb.Container(spec, "service279_db")
	allServices = append(allServices, service279_db)
	service279 := workflow.Service[large_scale_app.Service279](spec, "service279", service279_db)
	service279_ctr := applyDockerDefaults(spec, service279, "service279_proc", "service279_container")
	containers = append(containers, service279_ctr)
	allServices = append(allServices, service279)

	//----- service278 (terminal) ----------
	service278_db := mongodb.Container(spec, "service278_db")
	allServices = append(allServices, service278_db)
	service278 := workflow.Service[large_scale_app.Service278](spec, "service278", service278_db)
	service278_ctr := applyDockerDefaults(spec, service278, "service278_proc", "service278_container")
	containers = append(containers, service278_ctr)
	allServices = append(allServices, service278)

	//----- service277 (terminal) ----------
	service277_db := mongodb.Container(spec, "service277_db")
	allServices = append(allServices, service277_db)
	service277 := workflow.Service[large_scale_app.Service277](spec, "service277", service277_db)
	service277_ctr := applyDockerDefaults(spec, service277, "service277_proc", "service277_container")
	containers = append(containers, service277_ctr)
	allServices = append(allServices, service277)

	//----- service276 (terminal) ----------
	service276_db := mongodb.Container(spec, "service276_db")
	allServices = append(allServices, service276_db)
	service276 := workflow.Service[large_scale_app.Service276](spec, "service276", service276_db)
	service276_ctr := applyDockerDefaults(spec, service276, "service276_proc", "service276_container")
	containers = append(containers, service276_ctr)
	allServices = append(allServices, service276)

	//----- service275 (terminal) ----------
	service275_db := mongodb.Container(spec, "service275_db")
	allServices = append(allServices, service275_db)
	service275 := workflow.Service[large_scale_app.Service275](spec, "service275", service275_db)
	service275_ctr := applyDockerDefaults(spec, service275, "service275_proc", "service275_container")
	containers = append(containers, service275_ctr)
	allServices = append(allServices, service275)

	//----- service274 (terminal) ----------
	service274_db := mongodb.Container(spec, "service274_db")
	allServices = append(allServices, service274_db)
	service274 := workflow.Service[large_scale_app.Service274](spec, "service274", service274_db)
	service274_ctr := applyDockerDefaults(spec, service274, "service274_proc", "service274_container")
	containers = append(containers, service274_ctr)
	allServices = append(allServices, service274)

	//----- service273 (terminal) ----------
	service273_db := mongodb.Container(spec, "service273_db")
	allServices = append(allServices, service273_db)
	service273 := workflow.Service[large_scale_app.Service273](spec, "service273", service273_db)
	service273_ctr := applyDockerDefaults(spec, service273, "service273_proc", "service273_container")
	containers = append(containers, service273_ctr)
	allServices = append(allServices, service273)

	//----- service272 (terminal) ----------
	service272_db := mongodb.Container(spec, "service272_db")
	allServices = append(allServices, service272_db)
	service272 := workflow.Service[large_scale_app.Service272](spec, "service272", service272_db)
	service272_ctr := applyDockerDefaults(spec, service272, "service272_proc", "service272_container")
	containers = append(containers, service272_ctr)
	allServices = append(allServices, service272)

	//----- service271 (terminal) ----------
	service271_db := mongodb.Container(spec, "service271_db")
	allServices = append(allServices, service271_db)
	service271 := workflow.Service[large_scale_app.Service271](spec, "service271", service271_db)
	service271_ctr := applyDockerDefaults(spec, service271, "service271_proc", "service271_container")
	containers = append(containers, service271_ctr)
	allServices = append(allServices, service271)

	//----- service270 (terminal) ----------
	service270_db := mongodb.Container(spec, "service270_db")
	allServices = append(allServices, service270_db)
	service270 := workflow.Service[large_scale_app.Service270](spec, "service270", service270_db)
	service270_ctr := applyDockerDefaults(spec, service270, "service270_proc", "service270_container")
	containers = append(containers, service270_ctr)
	allServices = append(allServices, service270)

	//----- service269 (terminal) ----------
	service269_db := mongodb.Container(spec, "service269_db")
	allServices = append(allServices, service269_db)
	service269 := workflow.Service[large_scale_app.Service269](spec, "service269", service269_db)
	service269_ctr := applyDockerDefaults(spec, service269, "service269_proc", "service269_container")
	containers = append(containers, service269_ctr)
	allServices = append(allServices, service269)

	//----- service268 (terminal) ----------
	service268_db := mongodb.Container(spec, "service268_db")
	allServices = append(allServices, service268_db)
	service268 := workflow.Service[large_scale_app.Service268](spec, "service268", service268_db)
	service268_ctr := applyDockerDefaults(spec, service268, "service268_proc", "service268_container")
	containers = append(containers, service268_ctr)
	allServices = append(allServices, service268)

	//----- service267 (terminal) ----------
	service267_db := mongodb.Container(spec, "service267_db")
	allServices = append(allServices, service267_db)
	service267 := workflow.Service[large_scale_app.Service267](spec, "service267", service267_db)
	service267_ctr := applyDockerDefaults(spec, service267, "service267_proc", "service267_container")
	containers = append(containers, service267_ctr)
	allServices = append(allServices, service267)

	//----- service266 (terminal) ----------
	service266_db := mongodb.Container(spec, "service266_db")
	allServices = append(allServices, service266_db)
	service266 := workflow.Service[large_scale_app.Service266](spec, "service266", service266_db)
	service266_ctr := applyDockerDefaults(spec, service266, "service266_proc", "service266_container")
	containers = append(containers, service266_ctr)
	allServices = append(allServices, service266)

	//----- service265 (terminal) ----------
	service265_db := mongodb.Container(spec, "service265_db")
	allServices = append(allServices, service265_db)
	service265 := workflow.Service[large_scale_app.Service265](spec, "service265", service265_db)
	service265_ctr := applyDockerDefaults(spec, service265, "service265_proc", "service265_container")
	containers = append(containers, service265_ctr)
	allServices = append(allServices, service265)

	//----- service264 (terminal) ----------
	service264_db := mongodb.Container(spec, "service264_db")
	allServices = append(allServices, service264_db)
	service264 := workflow.Service[large_scale_app.Service264](spec, "service264", service264_db)
	service264_ctr := applyDockerDefaults(spec, service264, "service264_proc", "service264_container")
	containers = append(containers, service264_ctr)
	allServices = append(allServices, service264)

	//----- service263 (terminal) ----------
	service263_db := mongodb.Container(spec, "service263_db")
	allServices = append(allServices, service263_db)
	service263 := workflow.Service[large_scale_app.Service263](spec, "service263", service263_db)
	service263_ctr := applyDockerDefaults(spec, service263, "service263_proc", "service263_container")
	containers = append(containers, service263_ctr)
	allServices = append(allServices, service263)

	//----- service262 (terminal) ----------
	service262_db := mongodb.Container(spec, "service262_db")
	allServices = append(allServices, service262_db)
	service262 := workflow.Service[large_scale_app.Service262](spec, "service262", service262_db)
	service262_ctr := applyDockerDefaults(spec, service262, "service262_proc", "service262_container")
	containers = append(containers, service262_ctr)
	allServices = append(allServices, service262)

	//----- service261 (terminal) ----------
	service261_db := mongodb.Container(spec, "service261_db")
	allServices = append(allServices, service261_db)
	service261 := workflow.Service[large_scale_app.Service261](spec, "service261", service261_db)
	service261_ctr := applyDockerDefaults(spec, service261, "service261_proc", "service261_container")
	containers = append(containers, service261_ctr)
	allServices = append(allServices, service261)

	//----- service260 (terminal) ----------
	service260_db := mongodb.Container(spec, "service260_db")
	allServices = append(allServices, service260_db)
	service260 := workflow.Service[large_scale_app.Service260](spec, "service260", service260_db)
	service260_ctr := applyDockerDefaults(spec, service260, "service260_proc", "service260_container")
	containers = append(containers, service260_ctr)
	allServices = append(allServices, service260)

	//----- service259 (terminal) ----------
	service259_db := mongodb.Container(spec, "service259_db")
	allServices = append(allServices, service259_db)
	service259 := workflow.Service[large_scale_app.Service259](spec, "service259", service259_db)
	service259_ctr := applyDockerDefaults(spec, service259, "service259_proc", "service259_container")
	containers = append(containers, service259_ctr)
	allServices = append(allServices, service259)

	//----- service86 (terminal) ----------
	service86_db := mongodb.Container(spec, "service86_db")
	allServices = append(allServices, service86_db)
	service86 := workflow.Service[large_scale_app.Service86](spec, "service86", service86_db)
	service86_ctr := applyDockerDefaults(spec, service86, "service86_proc", "service86_container")
	containers = append(containers, service86_ctr)
	allServices = append(allServices, service86)

	//----- service87 (terminal) ----------
	service87_db := mongodb.Container(spec, "service87_db")
	allServices = append(allServices, service87_db)
	service87 := workflow.Service[large_scale_app.Service87](spec, "service87", service87_db)
	service87_ctr := applyDockerDefaults(spec, service87, "service87_proc", "service87_container")
	containers = append(containers, service87_ctr)
	allServices = append(allServices, service87)

	//----- service258 (terminal) ----------
	service258_db := mongodb.Container(spec, "service258_db")
	allServices = append(allServices, service258_db)
	service258 := workflow.Service[large_scale_app.Service258](spec, "service258", service258_db)
	service258_ctr := applyDockerDefaults(spec, service258, "service258_proc", "service258_container")
	containers = append(containers, service258_ctr)
	allServices = append(allServices, service258)

	//----- service89 (terminal) ----------
	service89_db := mongodb.Container(spec, "service89_db")
	allServices = append(allServices, service89_db)
	service89 := workflow.Service[large_scale_app.Service89](spec, "service89", service89_db)
	service89_ctr := applyDockerDefaults(spec, service89, "service89_proc", "service89_container")
	containers = append(containers, service89_ctr)
	allServices = append(allServices, service89)

	//----- service90 (terminal) ----------
	service90_db := mongodb.Container(spec, "service90_db")
	allServices = append(allServices, service90_db)
	service90 := workflow.Service[large_scale_app.Service90](spec, "service90", service90_db)
	service90_ctr := applyDockerDefaults(spec, service90, "service90_proc", "service90_container")
	containers = append(containers, service90_ctr)
	allServices = append(allServices, service90)

	//----- service91 (terminal) ----------
	service91_db := mongodb.Container(spec, "service91_db")
	allServices = append(allServices, service91_db)
	service91 := workflow.Service[large_scale_app.Service91](spec, "service91", service91_db)
	service91_ctr := applyDockerDefaults(spec, service91, "service91_proc", "service91_container")
	containers = append(containers, service91_ctr)
	allServices = append(allServices, service91)

	//----- service92 (terminal) ----------
	service92_db := mongodb.Container(spec, "service92_db")
	allServices = append(allServices, service92_db)
	service92 := workflow.Service[large_scale_app.Service92](spec, "service92", service92_db)
	service92_ctr := applyDockerDefaults(spec, service92, "service92_proc", "service92_container")
	containers = append(containers, service92_ctr)
	allServices = append(allServices, service92)

	//----- service93 (terminal) ----------
	service93_db := mongodb.Container(spec, "service93_db")
	allServices = append(allServices, service93_db)
	service93 := workflow.Service[large_scale_app.Service93](spec, "service93", service93_db)
	service93_ctr := applyDockerDefaults(spec, service93, "service93_proc", "service93_container")
	containers = append(containers, service93_ctr)
	allServices = append(allServices, service93)

	//----- service94 (terminal) ----------
	service94_db := mongodb.Container(spec, "service94_db")
	allServices = append(allServices, service94_db)
	service94 := workflow.Service[large_scale_app.Service94](spec, "service94", service94_db)
	service94_ctr := applyDockerDefaults(spec, service94, "service94_proc", "service94_container")
	containers = append(containers, service94_ctr)
	allServices = append(allServices, service94)

	//----- service95 (terminal) ----------
	service95_db := mongodb.Container(spec, "service95_db")
	allServices = append(allServices, service95_db)
	service95 := workflow.Service[large_scale_app.Service95](spec, "service95", service95_db)
	service95_ctr := applyDockerDefaults(spec, service95, "service95_proc", "service95_container")
	containers = append(containers, service95_ctr)
	allServices = append(allServices, service95)

	//----- service96 (terminal) ----------
	service96_db := mongodb.Container(spec, "service96_db")
	allServices = append(allServices, service96_db)
	service96 := workflow.Service[large_scale_app.Service96](spec, "service96", service96_db)
	service96_ctr := applyDockerDefaults(spec, service96, "service96_proc", "service96_container")
	containers = append(containers, service96_ctr)
	allServices = append(allServices, service96)

	//----- service97 (terminal) ----------
	service97_db := mongodb.Container(spec, "service97_db")
	allServices = append(allServices, service97_db)
	service97 := workflow.Service[large_scale_app.Service97](spec, "service97", service97_db)
	service97_ctr := applyDockerDefaults(spec, service97, "service97_proc", "service97_container")
	containers = append(containers, service97_ctr)
	allServices = append(allServices, service97)

	//----- service98 (terminal) ----------
	service98_db := mongodb.Container(spec, "service98_db")
	allServices = append(allServices, service98_db)
	service98 := workflow.Service[large_scale_app.Service98](spec, "service98", service98_db)
	service98_ctr := applyDockerDefaults(spec, service98, "service98_proc", "service98_container")
	containers = append(containers, service98_ctr)
	allServices = append(allServices, service98)

	//----- service99 (terminal) ----------
	service99_db := mongodb.Container(spec, "service99_db")
	allServices = append(allServices, service99_db)
	service99 := workflow.Service[large_scale_app.Service99](spec, "service99", service99_db)
	service99_ctr := applyDockerDefaults(spec, service99, "service99_proc", "service99_container")
	containers = append(containers, service99_ctr)
	allServices = append(allServices, service99)

	//----- service100 (terminal) ----------
	service100_db := mongodb.Container(spec, "service100_db")
	allServices = append(allServices, service100_db)
	service100 := workflow.Service[large_scale_app.Service100](spec, "service100", service100_db)
	service100_ctr := applyDockerDefaults(spec, service100, "service100_proc", "service100_container")
	containers = append(containers, service100_ctr)
	allServices = append(allServices, service100)

	//----- service101 (terminal) ----------
	service101_db := mongodb.Container(spec, "service101_db")
	allServices = append(allServices, service101_db)
	service101 := workflow.Service[large_scale_app.Service101](spec, "service101", service101_db)
	service101_ctr := applyDockerDefaults(spec, service101, "service101_proc", "service101_container")
	containers = append(containers, service101_ctr)
	allServices = append(allServices, service101)

	//----- service102 (terminal) ----------
	service102_db := mongodb.Container(spec, "service102_db")
	allServices = append(allServices, service102_db)
	service102 := workflow.Service[large_scale_app.Service102](spec, "service102", service102_db)
	service102_ctr := applyDockerDefaults(spec, service102, "service102_proc", "service102_container")
	containers = append(containers, service102_ctr)
	allServices = append(allServices, service102)

	//----- service103 (terminal) ----------
	service103_db := mongodb.Container(spec, "service103_db")
	allServices = append(allServices, service103_db)
	service103 := workflow.Service[large_scale_app.Service103](spec, "service103", service103_db)
	service103_ctr := applyDockerDefaults(spec, service103, "service103_proc", "service103_container")
	containers = append(containers, service103_ctr)
	allServices = append(allServices, service103)

	//----- service104 (terminal) ----------
	service104_db := mongodb.Container(spec, "service104_db")
	allServices = append(allServices, service104_db)
	service104 := workflow.Service[large_scale_app.Service104](spec, "service104", service104_db)
	service104_ctr := applyDockerDefaults(spec, service104, "service104_proc", "service104_container")
	containers = append(containers, service104_ctr)
	allServices = append(allServices, service104)

	//----- service105 (terminal) ----------
	service105_db := mongodb.Container(spec, "service105_db")
	allServices = append(allServices, service105_db)
	service105 := workflow.Service[large_scale_app.Service105](spec, "service105", service105_db)
	service105_ctr := applyDockerDefaults(spec, service105, "service105_proc", "service105_container")
	containers = append(containers, service105_ctr)
	allServices = append(allServices, service105)

	//----- service106 (terminal) ----------
	service106_db := mongodb.Container(spec, "service106_db")
	allServices = append(allServices, service106_db)
	service106 := workflow.Service[large_scale_app.Service106](spec, "service106", service106_db)
	service106_ctr := applyDockerDefaults(spec, service106, "service106_proc", "service106_container")
	containers = append(containers, service106_ctr)
	allServices = append(allServices, service106)

	//----- service107 (terminal) ----------
	service107_db := mongodb.Container(spec, "service107_db")
	allServices = append(allServices, service107_db)
	service107 := workflow.Service[large_scale_app.Service107](spec, "service107", service107_db)
	service107_ctr := applyDockerDefaults(spec, service107, "service107_proc", "service107_container")
	containers = append(containers, service107_ctr)
	allServices = append(allServices, service107)

	//----- service108 (terminal) ----------
	service108_db := mongodb.Container(spec, "service108_db")
	allServices = append(allServices, service108_db)
	service108 := workflow.Service[large_scale_app.Service108](spec, "service108", service108_db)
	service108_ctr := applyDockerDefaults(spec, service108, "service108_proc", "service108_container")
	containers = append(containers, service108_ctr)
	allServices = append(allServices, service108)

	//----- service109 (terminal) ----------
	service109_db := mongodb.Container(spec, "service109_db")
	allServices = append(allServices, service109_db)
	service109 := workflow.Service[large_scale_app.Service109](spec, "service109", service109_db)
	service109_ctr := applyDockerDefaults(spec, service109, "service109_proc", "service109_container")
	containers = append(containers, service109_ctr)
	allServices = append(allServices, service109)

	//----- service110 (terminal) ----------
	service110_db := mongodb.Container(spec, "service110_db")
	allServices = append(allServices, service110_db)
	service110 := workflow.Service[large_scale_app.Service110](spec, "service110", service110_db)
	service110_ctr := applyDockerDefaults(spec, service110, "service110_proc", "service110_container")
	containers = append(containers, service110_ctr)
	allServices = append(allServices, service110)

	//----- service111 (terminal) ----------
	service111_db := mongodb.Container(spec, "service111_db")
	allServices = append(allServices, service111_db)
	service111 := workflow.Service[large_scale_app.Service111](spec, "service111", service111_db)
	service111_ctr := applyDockerDefaults(spec, service111, "service111_proc", "service111_container")
	containers = append(containers, service111_ctr)
	allServices = append(allServices, service111)

	//----- service112 (terminal) ----------
	service112_db := mongodb.Container(spec, "service112_db")
	allServices = append(allServices, service112_db)
	service112 := workflow.Service[large_scale_app.Service112](spec, "service112", service112_db)
	service112_ctr := applyDockerDefaults(spec, service112, "service112_proc", "service112_container")
	containers = append(containers, service112_ctr)
	allServices = append(allServices, service112)

	//----- service113 (terminal) ----------
	service113_db := mongodb.Container(spec, "service113_db")
	allServices = append(allServices, service113_db)
	service113 := workflow.Service[large_scale_app.Service113](spec, "service113", service113_db)
	service113_ctr := applyDockerDefaults(spec, service113, "service113_proc", "service113_container")
	containers = append(containers, service113_ctr)
	allServices = append(allServices, service113)

	//----- service114 (terminal) ----------
	service114_db := mongodb.Container(spec, "service114_db")
	allServices = append(allServices, service114_db)
	service114 := workflow.Service[large_scale_app.Service114](spec, "service114", service114_db)
	service114_ctr := applyDockerDefaults(spec, service114, "service114_proc", "service114_container")
	containers = append(containers, service114_ctr)
	allServices = append(allServices, service114)

	//----- service115 (terminal) ----------
	service115_db := mongodb.Container(spec, "service115_db")
	allServices = append(allServices, service115_db)
	service115 := workflow.Service[large_scale_app.Service115](spec, "service115", service115_db)
	service115_ctr := applyDockerDefaults(spec, service115, "service115_proc", "service115_container")
	containers = append(containers, service115_ctr)
	allServices = append(allServices, service115)

	//----- service116 (terminal) ----------
	service116_db := mongodb.Container(spec, "service116_db")
	allServices = append(allServices, service116_db)
	service116 := workflow.Service[large_scale_app.Service116](spec, "service116", service116_db)
	service116_ctr := applyDockerDefaults(spec, service116, "service116_proc", "service116_container")
	containers = append(containers, service116_ctr)
	allServices = append(allServices, service116)

	//----- service117 (terminal) ----------
	service117_db := mongodb.Container(spec, "service117_db")
	allServices = append(allServices, service117_db)
	service117 := workflow.Service[large_scale_app.Service117](spec, "service117", service117_db)
	service117_ctr := applyDockerDefaults(spec, service117, "service117_proc", "service117_container")
	containers = append(containers, service117_ctr)
	allServices = append(allServices, service117)

	//----- service118 (terminal) ----------
	service118_db := mongodb.Container(spec, "service118_db")
	allServices = append(allServices, service118_db)
	service118 := workflow.Service[large_scale_app.Service118](spec, "service118", service118_db)
	service118_ctr := applyDockerDefaults(spec, service118, "service118_proc", "service118_container")
	containers = append(containers, service118_ctr)
	allServices = append(allServices, service118)

	//----- service119 (terminal) ----------
	service119_db := mongodb.Container(spec, "service119_db")
	allServices = append(allServices, service119_db)
	service119 := workflow.Service[large_scale_app.Service119](spec, "service119", service119_db)
	service119_ctr := applyDockerDefaults(spec, service119, "service119_proc", "service119_container")
	containers = append(containers, service119_ctr)
	allServices = append(allServices, service119)

	//----- service120 (terminal) ----------
	service120_db := mongodb.Container(spec, "service120_db")
	allServices = append(allServices, service120_db)
	service120 := workflow.Service[large_scale_app.Service120](spec, "service120", service120_db)
	service120_ctr := applyDockerDefaults(spec, service120, "service120_proc", "service120_container")
	containers = append(containers, service120_ctr)
	allServices = append(allServices, service120)

	//----- service121 (terminal) ----------
	service121_db := mongodb.Container(spec, "service121_db")
	allServices = append(allServices, service121_db)
	service121 := workflow.Service[large_scale_app.Service121](spec, "service121", service121_db)
	service121_ctr := applyDockerDefaults(spec, service121, "service121_proc", "service121_container")
	containers = append(containers, service121_ctr)
	allServices = append(allServices, service121)

	//----- service122 (terminal) ----------
	service122_db := mongodb.Container(spec, "service122_db")
	allServices = append(allServices, service122_db)
	service122 := workflow.Service[large_scale_app.Service122](spec, "service122", service122_db)
	service122_ctr := applyDockerDefaults(spec, service122, "service122_proc", "service122_container")
	containers = append(containers, service122_ctr)
	allServices = append(allServices, service122)

	//----- service123 (terminal) ----------
	service123_db := mongodb.Container(spec, "service123_db")
	allServices = append(allServices, service123_db)
	service123 := workflow.Service[large_scale_app.Service123](spec, "service123", service123_db)
	service123_ctr := applyDockerDefaults(spec, service123, "service123_proc", "service123_container")
	containers = append(containers, service123_ctr)
	allServices = append(allServices, service123)

	//----- service124 (terminal) ----------
	service124_db := mongodb.Container(spec, "service124_db")
	allServices = append(allServices, service124_db)
	service124 := workflow.Service[large_scale_app.Service124](spec, "service124", service124_db)
	service124_ctr := applyDockerDefaults(spec, service124, "service124_proc", "service124_container")
	containers = append(containers, service124_ctr)
	allServices = append(allServices, service124)

	//----- service125 (terminal) ----------
	service125_db := mongodb.Container(spec, "service125_db")
	allServices = append(allServices, service125_db)
	service125 := workflow.Service[large_scale_app.Service125](spec, "service125", service125_db)
	service125_ctr := applyDockerDefaults(spec, service125, "service125_proc", "service125_container")
	containers = append(containers, service125_ctr)
	allServices = append(allServices, service125)

	//----- service126 (terminal) ----------
	service126_db := mongodb.Container(spec, "service126_db")
	allServices = append(allServices, service126_db)
	service126 := workflow.Service[large_scale_app.Service126](spec, "service126", service126_db)
	service126_ctr := applyDockerDefaults(spec, service126, "service126_proc", "service126_container")
	containers = append(containers, service126_ctr)
	allServices = append(allServices, service126)

	//----- service127 (terminal) ----------
	service127_db := mongodb.Container(spec, "service127_db")
	allServices = append(allServices, service127_db)
	service127 := workflow.Service[large_scale_app.Service127](spec, "service127", service127_db)
	service127_ctr := applyDockerDefaults(spec, service127, "service127_proc", "service127_container")
	containers = append(containers, service127_ctr)
	allServices = append(allServices, service127)

	//----- service172 (terminal) ----------
	service172_db := mongodb.Container(spec, "service172_db")
	allServices = append(allServices, service172_db)
	service172 := workflow.Service[large_scale_app.Service172](spec, "service172", service172_db)
	service172_ctr := applyDockerDefaults(spec, service172, "service172_proc", "service172_container")
	containers = append(containers, service172_ctr)
	allServices = append(allServices, service172)

	//----- service129 (terminal) ----------
	service129_db := mongodb.Container(spec, "service129_db")
	allServices = append(allServices, service129_db)
	service129 := workflow.Service[large_scale_app.Service129](spec, "service129", service129_db)
	service129_ctr := applyDockerDefaults(spec, service129, "service129_proc", "service129_container")
	containers = append(containers, service129_ctr)
	allServices = append(allServices, service129)

	//----- service130 (terminal) ----------
	service130_db := mongodb.Container(spec, "service130_db")
	allServices = append(allServices, service130_db)
	service130 := workflow.Service[large_scale_app.Service130](spec, "service130", service130_db)
	service130_ctr := applyDockerDefaults(spec, service130, "service130_proc", "service130_container")
	containers = append(containers, service130_ctr)
	allServices = append(allServices, service130)

	//----- service131 (terminal) ----------
	service131_db := mongodb.Container(spec, "service131_db")
	allServices = append(allServices, service131_db)
	service131 := workflow.Service[large_scale_app.Service131](spec, "service131", service131_db)
	service131_ctr := applyDockerDefaults(spec, service131, "service131_proc", "service131_container")
	containers = append(containers, service131_ctr)
	allServices = append(allServices, service131)

	//----- service132 (terminal) ----------
	service132_db := mongodb.Container(spec, "service132_db")
	allServices = append(allServices, service132_db)
	service132 := workflow.Service[large_scale_app.Service132](spec, "service132", service132_db)
	service132_ctr := applyDockerDefaults(spec, service132, "service132_proc", "service132_container")
	containers = append(containers, service132_ctr)
	allServices = append(allServices, service132)

	//----- service133 (terminal) ----------
	service133_db := mongodb.Container(spec, "service133_db")
	allServices = append(allServices, service133_db)
	service133 := workflow.Service[large_scale_app.Service133](spec, "service133", service133_db)
	service133_ctr := applyDockerDefaults(spec, service133, "service133_proc", "service133_container")
	containers = append(containers, service133_ctr)
	allServices = append(allServices, service133)

	//----- service134 (terminal) ----------
	service134_db := mongodb.Container(spec, "service134_db")
	allServices = append(allServices, service134_db)
	service134 := workflow.Service[large_scale_app.Service134](spec, "service134", service134_db)
	service134_ctr := applyDockerDefaults(spec, service134, "service134_proc", "service134_container")
	containers = append(containers, service134_ctr)
	allServices = append(allServices, service134)

	//----- service135 (terminal) ----------
	service135_db := mongodb.Container(spec, "service135_db")
	allServices = append(allServices, service135_db)
	service135 := workflow.Service[large_scale_app.Service135](spec, "service135", service135_db)
	service135_ctr := applyDockerDefaults(spec, service135, "service135_proc", "service135_container")
	containers = append(containers, service135_ctr)
	allServices = append(allServices, service135)

	//----- service136 (terminal) ----------
	service136_db := mongodb.Container(spec, "service136_db")
	allServices = append(allServices, service136_db)
	service136 := workflow.Service[large_scale_app.Service136](spec, "service136", service136_db)
	service136_ctr := applyDockerDefaults(spec, service136, "service136_proc", "service136_container")
	containers = append(containers, service136_ctr)
	allServices = append(allServices, service136)

	//----- service137 (terminal) ----------
	service137_db := mongodb.Container(spec, "service137_db")
	allServices = append(allServices, service137_db)
	service137 := workflow.Service[large_scale_app.Service137](spec, "service137", service137_db)
	service137_ctr := applyDockerDefaults(spec, service137, "service137_proc", "service137_container")
	containers = append(containers, service137_ctr)
	allServices = append(allServices, service137)

	//----- service138 (terminal) ----------
	service138_db := mongodb.Container(spec, "service138_db")
	allServices = append(allServices, service138_db)
	service138 := workflow.Service[large_scale_app.Service138](spec, "service138", service138_db)
	service138_ctr := applyDockerDefaults(spec, service138, "service138_proc", "service138_container")
	containers = append(containers, service138_ctr)
	allServices = append(allServices, service138)

	//----- service139 (terminal) ----------
	service139_db := mongodb.Container(spec, "service139_db")
	allServices = append(allServices, service139_db)
	service139 := workflow.Service[large_scale_app.Service139](spec, "service139", service139_db)
	service139_ctr := applyDockerDefaults(spec, service139, "service139_proc", "service139_container")
	containers = append(containers, service139_ctr)
	allServices = append(allServices, service139)

	//----- service140 (terminal) ----------
	service140_db := mongodb.Container(spec, "service140_db")
	allServices = append(allServices, service140_db)
	service140 := workflow.Service[large_scale_app.Service140](spec, "service140", service140_db)
	service140_ctr := applyDockerDefaults(spec, service140, "service140_proc", "service140_container")
	containers = append(containers, service140_ctr)
	allServices = append(allServices, service140)

	//----- service141 (terminal) ----------
	service141_db := mongodb.Container(spec, "service141_db")
	allServices = append(allServices, service141_db)
	service141 := workflow.Service[large_scale_app.Service141](spec, "service141", service141_db)
	service141_ctr := applyDockerDefaults(spec, service141, "service141_proc", "service141_container")
	containers = append(containers, service141_ctr)
	allServices = append(allServices, service141)

	//----- service142 (terminal) ----------
	service142_db := mongodb.Container(spec, "service142_db")
	allServices = append(allServices, service142_db)
	service142 := workflow.Service[large_scale_app.Service142](spec, "service142", service142_db)
	service142_ctr := applyDockerDefaults(spec, service142, "service142_proc", "service142_container")
	containers = append(containers, service142_ctr)
	allServices = append(allServices, service142)

	//----- service143 (terminal) ----------
	service143_db := mongodb.Container(spec, "service143_db")
	allServices = append(allServices, service143_db)
	service143 := workflow.Service[large_scale_app.Service143](spec, "service143", service143_db)
	service143_ctr := applyDockerDefaults(spec, service143, "service143_proc", "service143_container")
	containers = append(containers, service143_ctr)
	allServices = append(allServices, service143)

	//----- service144 (terminal) ----------
	service144_db := mongodb.Container(spec, "service144_db")
	allServices = append(allServices, service144_db)
	service144 := workflow.Service[large_scale_app.Service144](spec, "service144", service144_db)
	service144_ctr := applyDockerDefaults(spec, service144, "service144_proc", "service144_container")
	containers = append(containers, service144_ctr)
	allServices = append(allServices, service144)

	//----- service145 (terminal) ----------
	service145_db := mongodb.Container(spec, "service145_db")
	allServices = append(allServices, service145_db)
	service145 := workflow.Service[large_scale_app.Service145](spec, "service145", service145_db)
	service145_ctr := applyDockerDefaults(spec, service145, "service145_proc", "service145_container")
	containers = append(containers, service145_ctr)
	allServices = append(allServices, service145)

	//----- service146 (terminal) ----------
	service146_db := mongodb.Container(spec, "service146_db")
	allServices = append(allServices, service146_db)
	service146 := workflow.Service[large_scale_app.Service146](spec, "service146", service146_db)
	service146_ctr := applyDockerDefaults(spec, service146, "service146_proc", "service146_container")
	containers = append(containers, service146_ctr)
	allServices = append(allServices, service146)

	//----- service147 (terminal) ----------
	service147_db := mongodb.Container(spec, "service147_db")
	allServices = append(allServices, service147_db)
	service147 := workflow.Service[large_scale_app.Service147](spec, "service147", service147_db)
	service147_ctr := applyDockerDefaults(spec, service147, "service147_proc", "service147_container")
	containers = append(containers, service147_ctr)
	allServices = append(allServices, service147)

	//----- service173 (terminal) ----------
	service173_db := mongodb.Container(spec, "service173_db")
	allServices = append(allServices, service173_db)
	service173 := workflow.Service[large_scale_app.Service173](spec, "service173", service173_db)
	service173_ctr := applyDockerDefaults(spec, service173, "service173_proc", "service173_container")
	containers = append(containers, service173_ctr)
	allServices = append(allServices, service173)

	//----- service149 (terminal) ----------
	service149_db := mongodb.Container(spec, "service149_db")
	allServices = append(allServices, service149_db)
	service149 := workflow.Service[large_scale_app.Service149](spec, "service149", service149_db)
	service149_ctr := applyDockerDefaults(spec, service149, "service149_proc", "service149_container")
	containers = append(containers, service149_ctr)
	allServices = append(allServices, service149)

	//----- service150 (terminal) ----------
	service150_db := mongodb.Container(spec, "service150_db")
	allServices = append(allServices, service150_db)
	service150 := workflow.Service[large_scale_app.Service150](spec, "service150", service150_db)
	service150_ctr := applyDockerDefaults(spec, service150, "service150_proc", "service150_container")
	containers = append(containers, service150_ctr)
	allServices = append(allServices, service150)

	//----- service151 (terminal) ----------
	service151_db := mongodb.Container(spec, "service151_db")
	allServices = append(allServices, service151_db)
	service151 := workflow.Service[large_scale_app.Service151](spec, "service151", service151_db)
	service151_ctr := applyDockerDefaults(spec, service151, "service151_proc", "service151_container")
	containers = append(containers, service151_ctr)
	allServices = append(allServices, service151)

	//----- service152 (terminal) ----------
	service152_db := mongodb.Container(spec, "service152_db")
	allServices = append(allServices, service152_db)
	service152 := workflow.Service[large_scale_app.Service152](spec, "service152", service152_db)
	service152_ctr := applyDockerDefaults(spec, service152, "service152_proc", "service152_container")
	containers = append(containers, service152_ctr)
	allServices = append(allServices, service152)

	//----- service153 (terminal) ----------
	service153_db := mongodb.Container(spec, "service153_db")
	allServices = append(allServices, service153_db)
	service153 := workflow.Service[large_scale_app.Service153](spec, "service153", service153_db)
	service153_ctr := applyDockerDefaults(spec, service153, "service153_proc", "service153_container")
	containers = append(containers, service153_ctr)
	allServices = append(allServices, service153)

	//----- service154 (terminal) ----------
	service154_db := mongodb.Container(spec, "service154_db")
	allServices = append(allServices, service154_db)
	service154 := workflow.Service[large_scale_app.Service154](spec, "service154", service154_db)
	service154_ctr := applyDockerDefaults(spec, service154, "service154_proc", "service154_container")
	containers = append(containers, service154_ctr)
	allServices = append(allServices, service154)

	//----- service155 (terminal) ----------
	service155_db := mongodb.Container(spec, "service155_db")
	allServices = append(allServices, service155_db)
	service155 := workflow.Service[large_scale_app.Service155](spec, "service155", service155_db)
	service155_ctr := applyDockerDefaults(spec, service155, "service155_proc", "service155_container")
	containers = append(containers, service155_ctr)
	allServices = append(allServices, service155)

	//----- service156 (terminal) ----------
	service156_db := mongodb.Container(spec, "service156_db")
	allServices = append(allServices, service156_db)
	service156 := workflow.Service[large_scale_app.Service156](spec, "service156", service156_db)
	service156_ctr := applyDockerDefaults(spec, service156, "service156_proc", "service156_container")
	containers = append(containers, service156_ctr)
	allServices = append(allServices, service156)

	//----- service157 (terminal) ----------
	service157_db := mongodb.Container(spec, "service157_db")
	allServices = append(allServices, service157_db)
	service157 := workflow.Service[large_scale_app.Service157](spec, "service157", service157_db)
	service157_ctr := applyDockerDefaults(spec, service157, "service157_proc", "service157_container")
	containers = append(containers, service157_ctr)
	allServices = append(allServices, service157)

	//----- service158 (terminal) ----------
	service158_db := mongodb.Container(spec, "service158_db")
	allServices = append(allServices, service158_db)
	service158 := workflow.Service[large_scale_app.Service158](spec, "service158", service158_db)
	service158_ctr := applyDockerDefaults(spec, service158, "service158_proc", "service158_container")
	containers = append(containers, service158_ctr)
	allServices = append(allServices, service158)

	//----- service159 (terminal) ----------
	service159_db := mongodb.Container(spec, "service159_db")
	allServices = append(allServices, service159_db)
	service159 := workflow.Service[large_scale_app.Service159](spec, "service159", service159_db)
	service159_ctr := applyDockerDefaults(spec, service159, "service159_proc", "service159_container")
	containers = append(containers, service159_ctr)
	allServices = append(allServices, service159)

	//----- service160 (terminal) ----------
	service160_db := mongodb.Container(spec, "service160_db")
	allServices = append(allServices, service160_db)
	service160 := workflow.Service[large_scale_app.Service160](spec, "service160", service160_db)
	service160_ctr := applyDockerDefaults(spec, service160, "service160_proc", "service160_container")
	containers = append(containers, service160_ctr)
	allServices = append(allServices, service160)

	//----- service161 (terminal) ----------
	service161_db := mongodb.Container(spec, "service161_db")
	allServices = append(allServices, service161_db)
	service161 := workflow.Service[large_scale_app.Service161](spec, "service161", service161_db)
	service161_ctr := applyDockerDefaults(spec, service161, "service161_proc", "service161_container")
	containers = append(containers, service161_ctr)
	allServices = append(allServices, service161)

	//----- service162 (terminal) ----------
	service162_db := mongodb.Container(spec, "service162_db")
	allServices = append(allServices, service162_db)
	service162 := workflow.Service[large_scale_app.Service162](spec, "service162", service162_db)
	service162_ctr := applyDockerDefaults(spec, service162, "service162_proc", "service162_container")
	containers = append(containers, service162_ctr)
	allServices = append(allServices, service162)

	//----- service163 (terminal) ----------
	service163_db := mongodb.Container(spec, "service163_db")
	allServices = append(allServices, service163_db)
	service163 := workflow.Service[large_scale_app.Service163](spec, "service163", service163_db)
	service163_ctr := applyDockerDefaults(spec, service163, "service163_proc", "service163_container")
	containers = append(containers, service163_ctr)
	allServices = append(allServices, service163)

	//----- service164 (terminal) ----------
	service164_db := mongodb.Container(spec, "service164_db")
	allServices = append(allServices, service164_db)
	service164 := workflow.Service[large_scale_app.Service164](spec, "service164", service164_db)
	service164_ctr := applyDockerDefaults(spec, service164, "service164_proc", "service164_container")
	containers = append(containers, service164_ctr)
	allServices = append(allServices, service164)

	//----- service165 (terminal) ----------
	service165_db := mongodb.Container(spec, "service165_db")
	allServices = append(allServices, service165_db)
	service165 := workflow.Service[large_scale_app.Service165](spec, "service165", service165_db)
	service165_ctr := applyDockerDefaults(spec, service165, "service165_proc", "service165_container")
	containers = append(containers, service165_ctr)
	allServices = append(allServices, service165)

	//----- service166 (terminal) ----------
	service166_db := mongodb.Container(spec, "service166_db")
	allServices = append(allServices, service166_db)
	service166 := workflow.Service[large_scale_app.Service166](spec, "service166", service166_db)
	service166_ctr := applyDockerDefaults(spec, service166, "service166_proc", "service166_container")
	containers = append(containers, service166_ctr)
	allServices = append(allServices, service166)

	//----- service167 (terminal) ----------
	service167_db := mongodb.Container(spec, "service167_db")
	allServices = append(allServices, service167_db)
	service167 := workflow.Service[large_scale_app.Service167](spec, "service167", service167_db)
	service167_ctr := applyDockerDefaults(spec, service167, "service167_proc", "service167_container")
	containers = append(containers, service167_ctr)
	allServices = append(allServices, service167)

	//----- service168 (terminal) ----------
	service168_db := mongodb.Container(spec, "service168_db")
	allServices = append(allServices, service168_db)
	service168 := workflow.Service[large_scale_app.Service168](spec, "service168", service168_db)
	service168_ctr := applyDockerDefaults(spec, service168, "service168_proc", "service168_container")
	containers = append(containers, service168_ctr)
	allServices = append(allServices, service168)

	//----- service169 (terminal) ----------
	service169_db := mongodb.Container(spec, "service169_db")
	allServices = append(allServices, service169_db)
	service169 := workflow.Service[large_scale_app.Service169](spec, "service169", service169_db)
	service169_ctr := applyDockerDefaults(spec, service169, "service169_proc", "service169_container")
	containers = append(containers, service169_ctr)
	allServices = append(allServices, service169)

	//----- service170 (terminal) ----------
	service170_db := mongodb.Container(spec, "service170_db")
	allServices = append(allServices, service170_db)
	service170 := workflow.Service[large_scale_app.Service170](spec, "service170", service170_db)
	service170_ctr := applyDockerDefaults(spec, service170, "service170_proc", "service170_container")
	containers = append(containers, service170_ctr)
	allServices = append(allServices, service170)

	//----- service88 (terminal) ----------
	service88_db := mongodb.Container(spec, "service88_db")
	allServices = append(allServices, service88_db)
	service88 := workflow.Service[large_scale_app.Service88](spec, "service88", service88_db)
	service88_ctr := applyDockerDefaults(spec, service88, "service88_proc", "service88_container")
	containers = append(containers, service88_ctr)
	allServices = append(allServices, service88)

	//----- service257 (terminal) ----------
	service257_db := mongodb.Container(spec, "service257_db")
	allServices = append(allServices, service257_db)
	service257 := workflow.Service[large_scale_app.Service257](spec, "service257", service257_db)
	service257_ctr := applyDockerDefaults(spec, service257, "service257_proc", "service257_container")
	containers = append(containers, service257_ctr)
	allServices = append(allServices, service257)

	//----- service148 (terminal) ----------
	service148_db := mongodb.Container(spec, "service148_db")
	allServices = append(allServices, service148_db)
	service148 := workflow.Service[large_scale_app.Service148](spec, "service148", service148_db)
	service148_ctr := applyDockerDefaults(spec, service148, "service148_proc", "service148_container")
	containers = append(containers, service148_ctr)
	allServices = append(allServices, service148)

	//----- service174 (terminal) ----------
	service174_db := mongodb.Container(spec, "service174_db")
	allServices = append(allServices, service174_db)
	service174 := workflow.Service[large_scale_app.Service174](spec, "service174", service174_db)
	service174_ctr := applyDockerDefaults(spec, service174, "service174_proc", "service174_container")
	containers = append(containers, service174_ctr)
	allServices = append(allServices, service174)

	//----- service175 (terminal) ----------
	service175_db := mongodb.Container(spec, "service175_db")
	allServices = append(allServices, service175_db)
	service175 := workflow.Service[large_scale_app.Service175](spec, "service175", service175_db)
	service175_ctr := applyDockerDefaults(spec, service175, "service175_proc", "service175_container")
	containers = append(containers, service175_ctr)
	allServices = append(allServices, service175)

	//----- service176 (terminal) ----------
	service176_db := mongodb.Container(spec, "service176_db")
	allServices = append(allServices, service176_db)
	service176 := workflow.Service[large_scale_app.Service176](spec, "service176", service176_db)
	service176_ctr := applyDockerDefaults(spec, service176, "service176_proc", "service176_container")
	containers = append(containers, service176_ctr)
	allServices = append(allServices, service176)

	//----- service177 (terminal) ----------
	service177_db := mongodb.Container(spec, "service177_db")
	allServices = append(allServices, service177_db)
	service177 := workflow.Service[large_scale_app.Service177](spec, "service177", service177_db)
	service177_ctr := applyDockerDefaults(spec, service177, "service177_proc", "service177_container")
	containers = append(containers, service177_ctr)
	allServices = append(allServices, service177)

	//----- service178 (terminal) ----------
	service178_db := mongodb.Container(spec, "service178_db")
	allServices = append(allServices, service178_db)
	service178 := workflow.Service[large_scale_app.Service178](spec, "service178", service178_db)
	service178_ctr := applyDockerDefaults(spec, service178, "service178_proc", "service178_container")
	containers = append(containers, service178_ctr)
	allServices = append(allServices, service178)

	//----- service179 (terminal) ----------
	service179_db := mongodb.Container(spec, "service179_db")
	allServices = append(allServices, service179_db)
	service179 := workflow.Service[large_scale_app.Service179](spec, "service179", service179_db)
	service179_ctr := applyDockerDefaults(spec, service179, "service179_proc", "service179_container")
	containers = append(containers, service179_ctr)
	allServices = append(allServices, service179)

	//----- service180 (terminal) ----------
	service180_db := mongodb.Container(spec, "service180_db")
	allServices = append(allServices, service180_db)
	service180 := workflow.Service[large_scale_app.Service180](spec, "service180", service180_db)
	service180_ctr := applyDockerDefaults(spec, service180, "service180_proc", "service180_container")
	containers = append(containers, service180_ctr)
	allServices = append(allServices, service180)

	//----- service181 (terminal) ----------
	service181_db := mongodb.Container(spec, "service181_db")
	allServices = append(allServices, service181_db)
	service181 := workflow.Service[large_scale_app.Service181](spec, "service181", service181_db)
	service181_ctr := applyDockerDefaults(spec, service181, "service181_proc", "service181_container")
	containers = append(containers, service181_ctr)
	allServices = append(allServices, service181)

	//----- service182 (terminal) ----------
	service182_db := mongodb.Container(spec, "service182_db")
	allServices = append(allServices, service182_db)
	service182 := workflow.Service[large_scale_app.Service182](spec, "service182", service182_db)
	service182_ctr := applyDockerDefaults(spec, service182, "service182_proc", "service182_container")
	containers = append(containers, service182_ctr)
	allServices = append(allServices, service182)

	//----- service183 (terminal) ----------
	service183_db := mongodb.Container(spec, "service183_db")
	allServices = append(allServices, service183_db)
	service183 := workflow.Service[large_scale_app.Service183](spec, "service183", service183_db)
	service183_ctr := applyDockerDefaults(spec, service183, "service183_proc", "service183_container")
	containers = append(containers, service183_ctr)
	allServices = append(allServices, service183)

	//----- service184 (terminal) ----------
	service184_db := mongodb.Container(spec, "service184_db")
	allServices = append(allServices, service184_db)
	service184 := workflow.Service[large_scale_app.Service184](spec, "service184", service184_db)
	service184_ctr := applyDockerDefaults(spec, service184, "service184_proc", "service184_container")
	containers = append(containers, service184_ctr)
	allServices = append(allServices, service184)

	//----- service185 (terminal) ----------
	service185_db := mongodb.Container(spec, "service185_db")
	allServices = append(allServices, service185_db)
	service185 := workflow.Service[large_scale_app.Service185](spec, "service185", service185_db)
	service185_ctr := applyDockerDefaults(spec, service185, "service185_proc", "service185_container")
	containers = append(containers, service185_ctr)
	allServices = append(allServices, service185)

	//----- service186 (terminal) ----------
	service186_db := mongodb.Container(spec, "service186_db")
	allServices = append(allServices, service186_db)
	service186 := workflow.Service[large_scale_app.Service186](spec, "service186", service186_db)
	service186_ctr := applyDockerDefaults(spec, service186, "service186_proc", "service186_container")
	containers = append(containers, service186_ctr)
	allServices = append(allServices, service186)

	//----- service187 (terminal) ----------
	service187_db := mongodb.Container(spec, "service187_db")
	allServices = append(allServices, service187_db)
	service187 := workflow.Service[large_scale_app.Service187](spec, "service187", service187_db)
	service187_ctr := applyDockerDefaults(spec, service187, "service187_proc", "service187_container")
	containers = append(containers, service187_ctr)
	allServices = append(allServices, service187)

	//----- service188 (terminal) ----------
	service188_db := mongodb.Container(spec, "service188_db")
	allServices = append(allServices, service188_db)
	service188 := workflow.Service[large_scale_app.Service188](spec, "service188", service188_db)
	service188_ctr := applyDockerDefaults(spec, service188, "service188_proc", "service188_container")
	containers = append(containers, service188_ctr)
	allServices = append(allServices, service188)

	//----- service189 (terminal) ----------
	service189_db := mongodb.Container(spec, "service189_db")
	allServices = append(allServices, service189_db)
	service189 := workflow.Service[large_scale_app.Service189](spec, "service189", service189_db)
	service189_ctr := applyDockerDefaults(spec, service189, "service189_proc", "service189_container")
	containers = append(containers, service189_ctr)
	allServices = append(allServices, service189)

	//----- service190 (terminal) ----------
	service190_db := mongodb.Container(spec, "service190_db")
	allServices = append(allServices, service190_db)
	service190 := workflow.Service[large_scale_app.Service190](spec, "service190", service190_db)
	service190_ctr := applyDockerDefaults(spec, service190, "service190_proc", "service190_container")
	containers = append(containers, service190_ctr)
	allServices = append(allServices, service190)

	//----- service191 (terminal) ----------
	service191_db := mongodb.Container(spec, "service191_db")
	allServices = append(allServices, service191_db)
	service191 := workflow.Service[large_scale_app.Service191](spec, "service191", service191_db)
	service191_ctr := applyDockerDefaults(spec, service191, "service191_proc", "service191_container")
	containers = append(containers, service191_ctr)
	allServices = append(allServices, service191)

	//----- service192 (terminal) ----------
	service192_db := mongodb.Container(spec, "service192_db")
	allServices = append(allServices, service192_db)
	service192 := workflow.Service[large_scale_app.Service192](spec, "service192", service192_db)
	service192_ctr := applyDockerDefaults(spec, service192, "service192_proc", "service192_container")
	containers = append(containers, service192_ctr)
	allServices = append(allServices, service192)

	//----- service193 (terminal) ----------
	service193_db := mongodb.Container(spec, "service193_db")
	allServices = append(allServices, service193_db)
	service193 := workflow.Service[large_scale_app.Service193](spec, "service193", service193_db)
	service193_ctr := applyDockerDefaults(spec, service193, "service193_proc", "service193_container")
	containers = append(containers, service193_ctr)
	allServices = append(allServices, service193)

	//----- service194 (terminal) ----------
	service194_db := mongodb.Container(spec, "service194_db")
	allServices = append(allServices, service194_db)
	service194 := workflow.Service[large_scale_app.Service194](spec, "service194", service194_db)
	service194_ctr := applyDockerDefaults(spec, service194, "service194_proc", "service194_container")
	containers = append(containers, service194_ctr)
	allServices = append(allServices, service194)

	//----- service195 (terminal) ----------
	service195_db := mongodb.Container(spec, "service195_db")
	allServices = append(allServices, service195_db)
	service195 := workflow.Service[large_scale_app.Service195](spec, "service195", service195_db)
	service195_ctr := applyDockerDefaults(spec, service195, "service195_proc", "service195_container")
	containers = append(containers, service195_ctr)
	allServices = append(allServices, service195)

	//----- service196 (terminal) ----------
	service196_db := mongodb.Container(spec, "service196_db")
	allServices = append(allServices, service196_db)
	service196 := workflow.Service[large_scale_app.Service196](spec, "service196", service196_db)
	service196_ctr := applyDockerDefaults(spec, service196, "service196_proc", "service196_container")
	containers = append(containers, service196_ctr)
	allServices = append(allServices, service196)

	//----- service197 (terminal) ----------
	service197_db := mongodb.Container(spec, "service197_db")
	allServices = append(allServices, service197_db)
	service197 := workflow.Service[large_scale_app.Service197](spec, "service197", service197_db)
	service197_ctr := applyDockerDefaults(spec, service197, "service197_proc", "service197_container")
	containers = append(containers, service197_ctr)
	allServices = append(allServices, service197)

	//----- service198 (terminal) ----------
	service198_db := mongodb.Container(spec, "service198_db")
	allServices = append(allServices, service198_db)
	service198 := workflow.Service[large_scale_app.Service198](spec, "service198", service198_db)
	service198_ctr := applyDockerDefaults(spec, service198, "service198_proc", "service198_container")
	containers = append(containers, service198_ctr)
	allServices = append(allServices, service198)

	//----- service199 (terminal) ----------
	service199_db := mongodb.Container(spec, "service199_db")
	allServices = append(allServices, service199_db)
	service199 := workflow.Service[large_scale_app.Service199](spec, "service199", service199_db)
	service199_ctr := applyDockerDefaults(spec, service199, "service199_proc", "service199_container")
	containers = append(containers, service199_ctr)
	allServices = append(allServices, service199)

	//----- service200 (terminal) ----------
	service200_db := mongodb.Container(spec, "service200_db")
	allServices = append(allServices, service200_db)
	service200 := workflow.Service[large_scale_app.Service200](spec, "service200", service200_db)
	service200_ctr := applyDockerDefaults(spec, service200, "service200_proc", "service200_container")
	containers = append(containers, service200_ctr)
	allServices = append(allServices, service200)

	//----- service201 (terminal) ----------
	service201_db := mongodb.Container(spec, "service201_db")
	allServices = append(allServices, service201_db)
	service201 := workflow.Service[large_scale_app.Service201](spec, "service201", service201_db)
	service201_ctr := applyDockerDefaults(spec, service201, "service201_proc", "service201_container")
	containers = append(containers, service201_ctr)
	allServices = append(allServices, service201)

	//----- service202 (terminal) ----------
	service202_db := mongodb.Container(spec, "service202_db")
	allServices = append(allServices, service202_db)
	service202 := workflow.Service[large_scale_app.Service202](spec, "service202", service202_db)
	service202_ctr := applyDockerDefaults(spec, service202, "service202_proc", "service202_container")
	containers = append(containers, service202_ctr)
	allServices = append(allServices, service202)

	//----- service203 (terminal) ----------
	service203_db := mongodb.Container(spec, "service203_db")
	allServices = append(allServices, service203_db)
	service203 := workflow.Service[large_scale_app.Service203](spec, "service203", service203_db)
	service203_ctr := applyDockerDefaults(spec, service203, "service203_proc", "service203_container")
	containers = append(containers, service203_ctr)
	allServices = append(allServices, service203)

	//----- service204 (terminal) ----------
	service204_db := mongodb.Container(spec, "service204_db")
	allServices = append(allServices, service204_db)
	service204 := workflow.Service[large_scale_app.Service204](spec, "service204", service204_db)
	service204_ctr := applyDockerDefaults(spec, service204, "service204_proc", "service204_container")
	containers = append(containers, service204_ctr)
	allServices = append(allServices, service204)

	//----- service205 (terminal) ----------
	service205_db := mongodb.Container(spec, "service205_db")
	allServices = append(allServices, service205_db)
	service205 := workflow.Service[large_scale_app.Service205](spec, "service205", service205_db)
	service205_ctr := applyDockerDefaults(spec, service205, "service205_proc", "service205_container")
	containers = append(containers, service205_ctr)
	allServices = append(allServices, service205)

	//----- service206 (terminal) ----------
	service206_db := mongodb.Container(spec, "service206_db")
	allServices = append(allServices, service206_db)
	service206 := workflow.Service[large_scale_app.Service206](spec, "service206", service206_db)
	service206_ctr := applyDockerDefaults(spec, service206, "service206_proc", "service206_container")
	containers = append(containers, service206_ctr)
	allServices = append(allServices, service206)

	//----- service207 (terminal) ----------
	service207_db := mongodb.Container(spec, "service207_db")
	allServices = append(allServices, service207_db)
	service207 := workflow.Service[large_scale_app.Service207](spec, "service207", service207_db)
	service207_ctr := applyDockerDefaults(spec, service207, "service207_proc", "service207_container")
	containers = append(containers, service207_ctr)
	allServices = append(allServices, service207)

	//----- service208 (terminal) ----------
	service208_db := mongodb.Container(spec, "service208_db")
	allServices = append(allServices, service208_db)
	service208 := workflow.Service[large_scale_app.Service208](spec, "service208", service208_db)
	service208_ctr := applyDockerDefaults(spec, service208, "service208_proc", "service208_container")
	containers = append(containers, service208_ctr)
	allServices = append(allServices, service208)

	//----- service209 (terminal) ----------
	service209_db := mongodb.Container(spec, "service209_db")
	allServices = append(allServices, service209_db)
	service209 := workflow.Service[large_scale_app.Service209](spec, "service209", service209_db)
	service209_ctr := applyDockerDefaults(spec, service209, "service209_proc", "service209_container")
	containers = append(containers, service209_ctr)
	allServices = append(allServices, service209)

	//----- service210 (terminal) ----------
	service210_db := mongodb.Container(spec, "service210_db")
	allServices = append(allServices, service210_db)
	service210 := workflow.Service[large_scale_app.Service210](spec, "service210", service210_db)
	service210_ctr := applyDockerDefaults(spec, service210, "service210_proc", "service210_container")
	containers = append(containers, service210_ctr)
	allServices = append(allServices, service210)

	//----- service211 (terminal) ----------
	service211_db := mongodb.Container(spec, "service211_db")
	allServices = append(allServices, service211_db)
	service211 := workflow.Service[large_scale_app.Service211](spec, "service211", service211_db)
	service211_ctr := applyDockerDefaults(spec, service211, "service211_proc", "service211_container")
	containers = append(containers, service211_ctr)
	allServices = append(allServices, service211)

	//----- service212 (terminal) ----------
	service212_db := mongodb.Container(spec, "service212_db")
	allServices = append(allServices, service212_db)
	service212 := workflow.Service[large_scale_app.Service212](spec, "service212", service212_db)
	service212_ctr := applyDockerDefaults(spec, service212, "service212_proc", "service212_container")
	containers = append(containers, service212_ctr)
	allServices = append(allServices, service212)

	//----- service213 (terminal) ----------
	service213_db := mongodb.Container(spec, "service213_db")
	allServices = append(allServices, service213_db)
	service213 := workflow.Service[large_scale_app.Service213](spec, "service213", service213_db)
	service213_ctr := applyDockerDefaults(spec, service213, "service213_proc", "service213_container")
	containers = append(containers, service213_ctr)
	allServices = append(allServices, service213)

	//----- service214 (terminal) ----------
	service214_db := mongodb.Container(spec, "service214_db")
	allServices = append(allServices, service214_db)
	service214 := workflow.Service[large_scale_app.Service214](spec, "service214", service214_db)
	service214_ctr := applyDockerDefaults(spec, service214, "service214_proc", "service214_container")
	containers = append(containers, service214_ctr)
	allServices = append(allServices, service214)

	//----- service215 (terminal) ----------
	service215_db := mongodb.Container(spec, "service215_db")
	allServices = append(allServices, service215_db)
	service215 := workflow.Service[large_scale_app.Service215](spec, "service215", service215_db)
	service215_ctr := applyDockerDefaults(spec, service215, "service215_proc", "service215_container")
	containers = append(containers, service215_ctr)
	allServices = append(allServices, service215)

	//----- service216 (terminal) ----------
	service216_db := mongodb.Container(spec, "service216_db")
	allServices = append(allServices, service216_db)
	service216 := workflow.Service[large_scale_app.Service216](spec, "service216", service216_db)
	service216_ctr := applyDockerDefaults(spec, service216, "service216_proc", "service216_container")
	containers = append(containers, service216_ctr)
	allServices = append(allServices, service216)

	//----- service217 (terminal) ----------
	service217_db := mongodb.Container(spec, "service217_db")
	allServices = append(allServices, service217_db)
	service217 := workflow.Service[large_scale_app.Service217](spec, "service217", service217_db)
	service217_ctr := applyDockerDefaults(spec, service217, "service217_proc", "service217_container")
	containers = append(containers, service217_ctr)
	allServices = append(allServices, service217)

	//----- service218 (terminal) ----------
	service218_db := mongodb.Container(spec, "service218_db")
	allServices = append(allServices, service218_db)
	service218 := workflow.Service[large_scale_app.Service218](spec, "service218", service218_db)
	service218_ctr := applyDockerDefaults(spec, service218, "service218_proc", "service218_container")
	containers = append(containers, service218_ctr)
	allServices = append(allServices, service218)

	//----- service219 (terminal) ----------
	service219_db := mongodb.Container(spec, "service219_db")
	allServices = append(allServices, service219_db)
	service219 := workflow.Service[large_scale_app.Service219](spec, "service219", service219_db)
	service219_ctr := applyDockerDefaults(spec, service219, "service219_proc", "service219_container")
	containers = append(containers, service219_ctr)
	allServices = append(allServices, service219)

	//----- service220 (terminal) ----------
	service220_db := mongodb.Container(spec, "service220_db")
	allServices = append(allServices, service220_db)
	service220 := workflow.Service[large_scale_app.Service220](spec, "service220", service220_db)
	service220_ctr := applyDockerDefaults(spec, service220, "service220_proc", "service220_container")
	containers = append(containers, service220_ctr)
	allServices = append(allServices, service220)

	//----- service221 (terminal) ----------
	service221_db := mongodb.Container(spec, "service221_db")
	allServices = append(allServices, service221_db)
	service221 := workflow.Service[large_scale_app.Service221](spec, "service221", service221_db)
	service221_ctr := applyDockerDefaults(spec, service221, "service221_proc", "service221_container")
	containers = append(containers, service221_ctr)
	allServices = append(allServices, service221)

	//----- service222 (terminal) ----------
	service222_db := mongodb.Container(spec, "service222_db")
	allServices = append(allServices, service222_db)
	service222 := workflow.Service[large_scale_app.Service222](spec, "service222", service222_db)
	service222_ctr := applyDockerDefaults(spec, service222, "service222_proc", "service222_container")
	containers = append(containers, service222_ctr)
	allServices = append(allServices, service222)

	//----- service223 (terminal) ----------
	service223_db := mongodb.Container(spec, "service223_db")
	allServices = append(allServices, service223_db)
	service223 := workflow.Service[large_scale_app.Service223](spec, "service223", service223_db)
	service223_ctr := applyDockerDefaults(spec, service223, "service223_proc", "service223_container")
	containers = append(containers, service223_ctr)
	allServices = append(allServices, service223)

	//----- service224 (terminal) ----------
	service224_db := mongodb.Container(spec, "service224_db")
	allServices = append(allServices, service224_db)
	service224 := workflow.Service[large_scale_app.Service224](spec, "service224", service224_db)
	service224_ctr := applyDockerDefaults(spec, service224, "service224_proc", "service224_container")
	containers = append(containers, service224_ctr)
	allServices = append(allServices, service224)

	//----- service225 (terminal) ----------
	service225_db := mongodb.Container(spec, "service225_db")
	allServices = append(allServices, service225_db)
	service225 := workflow.Service[large_scale_app.Service225](spec, "service225", service225_db)
	service225_ctr := applyDockerDefaults(spec, service225, "service225_proc", "service225_container")
	containers = append(containers, service225_ctr)
	allServices = append(allServices, service225)

	//----- service226 (terminal) ----------
	service226_db := mongodb.Container(spec, "service226_db")
	allServices = append(allServices, service226_db)
	service226 := workflow.Service[large_scale_app.Service226](spec, "service226", service226_db)
	service226_ctr := applyDockerDefaults(spec, service226, "service226_proc", "service226_container")
	containers = append(containers, service226_ctr)
	allServices = append(allServices, service226)

	//----- service227 (terminal) ----------
	service227_db := mongodb.Container(spec, "service227_db")
	allServices = append(allServices, service227_db)
	service227 := workflow.Service[large_scale_app.Service227](spec, "service227", service227_db)
	service227_ctr := applyDockerDefaults(spec, service227, "service227_proc", "service227_container")
	containers = append(containers, service227_ctr)
	allServices = append(allServices, service227)

	//----- service228 (terminal) ----------
	service228_db := mongodb.Container(spec, "service228_db")
	allServices = append(allServices, service228_db)
	service228 := workflow.Service[large_scale_app.Service228](spec, "service228", service228_db)
	service228_ctr := applyDockerDefaults(spec, service228, "service228_proc", "service228_container")
	containers = append(containers, service228_ctr)
	allServices = append(allServices, service228)

	//----- service229 (terminal) ----------
	service229_db := mongodb.Container(spec, "service229_db")
	allServices = append(allServices, service229_db)
	service229 := workflow.Service[large_scale_app.Service229](spec, "service229", service229_db)
	service229_ctr := applyDockerDefaults(spec, service229, "service229_proc", "service229_container")
	containers = append(containers, service229_ctr)
	allServices = append(allServices, service229)

	//----- service230 (terminal) ----------
	service230_db := mongodb.Container(spec, "service230_db")
	allServices = append(allServices, service230_db)
	service230 := workflow.Service[large_scale_app.Service230](spec, "service230", service230_db)
	service230_ctr := applyDockerDefaults(spec, service230, "service230_proc", "service230_container")
	containers = append(containers, service230_ctr)
	allServices = append(allServices, service230)

	//----- service231 (terminal) ----------
	service231_db := mongodb.Container(spec, "service231_db")
	allServices = append(allServices, service231_db)
	service231 := workflow.Service[large_scale_app.Service231](spec, "service231", service231_db)
	service231_ctr := applyDockerDefaults(spec, service231, "service231_proc", "service231_container")
	containers = append(containers, service231_ctr)
	allServices = append(allServices, service231)

	//----- service232 (terminal) ----------
	service232_db := mongodb.Container(spec, "service232_db")
	allServices = append(allServices, service232_db)
	service232 := workflow.Service[large_scale_app.Service232](spec, "service232", service232_db)
	service232_ctr := applyDockerDefaults(spec, service232, "service232_proc", "service232_container")
	containers = append(containers, service232_ctr)
	allServices = append(allServices, service232)

	//----- service233 (terminal) ----------
	service233_db := mongodb.Container(spec, "service233_db")
	allServices = append(allServices, service233_db)
	service233 := workflow.Service[large_scale_app.Service233](spec, "service233", service233_db)
	service233_ctr := applyDockerDefaults(spec, service233, "service233_proc", "service233_container")
	containers = append(containers, service233_ctr)
	allServices = append(allServices, service233)

	//----- service234 (terminal) ----------
	service234_db := mongodb.Container(spec, "service234_db")
	allServices = append(allServices, service234_db)
	service234 := workflow.Service[large_scale_app.Service234](spec, "service234", service234_db)
	service234_ctr := applyDockerDefaults(spec, service234, "service234_proc", "service234_container")
	containers = append(containers, service234_ctr)
	allServices = append(allServices, service234)

	//----- service235 (terminal) ----------
	service235_db := mongodb.Container(spec, "service235_db")
	allServices = append(allServices, service235_db)
	service235 := workflow.Service[large_scale_app.Service235](spec, "service235", service235_db)
	service235_ctr := applyDockerDefaults(spec, service235, "service235_proc", "service235_container")
	containers = append(containers, service235_ctr)
	allServices = append(allServices, service235)

	//----- service236 (terminal) ----------
	service236_db := mongodb.Container(spec, "service236_db")
	allServices = append(allServices, service236_db)
	service236 := workflow.Service[large_scale_app.Service236](spec, "service236", service236_db)
	service236_ctr := applyDockerDefaults(spec, service236, "service236_proc", "service236_container")
	containers = append(containers, service236_ctr)
	allServices = append(allServices, service236)

	//----- service237 (terminal) ----------
	service237_db := mongodb.Container(spec, "service237_db")
	allServices = append(allServices, service237_db)
	service237 := workflow.Service[large_scale_app.Service237](spec, "service237", service237_db)
	service237_ctr := applyDockerDefaults(spec, service237, "service237_proc", "service237_container")
	containers = append(containers, service237_ctr)
	allServices = append(allServices, service237)

	//----- service238 (terminal) ----------
	service238_db := mongodb.Container(spec, "service238_db")
	allServices = append(allServices, service238_db)
	service238 := workflow.Service[large_scale_app.Service238](spec, "service238", service238_db)
	service238_ctr := applyDockerDefaults(spec, service238, "service238_proc", "service238_container")
	containers = append(containers, service238_ctr)
	allServices = append(allServices, service238)

	//----- service239 (terminal) ----------
	service239_db := mongodb.Container(spec, "service239_db")
	allServices = append(allServices, service239_db)
	service239 := workflow.Service[large_scale_app.Service239](spec, "service239", service239_db)
	service239_ctr := applyDockerDefaults(spec, service239, "service239_proc", "service239_container")
	containers = append(containers, service239_ctr)
	allServices = append(allServices, service239)

	//----- service240 (terminal) ----------
	service240_db := mongodb.Container(spec, "service240_db")
	allServices = append(allServices, service240_db)
	service240 := workflow.Service[large_scale_app.Service240](spec, "service240", service240_db)
	service240_ctr := applyDockerDefaults(spec, service240, "service240_proc", "service240_container")
	containers = append(containers, service240_ctr)
	allServices = append(allServices, service240)

	//----- service241 (terminal) ----------
	service241_db := mongodb.Container(spec, "service241_db")
	allServices = append(allServices, service241_db)
	service241 := workflow.Service[large_scale_app.Service241](spec, "service241", service241_db)
	service241_ctr := applyDockerDefaults(spec, service241, "service241_proc", "service241_container")
	containers = append(containers, service241_ctr)
	allServices = append(allServices, service241)

	//----- service242 (terminal) ----------
	service242_db := mongodb.Container(spec, "service242_db")
	allServices = append(allServices, service242_db)
	service242 := workflow.Service[large_scale_app.Service242](spec, "service242", service242_db)
	service242_ctr := applyDockerDefaults(spec, service242, "service242_proc", "service242_container")
	containers = append(containers, service242_ctr)
	allServices = append(allServices, service242)

	//----- service243 (terminal) ----------
	service243_db := mongodb.Container(spec, "service243_db")
	allServices = append(allServices, service243_db)
	service243 := workflow.Service[large_scale_app.Service243](spec, "service243", service243_db)
	service243_ctr := applyDockerDefaults(spec, service243, "service243_proc", "service243_container")
	containers = append(containers, service243_ctr)
	allServices = append(allServices, service243)

	//----- service244 (terminal) ----------
	service244_db := mongodb.Container(spec, "service244_db")
	allServices = append(allServices, service244_db)
	service244 := workflow.Service[large_scale_app.Service244](spec, "service244", service244_db)
	service244_ctr := applyDockerDefaults(spec, service244, "service244_proc", "service244_container")
	containers = append(containers, service244_ctr)
	allServices = append(allServices, service244)

	//----- service245 (terminal) ----------
	service245_db := mongodb.Container(spec, "service245_db")
	allServices = append(allServices, service245_db)
	service245 := workflow.Service[large_scale_app.Service245](spec, "service245", service245_db)
	service245_ctr := applyDockerDefaults(spec, service245, "service245_proc", "service245_container")
	containers = append(containers, service245_ctr)
	allServices = append(allServices, service245)

	//----- service246 (terminal) ----------
	service246_db := mongodb.Container(spec, "service246_db")
	allServices = append(allServices, service246_db)
	service246 := workflow.Service[large_scale_app.Service246](spec, "service246", service246_db)
	service246_ctr := applyDockerDefaults(spec, service246, "service246_proc", "service246_container")
	containers = append(containers, service246_ctr)
	allServices = append(allServices, service246)

	//----- service247 (terminal) ----------
	service247_db := mongodb.Container(spec, "service247_db")
	allServices = append(allServices, service247_db)
	service247 := workflow.Service[large_scale_app.Service247](spec, "service247", service247_db)
	service247_ctr := applyDockerDefaults(spec, service247, "service247_proc", "service247_container")
	containers = append(containers, service247_ctr)
	allServices = append(allServices, service247)

	//----- service248 (terminal) ----------
	service248_db := mongodb.Container(spec, "service248_db")
	allServices = append(allServices, service248_db)
	service248 := workflow.Service[large_scale_app.Service248](spec, "service248", service248_db)
	service248_ctr := applyDockerDefaults(spec, service248, "service248_proc", "service248_container")
	containers = append(containers, service248_ctr)
	allServices = append(allServices, service248)

	//----- service249 (terminal) ----------
	service249_db := mongodb.Container(spec, "service249_db")
	allServices = append(allServices, service249_db)
	service249 := workflow.Service[large_scale_app.Service249](spec, "service249", service249_db)
	service249_ctr := applyDockerDefaults(spec, service249, "service249_proc", "service249_container")
	containers = append(containers, service249_ctr)
	allServices = append(allServices, service249)

	//----- service250 (terminal) ----------
	service250_db := mongodb.Container(spec, "service250_db")
	allServices = append(allServices, service250_db)
	service250 := workflow.Service[large_scale_app.Service250](spec, "service250", service250_db)
	service250_ctr := applyDockerDefaults(spec, service250, "service250_proc", "service250_container")
	containers = append(containers, service250_ctr)
	allServices = append(allServices, service250)

	//----- service251 (terminal) ----------
	service251_db := mongodb.Container(spec, "service251_db")
	allServices = append(allServices, service251_db)
	service251 := workflow.Service[large_scale_app.Service251](spec, "service251", service251_db)
	service251_ctr := applyDockerDefaults(spec, service251, "service251_proc", "service251_container")
	containers = append(containers, service251_ctr)
	allServices = append(allServices, service251)

	//----- service252 (terminal) ----------
	service252_db := mongodb.Container(spec, "service252_db")
	allServices = append(allServices, service252_db)
	service252 := workflow.Service[large_scale_app.Service252](spec, "service252", service252_db)
	service252_ctr := applyDockerDefaults(spec, service252, "service252_proc", "service252_container")
	containers = append(containers, service252_ctr)
	allServices = append(allServices, service252)

	//----- service253 (terminal) ----------
	service253_db := mongodb.Container(spec, "service253_db")
	allServices = append(allServices, service253_db)
	service253 := workflow.Service[large_scale_app.Service253](spec, "service253", service253_db)
	service253_ctr := applyDockerDefaults(spec, service253, "service253_proc", "service253_container")
	containers = append(containers, service253_ctr)
	allServices = append(allServices, service253)

	//----- service254 (terminal) ----------
	service254_db := mongodb.Container(spec, "service254_db")
	allServices = append(allServices, service254_db)
	service254 := workflow.Service[large_scale_app.Service254](spec, "service254", service254_db)
	service254_ctr := applyDockerDefaults(spec, service254, "service254_proc", "service254_container")
	containers = append(containers, service254_ctr)
	allServices = append(allServices, service254)

	//----- service255 (terminal) ----------
	service255_db := mongodb.Container(spec, "service255_db")
	allServices = append(allServices, service255_db)
	service255 := workflow.Service[large_scale_app.Service255](spec, "service255", service255_db)
	service255_ctr := applyDockerDefaults(spec, service255, "service255_proc", "service255_container")
	containers = append(containers, service255_ctr)
	allServices = append(allServices, service255)

	//----- service256 (terminal) ----------
	service256_db := mongodb.Container(spec, "service256_db")
	allServices = append(allServices, service256_db)
	service256 := workflow.Service[large_scale_app.Service256](spec, "service256", service256_db)
	service256_ctr := applyDockerDefaults(spec, service256, "service256_proc", "service256_container")
	containers = append(containers, service256_ctr)
	allServices = append(allServices, service256)

	//----- service24 depends on [94 95 96 97] ----------
	service24_db := mongodb.Container(spec, "service24_db")
	allServices = append(allServices, service24_db)
	service24 := workflow.Service[large_scale_app.Service24](
		spec,
		"service24",
		service24_db,
		service94,
		service95,
		service96,
		service97,
	)
	service24_ctr := applyDockerDefaults(spec, service24, "service24_proc", "service24_container")
	containers = append(containers, service24_ctr)
	allServices = append(allServices, service24)

	//----- service54 depends on [214 215 216 217] ----------
	service54_db := mongodb.Container(spec, "service54_db")
	allServices = append(allServices, service54_db)
	service54 := workflow.Service[large_scale_app.Service54](
		spec,
		"service54",
		service54_db,
		service214,
		service215,
		service216,
		service217,
	)
	service54_ctr := applyDockerDefaults(spec, service54, "service54_proc", "service54_container")
	containers = append(containers, service54_ctr)
	allServices = append(allServices, service54)

	//----- service85 depends on [338 339 340 341] ----------
	service85_db := mongodb.Container(spec, "service85_db")
	allServices = append(allServices, service85_db)
	service85 := workflow.Service[large_scale_app.Service85](
		spec,
		"service85",
		service85_db,
		service338,
		service339,
		service340,
		service341,
	)
	service85_ctr := applyDockerDefaults(spec, service85, "service85_proc", "service85_container")
	containers = append(containers, service85_ctr)
	allServices = append(allServices, service85)

	//----- service84 depends on [334 335 336 337] ----------
	service84_db := mongodb.Container(spec, "service84_db")
	allServices = append(allServices, service84_db)
	service84 := workflow.Service[large_scale_app.Service84](
		spec,
		"service84",
		service84_db,
		service334,
		service335,
		service336,
		service337,
	)
	service84_ctr := applyDockerDefaults(spec, service84, "service84_proc", "service84_container")
	containers = append(containers, service84_ctr)
	allServices = append(allServices, service84)

	//----- service83 depends on [330 331 332 333] ----------
	service83_db := mongodb.Container(spec, "service83_db")
	allServices = append(allServices, service83_db)
	service83 := workflow.Service[large_scale_app.Service83](
		spec,
		"service83",
		service83_db,
		service330,
		service331,
		service332,
		service333,
	)
	service83_ctr := applyDockerDefaults(spec, service83, "service83_proc", "service83_container")
	containers = append(containers, service83_ctr)
	allServices = append(allServices, service83)

	//----- service82 depends on [326 327 328 329] ----------
	service82_db := mongodb.Container(spec, "service82_db")
	allServices = append(allServices, service82_db)
	service82 := workflow.Service[large_scale_app.Service82](
		spec,
		"service82",
		service82_db,
		service326,
		service327,
		service328,
		service329,
	)
	service82_ctr := applyDockerDefaults(spec, service82, "service82_proc", "service82_container")
	containers = append(containers, service82_ctr)
	allServices = append(allServices, service82)

	//----- service81 depends on [322 323 324 325] ----------
	service81_db := mongodb.Container(spec, "service81_db")
	allServices = append(allServices, service81_db)
	service81 := workflow.Service[large_scale_app.Service81](
		spec,
		"service81",
		service81_db,
		service322,
		service323,
		service324,
		service325,
	)
	service81_ctr := applyDockerDefaults(spec, service81, "service81_proc", "service81_container")
	containers = append(containers, service81_ctr)
	allServices = append(allServices, service81)

	//----- service80 depends on [318 319 320 321] ----------
	service80_db := mongodb.Container(spec, "service80_db")
	allServices = append(allServices, service80_db)
	service80 := workflow.Service[large_scale_app.Service80](
		spec,
		"service80",
		service80_db,
		service318,
		service319,
		service320,
		service321,
	)
	service80_ctr := applyDockerDefaults(spec, service80, "service80_proc", "service80_container")
	containers = append(containers, service80_ctr)
	allServices = append(allServices, service80)

	//----- service44 depends on [174 175 176 177] ----------
	service44_db := mongodb.Container(spec, "service44_db")
	allServices = append(allServices, service44_db)
	service44 := workflow.Service[large_scale_app.Service44](
		spec,
		"service44",
		service44_db,
		service174,
		service175,
		service176,
		service177,
	)
	service44_ctr := applyDockerDefaults(spec, service44, "service44_proc", "service44_container")
	containers = append(containers, service44_ctr)
	allServices = append(allServices, service44)

	//----- service78 depends on [310 311 312 313] ----------
	service78_db := mongodb.Container(spec, "service78_db")
	allServices = append(allServices, service78_db)
	service78 := workflow.Service[large_scale_app.Service78](
		spec,
		"service78",
		service78_db,
		service310,
		service311,
		service312,
		service313,
	)
	service78_ctr := applyDockerDefaults(spec, service78, "service78_proc", "service78_container")
	containers = append(containers, service78_ctr)
	allServices = append(allServices, service78)

	//----- service77 depends on [306 307 308 309] ----------
	service77_db := mongodb.Container(spec, "service77_db")
	allServices = append(allServices, service77_db)
	service77 := workflow.Service[large_scale_app.Service77](
		spec,
		"service77",
		service77_db,
		service306,
		service307,
		service308,
		service309,
	)
	service77_ctr := applyDockerDefaults(spec, service77, "service77_proc", "service77_container")
	containers = append(containers, service77_ctr)
	allServices = append(allServices, service77)

	//----- service76 depends on [302 303 304 305] ----------
	service76_db := mongodb.Container(spec, "service76_db")
	allServices = append(allServices, service76_db)
	service76 := workflow.Service[large_scale_app.Service76](
		spec,
		"service76",
		service76_db,
		service302,
		service303,
		service304,
		service305,
	)
	service76_ctr := applyDockerDefaults(spec, service76, "service76_proc", "service76_container")
	containers = append(containers, service76_ctr)
	allServices = append(allServices, service76)

	//----- service75 depends on [298 299 300 301] ----------
	service75_db := mongodb.Container(spec, "service75_db")
	allServices = append(allServices, service75_db)
	service75 := workflow.Service[large_scale_app.Service75](
		spec,
		"service75",
		service75_db,
		service298,
		service299,
		service300,
		service301,
	)
	service75_ctr := applyDockerDefaults(spec, service75, "service75_proc", "service75_container")
	containers = append(containers, service75_ctr)
	allServices = append(allServices, service75)

	//----- service74 depends on [294 295 296 297] ----------
	service74_db := mongodb.Container(spec, "service74_db")
	allServices = append(allServices, service74_db)
	service74 := workflow.Service[large_scale_app.Service74](
		spec,
		"service74",
		service74_db,
		service294,
		service295,
		service296,
		service297,
	)
	service74_ctr := applyDockerDefaults(spec, service74, "service74_proc", "service74_container")
	containers = append(containers, service74_ctr)
	allServices = append(allServices, service74)

	//----- service73 depends on [290 291 292 293] ----------
	service73_db := mongodb.Container(spec, "service73_db")
	allServices = append(allServices, service73_db)
	service73 := workflow.Service[large_scale_app.Service73](
		spec,
		"service73",
		service73_db,
		service290,
		service291,
		service292,
		service293,
	)
	service73_ctr := applyDockerDefaults(spec, service73, "service73_proc", "service73_container")
	containers = append(containers, service73_ctr)
	allServices = append(allServices, service73)

	//----- service72 depends on [286 287 288 289] ----------
	service72_db := mongodb.Container(spec, "service72_db")
	allServices = append(allServices, service72_db)
	service72 := workflow.Service[large_scale_app.Service72](
		spec,
		"service72",
		service72_db,
		service286,
		service287,
		service288,
		service289,
	)
	service72_ctr := applyDockerDefaults(spec, service72, "service72_proc", "service72_container")
	containers = append(containers, service72_ctr)
	allServices = append(allServices, service72)

	//----- service71 depends on [282 283 284 285] ----------
	service71_db := mongodb.Container(spec, "service71_db")
	allServices = append(allServices, service71_db)
	service71 := workflow.Service[large_scale_app.Service71](
		spec,
		"service71",
		service71_db,
		service282,
		service283,
		service284,
		service285,
	)
	service71_ctr := applyDockerDefaults(spec, service71, "service71_proc", "service71_container")
	containers = append(containers, service71_ctr)
	allServices = append(allServices, service71)

	//----- service70 depends on [278 279 280 281] ----------
	service70_db := mongodb.Container(spec, "service70_db")
	allServices = append(allServices, service70_db)
	service70 := workflow.Service[large_scale_app.Service70](
		spec,
		"service70",
		service70_db,
		service278,
		service279,
		service280,
		service281,
	)
	service70_ctr := applyDockerDefaults(spec, service70, "service70_proc", "service70_container")
	containers = append(containers, service70_ctr)
	allServices = append(allServices, service70)

	//----- service69 depends on [274 275 276 277] ----------
	service69_db := mongodb.Container(spec, "service69_db")
	allServices = append(allServices, service69_db)
	service69 := workflow.Service[large_scale_app.Service69](
		spec,
		"service69",
		service69_db,
		service274,
		service275,
		service276,
		service277,
	)
	service69_ctr := applyDockerDefaults(spec, service69, "service69_proc", "service69_container")
	containers = append(containers, service69_ctr)
	allServices = append(allServices, service69)

	//----- service68 depends on [270 271 272 273] ----------
	service68_db := mongodb.Container(spec, "service68_db")
	allServices = append(allServices, service68_db)
	service68 := workflow.Service[large_scale_app.Service68](
		spec,
		"service68",
		service68_db,
		service270,
		service271,
		service272,
		service273,
	)
	service68_ctr := applyDockerDefaults(spec, service68, "service68_proc", "service68_container")
	containers = append(containers, service68_ctr)
	allServices = append(allServices, service68)

	//----- service67 depends on [266 267 268 269] ----------
	service67_db := mongodb.Container(spec, "service67_db")
	allServices = append(allServices, service67_db)
	service67 := workflow.Service[large_scale_app.Service67](
		spec,
		"service67",
		service67_db,
		service266,
		service267,
		service268,
		service269,
	)
	service67_ctr := applyDockerDefaults(spec, service67, "service67_proc", "service67_container")
	containers = append(containers, service67_ctr)
	allServices = append(allServices, service67)

	//----- service46 depends on [182 183 184 185] ----------
	service46_db := mongodb.Container(spec, "service46_db")
	allServices = append(allServices, service46_db)
	service46 := workflow.Service[large_scale_app.Service46](
		spec,
		"service46",
		service46_db,
		service182,
		service183,
		service184,
		service185,
	)
	service46_ctr := applyDockerDefaults(spec, service46, "service46_proc", "service46_container")
	containers = append(containers, service46_ctr)
	allServices = append(allServices, service46)

	//----- service65 depends on [258 259 260 261] ----------
	service65_db := mongodb.Container(spec, "service65_db")
	allServices = append(allServices, service65_db)
	service65 := workflow.Service[large_scale_app.Service65](
		spec,
		"service65",
		service65_db,
		service258,
		service259,
		service260,
		service261,
	)
	service65_ctr := applyDockerDefaults(spec, service65, "service65_proc", "service65_container")
	containers = append(containers, service65_ctr)
	allServices = append(allServices, service65)

	//----- service64 depends on [254 255 256 257] ----------
	service64_db := mongodb.Container(spec, "service64_db")
	allServices = append(allServices, service64_db)
	service64 := workflow.Service[large_scale_app.Service64](
		spec,
		"service64",
		service64_db,
		service254,
		service255,
		service256,
		service257,
	)
	service64_ctr := applyDockerDefaults(spec, service64, "service64_proc", "service64_container")
	containers = append(containers, service64_ctr)
	allServices = append(allServices, service64)

	//----- service63 depends on [250 251 252 253] ----------
	service63_db := mongodb.Container(spec, "service63_db")
	allServices = append(allServices, service63_db)
	service63 := workflow.Service[large_scale_app.Service63](
		spec,
		"service63",
		service63_db,
		service250,
		service251,
		service252,
		service253,
	)
	service63_ctr := applyDockerDefaults(spec, service63, "service63_proc", "service63_container")
	containers = append(containers, service63_ctr)
	allServices = append(allServices, service63)

	//----- service62 depends on [246 247 248 249] ----------
	service62_db := mongodb.Container(spec, "service62_db")
	allServices = append(allServices, service62_db)
	service62 := workflow.Service[large_scale_app.Service62](
		spec,
		"service62",
		service62_db,
		service246,
		service247,
		service248,
		service249,
	)
	service62_ctr := applyDockerDefaults(spec, service62, "service62_proc", "service62_container")
	containers = append(containers, service62_ctr)
	allServices = append(allServices, service62)

	//----- service61 depends on [242 243 244 245] ----------
	service61_db := mongodb.Container(spec, "service61_db")
	allServices = append(allServices, service61_db)
	service61 := workflow.Service[large_scale_app.Service61](
		spec,
		"service61",
		service61_db,
		service242,
		service243,
		service244,
		service245,
	)
	service61_ctr := applyDockerDefaults(spec, service61, "service61_proc", "service61_container")
	containers = append(containers, service61_ctr)
	allServices = append(allServices, service61)

	//----- service60 depends on [238 239 240 241] ----------
	service60_db := mongodb.Container(spec, "service60_db")
	allServices = append(allServices, service60_db)
	service60 := workflow.Service[large_scale_app.Service60](
		spec,
		"service60",
		service60_db,
		service238,
		service239,
		service240,
		service241,
	)
	service60_ctr := applyDockerDefaults(spec, service60, "service60_proc", "service60_container")
	containers = append(containers, service60_ctr)
	allServices = append(allServices, service60)

	//----- service59 depends on [234 235 236 237] ----------
	service59_db := mongodb.Container(spec, "service59_db")
	allServices = append(allServices, service59_db)
	service59 := workflow.Service[large_scale_app.Service59](
		spec,
		"service59",
		service59_db,
		service234,
		service235,
		service236,
		service237,
	)
	service59_ctr := applyDockerDefaults(spec, service59, "service59_proc", "service59_container")
	containers = append(containers, service59_ctr)
	allServices = append(allServices, service59)

	//----- service58 depends on [230 231 232 233] ----------
	service58_db := mongodb.Container(spec, "service58_db")
	allServices = append(allServices, service58_db)
	service58 := workflow.Service[large_scale_app.Service58](
		spec,
		"service58",
		service58_db,
		service230,
		service231,
		service232,
		service233,
	)
	service58_ctr := applyDockerDefaults(spec, service58, "service58_proc", "service58_container")
	containers = append(containers, service58_ctr)
	allServices = append(allServices, service58)

	//----- service57 depends on [226 227 228 229] ----------
	service57_db := mongodb.Container(spec, "service57_db")
	allServices = append(allServices, service57_db)
	service57 := workflow.Service[large_scale_app.Service57](
		spec,
		"service57",
		service57_db,
		service226,
		service227,
		service228,
		service229,
	)
	service57_ctr := applyDockerDefaults(spec, service57, "service57_proc", "service57_container")
	containers = append(containers, service57_ctr)
	allServices = append(allServices, service57)

	//----- service56 depends on [222 223 224 225] ----------
	service56_db := mongodb.Container(spec, "service56_db")
	allServices = append(allServices, service56_db)
	service56 := workflow.Service[large_scale_app.Service56](
		spec,
		"service56",
		service56_db,
		service222,
		service223,
		service224,
		service225,
	)
	service56_ctr := applyDockerDefaults(spec, service56, "service56_proc", "service56_container")
	containers = append(containers, service56_ctr)
	allServices = append(allServices, service56)

	//----- service55 depends on [218 219 220 221] ----------
	service55_db := mongodb.Container(spec, "service55_db")
	allServices = append(allServices, service55_db)
	service55 := workflow.Service[large_scale_app.Service55](
		spec,
		"service55",
		service55_db,
		service218,
		service219,
		service220,
		service221,
	)
	service55_ctr := applyDockerDefaults(spec, service55, "service55_proc", "service55_container")
	containers = append(containers, service55_ctr)
	allServices = append(allServices, service55)

	//----- service43 depends on [170 171 172 173] ----------
	service43_db := mongodb.Container(spec, "service43_db")
	allServices = append(allServices, service43_db)
	service43 := workflow.Service[large_scale_app.Service43](
		spec,
		"service43",
		service43_db,
		service170,
		service171,
		service172,
		service173,
	)
	service43_ctr := applyDockerDefaults(spec, service43, "service43_proc", "service43_container")
	containers = append(containers, service43_ctr)
	allServices = append(allServices, service43)

	//----- service53 depends on [210 211 212 213] ----------
	service53_db := mongodb.Container(spec, "service53_db")
	allServices = append(allServices, service53_db)
	service53 := workflow.Service[large_scale_app.Service53](
		spec,
		"service53",
		service53_db,
		service210,
		service211,
		service212,
		service213,
	)
	service53_ctr := applyDockerDefaults(spec, service53, "service53_proc", "service53_container")
	containers = append(containers, service53_ctr)
	allServices = append(allServices, service53)

	//----- service52 depends on [206 207 208 209] ----------
	service52_db := mongodb.Container(spec, "service52_db")
	allServices = append(allServices, service52_db)
	service52 := workflow.Service[large_scale_app.Service52](
		spec,
		"service52",
		service52_db,
		service206,
		service207,
		service208,
		service209,
	)
	service52_ctr := applyDockerDefaults(spec, service52, "service52_proc", "service52_container")
	containers = append(containers, service52_ctr)
	allServices = append(allServices, service52)

	//----- service51 depends on [202 203 204 205] ----------
	service51_db := mongodb.Container(spec, "service51_db")
	allServices = append(allServices, service51_db)
	service51 := workflow.Service[large_scale_app.Service51](
		spec,
		"service51",
		service51_db,
		service202,
		service203,
		service204,
		service205,
	)
	service51_ctr := applyDockerDefaults(spec, service51, "service51_proc", "service51_container")
	containers = append(containers, service51_ctr)
	allServices = append(allServices, service51)

	//----- service50 depends on [198 199 200 201] ----------
	service50_db := mongodb.Container(spec, "service50_db")
	allServices = append(allServices, service50_db)
	service50 := workflow.Service[large_scale_app.Service50](
		spec,
		"service50",
		service50_db,
		service198,
		service199,
		service200,
		service201,
	)
	service50_ctr := applyDockerDefaults(spec, service50, "service50_proc", "service50_container")
	containers = append(containers, service50_ctr)
	allServices = append(allServices, service50)

	//----- service49 depends on [194 195 196 197] ----------
	service49_db := mongodb.Container(spec, "service49_db")
	allServices = append(allServices, service49_db)
	service49 := workflow.Service[large_scale_app.Service49](
		spec,
		"service49",
		service49_db,
		service194,
		service195,
		service196,
		service197,
	)
	service49_ctr := applyDockerDefaults(spec, service49, "service49_proc", "service49_container")
	containers = append(containers, service49_ctr)
	allServices = append(allServices, service49)

	//----- service48 depends on [190 191 192 193] ----------
	service48_db := mongodb.Container(spec, "service48_db")
	allServices = append(allServices, service48_db)
	service48 := workflow.Service[large_scale_app.Service48](
		spec,
		"service48",
		service48_db,
		service190,
		service191,
		service192,
		service193,
	)
	service48_ctr := applyDockerDefaults(spec, service48, "service48_proc", "service48_container")
	containers = append(containers, service48_ctr)
	allServices = append(allServices, service48)

	//----- service47 depends on [186 187 188 189] ----------
	service47_db := mongodb.Container(spec, "service47_db")
	allServices = append(allServices, service47_db)
	service47 := workflow.Service[large_scale_app.Service47](
		spec,
		"service47",
		service47_db,
		service186,
		service187,
		service188,
		service189,
	)
	service47_ctr := applyDockerDefaults(spec, service47, "service47_proc", "service47_container")
	containers = append(containers, service47_ctr)
	allServices = append(allServices, service47)

	//----- service66 depends on [262 263 264 265] ----------
	service66_db := mongodb.Container(spec, "service66_db")
	allServices = append(allServices, service66_db)
	service66 := workflow.Service[large_scale_app.Service66](
		spec,
		"service66",
		service66_db,
		service262,
		service263,
		service264,
		service265,
	)
	service66_ctr := applyDockerDefaults(spec, service66, "service66_proc", "service66_container")
	containers = append(containers, service66_ctr)
	allServices = append(allServices, service66)

	//----- service22 depends on [86 87 88 89] ----------
	service22_db := mongodb.Container(spec, "service22_db")
	allServices = append(allServices, service22_db)
	service22 := workflow.Service[large_scale_app.Service22](
		spec,
		"service22",
		service22_db,
		service86,
		service87,
		service88,
		service89,
	)
	service22_ctr := applyDockerDefaults(spec, service22, "service22_proc", "service22_container")
	containers = append(containers, service22_ctr)
	allServices = append(allServices, service22)

	//----- service79 depends on [314 315 316 317] ----------
	service79_db := mongodb.Container(spec, "service79_db")
	allServices = append(allServices, service79_db)
	service79 := workflow.Service[large_scale_app.Service79](
		spec,
		"service79",
		service79_db,
		service314,
		service315,
		service316,
		service317,
	)
	service79_ctr := applyDockerDefaults(spec, service79, "service79_proc", "service79_container")
	containers = append(containers, service79_ctr)
	allServices = append(allServices, service79)

	//----- service23 depends on [90 91 92 93] ----------
	service23_db := mongodb.Container(spec, "service23_db")
	allServices = append(allServices, service23_db)
	service23 := workflow.Service[large_scale_app.Service23](
		spec,
		"service23",
		service23_db,
		service90,
		service91,
		service92,
		service93,
	)
	service23_ctr := applyDockerDefaults(spec, service23, "service23_proc", "service23_container")
	containers = append(containers, service23_ctr)
	allServices = append(allServices, service23)

	//----- service42 depends on [166 167 168 169] ----------
	service42_db := mongodb.Container(spec, "service42_db")
	allServices = append(allServices, service42_db)
	service42 := workflow.Service[large_scale_app.Service42](
		spec,
		"service42",
		service42_db,
		service166,
		service167,
		service168,
		service169,
	)
	service42_ctr := applyDockerDefaults(spec, service42, "service42_proc", "service42_container")
	containers = append(containers, service42_ctr)
	allServices = append(allServices, service42)

	//----- service41 depends on [162 163 164 165] ----------
	service41_db := mongodb.Container(spec, "service41_db")
	allServices = append(allServices, service41_db)
	service41 := workflow.Service[large_scale_app.Service41](
		spec,
		"service41",
		service41_db,
		service162,
		service163,
		service164,
		service165,
	)
	service41_ctr := applyDockerDefaults(spec, service41, "service41_proc", "service41_container")
	containers = append(containers, service41_ctr)
	allServices = append(allServices, service41)

	//----- service40 depends on [158 159 160 161] ----------
	service40_db := mongodb.Container(spec, "service40_db")
	allServices = append(allServices, service40_db)
	service40 := workflow.Service[large_scale_app.Service40](
		spec,
		"service40",
		service40_db,
		service158,
		service159,
		service160,
		service161,
	)
	service40_ctr := applyDockerDefaults(spec, service40, "service40_proc", "service40_container")
	containers = append(containers, service40_ctr)
	allServices = append(allServices, service40)

	//----- service39 depends on [154 155 156 157] ----------
	service39_db := mongodb.Container(spec, "service39_db")
	allServices = append(allServices, service39_db)
	service39 := workflow.Service[large_scale_app.Service39](
		spec,
		"service39",
		service39_db,
		service154,
		service155,
		service156,
		service157,
	)
	service39_ctr := applyDockerDefaults(spec, service39, "service39_proc", "service39_container")
	containers = append(containers, service39_ctr)
	allServices = append(allServices, service39)

	//----- service38 depends on [150 151 152 153] ----------
	service38_db := mongodb.Container(spec, "service38_db")
	allServices = append(allServices, service38_db)
	service38 := workflow.Service[large_scale_app.Service38](
		spec,
		"service38",
		service38_db,
		service150,
		service151,
		service152,
		service153,
	)
	service38_ctr := applyDockerDefaults(spec, service38, "service38_proc", "service38_container")
	containers = append(containers, service38_ctr)
	allServices = append(allServices, service38)

	//----- service37 depends on [146 147 148 149] ----------
	service37_db := mongodb.Container(spec, "service37_db")
	allServices = append(allServices, service37_db)
	service37 := workflow.Service[large_scale_app.Service37](
		spec,
		"service37",
		service37_db,
		service146,
		service147,
		service148,
		service149,
	)
	service37_ctr := applyDockerDefaults(spec, service37, "service37_proc", "service37_container")
	containers = append(containers, service37_ctr)
	allServices = append(allServices, service37)

	//----- service36 depends on [142 143 144 145] ----------
	service36_db := mongodb.Container(spec, "service36_db")
	allServices = append(allServices, service36_db)
	service36 := workflow.Service[large_scale_app.Service36](
		spec,
		"service36",
		service36_db,
		service142,
		service143,
		service144,
		service145,
	)
	service36_ctr := applyDockerDefaults(spec, service36, "service36_proc", "service36_container")
	containers = append(containers, service36_ctr)
	allServices = append(allServices, service36)

	//----- service35 depends on [138 139 140 141] ----------
	service35_db := mongodb.Container(spec, "service35_db")
	allServices = append(allServices, service35_db)
	service35 := workflow.Service[large_scale_app.Service35](
		spec,
		"service35",
		service35_db,
		service138,
		service139,
		service140,
		service141,
	)
	service35_ctr := applyDockerDefaults(spec, service35, "service35_proc", "service35_container")
	containers = append(containers, service35_ctr)
	allServices = append(allServices, service35)

	//----- service34 depends on [134 135 136 137] ----------
	service34_db := mongodb.Container(spec, "service34_db")
	allServices = append(allServices, service34_db)
	service34 := workflow.Service[large_scale_app.Service34](
		spec,
		"service34",
		service34_db,
		service134,
		service135,
		service136,
		service137,
	)
	service34_ctr := applyDockerDefaults(spec, service34, "service34_proc", "service34_container")
	containers = append(containers, service34_ctr)
	allServices = append(allServices, service34)

	//----- service33 depends on [130 131 132 133] ----------
	service33_db := mongodb.Container(spec, "service33_db")
	allServices = append(allServices, service33_db)
	service33 := workflow.Service[large_scale_app.Service33](
		spec,
		"service33",
		service33_db,
		service130,
		service131,
		service132,
		service133,
	)
	service33_ctr := applyDockerDefaults(spec, service33, "service33_proc", "service33_container")
	containers = append(containers, service33_ctr)
	allServices = append(allServices, service33)

	//----- service32 depends on [126 127 128 129] ----------
	service32_db := mongodb.Container(spec, "service32_db")
	allServices = append(allServices, service32_db)
	service32 := workflow.Service[large_scale_app.Service32](
		spec,
		"service32",
		service32_db,
		service126,
		service127,
		service128,
		service129,
	)
	service32_ctr := applyDockerDefaults(spec, service32, "service32_proc", "service32_container")
	containers = append(containers, service32_ctr)
	allServices = append(allServices, service32)

	//----- service31 depends on [122 123 124 125] ----------
	service31_db := mongodb.Container(spec, "service31_db")
	allServices = append(allServices, service31_db)
	service31 := workflow.Service[large_scale_app.Service31](
		spec,
		"service31",
		service31_db,
		service122,
		service123,
		service124,
		service125,
	)
	service31_ctr := applyDockerDefaults(spec, service31, "service31_proc", "service31_container")
	containers = append(containers, service31_ctr)
	allServices = append(allServices, service31)

	//----- service30 depends on [118 119 120 121] ----------
	service30_db := mongodb.Container(spec, "service30_db")
	allServices = append(allServices, service30_db)
	service30 := workflow.Service[large_scale_app.Service30](
		spec,
		"service30",
		service30_db,
		service118,
		service119,
		service120,
		service121,
	)
	service30_ctr := applyDockerDefaults(spec, service30, "service30_proc", "service30_container")
	containers = append(containers, service30_ctr)
	allServices = append(allServices, service30)

	//----- service29 depends on [114 115 116 117] ----------
	service29_db := mongodb.Container(spec, "service29_db")
	allServices = append(allServices, service29_db)
	service29 := workflow.Service[large_scale_app.Service29](
		spec,
		"service29",
		service29_db,
		service114,
		service115,
		service116,
		service117,
	)
	service29_ctr := applyDockerDefaults(spec, service29, "service29_proc", "service29_container")
	containers = append(containers, service29_ctr)
	allServices = append(allServices, service29)

	//----- service28 depends on [110 111 112 113] ----------
	service28_db := mongodb.Container(spec, "service28_db")
	allServices = append(allServices, service28_db)
	service28 := workflow.Service[large_scale_app.Service28](
		spec,
		"service28",
		service28_db,
		service110,
		service111,
		service112,
		service113,
	)
	service28_ctr := applyDockerDefaults(spec, service28, "service28_proc", "service28_container")
	containers = append(containers, service28_ctr)
	allServices = append(allServices, service28)

	//----- service27 depends on [106 107 108 109] ----------
	service27_db := mongodb.Container(spec, "service27_db")
	allServices = append(allServices, service27_db)
	service27 := workflow.Service[large_scale_app.Service27](
		spec,
		"service27",
		service27_db,
		service106,
		service107,
		service108,
		service109,
	)
	service27_ctr := applyDockerDefaults(spec, service27, "service27_proc", "service27_container")
	containers = append(containers, service27_ctr)
	allServices = append(allServices, service27)

	//----- service26 depends on [102 103 104 105] ----------
	service26_db := mongodb.Container(spec, "service26_db")
	allServices = append(allServices, service26_db)
	service26 := workflow.Service[large_scale_app.Service26](
		spec,
		"service26",
		service26_db,
		service102,
		service103,
		service104,
		service105,
	)
	service26_ctr := applyDockerDefaults(spec, service26, "service26_proc", "service26_container")
	containers = append(containers, service26_ctr)
	allServices = append(allServices, service26)

	//----- service25 depends on [98 99 100 101] ----------
	service25_db := mongodb.Container(spec, "service25_db")
	allServices = append(allServices, service25_db)
	service25 := workflow.Service[large_scale_app.Service25](
		spec,
		"service25",
		service25_db,
		service98,
		service99,
		service100,
		service101,
	)
	service25_ctr := applyDockerDefaults(spec, service25, "service25_proc", "service25_container")
	containers = append(containers, service25_ctr)
	allServices = append(allServices, service25)

	//----- service45 depends on [178 179 180 181] ----------
	service45_db := mongodb.Container(spec, "service45_db")
	allServices = append(allServices, service45_db)
	service45 := workflow.Service[large_scale_app.Service45](
		spec,
		"service45",
		service45_db,
		service178,
		service179,
		service180,
		service181,
	)
	service45_ctr := applyDockerDefaults(spec, service45, "service45_proc", "service45_container")
	containers = append(containers, service45_ctr)
	allServices = append(allServices, service45)

	//----- service13 depends on [50 51 52 53] ----------
	service13_db := mongodb.Container(spec, "service13_db")
	allServices = append(allServices, service13_db)
	service13 := workflow.Service[large_scale_app.Service13](
		spec,
		"service13",
		service13_db,
		service50,
		service51,
		service52,
		service53,
	)
	service13_ctr := applyDockerDefaults(spec, service13, "service13_proc", "service13_container")
	containers = append(containers, service13_ctr)
	allServices = append(allServices, service13)

	//----- service15 depends on [58 59 60 61] ----------
	service15_db := mongodb.Container(spec, "service15_db")
	allServices = append(allServices, service15_db)
	service15 := workflow.Service[large_scale_app.Service15](
		spec,
		"service15",
		service15_db,
		service58,
		service59,
		service60,
		service61,
	)
	service15_ctr := applyDockerDefaults(spec, service15, "service15_proc", "service15_container")
	containers = append(containers, service15_ctr)
	allServices = append(allServices, service15)

	//----- service21 depends on [82 83 84 85] ----------
	service21_db := mongodb.Container(spec, "service21_db")
	allServices = append(allServices, service21_db)
	service21 := workflow.Service[large_scale_app.Service21](
		spec,
		"service21",
		service21_db,
		service82,
		service83,
		service84,
		service85,
	)
	service21_ctr := applyDockerDefaults(spec, service21, "service21_proc", "service21_container")
	containers = append(containers, service21_ctr)
	allServices = append(allServices, service21)

	//----- service20 depends on [78 79 80 81] ----------
	service20_db := mongodb.Container(spec, "service20_db")
	allServices = append(allServices, service20_db)
	service20 := workflow.Service[large_scale_app.Service20](
		spec,
		"service20",
		service20_db,
		service78,
		service79,
		service80,
		service81,
	)
	service20_ctr := applyDockerDefaults(spec, service20, "service20_proc", "service20_container")
	containers = append(containers, service20_ctr)
	allServices = append(allServices, service20)

	//----- service19 depends on [74 75 76 77] ----------
	service19_db := mongodb.Container(spec, "service19_db")
	allServices = append(allServices, service19_db)
	service19 := workflow.Service[large_scale_app.Service19](
		spec,
		"service19",
		service19_db,
		service74,
		service75,
		service76,
		service77,
	)
	service19_ctr := applyDockerDefaults(spec, service19, "service19_proc", "service19_container")
	containers = append(containers, service19_ctr)
	allServices = append(allServices, service19)

	//----- service18 depends on [70 71 72 73] ----------
	service18_db := mongodb.Container(spec, "service18_db")
	allServices = append(allServices, service18_db)
	service18 := workflow.Service[large_scale_app.Service18](
		spec,
		"service18",
		service18_db,
		service70,
		service71,
		service72,
		service73,
	)
	service18_ctr := applyDockerDefaults(spec, service18, "service18_proc", "service18_container")
	containers = append(containers, service18_ctr)
	allServices = append(allServices, service18)

	//----- service17 depends on [66 67 68 69] ----------
	service17_db := mongodb.Container(spec, "service17_db")
	allServices = append(allServices, service17_db)
	service17 := workflow.Service[large_scale_app.Service17](
		spec,
		"service17",
		service17_db,
		service66,
		service67,
		service68,
		service69,
	)
	service17_ctr := applyDockerDefaults(spec, service17, "service17_proc", "service17_container")
	containers = append(containers, service17_ctr)
	allServices = append(allServices, service17)

	//----- service16 depends on [62 63 64 65] ----------
	service16_db := mongodb.Container(spec, "service16_db")
	allServices = append(allServices, service16_db)
	service16 := workflow.Service[large_scale_app.Service16](
		spec,
		"service16",
		service16_db,
		service62,
		service63,
		service64,
		service65,
	)
	service16_ctr := applyDockerDefaults(spec, service16, "service16_proc", "service16_container")
	containers = append(containers, service16_ctr)
	allServices = append(allServices, service16)

	//----- service12 depends on [46 47 48 49] ----------
	service12_db := mongodb.Container(spec, "service12_db")
	allServices = append(allServices, service12_db)
	service12 := workflow.Service[large_scale_app.Service12](
		spec,
		"service12",
		service12_db,
		service46,
		service47,
		service48,
		service49,
	)
	service12_ctr := applyDockerDefaults(spec, service12, "service12_proc", "service12_container")
	containers = append(containers, service12_ctr)
	allServices = append(allServices, service12)

	//----- service6 depends on [22 23 24 25] ----------
	service6_db := mongodb.Container(spec, "service6_db")
	allServices = append(allServices, service6_db)
	service6 := workflow.Service[large_scale_app.Service6](
		spec,
		"service6",
		service6_db,
		service22,
		service23,
		service24,
		service25,
	)
	service6_ctr := applyDockerDefaults(spec, service6, "service6_proc", "service6_container")
	containers = append(containers, service6_ctr)
	allServices = append(allServices, service6)

	//----- service7 depends on [26 27 28 29] ----------
	service7_db := mongodb.Container(spec, "service7_db")
	allServices = append(allServices, service7_db)
	service7 := workflow.Service[large_scale_app.Service7](
		spec,
		"service7",
		service7_db,
		service26,
		service27,
		service28,
		service29,
	)
	service7_ctr := applyDockerDefaults(spec, service7, "service7_proc", "service7_container")
	containers = append(containers, service7_ctr)
	allServices = append(allServices, service7)

	//----- service11 depends on [42 43 44 45] ----------
	service11_db := mongodb.Container(spec, "service11_db")
	allServices = append(allServices, service11_db)
	service11 := workflow.Service[large_scale_app.Service11](
		spec,
		"service11",
		service11_db,
		service42,
		service43,
		service44,
		service45,
	)
	service11_ctr := applyDockerDefaults(spec, service11, "service11_proc", "service11_container")
	containers = append(containers, service11_ctr)
	allServices = append(allServices, service11)

	//----- service14 depends on [54 55 56 57] ----------
	service14_db := mongodb.Container(spec, "service14_db")
	allServices = append(allServices, service14_db)
	service14 := workflow.Service[large_scale_app.Service14](
		spec,
		"service14",
		service14_db,
		service54,
		service55,
		service56,
		service57,
	)
	service14_ctr := applyDockerDefaults(spec, service14, "service14_proc", "service14_container")
	containers = append(containers, service14_ctr)
	allServices = append(allServices, service14)

	//----- service10 depends on [38 39 40 41] ----------
	service10_db := mongodb.Container(spec, "service10_db")
	allServices = append(allServices, service10_db)
	service10 := workflow.Service[large_scale_app.Service10](
		spec,
		"service10",
		service10_db,
		service38,
		service39,
		service40,
		service41,
	)
	service10_ctr := applyDockerDefaults(spec, service10, "service10_proc", "service10_container")
	containers = append(containers, service10_ctr)
	allServices = append(allServices, service10)

	//----- service9 depends on [34 35 36 37] ----------
	service9_db := mongodb.Container(spec, "service9_db")
	allServices = append(allServices, service9_db)
	service9 := workflow.Service[large_scale_app.Service9](
		spec,
		"service9",
		service9_db,
		service34,
		service35,
		service36,
		service37,
	)
	service9_ctr := applyDockerDefaults(spec, service9, "service9_proc", "service9_container")
	containers = append(containers, service9_ctr)
	allServices = append(allServices, service9)

	//----- service8 depends on [30 31 32 33] ----------
	service8_db := mongodb.Container(spec, "service8_db")
	allServices = append(allServices, service8_db)
	service8 := workflow.Service[large_scale_app.Service8](
		spec,
		"service8",
		service8_db,
		service30,
		service31,
		service32,
		service33,
	)
	service8_ctr := applyDockerDefaults(spec, service8, "service8_proc", "service8_container")
	containers = append(containers, service8_ctr)
	allServices = append(allServices, service8)

	//----- service2 depends on [6 7 8 9] ----------
	service2_db := mongodb.Container(spec, "service2_db")
	allServices = append(allServices, service2_db)
	service2 := workflow.Service[large_scale_app.Service2](
		spec,
		"service2",
		service2_db,
		service6,
		service7,
		service8,
		service9,
	)
	service2_ctr := applyDockerDefaults(spec, service2, "service2_proc", "service2_container")
	containers = append(containers, service2_ctr)
	allServices = append(allServices, service2)

	//----- service5 depends on [18 19 20 21] ----------
	service5_db := mongodb.Container(spec, "service5_db")
	allServices = append(allServices, service5_db)
	service5 := workflow.Service[large_scale_app.Service5](
		spec,
		"service5",
		service5_db,
		service18,
		service19,
		service20,
		service21,
	)
	service5_ctr := applyDockerDefaults(spec, service5, "service5_proc", "service5_container")
	containers = append(containers, service5_ctr)
	allServices = append(allServices, service5)

	//----- service4 depends on [14 15 16 17] ----------
	service4_db := mongodb.Container(spec, "service4_db")
	allServices = append(allServices, service4_db)
	service4 := workflow.Service[large_scale_app.Service4](
		spec,
		"service4",
		service4_db,
		service14,
		service15,
		service16,
		service17,
	)
	service4_ctr := applyDockerDefaults(spec, service4, "service4_proc", "service4_container")
	containers = append(containers, service4_ctr)
	allServices = append(allServices, service4)

	//----- service3 depends on [10 11 12 13] ----------
	service3_db := mongodb.Container(spec, "service3_db")
	allServices = append(allServices, service3_db)
	service3 := workflow.Service[large_scale_app.Service3](
		spec,
		"service3",
		service3_db,
		service10,
		service11,
		service12,
		service13,
	)
	service3_ctr := applyDockerDefaults(spec, service3, "service3_proc", "service3_container")
	containers = append(containers, service3_ctr)
	allServices = append(allServices, service3)

	//----- service1 (entry) ----------
	service1_db := mongodb.Container(spec, "service1_db")
	allServices = append(allServices, service1_db)
	service1 := workflow.Service[large_scale_app.Service1](
		spec,
		"service1",
		service1_db,
		service2,
		service3,
		service4,
		service5,
	)
	service1_ctr := applyHTTPDefaults(spec, service1, "service1_proc", "service1_container")
	containers = append(containers, service1_ctr)
	allServices = append(allServices, service1)

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
