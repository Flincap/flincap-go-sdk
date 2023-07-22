# Flincap SDK

[![npm shield](https://img.shields.io/npm/v/flincap-sdk)](https://www.npmjs.com/package/flincap-sdk)

The Flincap library provides access to the Flincap API from Go.

## Installation
To install see the lines below.

```sh
# using npm
npm install flincap-sdk

# using yarn
yarn add flincap-sdk
```

## Usage

[![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)](https://stackblitz.com/github/Flincap/flincap-sdk/tree/main/examples/backend)
<!-- [![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)](https://stackblitz.com/edit/raven-typescript-example-yrzyda?file=app.ts&view=editor) -->

```js
const baseURL = 'https://flincap.app/api';
const authToken = 'YOUR_AUTH_TOKEN'; // Optional, only if the API requires authentication

const apiClient = new FlincapApiClient(baseURL, authToken);

// Get rate
apiClient.getRate('USDT', 'NGN').then((response) => {
  console.log(response.data); // Process the rate data here
}).catch((error) => {
  console.error(error.message, error.code);
});

// Create transaction
const transactionData = {
  selectedCrypt: 'USDT',
  selectedFiat: 'NGN',
  email: 'test@example.com',
  bankName: 'Test Bank',
  bankCode: '12345',
  accNum: '9876543210',
  accName: 'John Doe',
  amountFiat: '5000',
  amountCrypt: '100',
  rate: '50',
};

apiClient.createTransaction(transactionData).then((response) => {
  console.log(response.data); // Process the transaction response here
}).catch((error) => {
  console.error(error.message, error.code);
});

// And so on for other API endpoints...
```

## Beta status

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning the package version to a specific version in your package.json file. This way, you can install the same version each time without breaking changes unless you are intentionally looking for the latest version.

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. Additions made directly to this library would have to be moved over to our generation code, otherwise they would be overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we will not be able to merge it as-is. We suggest [opening an issue](https://github.com/flincap/flincap-sdk) first to discuss with us!

On the other hand, contributions to the README are always very welcome!