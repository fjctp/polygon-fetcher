{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = [ 
      pkgs.go
      pkgs.just
      pkgs.miniserve];
}
