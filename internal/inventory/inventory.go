package inventory

func List() (*Inventory, error) {

	inv := &Inventory{}

	adapters, err := discover()
	if err != nil {
		return nil, err
	}

	adapters, err = enrichIW(adapters)
	if err != nil {
		return nil, err
	}

	adapters, err = enrichIP(adapters)
	if err != nil {
		return nil, err
	}

	adapters, err = enrichDriver(adapters)
	if err != nil {
		return nil, err
	}

	inv.Adapters = adapters

	return inv, nil
}
