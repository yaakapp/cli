const path = require("path");
const {execSync} = require("child_process");

module.exports.getBinaryName = function () {
  return process.platform === "win32" ? "yaakcli.exe" : "yaakcli";
}

module.exports.getBinaryPath = function () {
  return path.join(execSync("npm prefix -g").toString().trim(), "bin", module.exports.getBinaryName())
}
