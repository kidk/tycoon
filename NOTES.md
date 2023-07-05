
# Problemen

- FPS neemt al drop bij grotere grids. Mogelijk omdat hij alles rendert, maar heb twijfels bij camera performance.
- Animatie frames enzo worden nu steeds opnieuw gezet. Zou eigenlijk in spriteCache moeten zitten en /data/sprites.go

# Werk

## Phase 1: Een werkende game

Goal: Zo minimum mogelijk doen om een werkende game te hebben. Een werkende game is:

- [x] Mannetje hebben, en kunnen bewegen

  - [ ] Mannetje moet pixel per pixel bewegen, niet verspringen van positie

  - [ ] Voeg een direction + idle toe aan Player. UP, DOWN, LEFT, RIGHT, IDLE

- [ ] Tussen build mode en live mode kunnen wisselen (B build, G game, Q voor quickswitch)

- [ ] Orders kunnen geven voor bouwen (muur en vloer)

- [ ] Mannetje gebruiken om te bouwen (E voor bouwen)

- [ ] Zones kunnen toevoegen (build mode)

- [ ] Zone requirements (min bed, min bureau, ..)
