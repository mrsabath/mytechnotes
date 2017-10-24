The repo must be called `<your_name_or_org>.github.ibm.com`
use `index.html` or `README.md`
URL: pages.github.ibm.com/<your_name_or_org>

## install Jekyll
```
curl -sSL https://get.rvm.io | bash -s stable
rvm list known
rvm install ruby-2.4.2
ruby -v
ruby 2.4.2p198 (2017-09-14 revision 59899) [x86_64-darwin15]

gem install jekyll bundler
```
## init Jekyll
```
cd ~/workspace/go-work/src/github.ibm.com/kompass/kompass.github.ibm.com
jekyll new web
cd web/
bundle exec jekyll serve
 ```
http://127.0.0.1:4000/kompass/


Run Jekyll inside the container - had problems.. :(
```
cd ~/workspace/go-work/src/github.ibm.com/kompass/kompass.github.ibm.com
docker run -it --rm -v "$PWD":/usr/src/app -p "4000:4000" --entrypoint="/bin/sh" starefossen/github-pages
# initialize it:
/usr/local/bundle/bin/jekyll new blog
# this will create all the Jekyll required code and stubs in
~/workspace/go-work/src/github.ibm.com/kompass/kompass.github.ibm.com/blog
```
Now exit the Jekyll container and update the stubs (e.g._config.yml ),
reorganize your git directories,etc

Start Jekyll web service to render your files:
```
docker run -t --rm -v "$PWD":/usr/src/app -p "4000:4000" starefossen/github-pages
```
