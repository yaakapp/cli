#!/usr/bin/env node

const path = require("path");
const childProcess = require("child_process");
const {BINARY_NAME, PLATFORM_SPECIFIC_PACKAGE_NAME} = require("../common");

function getBinaryPath() {
  try {
    // Resolving will fail if the optionalDependency was not installed
    return require.resolve(`${PLATFORM_SPECIFIC_PACKAGE_NAME}/bin/${BINARY_NAME}`);
  } catch (e) {
    return path.join(__dirname, "..", BINARY_NAME);
  }
}

childProcess.execFileSync(getBinaryPath(), process.argv.slice(2), {
  stdio: "inherit",
});
