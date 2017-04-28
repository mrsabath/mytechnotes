# Procedures for Archiving:

1. Generate a token for the Alchemy team from this [link](https://api.slack.com/docs/oauth-test-tokens)
2. pip install [slacker](https://github.com/os/slacker)
3. git clone https://gist.github.com/fb7a070f52883849de35.git
4. cd fb7a070f52883849de35
5. (optional -dry run)
```
python slack_history.py --token [token from step 1]  --dryRun --skipChannels
```
6. execute:
```
python slack_history.py --token [token from step 1] --skipChannels
```
