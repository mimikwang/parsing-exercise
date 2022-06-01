# parsing-exercise
A simple json parsing exercise and solutions in various languages

## Problem Statement
You are given a folder of `json` files.  The schema can be either one of two different types (`A` or `B`) or neither.  The goal is to parse them and format them into a defined output format or ignore and log the filenames of the invalid format.  The solution should be a command line tool that takes in the following command:

```bash
$ parse <path to input folder> <path to output folder>
```

The output folder should contain the parsed outputs of the valid input files that are named the same and an `invalid.log` file that logs the names of the invalid files.

## Format Definitions
For both `A` and `B`, all fields are required to be present and contain the correct types.  Extra fields are allowed.

### Output Format
The following describes the output format.

```javascript
{
    "id": 1234,
    "name": "bob",
    "age": 38,
    "bmi": "normal"
}
```

### `A`
The following is an example of the `A` file format.

```javascript
{
    "id": 1234,
    "name": "bob",
    "country": "USA",
    "age": 38,
    "weight": 120.5,
    "height": 6.2
}
```

The `id`, `name`, and `age` fields map directly to the same fields in the output.  The `bmi` field is calculated by first calculating the numeric `bmi` by taking the `weight` / `height`^2 and assigning string value based on the following table.

| BMI Numeric | Label           |
| ----------- | --------------- |
| Below 18.5  | underweight     |
| 18.5 - 24.9 | normal          |
| 25.0 - 29.9 | overweight      |
| 30.0 and up | very overweight |

### `B`
The following is an example of the `B` file format.

```javascript
{
    "id": 1234,
    "info": {
        "name": "bob",
        "country": "USA"
    },
    "data": {
        "bmi": 18.5,
        "age": 38
    }
}
```

The `id` field maps directly to the same fields in the output.  The `info.name` field maps to the `name` field in the output, and the `data.age` field maps to the `age` field in the output.  The `data.bmi` field is subjected to the same thresholds as defined in the table above for the labeling and maps to the `bmi` field in the output.
