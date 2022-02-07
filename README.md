# Golang Gamelib

Provides functions that may or may not be useful when making games in Golang,
in particular ones using the [Ebiten](https://ebiten.org/) library.

## Overview

### `github.com/csmith/gamelib/math` (general purpose)

* Linear interpolation functions (lerp)
* Easing functions

### `github.com/csmith/gamelib/sprite` (Ebiten specific)

* Sprite sheets
* Text rendering for sprite-based fonts

## Potential future additions

- [ ] RNG with bad luck prevention
- [ ] Animation helper
- [ ] Input utilities (click hotspots, key binds)
- [ ] Sound playback (global volume control, looping, etc)
- [ ] Template for WASM games
- [ ] Cross-compilation helper + glibc shim
- [ ] Tiled importer
- [ ] Sprite sheets with non-standard sizes (Aseprite importer?)
