# Creates a development environment for the Advent of Code project.
let
  # Branch: nixos-unstable
  # Date of commit: 2025-11-30
  commit_ref = "2d293cbfa5a793b4c50d17c05ef9e385b90edf6c";
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/${commit_ref}";
  pkgs = import nixpkgs {
    config = { };
    overlays = [ ];
  };
in

pkgs.mkShellNoCC {
  packages = with pkgs; [
    go
    golangci-lint
    gopls
    tmux
  ];

  TMUX_SESSION = "[2025] Advent of Code";

  shellHook = ''
    export GOROOT=$( which go | xargs dirname | xargs dirname )/share/go
    tmux new-session -d -s "$TMUX_SESSION"
    exec tmux attach-session -t "$TMUX_SESSION"
  '';
}
