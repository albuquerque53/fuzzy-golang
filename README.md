# fuzzy-golang

![image](https://user-images.githubusercontent.com/57183466/199854372-dd0e39fc-0841-46d1-bb8b-faa25d806787.png)


> A simple project for my learnings in Golang's fuzz testing

## The application

This app is a simple calculator with sum, subtration, divison and multiplication that returns a JSON with the asked operation and result.<br>

To run, simply:<br>

```
make run
```

## Fuzz Testing


In the `math` package there are "simple" tests with manually specified sceneries but also each test file have a Fuzz Test.<br>

To run the tests:

```
make test
```

To run the fuzz tests:

```
make fuzz
```
