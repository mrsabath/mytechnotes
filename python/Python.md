# dump all attributes of the object
for slot in dir(result):
   logger.debug("%s = %s" % (slot, getattr(result, slot)))

# for runtime debugging, use:
source /os_config/choir_runtime.sh

Example:
dir()
import os
dir(os)
help(os)
help(os.wait)

###################
# debugging python
https://docs.python.org/2/library/pdb.html

l - context
n - next
s - step into
r - return

import pdb

# to set the breakpoint:
pdb.set_trace()


# during debugging:
print kwargs
(Pdb) import traceback
(Pdb) traceback.print_stack()

vi ./lib/python2.7/site-packages/heatclient/v1/stacks.py

(ccsapi-new)silesia-mac:ccsapi-new sabath$ heat stack-create -f /Users/sabath/workspace/alchemy/ccsapi/heat/
Simple_VM.yaml   asg.yaml         container.yaml   create_group.sh
(ccsapi-new)silesia-mac:ccsapi-new sabath$ heat stack-create -f /Users/sabath/workspace/alchemy/ccsapi/heat/Simple_VM.yaml MS-test2
> /Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/v1/stacks.py(118)create()
-> headers = self.client.credentials_headers()
(Pdb) l
113  	        return Stack(self, body['stack'])
114
115  	    def create(self, **kwargs):
116  	        """Create a stack."""
117  	        pdb.set_trace()
118  ->	        headers = self.client.credentials_headers()
119  	        resp, body = self.client.json_request('POST', '/stacks',
120  	                                              data=kwargs, headers=headers)
121  	        return body
122
123  	    def update(self, stack_id, **kwargs):
(Pdb) print kwargs
{'files': {}, 'disable_rollback': True, 'parameters': {}, 'stack_name': 'MS-test2', 'environment': {}, 'template': {u'outputs': None, u'heat_template_version': u'2013-05-23', u'description': u'A HOT template that holds a VM instance  The VM does nothing, it is only created.\n', u'parameters': {u'flavor': {u'default': u'm1.tiny', u'type': u'string', u'description': u'Flavor for the instance to be created', u'constraints': [{u'description': u"Value must be one of 'm1.tiny', 'm1.small', 'm1.medium', or 'm1.large'", u'allowed_values': [u'm1.tiny', u'm1.small', u'm1.medium', u'm1.large']}]}, u'image': {u'default': u'busybox:latest', u'type': u'string', u'description': u'Name or ID of the image to use for the instance. Any image should work since this template does not ask the VM to do anything.\n'}, u'network': {u'default': u'ext-net1', u'type': u'string', u'description': u'the network for the VM'}}, u'resources': {u'my_instance': {u'type': u'OS::Nova::Server', u'properties': {u'image': {u'get_param': u'image'}, u'networks': [{u'network': {u'get_param': u'network'}}], u'flavor': {u'get_param': u'flavor'}, u'availability_zone': u'docker'}}}}}
(Pdb) n
> /Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/v1/stacks.py(119)create()
-> resp, body = self.client.json_request('POST', '/stacks',
(Pdb) print headers
{'X-Auth-Key': 'e956b4c5b7f8bf4bf325bb7868d5b99a', 'X-Auth-User': '0b331950ace99163-dev_user10'}
(Pdb) s
> /Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/v1/stacks.py(120)create()
-> data=kwargs, headers=headers)
(Pdb) s
--Call--
> /Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/common/http.py(225)json_request()
-> def json_request(self, method, url, **kwargs):
(Pdb) l
220  	            creds['X-Auth-User'] = self.username
221  	        if self.password:
222  	            creds['X-Auth-Key'] = self.password
223  	        return creds
224
225  ->	    def json_request(self, method, url, **kwargs):
226  	        kwargs.setdefault('headers', {})
227  	        kwargs['headers'].setdefault('Content-Type', 'application/json')
228  	        kwargs['headers'].setdefault('Accept', 'application/json')
229
230  	        if 'data' in kwargs:
(Pdb) n
> /Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/common/http.py(226)json_request()
-> kwargs.setdefault('headers', {})
(Pdb) import traceback
(Pdb) traceback.print_stack()
  File "/Users/sabath/projects/virtualenvs/ccsapi-new/bin/heat", line 11, in <module>
    sys.exit(main())
  File "/Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/shell.py", line 443, in main
    HeatShell().main(args)
  File "/Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/shell.py", line 399, in main
    args.func(client, args)
  File "/Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/v1/shell.py", line 111, in do_stack_create
    hc.stacks.create(**fields)
  File "/Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/v1/stacks.py", line 120, in create
    data=kwargs, headers=headers)
  File "/Users/sabath/projects/virtualenvs/ccsapi-new/lib/python2.7/site-packages/heatclient/common/http.py", line 226, in json_request
    kwargs.setdefault('headers', {})
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/bdb.py", line 48, in trace_dispatch
    return self.dispatch_line(frame)
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/bdb.py", line 66, in dispatch_line
    self.user_line(frame)
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/pdb.py", line 158, in user_line
    self.interaction(frame, None)
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/pdb.py", line 210, in interaction
    self.cmdloop()
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/cmd.py", line 142, in cmdloop
    stop = self.onecmd(line)
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/pdb.py", line 279, in onecmd
    return cmd.Cmd.onecmd(self, line)
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/cmd.py", line 218, in onecmd
    return self.default(line)
  File "/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/pdb.py", line 234, in default
    exec code in globals, locals
  File "<stdin>", line 1, in <module>
(Pdb)
################

# Standard boilerplate to call the main() function.
if __name__ == '__main__':
  main()



# String:
s = 'hi'
print s[1]    # returns i
print len(s)

s.replace(stra, strb) # returns a version of string s
                      # where all instances of stra have been replaced by strb.

# replace %s with string
# repr function returns String represenation of the object.
print '%s got: %s expected: %s' % (prefix, repr(got), repr(expected))
print line,   # , removes /n
pi = 3.14
text = "this is a value of pi: " + str(pi)   # needs to convert int to str


s.lower(), s.upper() -- returns the lowercase or uppercase version of the string
s.strip() -- returns a string with whitespace removed from the start and end
s.isalpha()/s.isdigit()/s.isspace()... -- tests if all the string chars are in the various character classes
s.startswith('other'), s.endswith('other') -- tests if the string starts or ends with the given other string
s.find('other') -- searches for the given other string (not a regular expression) within s, and
				   returns the first index where it begins or -1 if not found
s.replace('old', 'new') -- returns a string where all occurrences of 'old' have been replaced by 'new'
s.split('delim') -- returns a list of substrings separated by the given delimiter. The delimiter is
                 -- not a regular expression, it's just text. 'aaa,bbb,ccc'.split(',') -> ['aaa', 'bbb', 'ccc'].
                 -- As a convenient special case s.split() (with no arguments) splits on all whitespace chars.
s.join(list) -- opposite of split(), joins the elements in the given list together using the string as
             -- the delimiter. e.g. '---'.join(['aaa', 'bbb', 'ccc']) -> aaa---bbb---ccc


# divsion
6/5    #  1
6.0/5  # 1.2
6.0//5 # 1  integer division

# <= == !=
# or, not, and


# define functions
def verbing(s):
  if len(s) < 3 :
    return s
  elif s[-3:] == 'ing' :
    return s + 'ly'
  else :
    return s + 'ing'

s.find("abc")   # returns index, -1 if not found, strings only
s.index("abc")  # returns index, exception if not found, all lists
s.index(4)      # value of position 4

# List:
#       are mutable, can change

a = [1,2,3, 'bbbb']
a = [1,2,3]
a[0] 	   # returns 1
b = a      # b points to the same array as a
b = a[:]   # copy array a into b
b = a[:-1] # copy array a without last element into b
a.pop(0)   # remove 1 and return
a.append(4) # appends 4 to the end of the list
a.insert(index, elem) # inserts the element at the given index, shifting elements to the right.

list = []
list.append(4)
list.append(5)
list.insert(1,6)    # list would be [4,6,5]

list.extend(list2)  #  adds the elements in list2 to the end of the list.
					#  Using + or += on a list is similar to using extend().

for val in a: print val    # iterate through all the elements of a
if (value in list)         # check if element in the array


# list slice
list = ['a', 'b', 'c', 'd']
print list[1:-1]   ## ['b', 'c']
list[0:2] = 'z'    ## replace ['a', 'b'] with ['z']
print list         ## ['z', 'c', 'd']

sorted(a)  # sorts numerically
help(sorted)

>>> a = [1,2,3,'aaa',4, 'bbb']
>>> sorted(a)
[1, 2, 3, 4, 'aaa', 'bbb']

>>> sorted(a,reverse=True)
['bbb', 'aaa', 4, 3, 2, 1]

sorted(a, key=len)  # sorted on the length of elements

# sorted by custom function
def Last(s): return s[-1]
sorted(a, key=Last)

>>> b = ':'.join(a)   # list all elements deliminated by :
>>> b
'bbb:aaa:1:2:3:4'

b.split(':') #  splits a string into array using ":" as deliminator
b.split()   # splits on white space

results = []

# FOR

for s in a: results.append(s)
range(20) # returns an array from 0-19
range(10,20) # returns array from 10-19

for s in range(20)


# DELETE
del a # delete a
del a[0]  # remove first item on the list
del a[-2] # remove last 2 items on the list
del dict['b']  # delete item with key 'b' from the dictionary

# list comprehensions:
# expression FOR var IN list IF expression2:

nums = [2, 8, 1, 6]
small = [ n*n for n in nums if n <= 2 ]  ## [4, 1]


# Tuple (non-mutable)

a = (1,2,3)
a = [(1,"b"), (2,"a"), (1,"a)]
(x,y) = (1,2)
>>> x
1
>>> y
2

# hashtable (dictionary)

d = {}
dict['a'] = 'alpha'
dict['g'] = 'gamma'
dict['o'] = 'omega'

# or
d = {'a':'alpha', 'b':'beta', 'g','gamma'}

print dict  ## {'a': 'alpha', 'o': 'omega', 'g': 'gamma'}

print dict['a']     ## Simple lookup, returns 'alpha'
dict['a'] = 6       ## Put new key/value into dict
'a' in dict         ## True

d.keys()
d.values()

if 'z' in dict: print dict['z']     ## Avoid KeyError
print dict.get('z')  ## None (instead of KeyError)



for k in sorted(d.keys()) : print "key:", k, "-->", d[k]

d.items() # lists all the Tuple (key, value)
for k, v in dict.items(): print k, '>', v



# file processing

def cat(filename) :
  f = open (filename, "rU")
  #  'r' is for reading, use 'w' for writing, 'r+' read-write and 'a' for append.
  #  The special mode 'rU' is the "Universal" option
  #  for text files where it's smart about converting different line-endings so they always come through as a
  #  simple '\n'.
  for line in f:
    print line,  # remove /n

  lines = f.readlines()  # reads all the lines into memory, returns a list of lines
  alltext = f.read()     # reads the whole file as a single string

  f.close()

  f = open(filename, 'w')
  f.write(string)   # for string only
  f.write(str('the answer is ', 42))

>>> f = open('workfile', 'r+')
>>> f.write('0123456789abcdef')
>>> f.seek(5)     # Go to the 6th byte in the file
>>> f.read(1)
'5'
>>> f.seek(-3, 2) # Go to the 3rd byte before the end
>>> f.read(1)
'd'
