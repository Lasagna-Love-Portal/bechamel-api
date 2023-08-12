# After installing nix, invoke nix-shell in the directory containing this file

{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-23.05.tar.gz") {} }:

pkgs.mkShell {
  buildInputs = [
    # Go
    pkgs.go
  ];
}
