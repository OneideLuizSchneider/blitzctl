#!/usr/bin/env sh
# POSIX shell installer for blitzctl
# - Detects OS/ARCH
# - Downloads the matching binary from GitHub Releases
# - Installs to /usr/local/bin or ~/.local/bin

set -eu

PROJECT_OWNER="OneideLuizSchneider"
PROJECT_REPO="blitzctl"
BIN_NAME="blitzctl"

# Defaults (can be overridden via flags)
BIN_DIR=""
VERSION="${BLITZCTL_VERSION:-}"

log() { printf "%s\n" "$*" 1>&2; }
err() { log "ERROR: $*"; exit 1; }

need_cmd() { command -v "$1" >/dev/null 2>&1 || err "Required command not found: $1"; }

detect_os() {
  os=$(uname -s 2>/dev/null || echo unknown)
  case "$os" in
    Linux) echo linux;;
    Darwin) echo darwin;;
    *) err "Unsupported OS: $os (supported: Linux, macOS)";;
  esac
}

detect_arch() {
  arch=$(uname -m 2>/dev/null || echo unknown)
  case "$arch" in
    x86_64|amd64) echo amd64;;
    aarch64|arm64) echo arm64;;
    *) err "Unsupported architecture: $arch (supported: amd64, arm64)";;
  esac
}

download() {
  url="$1"; out="$2"
  if command -v curl >/dev/null 2>&1; then
    curl -fL "$url" -o "$out"
  elif command -v wget >/dev/null 2>&1; then
    wget -O "$out" "$url"
  else
    err "Neither curl nor wget found to download $url"
  fi
}

choose_bindir() {
  # Honor explicit flag first
  if [ -n "$BIN_DIR" ]; then
    echo "$BIN_DIR"; return 0
  fi
  # Prefer /usr/local/bin when writable
  if [ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ]; then
    echo "/usr/local/bin"; return 0
  fi
  # Fallback to ~/.local/bin
  if [ -n "${HOME:-}" ]; then
    echo "$HOME/.local/bin"; return 0
  fi
  # Last resort: current directory
  pwd
}

usage() {
  cat 1>&2 <<EOF
Installer for $PROJECT_REPO

Usage:
  sh blitzctl.sh [--version vX.Y.Z] [--bin-dir DIR]

Environment:
  BLITZCTL_VERSION If set, install this tagged version (e.g., v0.0.1).

Examples:
  curl -fsSL https://raw.githubusercontent.com/$PROJECT_OWNER/$PROJECT_REPO/main/scripts/blitzctl.sh | sh -
  BLITZCTL_VERSION=v0.0.1 curl -fsSL https://raw.githubusercontent.com/$PROJECT_OWNER/$PROJECT_REPO/main/scripts/blitzctl.sh | sh -
EOF
}

parse_args() {
  while [ "$#" -gt 0 ]; do
    case "$1" in
      -h|--help)
        usage; exit 0;;
      -v|--version)
        [ "$#" -ge 2 ] || err "--version requires an argument"
        VERSION="$2"; shift 2;;
      -b|--bin-dir)
        [ "$#" -ge 2 ] || err "--bin-dir requires an argument"
        BIN_DIR="$2"; shift 2;;
      --)
        shift; break;;
      *)
        err "Unknown argument: $1";;
    esac
  done
}

main() {
  parse_args "$@"

  need_cmd uname
  need_cmd mktemp
  need_cmd chmod
  # curl or wget is checked by download()

  os=$(detect_os)
  arch=$(detect_arch)

  tmpdir=$(mktemp -d)
  trap 'rm -rf "$tmpdir"' EXIT INT HUP TERM

  file="$BIN_NAME"_${os}_${arch}

  base="https://github.com/$PROJECT_OWNER/$PROJECT_REPO/releases"
  if [ -n "$VERSION" ]; then
    # strip leading 'v' if user passed without it
    case "$VERSION" in
      v*) tag="$VERSION";;
      *) tag="v$VERSION";;
    esac
    url="$base/download/$tag/$file"
  else
    url="$base/latest/download/$file"
  fi

  dst="$tmpdir/$BIN_NAME"
  log "Downloading $BIN_NAME ($os/$arch) ..."
  download "$url" "$dst"
  chmod +x "$dst"

  bindir=$(choose_bindir)
  # Create bindir if needed
  if [ ! -d "$bindir" ]; then
    mkdir -p "$bindir"
  fi

  target="$bindir/$BIN_NAME"
  if mv "$dst" "$target" 2>/dev/null; then
    :
  else
    if command -v sudo >/dev/null 2>&1; then
      log "Elevating permissions to install into $bindir"
      sudo mv "$dst" "$target"
      sudo chmod +x "$target"
    else
      log "No permission to write to $bindir and sudo not available. Installing to ./"
      target="./$BIN_NAME"
      mv "$dst" "$target"
      chmod +x "$target"
    fi
  fi

  log "Installed: $target"
  if ! printf %s "${PATH}" | grep -q "$(dirname "$target")"; then
    log "Note: $(dirname "$target") is not in PATH. Add it to use '$BIN_NAME' globally."
  fi

  # Show version
  if "$target" --help >/dev/null 2>&1; then
    :
  fi
}

main "$@"

