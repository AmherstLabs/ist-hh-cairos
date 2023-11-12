# cairos
**cairos** is CairoVM-Cosmos Integration Module, an innovative project designed to integrate the Cairo virtual machine with the Cosmos blockchain ecosystem. The Cairo virtual machine, known for its powerful and efficient smart contract execution capabilities, is enabling developers to leverage Cairo's features in a Cosmos-based environment.

# Features
Seamless Integration: Easily integrate Cairo VM with existing Cosmos SDK-based blockchains.
Smart Contract Support: Utilize Cairo for writing and executing smart contracts within the Cosmos ecosystem.
Performance Optimization: Leverage Cairo's efficient computation model for improved performance and lower gas costs.
Security: Enhanced security features inherent to Cairo VM, now available for Cosmos applications.
Interoperability: Maintain Cosmos's interoperability features while using Cairo's smart contract capabilities.

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
