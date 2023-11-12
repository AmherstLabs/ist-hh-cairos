const { DirectSecp256k1HdWallet } = require("@cosmjs/proto-signing");
const { SigningStargateClient } = require("@cosmjs/stargate");

const rpcEndpoint = "http://0.0.0.0:26657"; // This should be the RPC endpoint of your node
const mnemonic =
  "demand cheap cancel version student hope team sorry chunk van party swarm belt deal convince define blue army bean flat avocado little deny snack"; // Replace with your actual mnemonic

async function createCosmosClient(rpcEndpoint, mnemonic) {
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: "cosmos",
  });
  const client = await SigningStargateClient.connectWithSigner(
    rpcEndpoint,
    wallet
  );
  return client;
}

async function getAccountDetails(address) {
  const client = await createCosmosClient(rpcEndpoint, mnemonic);
  const account = await client.getAccount(address);
  console.log(account);
}

async function main() {
  const address = "cosmos1apzfs6qgl9kj2mqym9rwl3xnxmd802gt0vdaky"; // Correct address for "alice"
  await getAccountDetails(address);
}

main().catch((error) => {
  console.error(error);
});
