package sht20

import (
	"atp-sht20/pkg/utils/domain"
	"context"
	"errors"
	"time"

	"github.com/simonvetter/modbus"
)

func (m shtStruct) Read(ctx context.Context, id uint8) (result domain.SHT20, err error) {
	url := "rtu://" + m.setting.Port
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      url,
		Speed:    uint(m.setting.Baudrate),
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  1 * time.Second,
	})
	if err != nil {
		err = errors.New("E0")
		return result, err
	}

	err = client.Open()
	if err != nil {
		err := errors.New("E1")
		return result, err
	}
	defer client.Close()

	client.SetUnitId(id)
	result.ID = id

	datas, err := client.ReadRegisters(1, 2, modbus.INPUT_REGISTER)
	//datas, err = client.ReadRegisters(1, 2, modbus.INPUT_REGISTER)
	if err != nil {
		err = errors.New("timeout")
		return result, err
	}

	for i, data := range datas {
		switch i {
		case 0: // temp
			result.Temp = float32(data) / 10.00
		case 1: // rh
			result.RH = float32(data) / 10.00
		}
	}

	return result, nil
}

func (m shtStruct) Write(ctx context.Context, prev, next uint8) (err error) {
	url := "rtu://" + m.setting.Port
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      url,
		Speed:    uint(m.setting.Baudrate),
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  3 * time.Second,
	})
	if err != nil {
		err = errors.New("E0")
		return err
	}

	err = client.Open()
	if err != nil {
		err := errors.New("E1")
		return err
	}
	defer client.Close()

	client.SetUnitId(prev)

	err = client.WriteRegister(uint16(0x0101), uint16(next))
	if err != nil {
		errN := errors.New("E2->" + err.Error())
		return errN
	}

	return nil
}
