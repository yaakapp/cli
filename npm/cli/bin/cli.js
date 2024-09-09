#!/usr/bin/env node

const path = require("path");
const childProcess = require("child_process");
const {BINARY_NAME} = require("../common");

// Lookup table for all platforms and binary distribution packages
const BINARY_DISTRIBUTION_PACKAGES = {
  darwin: "npm-binary-example-darwin",
  linux: "npm-binary-example-linux",
  freebsd: "npm-binary-example-linux",
  win32: "npm-binary-example-win32",
};

// Determine package name for this platform
const platformSpecificPackageName =
  BINARY_DISTRIBUTION_PACKAGES[process.platform];

function getBinaryPath() {
  try {
    // Resolving will fail if the optionalDependency was not installed
    return require.resolve(`${platformSpecificPackageName}/bin/${BINARY_NAME}`);
  } catch (e) {
    return path.join(__dirname, "..", BINARY_NAME);
  }
}

childProcess.execFileSync(getBinaryPath(), process.argv.slice(2), {
  stdio: "inherit",
});
