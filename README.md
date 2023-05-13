<div align='center'>

# [Color String](https://github.com/pandasoli/colorstring) <img width='32' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-blushing.png' alt='Gopher icon'/>

This lib is usable for console (terminal) apps when you need to manage strings with color codes.

I also have [goterm](https://github.com/pandasoli/goterm) that is a functions pack for console apps.  
Also I recommend you to use [Guffers](https://github.com/pandasoli/guffers) for a more complex CLI development.
</div>

<br/>

## Why colorstrings? <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-thinking.png' alt='Gopher icon'/>

I prefer to show you...

The problem it fixes:

```go
// Putting a \033[0m after 'hi'

/* --- Without colorstrings --- */

str := " \033[31mHey \033[1mhi! "
str = str[:7] + "\033[0m" + str[7:]   // ✗ ' \e[31mH\e[0mey \e[1mhi! '
str = str[:16] + "\033[0m" + str[16:] // ✔ ' \e[31mHey \e[1mhi\e[0m! '


/* --- With colorstrings --- */

str, _ := NewColorString(" \033[31mHey \033[1mhi! ")
str.Colorize("0", 7)
str.Join() // Get the result

// ✔ ' \e[31mHey \e[1mhi\e[0m! '
```

What is the problem, and how it fixes it:

```go
/*
 * in a pure string:
 *   ' Hey hi! '
 *           ^ 7
 *
 * in a colored string:
 *   ' \033[31mHey \033[1mhi! '
 *           ^ 7            ^ 16
 *
 * Colorstrings keep the positions.
 */
```

<br/>
<div align='right'>

## How to set up <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-pirate.png' alt='Gopher icon'/>
</div>

Run this code in your terminal inside your project directory:
```bash
$ go get -u github.com/pandasoli/colorstrings
```

And put this in your code:
```go
import . "github.com/pandasoli/colorstrings"
```

<br/>
<br/>

BTW: If you miss some features, please make an issue and I'll reply you as soon as possible.

<br/>

## <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-idea.png' alt='Gother icon'/> Initialize from an existing string

You don't have to create a new pure string (without colors), we can use a string already with them.

```go
str, err := NewColorString("  \033[1;31mHey hi, man!\033[0m")
if err != nil { panic(err) }
```
An error will be returned if found an unparsable color.  
PS: This lib doesn't recognize position codes!

To initialize a new pure string use the same function.  
If it's really pure you won't need the `err` return.

<br/>
<div align='right'>

## <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-insomnia.png' alt='Gopher icon'/> Warnings
</div>

### Joining colors

An important thing to say is that this library makes only one color by position, when you try to put a color in a position that there's already a color, they will be "joined".

So `\033[31m\033[42m` is impossible.  
The library is gonna try to join them like this `\033[31;42m`.

Maybe I am gonna change it, but not yet for now.

<br/>

### No duplicates

A color cannot have more than one background or foreground colors.

I mean `\033[31;94m` is impossible.  
The lib is gonna make the last one the only one: `\033[94m`.

<br/>

## Features <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-smiling-blushing.png' alt='Gopher icon'/>

### `Colorize(str_fields string, position uint) error`
> /* Add a new color to the string at the `position`  
> Ex: `str.Colorize("1;38;5;200", 2)` */

It returns an error if:
- the `position` is bigger than the string
- a field cannot be parsed
- a field non-RGB is bigger than 107
- a field that doesn't expect props has some
- no more fields after 38, 48 or 58
- missing fields in 8-bit or RGB color

<br/>

### `Join() string`
> // Join the string with its colors

<br/>

### `Substring(position, length uint) (res ColorString, err error)`
> // Cut the string from the `position` to `position + length`

<br/>

### `Append(strs ...ColorString)`
> // Append another string at the end of the main one

<br/>

### `Remove(position, length uint) error`
> // Remove a piece of the string

This is gonna return an error if `position` is bigger than the string.

<br/>

### `Connect(str ColorString, position uint) error`
> // Replaces from the `position` to the length of the `str` with the `str`

This is gonna return an error if the `str` is bigger the current string,  
or the `position` + the `str`'s length is bigger than the current string.

<br/>

### `AppendAt(str ColorString, position uint) error`
> // Append another string inside the main one at the told position

This returns an error if `position` is bigger than the string.

<br/>

### `Clear()`
> // Clear the string and the colors

<br/>

### `SplitString(at string) (res []colorstring, pos []uint)`
> /* Split the main string based on the piece of string passed and returns  
> the parts where the piece was in the middle, and  
> the positions where pieces were found */

<br/>
<div align='right'>

## <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-wink.png' alt='Gopher icon'/> Credits
</div>

Gopher icons by [Egon Elbre](https://github.com/egonelbre/gophers) <img width='20' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-mind-blown.png' alt='Gopher icon'/>  
Look at the [Dev Notes](./.docs/dev-notes.md) if you would like to contribute! <img width='20' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-pirate.png' alt='Gopher icon'/>
