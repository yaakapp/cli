const path = require("path");
const childProcess = require("child_process");
const {PLATFORM_SPECIFIC_PACKAGE_NAME, BINARY_NAME} = require("./common");

function getBinaryPath() {
  try {
    // Resolving will fail if the optionalDependency was not installed
    return require.resolve(`${PLATFORM_SPECIFIC_PACKAGE_NAME}/bin/${BINARY_NAME}`);
  } catch (e) {
    return path.join(__dirname, "..", BINARY_NAME);
  }
}

module.exports.runBinary = function (...args) {
  childProcess.execFileSync(getBinaryPath(), args, {
    stdio: "inherit",
  });
};
