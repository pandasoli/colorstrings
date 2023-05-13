<div align='center'>

# Dev Notes <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-pirate.png' alt='Gopher icon'/>

This page is for things that this library needs and you can help with.  
If you want to, just make a **Fork**, do the task and then make a **Pull Request**!

But **remember**! All the functions for `ColorString` are implemented for itself.  
In other words, there's not `SplitColorString(str, " ")`, but instead `str.SplitString(" ")`.

</div>
<br/>

- Take a look at `ColorString.AppendAt`

  Maybe we don't need to find a position.  
  Instead just `color.Position += len(self.string_)`

- Implement  
  `ColorString.SplitColor(fields string) (res []colorstring, pos []uint)`

  Instead of splitting based on the string, now based on the position of the color.  
  PS: `fields` is a color string (for example: `1;38;5;200;42`).
