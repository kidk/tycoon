
# Problemen

- FPS neemt al drop bij grotere grids. Mogelijk omdat hij alles rendert, maar heb twijfels bij camera performance.
- Animatie frames enzo worden nu steeds opnieuw gezet. Zou eigenlijk in spriteCache moeten zitten en /data/sprites.go

# Werk

## Phase 1: Een werkende game

Goal: Zo minimum mogelijk doen om een werkende game te hebben. Een werkende game is:

- [x] Mannetje hebben, en kunnen bewegen

  - [x] Mannetje moet pixel per pixel bewegen, niet verspringen van positie

  - [x] Voeg een direction + idle toe aan Player. UP, DOWN, LEFT, RIGHT, IDLE

- [ ] Orders kunnen geven voor bouwen (muur en vloer)

  - [ ] Bouw status toevoegen aan tiles

  - [ ] Te bouwen tiles kan je door wandelen

  - [ ] Pathfinding aanpassen nadat iets gebouwd is

  - [ ] Mannetje gebruiken om te bouwen (E voor bouwen)

- [ ] Zones kunnen toevoegen (build mode)

  - [ ] Zones zijn alleen zichtbaar tijdens bouwen zones (dus klik op zone button)

  - [ ] Zones detecteren en groeperen
 
  - [ ] Zone requirements (min bed, min bureau, ..) Optioneel

- [ ] Geld toevoegen en kosten voor muren/vloeren ..

- [ ] Guests toevoegen

  - [ ] Guests lopen naar receptie en blijven 10 seconden staan (progressbar?)

  - [ ] Guests lopen naar vrije kamer en blijven 60 seconden op bed staan (progressbar?)

  - [ ] Guests lopen naar receptie en blijven 10 seconden staan en we krijgen 100 geld bij (progressbar?)

  - [ ] Er zijn meerdere guests en we moeten ze verdelen per kamer

- [ ] UI verbeteren

  - [ ] Icoontjes toevoegen

  - [ ] Sub menu's