# Converting Markdown to HTML
One quick tool for converting markdown file into HTML
is [MS PowerShell](https://docs.microsoft.com/en-us/powershell/)

* get PowerShell
```console
# for mac:
brew install powershell
```
* write a Markdown document e.g. `readme.md`
* start powershell:
  ```console
  pwsh
  ```
* convert the document into variable `$md`
  ```console
  $md = ConvertFrom-Markdown -Path .\readme.md
  ```
* check the conversion:
  ```console
  $md | Get-Member
  ```
* export the converted HTML into a file:
  ```console
   $md.Html | Out-File -Encoding utf8 .\readme.html
   ```
* TADA!
