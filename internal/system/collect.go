package system

func Collect() (*Info, error) {

	info := &Info{}

	if err := collectOS(info); err != nil {
		return nil, err
	}

	if err := collectKernel(info); err != nil {
		return nil, err
	}

	if err := collectHashcat(info); err != nil {
		return nil, err
	}

	if err := collectCPU(info); err != nil {
		return nil, err
	}

	if err := collectMemory(info); err != nil {
		return nil, err
	}

	if err := collectDisk(info); err != nil {
		return nil, err
	}

	_ = collectGPU(info)

	_ = collectWiFi(info)

	collectTools(info)

	return info, nil
}
