#!/usr/bin/env node

const util = require("util");
const fs = require("fs");
const child_process = require("child_process");
const execFile = util.promisify(require("child_process").execFile);

const GENERATE = [
  {
    name: "Liquidation",
    pkg: "liquidate",
    go_file: "liq.go",
  },
  {
    name: "Vault",
    pkg: "vault",
    go_file: "v.go",
  },
];

const GEN_FILE_DIR = `custom_contracts`;

(async () => {
  await execFile("rm", ["-rf", GEN_FILE_DIR]);
  await execFile("mkdir", ["-p", GEN_FILE_DIR]);

  for (const gen of GENERATE) {
    await execFile("mkdir", ["-p", `${GEN_FILE_DIR}/${gen.pkg}`]);
    const bin = `--bin=abigenBindings/bin/${gen.name}.bin`;
    const abi = `--abi=abigenBindings/abi/${gen.name}.abi`;
    const pkg = `--pkg=${gen.pkg}`;
    const go_code = `--out=${GEN_FILE_DIR}/${gen.pkg}/${gen.go_file}`;
    const args = [bin, abi, pkg, go_code];
    console.log(args);
    await execFile("abigen", args);
  }
})();
