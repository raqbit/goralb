# Goral-B

A WIP Go library for connecting to and getting data from (Braun) Oral-B Bluetooth enabled electric toothbrushes.

## Todo

* [x] Parsing BLE advertisements
* [x] Connecting to brush
* [ ] Reading/Writing characteristics
    * [ ] Getting the battery status
    * [ ] Getting the color of the color ring
    * [x] Setting the color of the color ring
    * [ ] Pulling brushing session data from brush

### Debugging
Using `gatttool` it is possible to manually send commands to the brush.

Connecting to the brush: `gatttool --device=<ADDR> -I`

**Setting the color to green (#00ff00)**
```
connect                         # Connect to the brush
char-write-req 0x0071 00ff0000  # Set color to 00ff00
char-write-req 0x0052 1031      # Makes the light go on (and saves the change)
char-write-req 0x0052 372f      # Unknown status update
```

### Status values
These are the status values I've found being written to the brush.

|Value|Description|
|---|---|
|`1030`|Unknown|
|`1031`|Displays the current color value|
|`02fa..02fb`|Unknown|
|`0200..021d`|Unknown counter from 512 to 541 (29)|
|`311e`|Unknown|
|`31ff`|Unknown|
|`3351`|Unknown|
|`372f`|Unknown status update right after setting color value|
