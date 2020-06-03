# Code Search

My own implementation of Google's code search. This project can search through any code base of a public repository on GitHub. Currently, only searching through the master branch is supported.

## Installation
Using Docker:
```
docker build -t codesearch . && docker run -p 5000:5000 codesearch
```

Or natively:
```
npm install && yarn start
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://github.com/paulmj7/codesearch/blob/master/LICENSE)
