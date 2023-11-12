const express = require("express");
const { execSync } = require("child_process");
const app = express();
const port = 3000;

app.get("/v1/call_cairo/:fileName", (req, res) => {
  const { fileName } = req.params;
  const filePath = `./integration_tests/cairo_files/${fileName}.cairo`;

  try {
    const result = execSync(
      `cd cairo-vm-go && cairo-compile ${filePath} --proof_mode --output ./${fileName}_compiled.json`
    ).toString();

    res.send(result);
  } catch (error) {
    res.status(500).send("Error executing command");
  }
});

// Start the server
app.listen(port, () => {
  console.log(`Server is listening on port ${port}`);
});
