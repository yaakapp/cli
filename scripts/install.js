const fs = require("fs");
const path = require("path");
const zlib = require("zlib");
const https = require("https");
const {getBinaryPath, getBinaryName} = require("./util");

// Adjust the version you want to install. You can also make this dynamic.
const VERSION = "0.0.10";

// Windows binaries end with .exe so we need to special-case them.
const binaryName = getBinaryName();

// Compute the path we want to emit the binary to
const binaryInstallationPath = getBinaryPath()
fs.mkdirSync(path.dirname(binaryInstallationPath), {recursive: true});

function makeRequest(url) {
  return new Promise((resolve, reject) => {
    https
      .get(url, (response) => {
        if (response.statusCode >= 200 && response.statusCode < 300) {
          const chunks = [];
          response.on("data", (chunk) => chunks.push(chunk));
          response.on("end", () => {
            console.log("Finished downloading")
            resolve(Buffer.concat(chunks));
          });
        } else if (
          response.statusCode >= 300 &&
          response.statusCode < 400 &&
          response.headers.location
        ) {
          // Follow redirects
          makeRequest(response.headers.location).then(resolve, reject);
        } else {
          reject(
            new Error(
              `npm responded with status code ${response.statusCode}` +
              ` when downloading the package! ${url}`
            )
          );
        }
      })
      .on("error", (error) => {
        reject(error);
      });
  });
}

function extractFileFromTarball(tarballBuffer, filepath) {
  // Tar archives are organized in 512 byte blocks.
  // Blocks can either be header blocks or data blocks.
  // Header blocks contain file names of the archive in the first 100 bytes, terminated by a null byte.
  // The size of a file is contained in bytes 124-135 of a header block.
  // The following blocks will be data blocks containing the file.
  let offset = 0;
  while (offset < tarballBuffer.length) {
    const header = tarballBuffer.subarray(offset, offset + 512);
    offset += 512;

    const fileName = header.toString("utf-8", 0, 100).replace(/\0.*/g, "");
    const fileSize = parseInt(
      header.toString("utf-8", 124, 136).replace(/\0.*/g, ""),
      8
    );

    if (fileName === filepath) {
      return tarballBuffer.subarray(offset, offset + fileSize);
    }

    // Clamp offset to the uppoer multiple of 512
    offset = (offset + fileSize + 511) & ~511;
  }
}

async function downloadBinary() {
  // Download the tarball of the right binary distribution package
  const platform = process.platform === 'win32' ? 'windows' : process.platform;
  const arch = process.arch === 'x64' ? 'amd64' : process.arch;
  const url = `https://github.com/yaakapp/cli/releases/download/v${VERSION}/yaakcli_${VERSION}_${platform}_${arch}.tar.gz`;
  console.log(`Downloading ${url}`);
  const tarballDownloadBuffer = await makeRequest(url);

  console.log(`Extracting  ${tarballDownloadBuffer.length}`);
  const tarballBuffer = zlib.unzipSync(tarballDownloadBuffer);

  const dir = path.dirname(binaryInstallationPath);
  console.log(`Creating dir ${dir}`);
  fs.mkdirSync(dir, {recursive: true});

  // Extract binary from package and write to disk
  console.log(`Writing tarball to ${binaryInstallationPath}`);
  fs.writeFileSync(
    binaryInstallationPath,
    extractFileFromTarball(tarballBuffer, binaryName)
  );

  console.log(`Extracting tarball to ${binaryInstallationPath}`);
  // Make binary executable
  fs.chmodSync(binaryInstallationPath, "755");

  console.log("Done");
}

function isPlatformSpecificPackageInstalled() {
  return fs.existsSync(binaryInstallationPath)
}

// Skip downloading the binary if it was already installed via optionalDependencies
if (!isPlatformSpecificPackageInstalled()) {
  downloadBinary().catch(console.error);
} else {
  console.log(`yaakcli already exists at ${binaryInstallationPath}`);
}
