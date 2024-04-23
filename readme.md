
# medasdigital
**medasdigital** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Medasdigital is intended for future digital instantly payment purposes and tokenization of our company assets. We are runnin
g multiple validators and am planning rewards for delegators. Right now there is Sentinel, Mars, Bitcanna, Chihuahua and Fet
chAi (inactive)

We also host a Akash Providee Node


## technical Info

LCD
https://lcd.medas-digital.io:1317

RPC
https://rpc.medas-digital.io:26657

We cuurently have 5 own Validators running in 3 different datacenters

Uranus
Neptun
Jupiter
Mars
Saturn

Medas Digital Token is listed on Osmosis and integrated in Keplr Community Chain, we have to ressurect IBC Channel to Osmosis soon.

Unfortunately we had to do a hard fork to cosmos sdk v0.50, exporting the current state, but so old blockchain information is no longer present but the balances are.

For more technical visit our chain-registration at cosmos
https://github.com/cosmos/chain-registry/blob/master/medasdigital/chain.json

If anybody is interested supportig in DEV for the cosmos sdk we would be thankful. We are mainly infrastructure specialists.


## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

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
curl https://get.ignite.com/username/medasdigital@latest! | sudo bash
```
`username/medasdigital` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
