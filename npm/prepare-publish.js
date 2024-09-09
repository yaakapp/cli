const {copySync, readFileSync, writeFileSync} = require("node:fs");

console.log("Copying binary files to packages")
copySync('dist/cli_darwin_arm64/yaakcli', 'npm/cli-darwin-x64/bin/yaakcli');
copySync('dist/cli_darwin_amd64_v1/yaakcli', 'npm/cli-darwin-arm64/bin/yaakcli');
copySync('dist/cli_linux_amd64_v1/yaakcli', 'npm/cli-linux-x64/bin/yaakcli');
copySync('dist/cli_windows_amd64_v1/yaakcli.exe', 'npm/cli-win32-x64/bin/yaakcli.exe');

const version = process.env.YAAK_CLI_VERSINO
console.log(`Setting package versions to ${version}`);
replacePackageVersion('npm/cli', version);
replacePackageVersion('npm/cli-darwin-arm64', version);
replacePackageVersion('npm/cli-darwin-x64', version);
replacePackageVersion('npm/cli-linux-x64', version);
replacePackageVersion('npm/cli-win32-x64', version);

console.log("Done preparing for publish");

function replacePackageVersion(path, version) {
  const pkg = JSON.parse(readFileSync(path, 'utf-8'));
  pkg.version = version;
  if (pkg.name === '@yaakapp/cli') {
    pkg.optionalDependencies = {
      "@yaakapp/cli-darwin-x64": version,
      "@yaakapp/cli-darwin-arm64": version,
      "@yaakapp/cli-linux-x64": version,
      "@yaakapp/cli-win32-x64": version,
    }
  }
  writeFileSync(path, JSON.stringify(pkg, null, 2));
}
