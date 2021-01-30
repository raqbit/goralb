## Status characteristic

UUID: `a0f0ff21-5047-4d53-8208-4f72616c2d42`

The status characteristic is used to update settings or control functions of the brush.
There are currently two main types of values I've figured out: control values and configure values.

### Properties

- Notify
- Read
- Write

### Format

#### Control (0x10)

Control values are used to control a function on the brush, like turning the ring light on/off or turning the brush on.

|Value|Description|
|---|---|
|`0x10 0x30`|Disable the ring light|
|`0x10 0x31`|Enable the ring light|

#### Configure (0x37)

Configure values are used to update settings of the brush, like for instance making it save the value of the color characteristic.

|Value|Description|
|---|---|
|`0x37 0x2f`|Configure the color of the right light|


## Brushing modes characteristic

UUID: `a0f0ff25-5047-4d53-8208-4f72616c2d42`

Seems to have a list of [modes](./mode.go) which might be used to define the order of the modes that the brush 
cycles through when you press the mode-switch button. It might also be possible this way to activate modes which 
are not normally available on the specific brush model.

More research needed.

### Properties

- Read
- Write

## Color characteristic

UUID: `a0f0ff2b-5047-4d53-8208-4f72616c2d42`

The color characteristic can be used to read/write the color of the LED ring light.

### Properties

- Read
- Write

### Format

|byte|Purpose|
|---|---|
|1|Red value|
|2|Green value|
|3|Blue value|
|4|??|
