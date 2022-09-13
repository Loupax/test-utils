package utils

// Arguments is syntactic sugar for defining variadic arguments in a test without having to use that verbose array syntax
//
// Example using testify/mock:
// var testCases := []struct{
//   name string
//   mock_args []interface{}
//   mock_returns []interface{}
// } {
//  {
//   name: "Add returns the sum of its arguments",
//   mock_args: test_utils.Arguments(1,2,3,4),
//   mock_returns: test_utils.Arguments(10),
//  }
// }
//
// for tt := range testCases{
//   t.Run(tt.name, func(t *testing.T){
//    m := new(MyMock)
//    m.On("Add", tt.mock_args...).Return(tt.mock_returns...)
//    m.AssertExpectations(t)
//   })
// }
//
func Arguments(arg ...interface{}) []interface{} {
	return arg
}
