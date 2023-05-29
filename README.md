# Order API

This project is a lightweight API which accepts a number and determines what quantities of a range of pre-defined pack-sizes are required to fulfil the order.
The Supported packsizes are as follows:
`250, 500, 750, 1000, 2000, 5000`
With the intention to support addition and removal of packsizes in future.
As pack sizes cannot be adjusted, the role of the api is to return the minimum number items, broken down into the minimum number of packs in order to fulfill the order.
The calculation prioritises item numbers over pack numbers, meaning that a greater number of packs which delivers a fewer number of items is favoured over fewer packs which return more items.

Example test cases are as follows:

| Items ordered | Correct number of packs                | Incorrect number of packs |
| ------------- | -------------------------------------- | ------------------------- |
| 1             | 2 x 250                                | 1 x 500                   |
| 250           | 1 x 250                                | 1 x 500                   |
| 251           | 1 x 500                                | 2 x 250                   |
| 501           | 1 x 500, 1 x 250                       | 1 x 1000 OR 3 x 250       |
| 8500          | 1 x 5000, 1 x 2000, 1 x 1000 , 1 x 500 | 1 x 5000, 2 x 2000        |
| 12001         | 2 x 5000, 1 x 2000, 1 x 250            | 3 x 5000                  |

## Deployed Application

The deployed application can be queried at <a href="https://bqst6iv3sj.eu-west-1.awsapprunner.com/">https://bqst6iv3sj.eu-west-1.awsapprunner.com/</a>

### Accepted endpoints:

| Method | Endpoint       | Body           |
| ------ | -------------- | -------------- |
| POST   | "/place-order" | { "count": 1 } |

### Deployed UI

Additionally, this API can also be interacted via the user interface deployed at <a href="https://atcxanpfvm.eu-west-1.awsapprunner.com/">https://atcxanpfvm.eu-west-1.awsapprunner.com/</a>
<br>
You can also view the repository for the user interface <a href="https://github.com/laurawalters1/order-api-ui">here</a>

### Local Setup

If you would like to test this API locally, please clone the repository, and run `go get .` in the root of the repository. Once the dependencies are installed, you can run the application using `go run .`, additionally if you would like to run the test file which has been created for the pack quantity calculation, please run `cd calculatepacks` followed by `go test`.
