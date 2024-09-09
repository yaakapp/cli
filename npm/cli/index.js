const path = require("path");
const childProcess = require("child_process");
const {BINARY_DISTRIBUTION_PACKAGES} = require("./common");

// Windows binaries end with .exe so we need to special case them.
const binaryName = process.platform === "win32" ? "my-binary.exe" : "my-binary";

// Determine package name for this platform
const platformSpecificPackageName =
  BINARY_DISTRIBUTION_PACKAGES[process.platform];

function getBinaryPath() {
  try {
    // Resolving will fail if the optionalDependency was not installed
    return require.resolve(`${platformSpecificPackageName}/bin/${binaryName}`);
  } catch (e) {
    return path.join(__dirname, "..", binaryName);
  }
}

module.exports.runBinary = function (...args) {
  childProcess.execFileSync(getBinaryPath(), args, {
    stdio: "inherit",
  });
};
