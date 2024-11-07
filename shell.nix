{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.maven
    pkgs.docker
    pkgs.kubectl
    pkgs.skaffold
  ];

}

