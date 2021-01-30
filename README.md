# Goral-B

A WIP Go library for connecting to and getting data from (Braun) Oral-B Bluetooth enabled electric toothbrushes.

## Todo

* [x] Parsing BLE advertisements
* [x] Connecting to brush
* [x] Reading/Writing new right light colors
* [x] Enabling/Disabling ring light
* [ ] Change mode order / activate hidden modes

### Debugging
Using `gatttool` it is possible to manually send commands to the brush.

Connecting to the brush: `gatttool --device=<ADDR> -I`

#### Examples

**Setting the color to green (#00ff00)**
```
connect                         # Connect to the brush
char-write-req 0x0071 00ff0000  # Set color to 00ff00
char-write-req 0x0052 372f      # "Configure" color
char-write-req 0x0052 1031      # "Control" color enable
char-write-req 0x0052 1032      # "Control" color enable
```
