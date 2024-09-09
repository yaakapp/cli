// Lookup table for all platforms and binary distribution packages
const BINARY_DISTRIBUTION_PACKAGES = {
  darwin_arm64: "@yaakapp/cli-darwin-arm64",
  darwin_x64: "@yaakapp/cli-darwin-x64",
  linux_x64: "@yaakapp/cli-linux-x64",
  win32_x64: "@yaakapp/cli-win32-x64",
};

// Adjust the version you want to install. You can also make this dynamic.
const BINARY_DISTRIBUTION_VERSION = require('./package.json').version;

module.exports = {BINARY_DISTRIBUTION_PACKAGES, BINARY_DISTRIBUTION_VERSION};
