# 🚧 Note
Geregè is still under development and some features may not work at the moment. Soon we will deploy the testable version once we figure out the correct way to modify Cosmos SDK to work with Cairo VM. 

# About
**Cairos**, a pioneering project to integrate Cairo VM with the Cosmos SDK, represents a landmark fusion of two revolutionary blockchain technologies. This integration aims to leverage the Cairo VM's renowned capabilities in provable computation and efficient smart contract execution, introducing a new realm of inter-chain interoperability between Starknet and Cosmos SDK-based chains. Our goal is to unlock unprecedented potential in smart contract functionality and blockchain interoperability, elevating the standards of decentralized applications. By harnessing Cairo VM's unique features, including its robust security protocols and scalability solutions, Cairos is set to redefine the landscape of blockchain interactions and applications.

# Key Features
- **Provable Computation**: Utilize Cairo VM's state-of-the-art provable computation for enhanced security and trust in smart contracts.
- **Efficient Smart Contract Execution**: Experience improved performance and reduced operational costs with Cairo VM's efficient smart contract execution within the Cosmos ecosystem.
- **Inter-Chain Interoperability**: Explore new horizons in blockchain interactions with seamless cross-chain communications between Starknet and Cosmos SDK-based chains.
- **Enhanced Cosmos SDK Chains**: Leverage the upgraded capabilities of Cosmos SDK-based chains, unlocking new potential applications and use cases.

# Prerequisites
Before begin testing, ensure you have the following installed:
- Go version 1.18 or higher
- Latest version of Cosmos SDK
- Familiarity with Go programming language, as Cosmos SDK is written in Go.
- Familiarity with the Cosmos SDK and its module structure.
- Familiarity with the Cairo language and its VM.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite-hq/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/cvm-cosmos@latest! | sudo bash
```
`username/cvm-cosmos` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
