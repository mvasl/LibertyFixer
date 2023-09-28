# LibertyFixer

A tool to (improperly) fix Cyberpunk 2077 crashing after the Phantom Liberty update.

Phantom Liberty update introduced a problem with shader caching, it cannot load cached shaders anymore, so I hacked together a launcher that forces the game to recompile these shaders.

It was made in 30 minutes, was only tested on AMD Radeon RX 7900 XT, 
but should, in theory, work with any AMD GPU. 

Nvidia/Intel GPU's support may come at some later time, but I hope CDPR will fix their game before I do)

## How to use it

Basically, get the LibertyFixer.exe from the [releases](https://github.com/mvasl/LibertyFixer/releases) section or build it from source if you want.

Drop `LibertyFixer.exe` into your `<CP2077_Dir>\bin\x64` folder where `Cyberpunk2077.exe` is 
and launch `LibertyFixer.exe` instead of the game every time until there is a proper fix from CDPR. 

## How to build from source

Download Go toolchain at [go.dev](https://go.dev/dl), install it, then see `build.bat`, it's pretty straightforward.
You can run `build.bat` on your windows system after installing Go and in will produce the executable for you.
