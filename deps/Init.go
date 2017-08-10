package hblock

import (
	"fmt"

	"log"
)

func create_empty_template(obj InitParams, logger *log.Logger) (int, error) {

	//output := obj.name
	print_Log("Init hb directory.", logger)
	_, err := hb_Init()
	if err != nil {
		return FAIL, err
	}
	//if PathFileExists(obj.name) {
	if VerifyBackingFile(obj.name) == OK {
		return FAIL, fmt.Errorf("Already exist.")
		//	return FAIL, nil
	}

	print_Log("Create hyperlayer object...", logger)
	h, err := CreateHBM(&obj)
	if err != nil {
		return FAIL, err
	}
	// print_Log("Create backing file config file.", logger)

	// configPath := h.getImgConfigPath()
	// yamlConfig := YamlBackingFileConfig{
	// 	Name:        h.name,
	// 	VirtualSize: obj.size,
	// 	DefaultHead: "master",
	// 	Format:      obj.format,
	// }
	// err = WriteConfig(yamlConfig, &configPath)
	// if err != nil {
	// 	return FAIL, fmt.Errorf("write backingfile's config failed.")
	// }
	h.CreateDisk()
	msg := fmt.Sprintf("Create template '%s' finished.", obj.name)
	print_Log(Format_Success(msg), logger)
	if !obj.checkout {
		return OK, nil
	}

	// print_Log("Creating volume named "+obj.output, logger)

	// checkoutObj := CheckoutParams{layer: "", output: obj.output, template: obj.name}
	// ret, err := volume_checkout(&checkoutObj, logger)
	// if err != nil {
	// 	return FAIL, err
	// }
	// if ret == OK {
	// 	checkoutObj.branch = "master"
	// 	checkoutObj.volume = obj.output
	// 	return volume_checkout(&checkoutObj, logger)
	// }
	return OK, nil
}
