# How to generate your own!

## Step 1

1. In a truffle based project, add `@chainsafe/truffle-plugin-abigen`

`npm i --save @chainsafe/truffle-plugin-abigen`

2. then in truffle-config, add the plugin

```json
  plugins: ["@chainsafe/truffle-plugin-abigen"],
  compilers: {
...
```

3. now you run `truffle run abigen`
4. then use the script i added in this repo, `generate.js`, the `name` field is the name
   of your contract
5. Find the generated go files in `custom_contracts` directory

# defi in golang

getting the abigen version of defi contracts is a pain - but i already did it for you

# Includes

aave, compound, plain erc20, uniswap, chainlink price feed

# example

Example included shows how to watch for aave flashloans
