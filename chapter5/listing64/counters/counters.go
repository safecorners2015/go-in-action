// Define counters package
package counters

// Declare alertCounter as non-explict type
type alertCounter int

// Define Factory Function
// To Generate Value as a alertCounter type
func New(value int) alertCounter {
	return alertCounter(value)
}
