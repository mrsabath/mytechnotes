# Add a co-author to commit

## Get email of the author.
Find some old commit, clone the repo, switch to the
branch and execute `git log` to get name and email.

## Amend the last commit:
```console
git commit --amend
```

Then just add the following line:
```
Co-Authored-By: Name <name@email.com>
```

## Amend some older commit
First list all the commits:

```console
git log
```
and find the commit that needs modification, e.g. `abcdef`

```console
git rebase --interactive 'abcdef^'
```

Please note the caret ^ at the end of the command,
because you need actually to rebase back to the commit before the one you wish to modify.

In the default editor, modify `pick` to `edit` in the line mentioning 'abcdef'.

Save the file and exit: git will interpret and automatically execute the commands in the file.
You will find yourself in the previous situation in which you just had created commit `abcdef`.

At this point, `abcdef` is your last commit and you can easily amend it.
Make your changes and then commit them with the command:

```console
git commit --amend
```

Then just add the following line:
```
Co-Authored-By: Name <name@email.com>
```

After this, continue:
```
git rebase --continue
```

to return back to the previous HEAD commit.

When done, review the changes:
```console
git log
```

then push:
```console
git push --force
```
