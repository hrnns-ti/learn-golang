# Go — Variables
___

## Variables

### Using `var`
- Basic form:
    - `var name type`
- With initialization (type is optional if value is provided):
    - `var name string = "..." `
    - `var age uint8 = 18`
    - `var married bool = true`

### Using `:=` (short declaration)
- Only for the first declaration inside a function:
    - `name := "..." `
    - `age := uint8(18)`
    - `married := true`

### Multiple variables
- Declare multiple variables at once:
    - `var name1, name2, name3 = "A", "B", "C"`
    - `var x, y, z int = 1, 2, 3`

___

## Constants

### Single constant
- With type (optional if inferred):
    - `const name string = "..." `
    - `const name = "..."`

### Multiple constants
- Use a constant block:
    - `const (
      name1 = "A"
      name2 = "B"
      name3 = "C"
    )`
- Or single-line style:
    - `const (name1 = data, name2 = data, name3 = data)`
umber num = 64`