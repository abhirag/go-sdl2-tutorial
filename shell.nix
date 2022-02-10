let
  # Pinned nixpkgs, deterministic. Last updated: 2/12/21.
  pkgs = import (fetchTarball("https://github.com/NixOS/nixpkgs/archive/a58a0b5098f0c2a389ee70eb69422a052982d990.tar.gz")) {};

in pkgs.mkShell {
buildInputs = [
    pkgs.SDL2
    pkgs.SDL2_mixer
    pkgs.SDL2_image
    pkgs.SDL2_ttf
    pkgs.SDL2_gfx
    pkgs.pkg-config
  ];

  CGO_CFLAGS="-I/nix/store/0hr82y4x1gmcdi960xh4nsnq5yaiakif-SDL2-2.0.14-dev/include/SDL2 -I/nix/store/k35vgq5aryinxs7vcs76n583myxi7m3w-SDL2_mixer-2.0.4/include/SDL2 -I/nix/store/ds0si1zn7dblvf5ig5cg4i5iimsafrz8-SDL2_image-2.0.5/include/SDL2 -I/nix/store/i1nmgbwbkvahq6zifvrh7sg2n1pr4b21-SDL2_ttf-2.0.15/include/SDL2 -I/nix/store/anr1rajs6f67vkkip0x20z67d0cbqr8f-SDL2_gfx-1.0.4/include/SDL2";
}