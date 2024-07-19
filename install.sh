#!/bin/bash
# Copyright MontFerret Team 2023
# Licensed under the MIT license.

set -e

# Declare constants
readonly projectName="yaakapp"
readonly appName="cli"
readonly binName="yaak-cli"
readonly fullAppName="${projectName} $(echo ${appName} | awk '{print toupper(substr($0,0,1)) substr($0,2)}')"
readonly baseUrl="https://github.com/${projectName}/${appName}/releases/download"

# Declare default values
readonly defaultLocation="${HOME}/.yaakcli"
readonly defaultVersion="latest"

# Print a message to stdout
report() {
  command printf "%s\n" "$*" 2>/dev/null
}

# Check if a command is available
command_exists() {
  command -v "$1" >/dev/null 2>&1
}

# Check if a path exists
check_path() {
  if [ -z "${1-}" ] || [ ! -f "${1}" ]; then
    return 1
  fi

  report "${1}"
}

# Validate user input
validate_input() {
  local location="$1"
  local version="$2"

  if [ -z "$location" ]; then
    report "Invalid location: $location"
    exit 1
  fi

  # Check if location exists
  if [ ! -d "$location" ]; then
    mkdir -p "$location"
    report "Creating directory: $location"
  fi

  # Check if location is writable
  if [ ! -w "$location" ]; then
    report "Location is not writable: $location"
    exit 1
  fi

  if [ "$version" != "latest" ]; then
    # Remove leading 'v' if present
    version="${version#v}"

    # Check if version is valid using grep
    if ! echo "$version" | grep -qE "^[0-9]+\.[0-9]+\.[0-9]+$"; then
      report "Invalid version: $version"
      exit 1
    fi
  fi
}

# Detect the profile file
detect_profile() {
  local profile=""
  local detected_profile=""

  if [ "${PROFILE-}" = '/dev/null' ]; then
    # the user has specifically requested NOT to have nvm touch their profile
    return
  fi

  if [ -n "${PROFILE}" ] && [ -f "${PROFILE}" ]; then
    report "${PROFILE}"
    return
  fi

  if command_exists bash; then
    if [ -f "$HOME/.bashrc" ]; then
      detected_profile="$HOME/.bashrc"
    elif [ -f "$HOME/.bash_profile" ]; then
      detected_profile="$HOME/.bash_profile"
    fi
  elif command_exists zsh; then
    if [ -f "$HOME/.zshrc" ]; then
      detected_profile="$HOME/.zshrc"
    fi
  fi

  if [ -z "$detected_profile" ]; then
    for profile_name in ".profile" ".bashrc" ".bash_profile" ".zshrc"; do
      if detected_profile="$(check_path "${HOME}/${profile_name}")"; then
        break
      fi
    done
  fi

  if [ -n "$detected_profile" ]; then
    report "$detected_profile"
  fi
}

# Update the profile file
update_profile() {
  local location="$1"
  local profile="$(detect_profile)"

  if [[ ":$PATH:" == *":$location:"* ]]; then
    return
  fi

  report "Updating profile $profile"

  if [ -z "$profile" ]; then
    report "Profile not found. Tried ${DETECTED_PROFILE-} (as defined in \$PROFILE), ~/.bashrc, ~/.bash_profile, ~/.zshrc, and ~/.profile."
    report "Append the following lines to the correct file yourself:"
    report
    report "export PATH=\$PATH:${location}"
    report
  else
    if ! grep -q "${location}" "$profile"; then
      report "export PATH=\$PATH:${location}" >>"$profile"
    fi
  fi
}

# Get the platform-specific filename suffix
get_platform_suffix() {
  local platform_name="$(uname)"
  local arch_name="$(uname -m)"
  local platform=""
  local arch=""

  case "$platform_name" in
  "Darwin")
    platform="darwin"
    ;;
  "Linux")
    platform="linux"
    ;;
  "Windows")
    platform="windows"
    ;;
  *)
    report "$platform_name is not supported. Exiting..."
    exit 1
    ;;
  esac

  case "$arch_name" in
  "x86_64")
    arch="_x86_64"
    ;;
  "aarch64" | "arm64")
    arch="_arm64"
    ;;
  *)
    report "$arch_name is not supported. Exiting..."
    exit 1
    ;;
  esac

  echo "${platform}${arch}"
}

get_version_tag() {
  local version="$1"

  if [ "$version" = "latest" ]; then
    local url="https://api.github.com/repos/${projectName}/${appName}/releases/latest"

    curl -sSL "${url}" | grep "tag_name" | cut -d '"' -f 4
  else
    # Check if the version starts with a 'v'
    if [[ "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
      echo "$version"
    else
      echo "v$version"
    fi
  fi
}

# Install the package
install() {
  report "VERSION=$VERSION"
  local location="${LOCATION:-$defaultLocation}"
  local version=$(get_version_tag "${VERSION:-$defaultVersion}")
  report "VERSION2=$version"
  local tmp_dir="$(mktemp -d -t "${projectName}.${appName}.XXXXXXX")"

  validate_input "$location" "$version"

  report "Installing ${projectName} ${appName} ${version}..."

  # Download the archive to a temporary location
  local suffix="$(get_platform_suffix)"
  local file_name="${appName}_${version#v}_${suffix}"
  local download_dir="${tmp_dir}/${file_name}@${version}"

  mkdir -p "${download_dir}"

  local download_file="${download_dir}/${file_name}.tar.gz"
  local url="${baseUrl}/${version}/${file_name}.tar.gz"

  report "Downloading package $url as $download_file"

  curl -sSL "${url}" | tar xz --directory "${download_dir}"

  local downloaded_file="${download_dir}/${binName}"

  report "Copying ${downloaded_file} to ${location}"

  cp "${downloaded_file}" "${location}"

  local executable="${location}/${binName}"

  chmod +x "${executable}"

  update_profile "${location}"

  report "New version of ${fullAppName} installed to ${location}"

#  "$executable" version
}

# Call the main function
install
