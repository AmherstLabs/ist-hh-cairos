const axios = require("axios");

// Function to call compile cairo API
async function compileCairoContract(fileName) {
  try {
    const response = await axios.get(
      `http://localhost:3000/v1/compile_cairo/${fileName}`
    );
    console.log("Contract Compiled:", response.data);
    // Further processing or calling Cosmos SDK functions
  } catch (error) {
    console.error("Error compiling contract:", error);
  }
}

// Function to call execute cairo API
async function executeCairoContract(fileName) {
  try {
    const response = await axios.get(
      `http://localhost:3000/v1/call_cairo/${fileName}`
    );
    console.log("Contract Executed:", response.data);
    // Further processing or calling Cosmos SDK functions
  } catch (error) {
    console.error("Error executing contract:", error);
  }
}

// Assume you have a specific contract file name you want to compile and execute
const contractFileName = "simple";

compileCairoContract(contractFileName)
  .then(() => executeCairoContract(contractFileName))
  .catch((error) => console.error(error));
