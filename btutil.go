package goralb

import "github.com/muka/go-bluetooth/bluez/profile/gatt"

func EnableNotifications(char *gatt.GattCharacteristic1, callback func(buf []byte)) error {
	ch, err := char.WatchProperties()
	if err != nil {
		return err
	}
	go func() {
		for update := range ch {
			if update.Interface == "org.bluez.GattCharacteristic1" && update.Name == "Value" {
				callback(update.Value.([]byte))
			}
		}
	}()
	return char.StartNotify()
}
