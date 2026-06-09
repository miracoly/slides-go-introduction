{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-26.05";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
      go-version = pkgs.writeShellScript "go-version" ''
        set -euo pipefail
        go version | sed 's/go version go\([^ ]*\).*/\1/'
      '';
    in {
      devShells.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          buf
          delve
          go
          go-task
          golangci-lint
          gopls
          gotools
          grpcurl
          protoc-gen-go
          protoc-gen-go-grpc
          zed-editor
        ];
        shellHook = ''
          red=$(tput setaf 1 2>/dev/null || true)
          green=$(tput setaf 2 2>/dev/null || true)
          yellow=$(tput setaf 3 2>/dev/null || true)
          bold=$(tput bold 2>/dev/null || true)
          reset=$(tput sgr0 2>/dev/null || true)

          printf "$red$bold┌──────────────────────────────────────────────────────────┐\n"
          printf "│                        Go Shell                          │\n"
          printf "└──────────────────────────────────────────────────────────┘$reset\n\n"

          printf "$yellow Tooling $reset\n"
          printf "  $bold Go:                   $reset $(${go-version})\n"
          printf "  $bold buf:                  $reset $(buf --version)\n"
          printf "  $bold golangci-lint:        $reset $(golangci-lint version --short)\n"

          unset red green yellow bold reset
        '';
      };
    });
}
