# Contributing

## A quick examples of adding a new function to test

1. **Identify the function to test**: Determine which function you want to add a test for. 

2. **Define Test Cases**: Create an array of `testType` struct instances for the different test cases you want to cover. Each `testType` struct instance must have:

    - `name`: A descriptive string explaining what this case is testing.
    - `call`: A function that calls the function to be tested with the parameters for this case. This should be a closure (i.e., an anonymous function) that captures the parameters for the case.
    - `wantErr`: A boolean indicating whether an error is expected from this case.

   Example:
   ```go
   tests := []testType{
		{
			name: "Your Test Name",
			call: func() (model.LasagnaLoveUser, error) {
				return YourFunctionToTest(param1, param2)
			},
			wantErr: false,
		},
   }
   ```

3. **Call `runTests`**: Once you've defined the test cases, call `runTests()` function with the `testing.T` instance and your test cases array.

   Example:
   ```go
   runTests(t, tests)
   ```

The `runTests()` function will iterate over each of your test cases and assert that the expected and actual error states match.

## References
[internal/user_access_test.go](../internal/user_access_test.go)